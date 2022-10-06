package internal

import (
	"github.com/stretchr/testify/assert"
	"github/m-wrona/challenge-tik-tak-toe/internal/board"
	"github/m-wrona/challenge-tik-tak-toe/internal/players"
	"testing"
)

func Test_Game_ShouldStartAGameWithTwoDifferentPlayers(t *testing.T) {
	p1 := players.NewAiPlayer(1)
	p2 := players.NewAiPlayer(2)
	_, err := Start(p1, p2)
	assert.Nil(t, err, "two different players must start a game")
}

func Test_Game_ShouldNotStartAGameWithSinglePlayer(t *testing.T) {
	p1 := players.NewAiPlayer(1)
	_, err := Start(p1, p1)
	assert.NotNil(t, err, "only two different players can play a game")
	_, err = Start(nil, p1)
	assert.NotNil(t, err, "only two different players can play a game")
	_, err = Start(p1, nil)
	assert.NotNil(t, err, "only two different players can play a game")
}

func Test_Game_ShouldNotStartAGameWithPlayersIllegalID(t *testing.T) {
	invalid := players.NewAiPlayer(board.NoPlayer)
	valid := players.NewAiPlayer(1)
	_, err := Start(invalid, valid)
	assert.NotNil(t, err, "player X must have valid ID")
	_, err = Start(valid, invalid)
	assert.NotNil(t, err, "player Y must have valid ID")
}
