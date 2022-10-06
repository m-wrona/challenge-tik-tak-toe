package board

import (
	"context"
)

type PlayerID int

const (
	NoPlayer = PlayerID(0)
)

type Player interface {
	ID() PlayerID
	NextMove(ctx context.Context, b Board) (int, error)
}
