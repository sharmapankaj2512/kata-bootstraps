package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// xMasTree(-1) = ?
// xMasTree(0) = ["#", "#"]
// xMasTree(1) = ["#", #", "#"]
// xMasTree(2) = uneven number should not work
// xMasTree(3) = ["_#_", "###", "_#_", "_#_"]

func TestEmptyLevelsTree(t *testing.T) {
	assert.Equal(t, []string{"#", "#"}, xMasTree(0))
}

func TestTree1(t *testing.T) {
	assert.Equal(t, []string{"#", "#", "#"}, xMasTree(1))
}

//func TestGen(t *testing.T) {
//	tests := []struct {
//		name     string
//		input    string
//		expected int
//	}{
//		{
//			name:     "with a expected 0",
//			input:    "a",
//			expected: 0,
//		},
//		{
//			name:     "with b expected 42",
//			input:    "b",
//			expected: 42,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := doSomething(tt.input); got != tt.expected {
//				t.Errorf("doSomething() = %v, but expected %v", got, tt.expected)
//			}
//		})
//	}
//}

func xMasTree(size int) []string {
	tree := []string{"#", "#"}
	if size == 1 {
		return append(tree, "#")
	}
	return tree
}
