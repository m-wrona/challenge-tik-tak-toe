package players

import (
	"context"
	"github/m-wrona/challenge-tik-tak-toe/internal/board"
	"hash/maphash"
	"math/rand"
)

const (
	aiScoreLost    = -10
	aiScoreDraw    = 0
	aiScoreWin     = 10
	aiScoreDisturb = 10
)

type config struct {
	randomMoves bool //random moves in best found strategies
}

type AiOption func(cfg *config)

type AiPlayer struct {
	id   board.PlayerID
	cfg  *config
	rand *rand.Rand
}

// WithAIRandomMoves allows to make random moves among found best game strategies
func WithAIRandomMoves(v bool) AiOption {
	return func(cfg *config) {
		cfg.randomMoves = v
	}
}

func NewAiPlayer(id board.PlayerID, opts ...AiOption) board.Player {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &AiPlayer{
		id:   id,
		cfg:  cfg,
		rand: rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64()))),
	}
}

func (a *AiPlayer) ID() board.PlayerID {
	return a.id
}

func (a *AiPlayer) NextMove(ctx context.Context, b board.Board) (int, error) {
	bestScore := aiScoreLost
	bestNextMoves := make([]int, 0)
	for _, coordinates := range board.GetWinningCoordinates() {
		score, nextMove := a.evaluateNextMove(b, coordinates)
		if score > bestScore {
			bestScore = score
			bestNextMoves = []int{nextMove}
		} else if score == bestScore {
			bestNextMoves = append(bestNextMoves, nextMove)
		}
	}
	if len(bestNextMoves) == 0 {
		return board.NoMove, nil
	} else if !a.cfg.randomMoves {
		return bestNextMoves[0], nil
	}
	randomMoveIdx := a.rand.Intn(len(bestNextMoves))
	return bestNextMoves[randomMoveIdx], nil
}

func (a *AiPlayer) evaluateNextMove(b board.Board, coordinates []int) (score int, nextCoordinate int) {
	free := 0
	otherPlayer := 0
	score = aiScoreDraw
	nextCoordinate = board.NoMove
	for _, c := range coordinates {
		if b[c] == board.NoPlayer {
			if nextCoordinate == board.NoMove {
				nextCoordinate = c
			}
			free++
		} else if b[c] != a.ID() {
			score = aiScoreLost
			otherPlayer++
		}
	}
	if otherPlayer == 0 && free == 1 {
		//certain win
		score = aiScoreWin
	} else if score == aiScoreLost && nextCoordinate != board.NoMove {
		//lost or a draw
		if otherPlayer == 2 {
			//try to prevent other player to win
			score = aiScoreDisturb
		}
	} else if free > 0 {
		//still can win
		score = aiScoreWin / free
	}
	return
}
