package board_test

import (
	"github.com/stretchr/testify/assert"
	"github/m-wrona/challenge-tik-tak-toe/internal/board"
	"testing"
)

func Test_Board_ShouldNotChangeWinningCoordinates(t *testing.T) {
	wc := board.GetWinningCoordinates()
	assert.Equal(t, board.WinningCoordinate{0, 1, 2}, wc[0], "unexpected 1st winning coordinate")
	wc[0] = board.WinningCoordinate{11, 22, 33}
	assert.Equal(t, board.WinningCoordinate{0, 1, 2}, board.GetWinningCoordinates()[0], "winning cordinates must be immutable")
}
