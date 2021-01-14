package stack

// s[K] = s[k%size]
func decodeAtIndex(S string, K int) string {
	l := 0
	for i := 0; i < len(S); i++ {
		v := S[i]
		if '1' < v && v <= '9' {
			ll := int(v - '0')
			if ll*l <= K-1 {
				l *= ll
			} else {
				return decodeAtIndex(S, (K-1)%l+1)
			}

		} else {
			l++
			if K == l {
				return string(v)
			}
		}
	}
	return ""
}
