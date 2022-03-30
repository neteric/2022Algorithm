package mybinarytree

import (
	"testing"
)

func TestMakeTree(t *testing.T) {
	ele := []int{1, 4, 3, 2, 4, 2, 5, null, null, null, null, null, null, 4, 6}
	root := MakeTree(ele)
	inOderTraverse(root)

}
