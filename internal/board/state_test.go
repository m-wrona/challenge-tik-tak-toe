package board

import (
	"fmt"
	"github.com/stretchr/testify/assert"

	"testing"
)

func Test_State_ShouldMarkGameAsFinishedWithWinner(t *testing.T) {
	for _, coordinates := range GetWinningCoordinates() {
		t.Run(fmt.Sprintf("%+v", coordinates), func(t *testing.T) {
			var playerID PlayerID = 1
			s := GameState{
				board: New(),
			}
			s.board[coordinates[0]] = playerID
			s.board[coordinates[1]] = playerID
			s.board[coordinates[2]] = playerID
			winnerID, isFinished := s.IsFinished()
			assert.Equal(t, playerID, winnerID, "wrong winner ID returned in finished game")
			assert.True(t, isFinished, "player has won the game so it must be marked as finished")
		})
	}
}

func Test_State_ShouldMarkGameAsOnGoing(t *testing.T) {
	for _, coordinates := range GetWinningCoordinates() {
		t.Run(fmt.Sprintf("%+v", coordinates), func(t *testing.T) {
			var playerID PlayerID = 1
			var player2ID PlayerID = 2
			s := GameState{
				board: New(),
			}
			s.board[coordinates[0]] = playerID
			s.board[coordinates[1]] = playerID
			s.board[coordinates[2]] = player2ID
			winnerID, isFinished := s.IsFinished()
			assert.Equal(t, NoPlayer, winnerID, "nobody has won the game yet")
			assert.False(t, isFinished, "some fields are empty so game is still on-going")
		})
	}
}

func Test_State_ShouldMarkGameAsDraw(t *testing.T) {
	s := GameState{
		board: New(),
	}
	for idx := range s.board {
		s.board[idx] = PlayerID(idx + 1)
	}
	winnerID, isFinished := s.IsFinished()
	assert.Equal(t, NoPlayer, winnerID, "nobody has won the game yet")
	assert.True(t, isFinished, "no fields are empty so must be a draw")
}
