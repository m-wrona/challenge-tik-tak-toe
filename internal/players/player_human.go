package players

import (
	"bufio"
	"context"
	"github.com/pkg/errors"
	"github/m-wrona/challenge-tik-tak-toe/internal/board"
	"io"
	"strconv"
	"strings"
)

var inputRemoveChars = []byte{'-', ' ', ':'}

type HumanPlayer struct {
	id      board.PlayerID
	scanner *bufio.Scanner
}

type playerMove struct {
	move int
	err  error
}

func NewHumanPlayer(id board.PlayerID, in io.Reader) board.Player {
	return &HumanPlayer{
		id:      id,
		scanner: bufio.NewScanner(in),
	}
}

func (h *HumanPlayer) ID() board.PlayerID {
	return h.id
}

func (h *HumanPlayer) NextMove(ctx context.Context, b board.Board) (int, error) {
	input := readInput(h.scanner)
	select {
	case <-ctx.Done():
		return findFirstFreeWinningCoordinate(b)

	case line := <-input:
		if line.err != nil {
			return board.NoMove, errors.Wrapf(line.err, "couldn't read move of player %d", h.id)
		}
		return line.move, nil
	}
}

func readInput(scanner *bufio.Scanner) chan playerMove {
	ch := make(chan playerMove)
	go func() {
		defer close(ch)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			ch <- playerMove{err: err}
		}
		line := scanner.Text()
		for _, ch := range inputRemoveChars {
			line = strings.ReplaceAll(line, string(ch), "")
		}
		move, err := strconv.Atoi(line)
		if err != nil {
			ch <- playerMove{err: errors.Wrapf(err, "input is invalid number: %s", line)}
		}
		ch <- playerMove{move: move}
	}()
	return ch
}

func findFirstFreeWinningCoordinate(b board.Board) (int, error) {
	for _, coordinates := range board.GetWinningCoordinates() {
		for _, c := range coordinates {
			if b[c] == board.NoPlayer {
				return c, nil
			}
		}
	}
	return board.NoMove, errors.New("couldn't make any random move")
}
