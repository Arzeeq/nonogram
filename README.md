# nonogram

Nonogram solver written in pure go

## Problem

Nonograms, also known as Hanjie, Paint by Numbers, Griddlers, Pic-a-Pix, and Picross, are picture logic puzzles in which cells in a grid must be colored or left blank according to numbers at the edges of the grid to reveal a hidden picture.
[See](https://en.wikipedia.org/wiki/Nonogram#) for more details.
Example of solved nonogram:
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/9f/Nonogram_wiki.svg/1920px-Nonogram_wiki.svg.png" alt="drawing" width="600"/>

## Get started

```go
package main
import (
    "github.com/Arzeeq/nonogram"
)

func main() {
    rows := nonogram.FillPattern{{1}, {1, 1}, {1}}
	columns := nonogram.FillPattern{{1}, {1, 1}, {1}}

	var s nonogram.Solver
	err := s.Solve(rows, columns)
	if err != nil {
		panic(err)
	}
	print(s.String())
}
```
See [example](example/example.go) and solver [app](cmd/solver/main.go) for more details.

