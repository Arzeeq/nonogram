package nonogram

import (
	"errors"
	"strings"
)

var ErrInvalidSize = errors.New("nonogram has invalid size")
var ErrInvalidGrid = errors.New("grid has less/more data than needed")

const numBits = 64

type Nonogram struct {
	n, m int
	// each number from grid in binary form
	// represents encoded nonogram line by line
	// where 0 - empty cell, 1 - filled cell
	grid []uint64
}

func New(n, m int) *Nonogram {
	return &Nonogram{n: n, m: m, grid: make([]uint64, EncodedSize(n, m))}
}

func FromGrid(n, m int, grid []uint64) (*Nonogram, error) {
	if n <= 0 || m <= 0 {
		return nil, ErrInvalidSize
	}

	if len(grid) != EncodedSize(n, m) {
		return nil, ErrInvalidGrid
	}

	return &Nonogram{
		n:    n,
		m:    m,
		grid: grid,
	}, nil
}

func (n *Nonogram) String() string {
	return n.toString('#', '.', 0)
}

func (n *Nonogram) PrettyString() string {
	return n.toString('█', ' ', 0)
}

func (n *Nonogram) StringCaged(cage int) string {
	return n.toString('#', '.', cage)
}

func (n *Nonogram) PrettyStringCaged(cage int) string {
	return n.toString('█', ' ', cage)
}

func EncodedSize(n, m int) int {
	return (n*m + numBits - 1) / numBits
}

// cage=0 means no cage
func (n *Nonogram) toString(fill, blank rune, cage int) string {
	var b strings.Builder

	var bit int
	var idx int
	for i := range n.n {
		if cage != 0 && i != 0 && i%cage == 0 {
			for k := range n.m {
				if k != 0 && k%cage == 0 {
					b.WriteRune('┼')
				}
				b.WriteRune('─')
			}
			b.WriteRune('\n')
		}
		for j := range n.m {
			if cage != 0 && j != 0 && j%cage == 0 {
				b.WriteRune('│')
			}
			if n.grid[idx]>>bit&1 == 1 {
				b.WriteRune(fill)
			} else {
				b.WriteRune(blank)
			}
			bit = (bit + 1) % numBits
			if bit == 0 {
				idx++
			}
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func (n *Nonogram) Fill(i, j int) {
	if !(0 <= i && i < n.n) || !(0 <= j && j < n.m) {
		return
	}

	numberIdx := (i*n.n + j) / numBits
	bit := (i*n.n + j) % numBits

	n.grid[numberIdx] |= (1 << bit)
}

func (n *Nonogram) Clear(i, j int) {
	if !(0 <= i && i < n.n) || !(0 <= j && j < n.m) {
		return
	}

	numberIdx := (i*n.n + j) / numBits
	bit := (i*n.n + j) % numBits

	n.grid[numberIdx] &= (1<<numBits - 1) ^ (1 << bit)
}

func (n *Nonogram) Get(i, j int) bool {
	if !(0 <= i && i < n.n) || !(0 <= j && j < n.m) {
		return false
	}

	numberIdx := (i*n.n + j) / numBits
	bit := (i*n.n + j) % numBits

	return (n.grid[numberIdx]>>bit)&1 == 1
}
