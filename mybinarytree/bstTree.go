package mybinarytree

func isValidBST(root *TreeNode) bool {
	return helpValidBST(root, nil, nil)
}

func helpValidBST(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if max != nil && root.Val >= max.Val {
		return false
	}
	if min != nil && root.Val <= min.Val {
		return false
	}

	return helpValidBST(root.Left, min, root) && helpValidBST(root.Right, root, max)
}

func sumBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	return sumBST(root.Left) + sumBST(root.Right) + root.Val

}

func maxSumBST(root *TreeNode) int {

	if root == nil {
		return 0
	}
	if isValidBST(root) {
		return sumBST(root)
	}
	leftMax := maxSumBST(root.Left)
	rightMax := maxSumBST(root.Right)

	var maxSum int
	if leftMax > rightMax {
		maxSum = leftMax
	} else {
		maxSum = rightMax
	}
	if maxSum < 0 {
		return 0
	}
	return maxSum
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	convertBST(root.Right)
	sum += root.Val
	// fmt.Println(sum)
	root.Val = sum
	convertBST(root.Left)

	return root
}

func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	} else if val > root.Val {
		return searchBST(root.Right, val)
	}

	return root

}

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		leaf := new(TreeNode)
		leaf.Val = val
		return leaf
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key == root.Val {

		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minLeafTree := getMinLeafTree(root.Right)
		// fmt.Println("minLeafTree:", minLeafTree.Val)
		root.Val = minLeafTree.Val
		root.Right = deleteNode(root.Right, minLeafTree.Val)

	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
