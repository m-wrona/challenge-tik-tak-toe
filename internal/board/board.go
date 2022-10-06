package board

import (
	"fmt"
)

const (
	boardSize         = 9 //3x3
	coordinatesLength = 3
	NoMove            = -1
)

type Board []PlayerID
type WinningCoordinate []int
type WinningCoordinates []WinningCoordinate

var winningCoordinates = WinningCoordinates{
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

func init() {
	for _, c := range winningCoordinates {
		validateCoordinates(c)
	}
}

func New() Board {
	b := make([]PlayerID, boardSize)
	return b
}

func copyBoard(src Board) Board {
	clone := New()
	copy(clone, src)
	return clone
}

func validateCoordinates(c WinningCoordinate) {
	if len(c) != coordinatesLength {
		panic(fmt.Sprintf("unexpected board cordinates - got: %+v, expected: %d elements", c, coordinatesLength))
	}
}

func GetWinningCoordinates() WinningCoordinates {
	c := make(WinningCoordinates, len(winningCoordinates))
	copy(c, winningCoordinates)
	return c
}
