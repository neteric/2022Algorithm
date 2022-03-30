package mybinarytree

import (
	"fmt"
	"testing"
)

func TestMaxSumBST(*testing.T) {
	// head := new(TreeNode)
	// head.Init()
	var x = []int{1, 2, 3, null}
	root := MakeTree(x)
	fmt.Println("result:", maxSumBST(root))
}

func TestBST(t *testing.T) {
	var x = []int{1, 2, 3, null}
	root := MakeTree(x)
	inOderTraverse(root)
	fmt.Println(isValidBST(root))
	// result := searchBST(head, 8)
	// if result == nil {
	// 	fmt.Println("不存在")
	// } else {
	// 	fmt.Println(result.Val)
	// }

	deleteNode(root, 7)
	inOderTraverse(root)
	fmt.Println(isValidBST(root))

}
