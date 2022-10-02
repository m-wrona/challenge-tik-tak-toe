package challenge_tik_tak_toe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Game_ShouldStartAGameWithTwoDifferentPlayers(t *testing.T) {
	p1 := HumanPlayer{id: PlayerID(1)}
	p2 := AiPlayer{id: PlayerID(2)}
	_, err := Start(p1, p2)
	assert.Nil(t, err, "two different players must start a game")
}

func Test_Game_ShouldNotStartAGameWithSinglePlayer(t *testing.T) {
	p1 := HumanPlayer{id: PlayerID(1)}
	_, err := Start(p1, p1)
	assert.NotNil(t, err, "only two different players can play a game")
	_, err = Start(nil, p1)
	assert.NotNil(t, err, "only two different players can play a game")
	_, err = Start(p1, nil)
	assert.NotNil(t, err, "only two different players can play a game")
}

func Test_Game_ShouldNotStartAGameWithPlayersIllegalID(t *testing.T) {
	invalid := HumanPlayer{id: NoPlayer}
	valid := AiPlayer{id: PlayerID(1)}
	_, err := Start(invalid, valid)
	assert.NotNil(t, err, "player X must have valid ID")
	_, err = Start(valid, invalid)
	assert.NotNil(t, err, "player Y must have valid ID")
}

//func Test_Game_PlayerShouldMakeAMove(t *testing.T) {
//	p1 := HumanPlayer{id: PlayerID(1)}
//	g := Game{}
//	assert.Nil(t, g.Join(p1), "player couldn't join a game")
//	s0, err := g.Start()
//	assert.Nil(t, err, "couldn't start a game")
//
//	s1, err := s0.Move(context.Background(), p1)
//	assert.Nil(t, err, "player couldn't make a move")
//	assert.NotNil(t, s1, "new game state not present after move")
//	assert.Equal(t, NoPlayer, s1.HasPlayerWon(), "nobody mustn't win the game after 1st move")
//}
//
//func Test_GameState_ShouldMarkGameAsWonByMarkingWholeRow(t *testing.T) {
//	p1 := HumanPlayer{id: PlayerID(1)}
//	s0 := GameState{
//		board: newBoard(),
//	}
//	s0.board[0][0] = p1.ID()
//	s0.board[0][1] = p1.ID()
//	s0.board[0][2] = p1.ID()
//	assert.Equal(t, p1.ID(), s0.HasPlayerWon())
//}
//
//func Test_GameState_ShouldMarkGameAsWonByMarkingWholeColumn(t *testing.T) {
//	p1 := HumanPlayer{id: PlayerID(1)}
//	s0 := GameState{
//		board: newBoard(),
//	}
//	s0.board[0][0] = p1.ID()
//	s0.board[1][0] = p1.ID()
//	s0.board[2][0] = p1.ID()
//	assert.Equal(t, p1.ID(), s0.HasPlayerWon())
//}
//
//func Test_GameState_ShouldMarkGameAsWonByMarkingWholeLeftDiagonal(t *testing.T) {
//	p1 := HumanPlayer{id: PlayerID(1)}
//	s0 := GameState{
//		board: newBoard(),
//	}
//	s0.board[0][0] = p1.ID()
//	s0.board[1][1] = p1.ID()
//	s0.board[2][2] = p1.ID()
//	assert.Equal(t, p1.ID(), s0.HasPlayerWon())
//}
//
//func Test_GameState_ShouldMarkGameAsWonByMarkingWholeRightDiagonal(t *testing.T) {
//	p1 := HumanPlayer{id: PlayerID(1)}
//	s0 := GameState{
//		board: newBoard(),
//	}
//	s0.board[0][2] = p1.ID()
//	s0.board[1][1] = p1.ID()
//	s0.board[2][0] = p1.ID()
//	assert.Equal(t, p1.ID(), s0.HasPlayerWon())
//}
//
//func Test_GamePlay_PlayerShouldWinGameByMarkingWholeRow(t *testing.T) {
//	p1 := HumanPlayer{id: PlayerID(1)}
//	s0 := GameState{
//		board: newBoard(),
//	}
//	s0.board[0][0] = p1.ID()
//	s0.board[0][1] = p1.ID()
//	assert.Equal(t, NoPlayer, s0.HasPlayerWon(), "cannot make next move since game is won already")
//
//	s1, err := s0.Move(context.Background(), p1)
//	assert.Nil(t, err, "player couldn't make a move")
//	assert.NotNil(t, s1, "new game state not present after move")
//	assert.Equal(t, p1.ID(), s1.HasPlayerWon(), "player should win the game")
//}
//
//func Test_GamePlay_PlayerShouldWinGameByMarkingWholeColumn(t *testing.T) {
//	p1 := HumanPlayer{id: PlayerID(1)}
//	p2 := HumanPlayer{id: PlayerID(2)}
//	s0 := GameState{
//		board: newBoard(),
//	}
//	s0.board[0][0] = p1.ID()
//	s0.board[1][0] = p1.ID()
//
//	s0.board[0][1] = p2.ID()
//	assert.Equal(t, NoPlayer, s0.HasPlayerWon(), "cannot make next move since game is won already")
//
//	s1, err := s0.Move(context.Background(), p1)
//	assert.Nil(t, err, "player couldn't make a move")
//	assert.NotNil(t, s1, "new game state not present after move")
//	assert.Equal(t, p1.ID(), s1.HasPlayerWon(), "player should win the game")
//}
//
//func Test(t *testing.T) {
//	a := []int{1, 2, 3}
//	fmt.Printf("%+v\n", a)
//	change(a)
//	fmt.Printf("%+v\n", a)
//}
//
//func change(a []int) {
//	a[1] = a[1] * 10
//}
