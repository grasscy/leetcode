package stack

func decodeAtIndex(S string, K int) string {
	s := ""
	l := 0
	for i := 0; i < len(S); i++ {
		v := S[i]
		if '1' < v && v <= '9' {
			l = l * (int(v-'1') + 1)
			ss := s
			if l == K {
				return string(s[len(s)-1])
			}
			for j := 1; j < (int(v-'1') + 1); j++ {
				s += ss
			}
			if l > K {
				break
			}
		} else {
			l++
			if K == l {
				return string(v)
			}
			s += string(v)
		}
	}
	return string(s[K-1])
}
