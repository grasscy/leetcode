package stack

func postorderTraversal(root *TreeNode) []int {
	var s []*TreeNode
	var res []int
	for root != nil || len(s) > 0 {
		for root != nil {
			s = append(s, root)
			s = append(s, root.Right)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]
		res = append(res, root.Val)
	}
	return res
}
