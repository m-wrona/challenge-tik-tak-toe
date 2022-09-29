package challenge_tik_tak_toe

import (
	"context"
	"github.com/pkg/errors"
)

type PlayerID int

const (
	NoPlayer = PlayerID(0)
)

type Player interface {
	ID() PlayerID
	NextMove(ctx context.Context, b Board) (int, int, error)
}

type HumanPlayer struct {
	id PlayerID
}

type AiPlayer struct {
	id PlayerID
}

func (h HumanPlayer) ID() PlayerID {
	return h.id
}

func (h HumanPlayer) NextMove(ctx context.Context, b Board) (int, int, error) {
	//rows
	for r := range b {
		for c := range b[r] {
			if b[r][c] == NoPlayer {
				return r, c, nil
			} else if b[r][c] != h.ID() {
				break
			}
		}
	}
	//columns
	for c := 0; c < 3; c++ {
		for r := 0; r < 3; r++ {
			if b[r][c] == NoPlayer {
				return r, c, nil
			} else if b[r][c] != h.ID() {
				break
			}
		}
	}
	return 0, 0, errors.Errorf("player %d couldn't choose next move for board: %+v", h.id, b)
}

func (a AiPlayer) ID() PlayerID {
	return a.id
}

func (a AiPlayer) NextMove(ctx context.Context, b Board) (int, int, error) {
	return 0, 0, notImplemented
}
