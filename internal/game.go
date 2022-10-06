package internal

import (
	"github.com/pkg/errors"
	"github/m-wrona/challenge-tik-tak-toe/internal/board"
)

func Start(x board.Player, y board.Player) (*board.GameState, error) {
	if x == nil || y == nil || x.ID() == y.ID() {
		return nil, errors.New("two players needed to start a game")
	} else if x.ID() == board.NoPlayer {
		return nil, errors.Errorf("illegal ID of player X: %d", board.NoPlayer)
	} else if y.ID() == board.NoPlayer {
		return nil, errors.Errorf("illegal ID of player Y: %d", board.NoPlayer)
	}
	return board.NewState(x, y), nil
}
