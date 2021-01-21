package sort

import "sort"

func merge(intervals [][]int) [][]int {
	res := make([][]int, len(intervals))
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	index := 0
	lend := intervals[0][1]
	res[0] = intervals[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > lend {
			res[index+1] = intervals[i]
			res[index][1] = lend
			index++
		}
		if intervals[i][0] > lend || intervals[i][1] > lend {
			lend = intervals[i][1]
		}
	}

	res[index][1] = lend

	return res[:index+1]
}
