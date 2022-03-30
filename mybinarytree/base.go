package mybinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	null = -1
)

func MakeTree(element []int) *TreeNode {

	var tree []*TreeNode

	for j := 0; j < len(element); j++ {
		tree = append(tree, new(TreeNode))
	}

	for i := 0; i < len(element); i++ {

		if element[i] != null {
			tree[i].Val = element[i]
		} else {
			tree[i] = nil
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
