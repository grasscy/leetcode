package stack

func find132pattern(nums []int) bool {
	mstk := make([]int, 0, len(nums))

	for i, v := range nums {
		if len(mstk) == 0 {
			mstk = append(mstk, i)
			continue
		}

		if v > nums[mstk[len(mstk)-1]] {
			mstk = append(mstk, i)
		} else if v < nums[mstk[len(mstk)-1]] {
			for j := 0; nums[j] < nums[mstk[len(mstk)-1]]; j++ {
				if nums[j] > v {
					return true
				}
			}
		}

	}
	return false
}
