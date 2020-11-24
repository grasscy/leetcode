package stack

func largestRectangleArea(heights []int) int {
	var s1 = []int{-1}
	var s2 []int
	max := -1

	for _, v := range heights {
		if v >= s1[len(s1)-1] {
			s1 = append(s1, v)
		} else {
			s2 = append(s2, v)
		}
	}

}
