package challenge_tik_tak_toe

import (
	"github.com/pkg/errors"
)

var notImplemented = errors.New("not implemented")

func Start(x Player, y Player) (*GameState, error) {
	if x == nil || y == nil || x.ID() == y.ID() {
		return nil, errors.New("two players needed to start a game")
	} else if x.ID() == NoPlayer {
		return nil, errors.Errorf("illegal ID of player X: %d", NoPlayer)
	} else if y.ID() == NoPlayer {
		return nil, errors.Errorf("illegal ID of player Y: %d", NoPlayer)
	}
	newGameState := GameState{
		board:   newBoard(),
		players: []Player{x, y},
	}
	return &newGameState, nil
}
