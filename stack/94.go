package stack

func inorderTraversal(root *TreeNode) []int {
	var s []*TreeNode
	var res []int

	cur := root

	for cur != nil {
		for cur.Left != nil {
			s = append(s, cur)
			cur = cur.Left
		}
		res = append(res, cur.Val)
		if cur.Right != nil {
			cur = cur.Right
		} else {
			if len(s) == 0 {
				break
			}
			cur = s[len(s)-1]
			cur.Left = nil
			s = s[:len(s)-1]
		}
	}
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
