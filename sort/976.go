package sort

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)

	k := len(A) - 1
	j := k - 1
	i := k - 2

	res := 0

	for {
		if A[i]+A[j] > A[k] {
			return A[i] + A[j] + A[k]
		}
		if A[i]+A[j] == k {
			return
		}

	}

	return 0
}
