package main

import (
	"fmt"
	"nonogram"
	"time"
)

func main() {
	rows := nonogram.FillPattern{
		{3, 3},
		{5, 7, 5},
		{2, 12, 2},
		{2, 12, 2},
		{18},
		{17},
		{18},
		{4, 6, 6},
		{4, 6, 6},
		{3, 1, 2, 5},
		{5, 3, 7},
		{4, 3, 6},
		{4, 1, 1, 1, 5},
		{3, 5, 5},
		{3, 7},
		{2, 2, 5},
		{10, 6},
		{4, 6},
		{4, 7},
		{4, 6},
		{4, 7},
		{7, 3, 3, 2},
		{9, 6, 2, 2},
		{2, 9, 6, 3, 3},
		{1, 7, 14},
		{1, 7, 13},
		{2, 8, 10},
		{11, 10},
		{6, 13},
		{10},
	}

	columns := nonogram.FillPattern{
		{4},
		{2, 2},
		{2, 2},
		{3, 3},
		{8},
		{3, 6, 10},
		{5, 8, 11},
		{2, 11, 12},
		{2, 5, 5, 12},
		{6, 1, 5, 6},
		{5, 1, 1, 3, 6},
		{7, 1, 1, 2, 3},
		{8, 2, 1, 1, 2, 2},
		{8, 4, 1, 2, 2},
		{8, 2, 1, 1, 3, 3},
		{8, 1, 1, 8},
		{9, 1, 1, 8},
		{6, 1, 2, 7},
		{6, 1, 2, 7},
		{7, 5, 7},
		{14, 6},
		{2, 13, 5},
		{2, 15, 6},
		{28},
		{3, 5, 13},
		{7, 2},
		{5, 2},
		{4, 3},
		{7},
		{4},
	}

	start := time.Now()

	var s nonogram.Solver
	err := s.Solve(rows, columns)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("solving nonogram took %s time\n", time.Since(start))

	// show solving result
	fmt.Println(s.StringCaged(5))
	// save result to png
	s.SavePNG("output.png", 10)
}
