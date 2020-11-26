package stack

func trap(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}

	height = append([]int{111111}, height...)
	height = append(height, 111111)

	sum := 0
	s := make([]int, 0, len(height))

	for i, v := range height {
		if len(s) == 0 {
			s = append(s, i)
			continue
		}
		for v > height[s[len(s)-1]] {
			s = s[:len(s)-1]
			h := -height[s[len(s)-1]]
			if v > height[s[len(s)-1]] {
				h += height[s[len(s)-1]]
			} else {
				h += v
			}
			sum += h * (i - s[len(s)-1])
		}
		s = append(s, i)
	}
	return sum
}
