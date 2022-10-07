package players_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github/m-wrona/challenge-tik-tak-toe/internal/board"
	"github/m-wrona/challenge-tik-tak-toe/internal/players"
	"testing"
	"time"
)

func Test_AI_ShouldReturnNoMove(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	for idx := range b {
		b[idx] = board.PlayerID(2)
	}
	m1, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, board.NoMove, m1, "no move is possible")
}

func Test_AI_ShouldMarkFirstRow(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()

	m1, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 0, m1, "unexpected 1s move")
	b[m1] = ai.ID()

	m2, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 1, m2, "unexpected 2nd move")
	b[m2] = ai.ID()

	m3, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 2, m3, "unexpected 3rd move")
}

func Test_AI_ShouldMarkSecondRow(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	b[0] = board.PlayerID(2)

	m1, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 3, m1, "unexpected 1s move")
	b[m1] = ai.ID()

	m2, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 4, m2, "unexpected 2nd move")
	b[m2] = ai.ID()

	m3, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 5, m3, "unexpected 3rd move")
}

func Test_AI_ShouldMarkThirdRow(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	b[0] = board.PlayerID(2)
	b[3] = board.PlayerID(2)

	m1, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 6, m1, "unexpected 1s move")
	b[m1] = ai.ID()

	m2, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 7, m2, "unexpected 2nd move")
	b[m2] = ai.ID()

	m3, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 8, m3, "unexpected 3rd move")
}

func Test_AI_ShouldMarkFirstColumn(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	b[0] = ai.ID()
	b[1] = board.PlayerID(2)
	b[5] = board.PlayerID(2)

	m2, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 3, m2, "unexpected 2nd move")
	b[m2] = ai.ID()

	m3, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 6, m3, "unexpected 3rd move")
}

func Test_AI_ShouldMarkSecondColumn(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	b[1] = ai.ID()
	b[0] = board.PlayerID(2)
	b[5] = board.PlayerID(2)

	m2, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 4, m2, "unexpected 2nd move")
	b[m2] = ai.ID()

	m3, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 7, m3, "unexpected 3rd move")
}

func Test_AI_ShouldMarkThirdColumn(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	b[0] = board.PlayerID(2)
	b[1] = board.PlayerID(2)
	b[2] = ai.ID()

	b[3] = board.PlayerID(2)
	b[4] = board.PlayerID(2)

	b[6] = board.PlayerID(2)
	b[7] = board.PlayerID(2)

	m2, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 5, m2, "unexpected 2nd move")
	b[m2] = ai.ID()

	m3, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 8, m3, "unexpected 3rd move")
}

func Test_AI_ShouldMarkFirstDiagonal(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	b[0] = ai.ID()
	b[1] = board.PlayerID(2)
	b[2] = board.PlayerID(2)

	b[3] = board.PlayerID(2)
	b[5] = board.PlayerID(2)

	b[6] = board.PlayerID(2)
	b[7] = board.PlayerID(2)

	m2, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 4, m2, "unexpected 2nd move")
	b[m2] = ai.ID()

	m3, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 8, m3, "unexpected 3rd move")
}

func Test_AI_ShouldMarkSecondDiagonal(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	b[0] = board.PlayerID(2)
	b[1] = board.PlayerID(2)
	b[3] = board.PlayerID(2)
	b[5] = board.PlayerID(2)
	b[7] = board.PlayerID(2)
	b[8] = board.PlayerID(2)

	m1, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 2, m1, "unexpected 1s move")
	b[m1] = ai.ID()

	m2, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 4, m2, "unexpected 2nd move")
	b[m2] = ai.ID()

	m3, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 6, m3, "unexpected 3rd move")
}

func Test_AI_ShouldDisturbRowMarking(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	b[0] = board.PlayerID(2)
	b[1] = board.PlayerID(2)

	m1, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 2, m1, "unexpected disturb move")
}

func Test_AI_ShouldDisturbColumnMarking(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	b[0] = board.PlayerID(2)
	b[3] = board.PlayerID(2)

	m1, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 6, m1, "unexpected disturb move")
}

func Test_AI_ShouldDisturbDiagonalMarking(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	b[0] = board.PlayerID(2)
	b[4] = board.PlayerID(2)

	m1, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 8, m1, "unexpected disturb move")
}

func Test_AI_ShouldMakeWinningMoveOverDistrubOne(t *testing.T) {
	ai := players.NewAiPlayer(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	b := board.New()
	b[0] = ai.ID()
	b[1] = ai.ID()

	b[3] = board.PlayerID(2)
	b[4] = board.PlayerID(2)

	m1, err := ai.NextMove(ctx, b)
	assert.Nil(t, err, "couldn't make next move")
	assert.Equal(t, 2, m1, "should make winning move and finish the game")
	assert.NotEqual(t, 5, m1, "should not make disturb move since player could win a game")
}
