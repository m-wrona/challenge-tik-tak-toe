package challenge_tik_tak_toe

import (
	"context"
)

type HumanPlayer struct {
	id PlayerID
}

func (h HumanPlayer) ID() PlayerID {
	return h.id
}

func (h HumanPlayer) NextMove(ctx context.Context, b Board) (int, error) {
	return 0, notImplemented
}
