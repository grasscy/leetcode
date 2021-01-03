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
			l := len(stk)
			for l > 0 && stk[l-1] > v && ns[stk[l-1]] > 0 {
				mark[stk[l-1]] = false
				stk = stk[:l-1]
				l = len(stk)
			}
			stk = append(stk, v)
			mark[v] = true
		}
		ns[v]--
	}

	return string(stk)
}
