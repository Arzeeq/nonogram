package nonogram

import (
	"math/rand"
)

func Gen(n, m int) *Nonogram {
	size := EncodedSize(n, m)

	data := make([]uint64, size)
	for i := range size {
		data[i] = rand.Uint64()
	}

	res, _ := FromGrid(n, m, data)

	return res
}
