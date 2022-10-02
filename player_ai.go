package challenge_tik_tak_toe

import (
	"context"
)

const (
	aiScoreLost    = -10
	aiScoreDraw    = 0
	aiScoreWin     = 10
	aiScoreDisturb = 20
)

type AiPlayer struct {
	id PlayerID
}

func (a AiPlayer) ID() PlayerID {
	return a.id
}

func (a AiPlayer) NextMove(ctx context.Context, b Board) (int, error) {
	bestScore := aiScoreLost
	aiNextMove := noMove
	for _, coordinates := range boardWinningCoordinates {
		validateCoordinates(coordinates)
		score, nextMove := a.evaluateNextMove(b, coordinates)
		if score > bestScore {
			bestScore = score
			aiNextMove = nextMove
		}
	}
	return aiNextMove, nil
}

func (a AiPlayer) evaluateNextMove(b Board, coordinates []int) (score int, nextCoordinate int) {
	free := 0
	otherPlayer := 0
	score = aiScoreDraw
	nextCoordinate = noMove
	for _, c := range coordinates {
		if b[c] == NoPlayer {
			if nextCoordinate == noMove {
				nextCoordinate = c
			}
			free++
		} else if b[c] != a.ID() {
			score = aiScoreLost
			otherPlayer++
		}
	}
	if score == aiScoreLost && nextCoordinate != noMove {
		if otherPlayer == 2 {
			//try to prevent other player to win
			score = aiScoreDisturb
		}
	} else if free > 0 {
		score = aiScoreWin / free
	}
	return
}
