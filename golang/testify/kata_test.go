package kata

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

//             XMasTree(5) will result in:
//             ____#____              1
//             ___###___              2
//             __#####__              3
//             _#######_              4
//             #########       -----> 5 - Height of Tree
//             ____#____        1
//             ____#____        2 - Trunk/Stem of Tree
//
//             XMasTree(3) will result in:
//             __#__                  1
//             _###_                  2
//             #####          ----->  3 - Height of Tree
//             __#__           1
//             __#__           2 - Trunk/Stem of Tree
func TestTrunkOnly(t *testing.T) {
	tree, err := XMasTree_(0)
	require.NoError(t, err)
	assert.Equal(t, []string{"#", "#"}, tree)
}

func TestTree_1(t *testing.T) {
	assert.Equal(t, []string{"#", "#", "#"}, XMasTree(1))
}

func TestXMasTree_2(t *testing.T) {
	assert.Equal(t, []string{
		"_#_",
		"###",
		"_#_",
		"_#_"}, XMasTree(2))
}

func TestTree_5(t *testing.T) {
	assert.Equal(t, []string{
		"____#____",
		"___###___",
		"__#####__",
		"_#######_",
		"#########",
		"____#____",
		"____#____"}, XMasTree(5))
}
func TestTree_7(t *testing.T) {
	assert.Equal(t, []string{
		"______#______",
		"_____###_____",
		"____#####____",
		"___#######___",
		"__#########__",
		"_###########_",
		"#############",
		"______#______",
		"______#______"}, XMasTree(7))
}

func TestTree_3(t *testing.T) {
	assert.Equal(t, []string{
		"__#__",
		"_###_",
		"#####",
		"__#__",
		"__#__"}, XMasTree(3))
}

func TestXMasTree_2_Top_Level(t *testing.T) {
	assert.Equal(t, "_#_", XMasTree(2)[0])
}
func TestXMasTree_2_Trunk_Level(t *testing.T) {
	assert.Equal(t, "_#_", XMasTree(2)[2])
	assert.Equal(t, "_#_", XMasTree(2)[3])
}

func TestXMasTree_3_TopLevel(t *testing.T) {
	assert.Equal(t, "__#__", XMasTree(3)[0])
}

func TestXMasTree_2_2nd_Level(t *testing.T) {
	assert.Equal(t, "###", XMasTree(2)[1])
}

func TestNegativeHeightsDontPanic(t *testing.T) {
	assert.NotPanics(t, func() { XMasTree(-1) })
	assert.NotPanics(t, func() { treeBody(-1) })
	assert.NotPanics(t, func() { sidePadding(-1, -1) })
}

func TestError_Message(t *testing.T) {
	_, err := XMasTree_(-1)
	require.Error(t, err)
}

func XMasTree_(height int) ([]string, error) {
	if height < 0 {
		return nil, errors.New("negative height")
	}
	return XMasTree(height), nil
}

func XMasTree(height int) []string {
	var tree []string

	for level := 0; level < height; level++ {
		padding := sidePadding(height, level)
		tree = append(tree, padding+treeBody(level)+padding)
	}

	return addTreeTrunk(tree, height)
}

func treeBody(level int) string {
	if level < 0 {
		return ""
	}
	return strings.Repeat("#", 2*level+1)
}

func addTreeTrunk(tree []string, height int) []string {
	padding := sidePadding(height, 0)
	trunk := padding + "#" + padding
	return append(tree, trunk, trunk)
}

func sidePadding(height int, level int) string {
	count := height - level - 1
	if count < 1 {
		return ""
	}
	return strings.Repeat("_", count)
}
