package stack

func scoreOfParentheses(S string) int {
	stk := make([]int32, 0, len(S))

	for _, v := range S {
		if v == ')' {
			if stk[len(stk)-1] == -1 {
				stk[len(stk)-1] = 1
			} else {
				sc := int32(0)

				for len(stk) > 0 && stk[len(stk)-1] != -1 {
					p := stk[len(stk)-1]
					stk = stk[:len(stk)-1]
					sc += p
				}
				stk[len(stk)-1] = sc * 2
			}

		} else {
			stk = append(stk, -1)
		}
	}

	ss := int32(0)
	for _, v := range stk {
		ss += v
	}

	return int(ss)
}
