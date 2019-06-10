package solver

import (
	"fmt"
	g "n-puzzle/golib"
)

// look at current board for potential moves
// count n potential moves
// check if move is valid
// generate n boards for n moves
// ... how to check if board is duplicate?

// Checks which directions the zero tile can move in
func checkMoves(width int, i int) (int, int, int, int) {
	up, down, left, right := 0, 0, 0, 0
	length := width * width
	// if zero is in top row
	if i < length-width {
		down = 1
	}
	// if zero is not in top row
	if i >= width && i < length {
		up = 1
	}
	// if zero is in right column
	maxMod := (length - 1) % width
	// if zero is in left column
	minMod := 0
	if i%width != minMod {
		left = 1
	}
	if i%width != maxMod {
		right = 1
	}
	return up, down, left, right
}

func MovePieces(puzzle []int, size int) {

	empty := g.FindIndexSlice(puzzle, 0)
	up, down, left, right := checkMoves(size, empty)
	fmt.Printf("up %d, down %d, left %d, right %d\n", up, down, left, right)

	moves := up + down + left + right
	fmt.Printf("\nthere are %d possible moves\n", moves)
	//return new

	//	for i := 0; i <= moves; i++ {
	//		new := make([]int, size*size)
	// create new board according to each possible move,
	// to store in decision tree
	//	}
}
