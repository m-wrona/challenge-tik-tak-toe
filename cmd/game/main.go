package main

import (
	"context"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github/m-wrona/challenge-tik-tak-toe/internal"
	"github/m-wrona/challenge-tik-tak-toe/internal/board"
	"github/m-wrona/challenge-tik-tak-toe/internal/players"
	"os"
	"time"
)

func main() {
	h := players.NewHumanPlayer(1, os.Stdin)
	ai := players.NewAiPlayer(2)
	s, err := internal.Start(h, ai)
	if err != nil {
		panic(err)
	}
	players := []board.Player{h, ai}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"#", "0", "1", "2"})
	for {
		for _, p := range players {
			printGameState(t, s, p)
			if _, finished := s.IsFinished(); finished {
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			s, err = s.Move(ctx, p)
			if err != nil {
				panic(err)
			}
			cancel()
		}
	}
}

func printGameState(t table.Writer, s *board.GameState, p board.Player) {
	b := s.GetBoard()
	t.ResetRows()
	t.ResetFooters()
	t.AppendRows([]table.Row{
		{"0", b[0], b[1], b[2]},
		{"1", b[3], b[4], b[5]},
		{"2", b[6], b[7], b[8]},
	})
	if winnerID, finished := s.IsFinished(); finished {
		msg := fmt.Sprintf("Game finished - player %d has won!\n", winnerID)
		if winnerID == board.NoPlayer {
			msg = "Game finished - it's a draw!"
		}
		t.AppendFooter(table.Row{msg})
	} else {
		t.AppendFooter(table.Row{"Note: type 0-8 to make a move."})
		t.AppendFooter(table.Row{fmt.Sprintf("Waiting for player's %d move...", p.ID())})
	}
	t.Render()
}
