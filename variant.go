package nonogram

import "sync"

type variant struct {
	once  sync.Once
	ch    chan []int
	n     int
	cur   []int
	pref  []int
	block []int
}

func (v *variant) Provide(n int, block []int) chan []int {
	v.once.Do(func() {
		v.ch = make(chan []int)
		v.pref = make([]int, len(block)+1)
		v.block = block
		v.n = n

		for i := 1; i <= len(block); i++ {
			v.pref[i] = v.pref[i-1] + block[i-1]
		}

		go func() {
			v.generateVariant(0)
			close(v.ch)
		}()
	})

	return v.ch
}

func (v *variant) generateVariant(idx int) {
	if idx == len(v.block) {
		cp := make([]int, len(v.cur))
		copy(cp, v.cur)
		v.ch <- cp
		return
	}

	begin := 0
	end := v.n - (v.pref[len(v.block)] - v.pref[idx]) - (len(v.block) - idx - 1)
	if idx > 0 {
		begin = v.cur[idx-1] + v.block[idx-1] + 1
	}

	for i := begin; i <= end; i++ {
		v.cur = append(v.cur, i)
		v.generateVariant(idx + 1)
		v.cur = v.cur[:len(v.cur)-1]
	}
}
