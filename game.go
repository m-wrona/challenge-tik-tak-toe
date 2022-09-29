package challenge_tik_tak_toe

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

var notImplemented = errors.New("not implemented")

type Board = [][]PlayerID

type Game struct {
	players   []Player
	isStarted bool
}

type GameState struct {
	board     Board
	game      *Game
	iteration int
}

func newBoard() Board {
	b := make([][]PlayerID, 3)
	for idx := range b {
		b[idx] = make([]PlayerID, 3)
	}
	return b
}

func (g *Game) Join(p Player) error {
	for _, gp := range g.players {
		if gp.ID() == p.ID() {
			return errors.Errorf("player %d has already joined the game", p.ID())
		}
	}
	g.players = append(g.players, p)
	return nil
}

func (g *Game) Start() (*GameState, error) {
	if g.isStarted {
		return nil, errors.New("game has already started")
	}
	g.isStarted = true
	newGameState := GameState{
		board: newBoard(),
		game:  g,
	}
	return &newGameState, nil
}

func (s *GameState) copy() *GameState {
	newState := GameState{
		board: newBoard(),
		game:  s.game,
	}
	copy(newState.board, s.board)
	newState.iteration = s.iteration + 1
	return &newState
}

func (s *GameState) Move(ctx context.Context, p Player) (*GameState, error) {
	x, y, err := p.NextMove(ctx, s.board)
	if err != nil {
		return nil, errors.Wrapf(err, "player %d couldn't make next move", p.ID())
	}
	if s.board[x][y] != NoPlayer {
		return nil, errors.Errorf(
			"player %d cannot mark fields x:%d & y:%d since they are already used by player %d",
			p.ID(), x, y, s.board[x][y],
		)
	} else {
		newState := s.copy()
		id := p.ID()
		newState.board[x][y] = id
		return newState, nil
	}
}

func (s *GameState) HasPlayerWon() PlayerID {
	//check rows
	println("rows...")
	for r := 0; r < 3; r++ {
		rP := s.board[r][0]
		if rP != NoPlayer {
			m := 1
			for c := 1; c < 3; c++ {
				println(fmt.Sprintf("%d: [%d][%d]==%d", s.iteration, r, c, rP))
				if rP == s.board[r][c] {
					println(fmt.Sprintf("%d: yes", s.iteration))
					m++
				}
			}
			if m == 3 {
				println("row won in interation: ", s.iteration)
				return rP
			}
		}
	}
	//check columns
	println("columns...")
	for c := 0; c < 3; c++ {
		cP := s.board[0][c]
		if cP != NoPlayer {
			m := 1
			for r := 1; r < 3; r++ {
				println(fmt.Sprintf("%d: [%d][%d]==%d", s.iteration, r, c, cP))
				if cP == s.board[r][c] {
					println(fmt.Sprintf("%d: yes", s.iteration))
					m++
				}
			}
			if m == 3 {
				println("column won in interation: ", s.iteration)
				return cP
			}
		}
	}
	println("diagonals...")
	//diagonal 1
	if s.board[0][0] != NoPlayer && s.board[0][0] == s.board[1][1] && s.board[0][0] == s.board[2][2] {
		return s.board[0][0]
	}
	//diagonal 2
	if s.board[0][2] != NoPlayer && s.board[0][2] == s.board[1][1] && s.board[0][2] == s.board[2][0] {
		return s.board[0][2]
	}
	return NoPlayer
}
