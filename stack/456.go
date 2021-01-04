package stack

func find132pattern(nums []int) bool {
	mstk := make([]int, 0, len(nums))

	for i, v := range nums {
		l := len(mstk)
		if l == 0 {
			mstk = append(mstk, v)
			continue
		}

		if v > mstk[l-1] {
			mstk = append(mstk, v)
		} else if v < mstk[l-1] {
			if v > mstk[0] {
				return true
			}
			for j := i + 1; j < len(nums); j++ {
				if nums[j] < mstk[l-1] && nums[j] > mstk[0] {
					return true
				}
			}
			for l > 0 && mstk[l-1] > v {
				mstk = mstk[:l-1]
				l--
			}
			mstk = append(mstk, v)
		}

	}
	return false
}

// 当存在某个元素小于其右边的次大元素 是 数组存在132模式的充要条件
func find132pattern1(nums []int) bool {
	stk := make([]int, 0, len(nums))
	cd := ^int(^uint(0) >> 1)
	for i := len(nums) - 1; i > -1; i-- {
		if cd > nums[i] {
			return true
		}
		for len(stk) > 0 && stk[len(stk)-1] < nums[i] {
			cd = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, nums[i])
	}
	return false
}
