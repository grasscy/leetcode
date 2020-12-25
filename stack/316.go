package stack

func removeDuplicateLetters(s string) string {
	empty := make([]int32, 0, len(s))
	stk := empty

	for _, v := range s {
		for len(stk) > 0 && stk[0] > v {
			stk = empty
		}
		if len(stk) == 0 {
			stk = append(stk, v)
		} else if stk[len(stk)-1] < v {
			stk = append(stk, v)
		}
	}

	return string(stk)
}
