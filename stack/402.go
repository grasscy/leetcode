package stack

func removeKdigits(num string, k int) string {
	res := make([]uint8, 0, len(num))
	for i := 0; i < len(num); i++ {
		if len(res) > 0 && res[len(res)-1] > num[i] {
			if k > 0 {
				res = res[:len(res)-1]
				k--
			}
		}
		res = append(res, num[i])
	}
	return res
}
