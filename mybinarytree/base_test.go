package mybinarytree

import (
	"testing"
)

func TestMakeTree(t *testing.T) {
	ele := []int{1, 4, 3, 2, 4, 2, 5, null, null, null, null, null, null, 4, 6}
	root := MakeTree(ele)
	inOderTraverse(root)

}
func TestMakeTreeByLevel(*testing.T) {
	ele := []int{2, null, 3, null, 4, null, 5, null, 6}
	root := MakeTreeByLevelList(ele)
	inOderTraverse(root)
}
