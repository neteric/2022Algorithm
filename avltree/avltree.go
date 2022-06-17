package avltree

const (
	null = -1
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 此方法仅适用于完全二叉树
func MakeTree(element []int) *TreeNode {

	var tree []*TreeNode

	for j := 0; j < len(element); j++ {
		tree = append(tree, new(TreeNode))
	}

	for i := 0; i < len(element); i++ {

		if element[i] != null {
			tree[i].Val = element[i]
		} else {
			continue
		}
		if 2*i+1 < len(element) {
			if element[2*i+1] != null {
				tree[i].Left = tree[2*i+1]
			} else {
				tree[i].Left = nil
			}
		}

		if 2*i+2 < len(element) {
			if element[2*i+2] != null {
				tree[i].Right = tree[2*i+2]
			} else {
				tree[i].Right = nil
			}
		}
	}
	return tree[0]
}

// number: 110
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1

	if left > right {
		return left
	}
	return right

}
func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if AbsInt(left-right) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)

}

func AbsInt(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}
