package stack

func postorderTraversal(root *TreeNode) []int {
	var res []int
	postorder(root, &res)
	return res
}

func postorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, res)
	postorder(root.Right, res)
	*res = append(*res, root.Val)
}
