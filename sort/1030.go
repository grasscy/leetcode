package sort

import (
	"math"
	"sort"
)

var r int
var c int

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	r = r0
	c = c0

	res := make(xy, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			res = append(res, []int{i, j})
		}
	}
	sort.Sort(res)

	return res
}

type xy [][]int

func (a xy) Len() int      { return len(a) }
func (a xy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a xy) Less(i, j int) bool {
	return math.Abs(float64(a[i][0]-r))+math.Abs(float64(a[i][1]-c)) < math.Abs(float64(a[j][0]-r))+math.Abs(float64(a[j][1]-c))
}
