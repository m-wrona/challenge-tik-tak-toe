package challenge_tik_tak_toe

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	boardSize         = 9 //3x3
	coordinatesLength = 3
	noMove            = -1
)

type Board = []PlayerID

var boardWinningCoordinates = [][]int{
	//rows
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	//columns
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	//diagonals
	{0, 4, 8},
	{2, 4, 6},
}

func newBoard() Board {
	b := make([]PlayerID, boardSize)
	return b
}

func copyBoard(src Board) Board {
	clone := newBoard()
	copy(clone, src)
	return clone
}

func validateCoordinates(c []int) {
	if len(c) != coordinatesLength {
		panic(fmt.Sprintf("unexpected board cordinates - got: %+v, expected: %d elements", c, coordinatesLength))
	}
}

func findFirstFreeWinningCoordinate(b Board) (int, error) {
	for _, coordinates := range boardWinningCoordinates {
		validateCoordinates(coordinates)
		for _, c := range coordinates {
			if b[c] == NoPlayer {
				return c, nil
			}
		}
	}
	return noMove, errors.New("couldn't make any random move")
}
