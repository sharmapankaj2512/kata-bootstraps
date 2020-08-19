package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame(t *testing.T) {
  game := &game{}

  game.Roll(4)
  assert.Equal(t, 4, game.Score())


}
// TODO: validate the amount 0-10
// TODO: keep track of the score