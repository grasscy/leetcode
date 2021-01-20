package sort

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sortst := make(sortByStart, 0, len(points))
	sortst = append(sortst, points...)
	sort.Sort(sortst)

	res := 0

	lend := -2<<31 - 1
	for _, v := range sortst {
		if v[0] > lend {
			res++
			lend = v[1]
			continue
		}
		if lend > v[1] { // 按右端升序排则无需处理此情况
			lend = v[1]
		}
	}

	return res
}

type sortByStart [][]int

func (a sortByStart) Len() int      { return len(a) }
func (a sortByStart) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortByStart) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}
