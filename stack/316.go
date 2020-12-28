package stack

func removeDuplicateLetters(s string) string {

	ns := make([]int32, 'z'+1)
	stk := make([]int32, 0, len(s))
	mark := make([]bool, 'z'+1)
	for _, v := range s {
		ns[v]++
	}

	for _, v := range s {
		if !mark[v] {
			for len(stk) > 0 && stk[len(stk)-1] > v && ns[v] > 0 {
				mark[len(stk)-1] = false
				stk = stk[:len(stk)-1]
			}
			stk = append(stk, v)
			mark[v] = true
			ns[v]--
		}
	}

	return string(stk)
}
