package stack

func preorderTraversal(root *TreeNode) []int {
	var s []*TreeNode
	var res []int
	if root == nil {
		return res
	}
	s = append(s, root)
	for len(s) > 0 {
		root = s[len(s)-1]
		s = s[:len(s)-1]
		res = append(res, root.Val)
		if root.Right != nil {
			s = append(s, root.Right)
		}
		if root.Left != nil {
			s = append(s, root.Left)
		}
	}
	return res
}
