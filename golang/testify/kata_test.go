package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = assert.Equal
var _ = require.Equal

func TestBirth(t *testing.T) {
	g := new(game)

	g.field = [][]byte{
		[]byte{ 0, 0, 1 },
		[]byte{ 0, 0, 1 },
		[]byte{ 0, 0, 1 },
	}

	expect := [][]byte{
		[]byte{ 0, 0, 1 },
		[]byte{ 0, 1, 1 },
		[]byte{ 0, 0, 1 },
	}
	g.iterate()
	assert.Equal(t, expect, g.field)
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
