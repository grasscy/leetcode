package stack

func largestRectangleArea(heights []int) int {
	max := 0
	for i, v := range heights {
		countH := 1
		for j := i + 1; j < len(heights) && heights[j] >= v; j++ {
			countH++
		}
		for j := i - 1; j >= 0 && heights[j] >= v; j-- {
			countH++
		}

		if max < (countH * v) {
			max = countH * v
		}
	}
	return max
}

func largestRectangleArea1(heights []int) int {

	if len(heights) == 0 {
		return 0
	}

	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	s := make([]int, 0, len(heights)+1)
	s = append([]int{0}, s...)

	max := 0

	for i, v := range heights {
		for v < heights[s[len(s)-1]] {
			h := heights[s[len(s)-1]]
			s = s[:len(s)-1]
			w := i - s[len(s)-1] - 1
			if max < h*w {
				max = h * w
			}
		}
		s = append(s, i)

	}
	return max
}
