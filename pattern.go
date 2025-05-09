package nonogram

// FillPattern is a numbers written at the edge of puzzle.
// These numbers show the len of unbroken lines of filled-in
// squares there are in any given row or column.
// There are at least one blank square between filled-in lines.
//         22
//      09922440
//     ┌────────┐
//    0│........│
//    4│.####...│
//    6│.######.│
//    6│.######.│
//  2 2│.##..##.│
//  2 2│.##..##.│
//    4│.####...│
//    2│.##.....│
//    2│.##.....│
//    2│.##.....│
//    0│........│
//     └────────┘
type FillPattern [][]int

// return rows and columns FillPattern respectively
func (n *Nonogram) FillPatterns() (FillPattern, FillPattern) {
	return n.rowFillPattern(), n.columnFillPattern()
}

func (n *Nonogram) rowFillPattern() FillPattern {
	p := make(FillPattern, n.n)
	for row := range n.n {
		blockSize := 0
		for column := range n.m {
			if n.Get(row, column) {
				blockSize++
			} else {
				if blockSize > 0 {
					p[row] = append(p[row], blockSize)
				}
				blockSize = 0
			}
		}

		if blockSize != 0 || len(p[row]) == 0 {
			p[row] = append(p[row], blockSize)
		}
	}

	return p
}

func (n *Nonogram) columnFillPattern() FillPattern {
	p := make(FillPattern, n.m)
	for column := range n.m {
		blockSize := 0
		for row := range n.n {
			if n.Get(row, column) {
				blockSize++
			} else {
				if blockSize > 0 {
					p[column] = append(p[column], blockSize)
				}
				blockSize = 0
			}
		}

		if blockSize != 0 || len(p[column]) == 0 {
			p[column] = append(p[column], blockSize)
		}
	}

	return p
}
