package stack

import "strconv"

func evalRPN(tokens []string) int {
	s := make([]int, 0, len(tokens))
	for _, v := range tokens {
		switch v {
		case "+":
			a := s[len(s)-1]
			s = s[:len(s)-1]
			b := s[len(s)-1]
			s = s[:len(s)-1]
			r := a + b
			s = append(s, r)
		case "-":
			a := s[len(s)-1]
			s = s[:len(s)-1]
			b := s[len(s)-1]
			s = s[:len(s)-1]
			r := b - a
			s = append(s, r)
		case "*":
			a := s[len(s)-1]
			s = s[:len(s)-1]
			b := s[len(s)-1]
			s = s[:len(s)-1]
			r := a * b
			s = append(s, r)
		case "/":
			a := s[len(s)-1]
			s = s[:len(s)-1]
			b := s[len(s)-1]
			s = s[:len(s)-1]
			r := b / a
			s = append(s, r)
		default:
			atoi, _ := strconv.Atoi(v)
			s = append(s, atoi)
		}
	}
	return s[len(s)-1]

}
