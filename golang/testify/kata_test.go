package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 roll(int) ->  The argument is the number of pins knocked down.
 score()
*/

// Arrange
// Act
// Assert

func TestTwoRolls_Success(t *testing.T) {
	g := Game{}
	g.Roll(4)
	g.Roll(10)
	assert.Equal(t, 14, g.Score())
}

func Test_SpareCountsNextRollTwice(t *testing.T) {
	g := Game{}
	g.Roll(4)
	g.Roll(6)
	assert.Equal(t, 10, g.Score())
	g.Roll(3)
	assert.Equal(t, 16, g.Score())
}

func TestSomething(t *testing.T) {
	assert.Equal(t, 42, doSomething("b"), "Answer to the Ultimate Question of Life, the Universe, and Everything")
}

func TestGen(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "with a expected 0",
			input:    "a",
			expected: 0,
		},
		{
			name:     "with b expected 42",
			input:    "b",
			expected: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doSomething(tt.input); got != tt.expected {
				t.Errorf("doSomething() = %v, but expected %v", got, tt.expected)
			}
		})
	}
}

func doSomething(input string) interface{} {
	if input == "b" {
		return 42
	}
	return 0
}
