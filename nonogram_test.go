package nonogram_test

import (
	"math/rand"
	"nonogram"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		n, m     int
		grid     []uint64
		expected string
	}{
		{
			name: "small size 1",
			n:    4,
			m:    4,
			grid: []uint64{12281711624147758220},
			expected: `..##
...#
..##
.###
`,
		},
		{
			name: "small size 2",
			n:    4,
			m:    4,
			grid: []uint64{13832420446842977395},
			expected: `##..
###.
..##
####
`,
		},
		{
			name: "small size 3",
			n:    4,
			m:    4,
			grid: []uint64{2445733788184770683},
			expected: `##.#
###.
...#
###.
`,
		},
		{
			name: "small size 4",
			n:    4,
			m:    4,
			grid: []uint64{9133373410999700124},
			expected: `..##
#..#
.#.#
...#
`,
		},
		{
			name: "small size 5",
			n:    4,
			m:    4,
			grid: []uint64{7740365830381069877},
			expected: `#.#.
##..
.##.
.##.
`,
		},
		{
			name: "big size 1",
			n:    15,
			m:    15,
			grid: []uint64{4802168399298677416, 8946851233348289658, 2462032354927984822, 289766607327246001},
			expected: `...#.#.#.#.#..#
##.##..#####.#.
#.###.##.#.#.##
#.#..#..#.#.#..
..#..#.####...#
....##.#######.
#.#.#.##.#..###
..##..##..#.#..
..#####..##.##.
#......#.......
..#.#..####....
##.#.#..###.#.#
.#...#...#..#..
.##.#.##.##.#.#
.....##...##.##
`,
		},
		{
			name: "big size 2",
			n:    15,
			m:    15,
			grid: []uint64{11907156946780570069, 11174566804707492334, 17606257748242159461, 7889916177122246429},
			expected: `#.#.#.###.....#
..###.#.##.....
.....#.#.....##
#.#.#####..#.#.
.#.#.###.####..
...#..##.###.#.
..##.##.##..#.#
..#......#.#...
##.##..##.#..##
.###.#.#.#...##
####.##.####...
#...########.#.
#.#...#.#####.#
##...######.#.#
.#####..#..####
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gram, err := nonogram.FromGrid(tt.n, tt.m, tt.grid)

			require.NoError(t, err)
			require.Equal(t, tt.expected, gram.String())
		})
	}
}

func TestStringCaged(t *testing.T) {
	tests := []struct {
		name     string
		n, m     int
		grid     []uint64
		expected string
		cage     int
	}{
		{
			name: "small size caged 1",
			n:    4,
			m:    4,
			grid: []uint64{12281711624147758220},
			cage: 2,
			expected: `..│##
..│.#
──┼──
..│##
.#│##
`,
		},
		{
			name: "small size caged 2",
			n:    4,
			m:    4,
			grid: []uint64{13832420446842977395},
			cage: 2,
			expected: `##│..
##│#.
──┼──
..│##
##│##
`,
		},
		{
			name: "small size caged 3",
			n:    4,
			m:    4,
			grid: []uint64{2445733788184770683},
			cage: 2,
			expected: `##│.#
##│#.
──┼──
..│.#
##│#.
`,
		},
		{
			name: "small size caged 4",
			n:    4,
			m:    4,
			grid: []uint64{9133373410999700124},
			cage: 2,
			expected: `..│##
#.│.#
──┼──
.#│.#
..│.#
`,
		},
		{
			name: "small size caged 5",
			n:    4,
			m:    4,
			grid: []uint64{7740365830381069877},
			cage: 2,
			expected: `#.│#.
##│..
──┼──
.#│#.
.#│#.
`,
		},
		{
			name: "big size caged 1",
			n:    15,
			m:    15,
			grid: []uint64{4802168399298677416, 8946851233348289658, 2462032354927984822, 289766607327246001},
			cage: 5,
			expected: `...#.│#.#.#│.#..#
##.##│..###│##.#.
#.###│.##.#│.#.##
#.#..│#..#.│#.#..
..#..│#.###│#...#
─────┼─────┼─────
....#│#.###│####.
#.#.#│.##.#│..###
..##.│.##..│#.#..
..###│##..#│#.##.
#....│..#..│.....
─────┼─────┼─────
..#.#│..###│#....
##.#.│#..##│#.#.#
.#...│#...#│..#..
.##.#│.##.#│#.#.#
.....│##...│##.##
`,
		},
		{
			name: "big size caged 2",
			n:    15,
			m:    15,
			grid: []uint64{11907156946780570069, 11174566804707492334, 17606257748242159461, 7889916177122246429},
			cage: 5,
			expected: `#.#.#│.###.│....#
..###│.#.##│.....
.....│#.#..│...##
#.#.#│####.│.#.#.
.#.#.│###.#│###..
─────┼─────┼─────
...#.│.##.#│##.#.
..##.│##.##│..#.#
..#..│....#│.#...
##.##│..##.│#..##
.###.│#.#.#│...##
─────┼─────┼─────
####.│##.##│##...
#...#│#####│##.#.
#.#..│.#.##│###.#
##...│#####│#.#.#
.####│#..#.│.####
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gram, err := nonogram.FromGrid(tt.n, tt.m, tt.grid)

			require.NoError(t, err)
			require.Equal(t, tt.expected, gram.StringCaged(tt.cage))
		})
	}
}

