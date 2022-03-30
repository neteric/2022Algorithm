package mybinarytree

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	var x = []int{1, 2, 3, null}
	root := MakeTree(x)
	kthSmallest(root, 3)
}
func TestTreeNode(t *testing.T) {
	// preOrder := []int{3, 9, 20, 15, 7}
	// inOrder := []int{9, 3, 15, 20, 7}
	// root := buildTree(preOrder, inOrder)
	// preOderTraverse(root)
	var x = []int{1, 2, 3, null}
	root := MakeTree(x)
	// traverse(head)
	// fmt.Println(maxDepth(head))
	// invertTree(head)
	// helpInvertTree(head)
	// levelTraverse(head)
	// preOderTraverse(head)
	// flatten(head)
	// levelTraverse(head)

	// LeftLeafNode(head)
	convertBST(root)
	inOderTraverse(root)

}

func TestInorderTraversal(*testing.T) {
	var x = []int{1, 2, 3, null}
	root := MakeTree(x)
	fmt.Println(inorderTraversal(root))

}

func TestIsSameTree(*testing.T) {
	var p = []int{1, 2, 1}
	rootP := MakeTree(p)
	var q = []int{1, 2, 1}
	rootQ := MakeTree(q)
	fmt.Println(isSameTree(rootP, rootQ))
}
