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

func TestGame_Spare(t *testing.T) {
	game := &game{}

	game.Roll(6) //6
	game.Roll(4) //6 + 4 = 10 (spare)
	game.Roll(3) //6 + 4 + 3 = 13
	assert.Equal(t, (6 + 4 + 3), game.Score())
}

func TestGame_SpareWithNextFrame(t *testing.T) {
	game := &game{}

	game.Roll(6) //
	game.Roll(4) // spare , now we store 10 in frame
	game.Roll(1) // add 1 to the same frame
	game.Roll(5) // now we have a new frame
	assert.Equal(t, (6+4+1)+(1+5), game.Score())
}

// TODO: validate the amount 0-10
// TODO: keep track of the score
