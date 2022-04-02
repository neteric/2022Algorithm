package mybinarytree

import (
	"fmt"
	"strings"
)

var res = 0
var depth = 0

func traverse(node *TreeNode) {
	if node == nil {
		if res < depth {
			res = depth
		}
		return
	}
	// fmt.Println(node.Val)

	depth++
	traverse(node.Left)
	traverse(node.Right)
	depth--

}

func maxDepth(node *TreeNode) int {
	traverse(node)
	return res
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root

}
func levelTraverse(node *TreeNode) {

	var level []int
	var q []*TreeNode
	q = append(q, node)
	level = append(level, 1)
	for {
		out := q[0]
		if out.Left != nil {
			q = append(q, out.Left)
		}
		if out.Right != nil {
			q = append(q, out.Right)
		}
		fmt.Println(out.Val)
		if len(q) <= 1 {
			break
		}
		q[0] = nil
		q = q[1:]
	}

}
func helpInvertTree(node *TreeNode) {
	fmt.Println("before")
	levelTraverse(node)
	root := invertTree(node)
	fmt.Println("after")

	levelTraverse(root)
}

func preOderTraverse(head *TreeNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	preOderTraverse(head.Left)
	preOderTraverse(head.Right)
}

func inOderTraverse(head *TreeNode) {
	if head == nil {
		return
	}
	inOderTraverse(head.Left)
	fmt.Println(head.Val)
	inOderTraverse(head.Right)
}

func postorderTraversal(root *TreeNode) []int {

	var result []int

	helpPostOderTraverse(root, &result)

	return result

}
func helpPostOderTraverse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	helpPostOderTraverse(root.Left, result)
	helpPostOderTraverse(root.Right, result)
	*result = append(*result, root.Val)
}

func preorderTraversal(root *TreeNode) []int {

	var result []int

	helpPrePostOderTraverse(root, &result)

	return result

}
func helpPrePostOderTraverse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	helpPrePostOderTraverse(root.Left, result)
	helpPrePostOderTraverse(root.Right, result)

}

func LeftLeafNode(node *TreeNode) *TreeNode {
	for {
		if node.Left == nil {
			return node
		}
		node = node.Left
	}
}

func TreeToLink(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	flatten(root.Left)
	flatten(root.Right)

	if root.Left == nil {
		root.Left = root.Right
		root.Right = nil
	} else {
		leftleaf := LeftLeafNode(root)
		leftleaf.Left = root.Right
		root.Right = nil
	}
	return root
}

func flatten(root *TreeNode) {
	linkTree := TreeToLink(root)
	for {
		if linkTree == nil {
			return
		}

		linkTree.Right = linkTree.Left
		linkTree.Left = nil
		linkTree = linkTree.Right

	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	// 递归 base case
	if len(preorder) == 1 || len(inorder) == 1 {
		leaf := new(TreeNode)
		leaf.Val = preorder[0]
		return leaf
	} else if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 找到前序遍历的根节点在中序遍历中的位置
	var midIndex int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			midIndex = i
			break
		}
	}
	// 根据中序遍历把树切分成左右子树两部分
	leftChildLength := len(inorder[0 : midIndex+1])
	preorderLeftChild := preorder[1:leftChildLength]
	preorderRightChild := preorder[leftChildLength:]

	inorderLeftChild := inorder[0:midIndex]
	inorderRightChild := inorder[midIndex+1:]
	// 构造出根节点
	root := new(TreeNode)
	root.Val = inorder[midIndex]

	// 根据前序和中序左右子树序列递归
	leftChild := buildTree(preorderLeftChild, inorderLeftChild)
	rightChild := buildTree(preorderRightChild, inorderRightChild)
	root.Left = leftChild
	root.Right = rightChild
	return root
}

var sum int

