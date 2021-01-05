package stack

func nextGreaterElements(nums []int) []int {
	INT_MAX := int(^uint(0) >> 1)

	stk := make([]int, 0, len(nums))
	//stk = append(stk, INT_MAX)

	res := make([]int, len(nums)+1)

	nums = append([]int{INT_MAX}, nums...)

	l := len(stk)
	for i, v := range nums {
		if l == 0 {
			stk = append(stk, i)
			l++
			continue
		}

		if nums[stk[l-1]] < v {
			for nums[stk[l-1]] < v {
				res[stk[l-1]] = v
				stk = stk[:l-1]
				l--
			}
		}
		stk = append(stk, i)
		l++
	}

	for i := 1; i < len(nums)-1; i++ {
		if l == 1 {
			break
		}

		for nums[stk[l-1]] < nums[i] {
			res[stk[l-1]] = nums[i]
			stk = stk[:l-1]
			l--
		}
	}

	for l > 1 {
		res[stk[l-1]] = -1
		stk = stk[:l-1]
		l--
	}

	return res[1:]

}
