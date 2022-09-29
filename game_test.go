package challenge_tik_tak_toe

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PLayer_ShouldRegisterNewPlayers(t *testing.T) {
	p1 := HumanPlayer{id: PlayerID(1)}
	p2 := AiPlayer{id: PlayerID(2)}
	g := Game{}
	assert.Nil(t, g.Join(p1), "human player couldn't join a game")
	assert.Nil(t, g.Join(p2), "AI player couldn't join a game")
}

func Test_Player_ShouldNotRegisterTheSamePlayerTwice(t *testing.T) {
	p1 := HumanPlayer{id: PlayerID(1)}
	g := Game{}
	assert.Nil(t, g.Join(p1), "human player couldn't join a game")
	assert.NotNil(t, g.Join(p1), "the same player has joined the game twice")
}

func Test_Game_PlayerShouldMakeAMove(t *testing.T) {
	p1 := HumanPlayer{id: PlayerID(1)}
	g := Game{}
	assert.Nil(t, g.Join(p1), "player couldn't join a game")
	s0, err := g.Start()
	assert.Nil(t, err, "couldn't start a game")

	s1, err := s0.Move(context.Background(), p1)
	assert.Nil(t, err, "player couldn't make a move")
	assert.NotNil(t, s1, "new game state not present after move")
	assert.Equal(t, NoPlayer, s1.HasPlayerWon(), "nobody mustn't win the game after 1st move")
}

func Test_GameState_ShouldMarkGameAsWonByMarkingWholeRow(t *testing.T) {
	p1 := HumanPlayer{id: PlayerID(1)}
	s0 := GameState{
		board: newBoard(),
	}
	s0.board[0][0] = p1.ID()
	s0.board[0][1] = p1.ID()
	s0.board[0][2] = p1.ID()
	assert.Equal(t, p1.ID(), s0.HasPlayerWon())
}

func Test_GameState_ShouldMarkGameAsWonByMarkingWholeColumn(t *testing.T) {
	p1 := HumanPlayer{id: PlayerID(1)}
	s0 := GameState{
		board: newBoard(),
	}
	s0.board[0][0] = p1.ID()
	s0.board[1][0] = p1.ID()
	s0.board[2][0] = p1.ID()
	assert.Equal(t, p1.ID(), s0.HasPlayerWon())
}

func Test_GameState_ShouldMarkGameAsWonByMarkingWholeLeftDiagonal(t *testing.T) {
	p1 := HumanPlayer{id: PlayerID(1)}
	s0 := GameState{
		board: newBoard(),
	}
	s0.board[0][0] = p1.ID()
	s0.board[1][1] = p1.ID()
	s0.board[2][2] = p1.ID()
	assert.Equal(t, p1.ID(), s0.HasPlayerWon())
}

func Test_GameState_ShouldMarkGameAsWonByMarkingWholeRightDiagonal(t *testing.T) {
	p1 := HumanPlayer{id: PlayerID(1)}
	s0 := GameState{
		board: newBoard(),
	}
	s0.board[0][2] = p1.ID()
	s0.board[1][1] = p1.ID()
	s0.board[2][0] = p1.ID()
	assert.Equal(t, p1.ID(), s0.HasPlayerWon())
}

func Test_GamePlay_PlayerShouldWinGameByMarkingWholeRow(t *testing.T) {
	p1 := HumanPlayer{id: PlayerID(1)}
	s0 := GameState{
		board: newBoard(),
	}
	s0.board[0][0] = p1.ID()
	s0.board[0][1] = p1.ID()
	assert.Equal(t, NoPlayer, s0.HasPlayerWon(), "cannot make next move since game is won already")

	s1, err := s0.Move(context.Background(), p1)
	assert.Nil(t, err, "player couldn't make a move")
	assert.NotNil(t, s1, "new game state not present after move")
	assert.Equal(t, p1.ID(), s1.HasPlayerWon(), "player should win the game")
}

func Test_GamePlay_PlayerShouldWinGameByMarkingWholeColumn(t *testing.T) {
	p1 := HumanPlayer{id: PlayerID(1)}
	p2 := HumanPlayer{id: PlayerID(2)}
	s0 := GameState{
		board: newBoard(),
	}
	s0.board[0][0] = p1.ID()
	s0.board[1][0] = p1.ID()

	s0.board[0][1] = p2.ID()
	assert.Equal(t, NoPlayer, s0.HasPlayerWon(), "cannot make next move since game is won already")

	s1, err := s0.Move(context.Background(), p1)
	assert.Nil(t, err, "player couldn't make a move")
	assert.NotNil(t, s1, "new game state not present after move")
	assert.Equal(t, p1.ID(), s1.HasPlayerWon(), "player should win the game")
}
