package nonogram

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type State int

const (
	Unknown State = iota
	Filled
	Blank
)

var ErrNilPattern = errors.New("called solve with nil pattern")
var ErrContradiction = errors.New("found contradiction")
var ErrCanNotSolve = errors.New("can not solve this puzzle completely")

type Solver struct {
	n, m    int
	grid    [][]State
	rows    FillPattern
	columns FillPattern
}

func (s *Solver) Solve(rows FillPattern, columns FillPattern) error {
	if rows == nil || columns == nil {
		return ErrNilPattern
	}

	s.n = len(rows)
	s.m = len(columns)
	s.rows = rows
	s.columns = columns
	s.grid = make([][]State, s.n)
	for i := range s.n {
		s.grid[i] = make([]State, s.m)
	}

	return s.solve()
}

func (s *Solver) solve() error {
	for {
		rowChanges, err := s.tryRows()
		if err != nil {
			return err
		}

		columnChanges, err := s.tryColumns()
		if err != nil {
			return err
		}

		if rowChanges == 0 && columnChanges == 0 {
			if s.isSolved() {
				return nil
			}

			if !s.tryToGuess() {
				return ErrCanNotSolve
			} else {
				return nil
			}
		}
	}
}

// this function tries to Fill all unknown cells from the grid
// and check if contradiction is occured
func (s *Solver) tryToGuess() bool {
	for i := range s.n {
		for j := range s.m {
			if s.grid[i][j] != Unknown {
				continue
			}

			newSolver := copySolver(s)
			newSolver.grid[i][j] = Filled
			if err := newSolver.solve(); err != nil {
				continue
			}

			for i1 := range s.n {
				for j1 := range s.m {
					s.grid[i1][j1] = newSolver.grid[i1][j1]
				}
			}

			return s.isSolved()
		}
	}

	return false
}

// returns count of changes done and error if contradiction is found
func (s *Solver) tryRows() (int, error) {
	changesCount := 0

	row := make([]State, s.n)
	for rowIdx := range s.n {
		var v variant

		rowFills := make([]int, s.n)
		varCount := 0
		for curVar := range v.Provide(s.n, s.rows[rowIdx]) {
			fillWith(row, curVar, s.rows[rowIdx])

			isSuitable := true
			for column := range s.m {
				if !isPossible(s.grid[rowIdx][column], row[column]) {
					isSuitable = false
					break
				}
			}

			if !isSuitable {
				continue
			}

			varCount++
			for i := range s.n {
				if row[i] == Filled {
					rowFills[i]++
				}
			}
		}

		if varCount == 0 {
			return 0, ErrContradiction
		}

		for column := range rowFills {
			if rowFills[column] == varCount && s.grid[rowIdx][column] == Unknown {
				s.grid[rowIdx][column] = Filled
				changesCount++
			} else if rowFills[column] == 0 && s.grid[rowIdx][column] == Unknown {
				s.grid[rowIdx][column] = Blank
				changesCount++
			}
		}
	}
	return changesCount, nil
}

// returns count of changes done and error if contradiction is found
func (s *Solver) tryColumns() (int, error) {
	changesCount := 0

	column := make([]State, s.m)
	for columnIdx := range s.m {
		var v variant

		columnFills := make([]int, s.m)
		varCount := 0
		for curVar := range v.Provide(s.m, s.columns[columnIdx]) {
			fillWith(column, curVar, s.columns[columnIdx])

			isSuitable := true
			for row := range s.n {
				if !isPossible(s.grid[row][columnIdx], column[row]) {
					isSuitable = false
					break
				}
			}

			if !isSuitable {
				continue
			}

			varCount++
			for i := range s.m {
				if column[i] == Filled {
					columnFills[i]++
				}
			}
		}

		if varCount == 0 {
			return 0, ErrContradiction
		}

		for row := range columnFills {
			if columnFills[row] == varCount && s.grid[row][columnIdx] == Unknown {
				s.grid[row][columnIdx] = Filled
				changesCount++
			} else if columnFills[row] == 0 && s.grid[row][columnIdx] == Unknown {
				s.grid[row][columnIdx] = Blank
				changesCount++
			}
		}
	}
	return changesCount, nil
}

// this function fills [arr] array with unbroken blocks
// which starts in indexes stored in start
// with length stored in [block]
func fillWith(arr []State, start []int, block []int) {
	for i := range arr {
		arr[i] = Blank
	}
	for i := range start {
		for j := range block[i] {
			arr[start[i]+j] = Filled
		}
	}
}

// [current] is a state of certain cell in grid and
// [possible] is a state of the same cell, but we check if
// we can set that cell to [possible] state
func isPossible(current, possible State) bool {
	if current == Unknown {
		return true
	}

	return current == possible
}

func (s *Solver) isSolved() bool {
	for i := range s.n {
		for j := range s.m {
			if s.grid[i][j] == Unknown {
				return false
			}
		}
	}

	return true
}

// cage=0 means no cage
func (s *Solver) toString(fill, empty, unknown rune, cage int) string {
	var b strings.Builder

	for i := range s.n {
		if cage != 0 && i != 0 && i%cage == 0 {
			for k := range s.m {
				if k != 0 && k%cage == 0 {
					b.WriteRune('┼')
				}
				b.WriteRune('─')
			}
			b.WriteRune('\n')
		}
		for j := range s.m {
			if cage != 0 && j != 0 && j%cage == 0 {
				b.WriteRune('│')
			}
			if s.grid[i][j] == Filled {
				b.WriteRune(fill)
			} else if s.grid[i][j] == Blank {
				b.WriteRune(empty)
			} else if s.grid[i][j] == Unknown {
				b.WriteRune(unknown)
			}
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func (s *Solver) String() string {
	return s.toString('#', 'x', '.', 0)
}

func (s *Solver) PrettyString() string {
	return s.toString('█', '╳', ' ', 0)
}

func (s *Solver) StringCaged(cage int) string {
	return s.toString('#', 'x', '.', cage)
}

func (s *Solver) PrettyStringCaged(cage int) string {
	return s.toString('█', '╳', ' ', cage)
}

// When solver saves png, it paints cells in black if it's Filled,
// in white if it's Blank and in red if it's Unknown
func (s *Solver) SavePNG(name string, scale int) error {
	img := image.NewRGBA(image.Rect(0, 0, s.m*scale, s.n*scale))

	for i := 0; i < s.n*scale; i++ {
		for j := 0; j < s.m*scale; j++ {
			var c color.RGBA
			if s.grid[i/scale][j/scale] == Filled {
				c = color.RGBA{0, 0, 0, 255}
			} else if s.grid[i/scale][j/scale] == Blank {
				c = color.RGBA{255, 255, 255, 255}
			} else {
				c = color.RGBA{255, 0, 0, 255}
			}
			img.Set(j, i, c)
		}
	}

	f, err := os.Create(name)
	if err != nil {
		return err
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Solver) ToNonogram() *Nonogram {
	nono := New(s.n, s.m)

	for i := 0; i < s.n; i++ {
		for j := 0; j < s.m; j++ {
			if s.grid[i][j] == Filled {
				nono.Fill(i, j)
			}
		}
	}

	return nono
}

func copySolver(s *Solver) *Solver {
	newSolver := Solver{
		n:       s.n,
		m:       s.m,
		rows:    s.rows,
		columns: s.columns,
		grid:    make([][]State, s.n),
	}

	for i := range s.n {
		newSolver.grid[i] = make([]State, s.m)
	}

	for i := range s.n {
		for j := range s.m {
			newSolver.grid[i][j] = s.grid[i][j]
		}
	}

	return &newSolver
}
