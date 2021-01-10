package stack

func dailyTemperatures(T []int) []int {
	stk := make([]int, 0, len(T))
	res := make([]int, len(T))

	l := len(stk)
	for i, v := range T {
		for l > 0 && v > T[stk[l-1]] {
			pop := stk[l-1]
			stk = stk[:l-1]
			l--
			res[pop] = i - pop
		}
		stk = append(stk, i)
		l++
	}
	return res

}
