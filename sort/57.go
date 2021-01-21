package sort

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	res := make([][]int, 0, len(intervals)+1)

	var i int
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		res = append(res, intervals[i])
	}

	if i == len(intervals) {
		res = append(res, newInterval)
		return res
	}

	res = append(res, []int{intervals[i][0], newInterval[1]})
	if intervals[i][0] > newInterval[0] {
		res[len(res)-1][0] = newInterval[0]
	}
	j := i
	for ; j < len(intervals) && intervals[j][0] <= newInterval[1]; j++ {

	}
	if j > 0 && intervals[j-1][1] > newInterval[1] {
		res[len(res)-1][1] = intervals[j-1][1]
	}

	for k := j; k < len(intervals); k++ {
		res = append(res, intervals[k])
	}
	return res
}

func insert2(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	l := len(intervals)
	i := 0
	for i < l && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < l && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	for i < l {
		res = append(res, intervals[i])
		i++
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
