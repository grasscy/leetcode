package sort

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	for _, v := range nums1 {
		b, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] = b + 1
		}
	}

	var res []int

	for _, v := range nums2 {
		b, ok := m[v]
		if ok && b > 0 {
			res = append(res, v)
			m[v] = b - 1
		}
	}
	return res
}