func TestPrettyString(t *testing.T) {
	tests := []struct {
		name     string
		n, m     int
		grid     []uint64
		expected string
	}{
		{
			name: "small size",
			n:    4,
			m:    4,
			grid: []uint64{14469453583798471129},
			expected: `█  █
█ ██
█  █
   █
`,
		},
		{
			name: "big size",
			n:    15,
			m:    15,
			grid: []uint64{3854656256588814351, 5899494372590017579, 13723995212963435137, 9560539225142963483},
			expected: `████       ██  
█ ██ ██  █  █  
█ █   ███ █ ███
██  ██████ █ █ 
██  ██ █ █    █
██ ██  █  █ █  
           ███ 
██ ██  █████ ██
█   █ █ █      
█ █  █ ██   ██ 
██████ ███ ███ 
███  ██ ██ █ █ 
███  █████ ███ 
██   █ ██ █ ██ 
█ █ █████ ███ █
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gram, err := nonogram.FromGrid(tt.n, tt.m, tt.grid)

			require.NoError(t, err)
			require.Equal(t, tt.expected, gram.PrettyString())
		})
	}
}

func TestPrettyStringCaged(t *testing.T) {
	tests := []struct {
		name     string
		n, m     int
		grid     []uint64
		cage     int
		expected string
	}{
		{
			name: "small size caged",
			n:    4,
			m:    4,
			grid: []uint64{14469453583798471129},
			cage: 2,
			expected: `█ │ █
█ │██
──┼──
█ │ █
  │ █
`,
		},
		{
			name: "big size caged",
			n:    15,
			m:    15,
			grid: []uint64{3854656256588814351, 5899494372590017579, 13723995212963435137, 9560539225142963483},
			cage: 5,
			expected: `████ │     │ ██  
█ ██ │██  █│  █  
█ █  │ ███ │█ ███
██  █│█████│ █ █ 
██  █│█ █ █│    █
─────┼─────┼─────
██ ██│  █  │█ █  
     │     │ ███ 
██ ██│  ███│██ ██
█   █│ █ █ │     
█ █  │█ ██ │  ██ 
─────┼─────┼─────
█████│█ ███│ ███ 
███  │██ ██│ █ █ 
███  │█████│ ███ 
██   │█ ██ │█ ██ 
█ █ █│████ │███ █
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gram, err := nonogram.FromGrid(tt.n, tt.m, tt.grid)

			require.NoError(t, err)
			require.Equal(t, tt.expected, gram.PrettyStringCaged(tt.cage))
		})
	}
}

func TestFill(t *testing.T) {
	tests := []struct {
		name     string
		n, m     int
		i, j     int
		expected string
	}{
		{
			name: "small size 1",
			n:    4,
			m:    4,
			i:    0,
			j:    0,
			expected: `#...
....
....
....
`,
		},
		{
			name: "small size 2",
			n:    4,
			m:    4,
			i:    3,
			j:    3,
			expected: `....
....
....
...#
`,
		},
		{
			name: "small size 3",
			n:    4,
			m:    4,
			i:    1,
			j:    3,
			expected: `....
...#
....
....
`,
		},
		{
			name: "small size 4",
			n:    4,
			m:    4,
			i:    2,
			j:    1,
			expected: `....
....
.#..
....
`,
		},
		{
			name: "big size",
			n:    15,
			m:    15,
			i:    9,
			j:    7,
			expected: `...............
...............
...............
...............
...............
...............
...............
...............
...............
.......#.......
...............
...............
...............
...............
...............
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gram := nonogram.New(tt.n, tt.m)

			gram.Fill(tt.i, tt.j)
			require.Equal(t, tt.expected, gram.String())
		})
	}
}

func TestFillClear(t *testing.T) {
	gram := nonogram.New(15, 15)
	var board [15][15]int
	for range 10000 {
		i := rand.Int() % 15
		j := rand.Int() % 15
		if rand.Int()%2 == 0 {
			gram.Fill(i, j)
			board[i][j] = 1
		} else {
			gram.Clear(i, j)
			board[i][j] = 0
		}
	}

	for i := range 15 {
		for j := range 15 {
			require.Equal(t, board[i][j] == 1, gram.Get(i, j))
		}
	}
}
