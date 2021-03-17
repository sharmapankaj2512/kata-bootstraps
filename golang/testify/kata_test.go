package kata

import (
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, []string{"#", "#"}, XMasTree(0))
}

func TestTree_1(t *testing.T) {
	assert.Equal(t, []string{"#", "#", "#"}, XMasTree(1))
}

func TestTree_2(t *testing.T) {
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

func TestTree_2_Top_Level(t *testing.T) {
	assert.Equal(t, "_#_", XMasTree(2)[0])
}
func TestTree_2_Trunk_Level(t *testing.T) {
	assert.Equal(t, "_#_", XMasTree(2)[2])
	assert.Equal(t, "_#_", XMasTree(2)[3])
}

func TestTree_3_Top_Level(t *testing.T) {
	assert.Equal(t, "__#__", XMasTree(3)[0])
}

func TestTree_2_2nd_Level(t *testing.T) {
	assert.Equal(t, "###", XMasTree(2)[1])
}
func XMasTree(height int) []string {
	tree := []string{}
	for level := 0; level < height; level++ {
		tree = append(tree, sidePadding(height, level)+treeBody(level)+sidePadding(height, level))
	}

	return addTreeTrunk(tree, height)
}

func treeBody(level int) string {
	tree := "#"
	for i := 0; i < level; i++ {
		tree += "##"
	}
	return tree
}

func addTreeTrunk(tree []string, height int) []string {
	tree = append(tree, sidePadding(height, 0)+"#"+sidePadding(height, 0))
	tree = append(tree, sidePadding(height, 0)+"#"+sidePadding(height, 0))
	return tree
}

func sidePadding(height int, level int) string {
	padding := ""
	for i := 0; i < height-level-1; i++ {
		padding += "_"
	}
	return padding
}
