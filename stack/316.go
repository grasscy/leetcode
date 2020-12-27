package stack

func removeDuplicateLetters(s string) string {
	//empty := make([]int32, 0, len(s))
	//stk := empty

	res := make([]int32, 0, len(s))
	masks := make([]bool, 'z'+1)
	for _, v := range s {
		if !masks[v] {
			res = append(res, v)
			masks[v] = true
		} else {

		}

	}

	return res
}
