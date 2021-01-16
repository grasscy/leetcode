package sort

func sortArrayByParityII(A []int) []int {
	res := make([]int, len(A))
	i := 0
	j := 1
	for _, v := range A {
		if v%2 == 0 {
			res[i] = v
			i += 2
		} else {
			res[j] = v
			j += 2
		}
	}
	return res

}
