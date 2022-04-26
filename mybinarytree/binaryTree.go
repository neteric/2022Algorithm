package mybinarytree

//
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

func rightSideView(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}
	var q []*TreeNode
	q = append(q, root)

	for len(q) > 0 {
		sizeq := len(q)
		for j := 0; j < sizeq; j++ {
			out := q[0]
			q[0] = nil
			q = q[1:]
			if j == sizeq-1 {
				result = append(result, out.Val)
			}
			if out.Left != nil {
				q = append(q, out.Left)
			}
			if out.Right != nil {
				q = append(q, out.Right)
			}
		}
	}

	return result
}

// 给你一个二叉树的根节点 root ，计算并返回 整个树 的坡度 。
// 一个树的 节点的坡度 定义即为，该节点左子树的节点之和和右子树节点之和的 差的绝对值 。如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是 0 。
// 整个树 的坡度就是其所有节点的坡度之和

func findTilt(root *TreeNode) int {
	var sum int
	helpFindTilt(root, &sum)
	return sum
}

func helpFindTilt(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}

	left := helpFindTilt(root.Left, sum)
	rigth := helpFindTilt(root.Right, sum)

	*sum += IntAbs(left - rigth)
	return left + rigth + root.Val
}
func IntAbs(in int) int {
	if in < 0 {
		return in * (-1)
	}
	return in
}

// 此思路为前序遍历，但无法满足题目要求（左子树为空时保留括号，右子树为空时不保留括号）
func tree2str(root *TreeNode) string {

	var s string
	helpTree2str(root, &s)
	return strings.TrimPrefix(strings.TrimSuffix(s, ")"), "(")

}

func helpTree2str(root *TreeNode, s *string) {

	if root == nil {
		return
	}
	*s += "("
	*s += fmt.Sprint(root.Val)

	helpTree2str(root.Left, s)
	helpTree2str(root.Right, s)

	*s += ")"

}

func helpTree2str2(root *TreeNode) string {

	if root == nil {
		return ""
	}

	if (root.Left == nil) && (root.Right == nil) {
		return fmt.Sprint(root.Val)
	}
	leftStr := helpTree2str2(root.Left)
	rightStr := helpTree2str2(root.Right)
	// 后序遍历代码位置
	// 根据左右子树字符串组装出前序遍历的顺序
	// 按题目要求处理 root 只有一边有子树的情况
	if root.Left != nil && root.Right == nil {
		return fmt.Sprint(root.Val) + "(" + leftStr + ")"
	}
	if root.Left == nil && root.Right != nil {
		return fmt.Sprint(root.Val) + "()" + "(" + rightStr + ")"
	}
	return fmt.Sprint(root.Val) + "(" + leftStr + ")" + "(" + rightStr + ")"
}

// number:617
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 != nil && root2 == nil {
		return root1
	}
	if root1 == nil && root2 != nil {
		return root2
	}

	leftTree := mergeTrees(root1.Left, root2.Left)
	rightTree := mergeTrees(root1.Right, root2.Right)

	t := new(TreeNode)
	t.Val = root1.Val + root2.Val
	t.Left = leftTree
	t.Right = rightTree

	return t

}

// number:897
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// if root.Left == nil && root.Right == nil {
	// 	return root
	// }
	l := increasingBST(root.Left)
	root.Left = nil
	r := increasingBST(root.Right)
	root.Right = r
	if l == nil {
		return root
	}
	x := l
	for x.Right != nil {
		x = x.Right
	}
	x.Right = root

	return l

}

// number: 103
func zigzagLevelOrder(root *TreeNode) [][]int {

	var result [][]int
	if root == nil {
		return result
	}
	var q []*TreeNode
	q = append(q, root)
	var level bool
	for len(q) > 0 {
		var sizeq = len(q)
		var levelValue = make([]int, 0, 0)
		for i := 0; i < sizeq; i++ {

			out := q[0]
			q = q[1:]
			levelValue = append(levelValue, out.Val)
			if out.Left != nil {
				q = append(q, out.Left)
			}
			if out.Right != nil {
				q = append(q, out.Right)
			}
		}
		if level {
			for i, j := 0, len(levelValue)-1; i < j; i, j = i+1, j-1 {
				levelValue[i], levelValue[j] = levelValue[j], levelValue[i]
			}
			level = false
		} else {
			level = true

		}
		result = append(result, levelValue)

	}

	return result

}

// number: 108
func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		root := new(TreeNode)
		root.Val = nums[0]
		return root
	}

	root := new(TreeNode)
	mid := len(nums) / 2
	root.Val = nums[mid]
	left := nums[:mid]
	right := nums[mid+1:]

	leftTree := sortedArrayToBST(left)
	rightTree := sortedArrayToBST(right)
	root.Left = leftTree
	root.Right = rightTree

	return root
}

// number:110
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var r bool = true
	max(root, &r)
	return r

}

func max(root *TreeNode, r *bool) int {

	if root == nil {
		return 0
	}

	left := max(root.Left, r) + 1
	right := max(root.Right, r) + 1

	if IntAbs(left-right) > 1 {
		*r = false
	}

	if left > right {
		return left
	}
	return right
}

// number:543
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var max int
	helpDiameterOfBinaryTree(root, &max)
	fmt.Println("x", max)
	return max - 2

}

func helpDiameterOfBinaryTree(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := helpDiameterOfBinaryTree(root.Left, max) + 1
	right := helpDiameterOfBinaryTree(root.Right, max) + 1
	// fmt.Println("left", left)
	// fmt.Println("right", right)
	// fmt.Println("z", root.Val)
	if left+right > *max {
		*max = left + right
	}

	if left > right {
		return left
	}
	return right

}

// number: 剑指 Offer 27
// 本题如果交换节点的value，就是误入歧途了
// func mirrorTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	if root.Left == nil && root.Right == nil {
// 		return root
// 	}
// 	helpMirrorTree(root.Left, root.Right)
// 	return root

// }

// func helpMirrorTree(l, r *TreeNode) {

// 	if l == nil && r == nil {
// 		return
// 	}

// 	if l == nil && r != nil {
// 		l := new(TreeNode)
// 		l.Val = r.Val
// 		r = nil
// 	} else if r == nil && l != nil {
// 		r := new(TreeNode)
// 		r.Val = l.Val
// 		l = nil
// 	} else {
// 		l.Val, r.Val = r.Val, l.Val
// 	}

// 	helpMirrorTree(l.Left, r.Right)
// 	helpMirrorTree(l.Right, r.Left)

// }

func mirrorTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)

	return root

}
