package stack

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	max := 0
	rowN := len(matrix)
	colN := len(matrix[0])

	hs := make([]int, colN)
	for i := 0; i < rowN; i++ {
		for j := 0; j < colN; j++ {
			if matrix[i][j] == 1 {
				hs[j]++
			} else {
				hs[j] = 0
			}
		}
		r := largestRectangleArea1(hs)
		if r > max {
			max = r
		}
	}

	return max
}