func getMinLeafTree(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

var count = 0

func kthSmallestInOrderTraverse(root *TreeNode, key int, value *int) {
	if root == nil {
		return
	}

	kthSmallestInOrderTraverse(root.Left, key, value)
	count++
	if count == key {
		*value = root.Val
		return
	}
	kthSmallestInOrderTraverse(root.Right, key, value)

}

func kthSmallest(root *TreeNode, k int) int {
	var value int
	kthSmallestInOrderTraverse(root, k, &value)
	return value
}

func inorderTraversal(root *TreeNode) []int {

	var result []int
	helpInorderTraversal(root, &result)

	return result

}

func helpInorderTraversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	helpInorderTraversal(root.Left, result)
	*result = append(*result, root.Val)
	helpInorderTraversal(root.Right, result)

}

//思路：
// 1. 层序遍历两棵树，看结果是否相同
// 2. 前序和中序遍历能确定一颗二叉树，对边两颗树的前序和中序遍历结果是否都相同
// 3. 同时遍历两棵树，对比结果
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSymmetric(root *TreeNode) bool {

	return helpIsSymmetric(root.Left, root.Right)

}

func helpIsSymmetric(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return helpIsSymmetric(l.Left, r.Right) && helpIsSymmetric(l.Right, r.Left)

}

func maxDepth2(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftMax := maxDepth2(root.Left) + 1
	rightMax := maxDepth2(root.Right) + 1

	if leftMax > rightMax {
		return leftMax
	}
	return rightMax
}

// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
// 叶子节点 是指没有子节点的节点

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}
	if (root.Left == root.Right) && root.Val == targetSum {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func pathSum(root *TreeNode, targetSum int) [][]int {

	var result [][]int
	var one []int
	helpPathSum(root, 0, targetSum, &result, one)
	return result
}

func helpPathSum(root *TreeNode, sum int, targetSum int, result *[][]int, one []int) {

	if root == nil {
		return
	}

	one = append(one, root.Val)
	sum += root.Val

	// 除了考虑sum == targetSum,还需要考虑边界条件，必须是root到叶子节点的和才满足要求
	if sum == targetSum && root.Left == nil && root.Right == nil {
		// 找到一个
		r := make([]int, len(one), len(one))
		// golang slice为引用类型，在此需要深拷贝一下
		copy(r, one)
		*result = append(*result, r)
	}

	helpPathSum(root.Left, sum, targetSum, result, one)
	helpPathSum(root.Right, sum, targetSum, result, one)
	sum -= root.Val
	one = one[:len(one)-1]

}

func sumOfLeftLeaves(root *TreeNode) int {

	return helpSumOfLeftLeaves(root, false)
}

func helpSumOfLeftLeaves(root *TreeNode, option bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil && option {
		return root.Val
	}

	return helpSumOfLeftLeaves(root.Left, true) + helpSumOfLeftLeaves(root.Right, false)

}

func binaryTreePaths(root *TreeNode) []string {

	var result []string
	helpBinaryTreePaths(root, &result, "")
	return result
}

func helpBinaryTreePaths(root *TreeNode, result *[]string, one string) {

	if root == nil {
		return
	}
	one += fmt.Sprint(root.Val)
	if root.Left == nil && root.Right == nil {
		// 其中一个路径
		*result = append(*result, one)
	}
	one += "->"

	helpBinaryTreePaths(root.Left, result, one)
	helpBinaryTreePaths(root.Right, result, one)

	out := strings.Split(one, "->")
	out = out[:len(out)-1]
	strings.Join(out, "->")

}

func levelOrder(root *TreeNode) [][]int {

	var result [][]int

	if root == nil {
		return result
	}

	var q []*TreeNode
	q = append(q, root)

	for len(q) > 0 {
		sizeq := len(q)
		var level = make([]int, 0, 0)
		for j := 0; j < sizeq; j++ {
			out := q[0]
			q[0] = nil
			q = q[1:]
			level = append(level, out.Val)
			if out.Left != nil {
				q = append(q, out.Left)
			}
			if out.Right != nil {
				q = append(q, out.Right)
			}
		}
		result = append(result, level)

	}

	return result
}
