package stack

import "container/list"

func nextGreaterElements(nums []int) []int {
	stk := list.New()
	cd := 0
	res := make([]int, 0, len(nums))
	for _, v := range nums {
		if stk.Len() == 0 {
			stk.PushBack(v)
			cd = v
			continue
		}
	}
}
