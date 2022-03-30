package mybinarytree

import "fmt"

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

func postOderTraverse(head *TreeNode) {
	if head == nil {
		return
	}
	postOderTraverse(head.Left)
	postOderTraverse(head.Right)
	fmt.Println(head.Val)

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
