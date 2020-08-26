package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame_Initial(t *testing.T) {
	game := &game{}

	assert.Equal(t, 0, game.Score())
}

func TestGame_FirstRound(t *testing.T) {
	game := &game{}

	game.Roll(5)
	assert.Equal(t, 5, game.Score())
}

func TestGame_SumScores(t *testing.T) {
	game := &game{}

	game.Roll(5)
	game.Roll(4)
	assert.Equal(t, 9, game.Score())
}

func TestGame_SkipUpdatingOutOfUpperRange(t *testing.T) {
	game := &game{}

	game.Roll(11)
	assert.Equal(t, 0, game.Score())
}

func TestGame_SkipUpdatingOutOfLowerRange(t *testing.T) {
	game := &game{}

	game.Roll(-1)
	assert.Equal(t, 0, game.Score())
}

func TestGame_SumInvalidScores(t *testing.T) {
	game := &game{}

	game.Roll(5)
	game.Roll(4)
	assert.Equal(t, 9, game.Score())
}

func TestGame_StrikeAsZeroScore(t *testing.T) {
	game := &game{}

	game.Roll(10)
	assert.Equal(t, 0, game.Score())
}

func TestGame_StrikeAddUpBonusFromNextRound(t *testing.T) {
	game := &game{}

	game.Roll(10)
	game.Roll(3)
	game.Roll(4)
	assert.Equal(t, (10+3+4)+(3+4), game.Score())
}

// TODO: validate the amount 0-10
// TODO: keep track of the score
