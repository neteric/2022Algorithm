package avltree

import (
	"fmt"
	"testing"
)

func TestIsBalanced(*testing.T) {
	var x = []int{1, 2, 2, 3, null, null, 3, 4, null, null, 4}
	root := MakeTree(x)
	fmt.Println(isBalanced(root))

}

func TestMaxDepth(*testing.T) {
	var x = []int{3, 9, 20, null, null, 15, 7}
	root := MakeTree(x)
	fmt.Println(maxDepth(root))
}
