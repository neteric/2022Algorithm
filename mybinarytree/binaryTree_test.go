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

func TestIsSymmetric(*testing.T) {
	var ele = []int{1, 2, 2, null, 3, null, 3}
	root := MakeTree(ele)
	fmt.Println(isSymmetric(root))

}

func TestMaxDepth2(*testing.T) {
	var ele = []int{3, 9, 20, null, null, 15, 7}
	root := MakeTree(ele)
	fmt.Println(maxDepth2(root))
}

func TestHasPathSum(*testing.T) {
	var ele = []int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1}
	root := MakeTreeByLevelList(ele)
	fmt.Println(hasPathSum(root, 22))
}

func TestPathSum(*testing.T) {
	var ele = []int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1}
	root := MakeTreeByLevelList(ele)
	fmt.Println(pathSum(root, 22))
}

func TestSumOfLeftLeaves(*testing.T) {
	var ele = []int{3, 9, 20, null, null, 15, 7}
	root := MakeTreeByLevelList(ele)
	fmt.Println(sumOfLeftLeaves(root))
}

func TestBinaryTreePaths(*testing.T) {
	var ele = []int{1, 2, 3, null, 5}
	root := MakeTreeByLevelList(ele)
	fmt.Println(binaryTreePaths(root))
}

func TestLevelOrder(*testing.T) {
	var ele = []int{3, 9, 20, null, null, 15, 7}
	root := MakeTreeByLevelList(ele)
	fmt.Println(levelOrder(root))
}

func TestRightSideView(*testing.T) {
	var ele = []int{}
	root := MakeTreeByLevelList(ele)
	fmt.Println(rightSideView(root))
}

func TestSumTree(*testing.T) {
	var ele = []int{21, 7, 14, 1, 1, 2, 2, 3, 3}
	root := MakeTreeByLevelList(ele)
	fmt.Println(findTilt(root))
}

func TestTree2str(*testing.T) {
	var ele = []int{1, 2, 3, null, 4}
	root := MakeTreeByLevelList(ele)
	fmt.Println("s:", tree2str(root))
}

func TestMergeTrees(*testing.T) {
	var ele1 = []int{1}
	root1 := MakeTreeByLevelList(ele1)

	var ele2 = []int{1, 2}
	root2 := MakeTreeByLevelList(ele2)

	levelTraverse(mergeTrees(root1, root2))
}

func TestIncreasingBST(*testing.T) {
	var ele = []int{5, 3, 6, 2, 4, null, 8, 1, null, null, null, 7, 9}
	root := MakeTreeByLevelList(ele)
	levelTraverse(increasingBST(root))
}

func TestZigzagLevelOrder(*testing.T) {
	var ele = []int{3, 9, 20, null, null, 15, 7}
	root := MakeTreeByLevelList(ele)
	fmt.Println(zigzagLevelOrder(root))
}

func TestSortedArrayToBST(*testing.T) {
	var ele = []int{-10, -3, 0, 5, 9}
	levelTraverse(sortedArrayToBST(ele))
}

func TestIsBalanced(*testing.T) {
	var ele = []int{1, 2, 2, 3, 3, null, null, 4, 4}
	root := MakeTreeByLevelList(ele)
	fmt.Println(isBalanced(root))
}

func TestDiameterOfBinaryTree(*testing.T) {
	var ele = []int{1, 2, 3, 4, 5}
	root := MakeTreeByLevelList(ele)
	fmt.Println(diameterOfBinaryTree(root))
}

func TestMirrorTree(*testing.T) {
	var ele = []int{1, 2, 3, 4, null, null, 7}
	root := MakeTreeByLevelList(ele)
	levelTraverse(mirrorTree(root))

}

func TestBuildTree1(*testing.T) {
	tree := buildTree1([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7})
	levelTraverse(tree)
}
func TestMinDepth(*testing.T) {
	var ele = []int{2, null, 3, null, 4, null, 5, null, 6}
	root := MakeTreeByLevelList(ele)
	fmt.Println(minDepth(root))
}

func TestMaxPathSum(*testing.T) {
	var ele = []int{9, 6, -3, null, null, -6, 2, null, null, 2, null, -6, -6, -6}
	root := MakeTreeByLevelList(ele)
	fmt.Println(maxPathSum(root))

}

func TestLowestCommonAncestor(*testing.T) {
	var ele = []int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}
	var pe = []int{2}
	var qe = []int{4}
	root := MakeTreeByLevelList(ele)
	p := MakeTreeByLevelList(pe)
	q := MakeTreeByLevelList(qe)
	x := lowestCommonAncestor(root, p, q)
	if x == nil {
		fmt.Println("empty")
	} else {
		fmt.Println(x.Val)
	}
}
