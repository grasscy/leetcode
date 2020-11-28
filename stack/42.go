package stack

func trap(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}

	sum := 0
	s := make([]int, 0, len(height))

	for i, v := range height {
		l := len(s)
		if l == 0 {
			s = append(s, i)
			continue
		}
		for l > 0 && v > height[s[l-1]] {
			h := -height[s[l-1]]
			s = s[:l-1]
			l = len(s)
			if l < 1 {
				continue
			}
			if v > height[s[l-1]] {
				h += height[s[l-1]]
			} else {
				h += v
			}
			sum += h * (i - s[l-1] - 1)
		}
		s = append(s, i)
	}
	return sum
}
