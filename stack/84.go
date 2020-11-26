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
