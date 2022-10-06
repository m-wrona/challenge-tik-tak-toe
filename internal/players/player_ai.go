package players

import (
	"context"
	"github/m-wrona/challenge-tik-tak-toe/internal/board"
)

const (
	aiScoreLost    = -10
	aiScoreDraw    = 0
	aiScoreWin     = 10
	aiScoreDisturb = 20
)

type AiPlayer struct {
	id board.PlayerID
}

func NewAiPlayer(id board.PlayerID) board.Player {
	return &AiPlayer{id: id}
}

func (a AiPlayer) ID() board.PlayerID {
	return a.id
}

func (a AiPlayer) NextMove(ctx context.Context, b board.Board) (int, error) {
	bestScore := aiScoreLost
	aiNextMove := board.NoMove
	for _, coordinates := range board.GetWinningCoordinates() {
		score, nextMove := a.evaluateNextMove(b, coordinates)
		if score > bestScore {
			bestScore = score
			aiNextMove = nextMove
		}
	}
	return aiNextMove, nil
}

func (a AiPlayer) evaluateNextMove(b board.Board, coordinates []int) (score int, nextCoordinate int) {
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
	if score == aiScoreLost && nextCoordinate != board.NoMove {
		if otherPlayer == 2 {
			//try to prevent other player to win
			score = aiScoreDisturb
		}
	} else if free > 0 {
		score = aiScoreWin / free
	}
	return
}
