package nonogram

import (
	"errors"
	"strings"
)

type State int

const (
	Unknown State = iota
	Filled
	Blank
)

var ErrNilPattern = errors.New("called solve with nil pattern")
var ErrCanNotSolve = errors.New("could not solve puzzle completely")

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

	for s.tryRows() != 0 || s.tryColumns() != 0 {

	}

	if s.isSolved() {
		return nil
	}

	return ErrCanNotSolve
}

func (s *Solver) tryRows() int {
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
	return changesCount
}

func (s *Solver) tryColumns() int {
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
	return changesCount
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
