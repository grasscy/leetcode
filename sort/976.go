package sort

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)

	for k := len(A) - 1; k > 1; k-- {
		if A[k-2]+A[k-1] > A[k] {
			return A[k-2] + A[k-1] + A[k]
		}
	}

	return 0
}
