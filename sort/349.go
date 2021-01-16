package sort

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)

	for _, v := range nums1 {
		m[v] = true
	}

	var res []int

	for _, v := range nums2 {
		b, ok := m[v]
		if ok && b {
			res = append(res, v)
			m[v] = false
		}
	}
	return res
}
