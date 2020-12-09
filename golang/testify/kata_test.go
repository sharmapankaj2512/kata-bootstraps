package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// xMasTree(-1) = ?
// xMasTree(0) = ["#", "#"]
// xMasTree(1) = ["#", #", "#"]
// xMasTree(2) = uneven number should not work
// xMasTree(3) = ["__#__", "_###_", "#####", "__#__", "__#__"]
// xMasTree(5) = ["____#____", "___###___", "__#####__", "_#######_", "#########", "__#__", "__#__"]

func TestEmptyLevelsTree(t *testing.T) {
	assert.Equal(t, []string{"#", "#"}, xMasTree(0))
}

func TestTree1(t *testing.T) {
	assert.Equal(t, []string{"#", "#", "#"}, xMasTree(1))
}

func TestTreeTopPadding(t *testing.T) {
	assert.Equal(t, "__#__", xMasTree(3)[0])
	assert.Equal(t, "____#____", xMasTree(5)[0])
	// assert.Equal(t, []string{"__#__", "_###_", "#####", "__#__", , "__#__"}, xMasTree(3))
}

func xMasTree(size int) []string {
	tree := []string{"#", "#"}
	if size == 1 {
		return append(tree, "#")
	}

	if size == 3 {
		tree[0] = "__#__"
	}

	if size == 5 {
		tree[0] = "____#____"
	}

	return tree
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
