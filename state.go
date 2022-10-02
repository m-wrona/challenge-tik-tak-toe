package challenge_tik_tak_toe

import (
	"context"
	"github.com/pkg/errors"
)

type GameState struct {
	board   Board
	players []Player
}

func (s *GameState) copy() *GameState {
	return &GameState{
		board:   copyBoard(s.board),
		players: s.players,
	}
}

func (s *GameState) Move(ctx context.Context, p Player) (*GameState, error) {
	winnerID, finished := s.IsFinished()
	if finished {
		if winnerID != NoPlayer {
			return nil, errors.Errorf("player %d has already won the game", winnerID)
		} else {
			return nil, errors.Errorf("game has finished with a draw")
		}
	}

	copy := s.copy() //to prevent cheating
	x, err := p.NextMove(ctx, copy.board)
	if err != nil {
		return nil, errors.Wrapf(err, "player %d couldn't make next move", p.ID())
	}
	if s.board[x] != NoPlayer {
		return nil, errors.Errorf(
			"player %d cannot mark field %d since it's already taken by player %d",
			p.ID(), x, s.board[x],
		)
	} else {
		newState := s.copy()
		id := p.ID()
		newState.board[x] = id
		return newState, nil
	}
}

func (s *GameState) IsFinished() (PlayerID, bool) {
	isFinished := true
	for _, coordinates := range boardWinningCoordinates {
		validateCoordinates(coordinates)
		x := s.board[coordinates[0]]
		y := s.board[coordinates[1]]
		z := s.board[coordinates[2]]
		if x == NoPlayer || y == NoPlayer || z == NoPlayer {
			isFinished = false
		} else if x == y && x == z && y == z {
			return x, true
		}
	}
	return NoPlayer, isFinished
}
