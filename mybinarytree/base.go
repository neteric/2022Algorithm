package mybinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	null = -1
)

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

//此方法适用于所有二叉树
func MakeTreeByLevelList(element []int) *TreeNode {
	// 空树情况
	if len(element) == 0 {
		root := new(TreeNode)
		return root
	}

	var q []*TreeNode
	root := new(TreeNode)
	q = append(q, root)
	root.Val = element[0]
	var j = 0
	for len(q) > 0 {
		// 从队列中取一个元素
		out := q[0]
		q[0] = nil
		q = q[1:]

		// 从层序列表中取一个元素作为取出来元素的左子树的值
		j++
		if j >= len(element) {
			break
		}
		if element[j] != null {
			left := new(TreeNode)
			left.Val = element[j]
			//将新创建出来的左子树加入队列，以便下次处理
			q = append(q, left)
			out.Left = left
		} else {
			out.Left = nil
		}

		// 从层序列表中取一个元素作为取出来元素的右子树的值
		j++
		if j >= len(element) {
			break
		}
		if element[j] != null {
			right := new(TreeNode)
			right.Val = element[j]
			//将新创建出来的左子树加入队列，以便下次处理
			q = append(q, right)
			out.Right = right
		} else {
			out.Right = nil
		}
	}
	return root

}
