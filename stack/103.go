package stack

func zigzagLevelOrder(root *TreeNode) [][]int {
	var s []*TreeNode
	var res [][]int
	if root == nil {
		return res
	}
	s = append(s, root)
	i := 0
	for len(s) > 0 {
		root = s[len(s)-1]
		var r []int
		var tmp []*TreeNode
		for len(s) > 0 {
			root = s[len(s)-1]
			r = append(r, root.Val)
			s = s[:len(s)-1]
			if i%2 == 0 {
				if root.Left != nil {
					tmp = append(tmp, root.Left)
				}
				if root.Right != nil {
					tmp = append(tmp, root.Right)
				}
			} else {
				if root.Right != nil {
					tmp = append(tmp, root.Right)
				}
				if root.Left != nil {
					tmp = append(tmp, root.Left)
				}
			}
		}
		i++
		s = append(s, tmp...)
		res = append(res, r)
	}
	return res
}
