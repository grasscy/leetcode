package stack

import (
	"strings"
)

func simplifyPath(path string) string {
	split := strings.Split(path, "/")
	res := make([]string, 0, len(split))
	for _, v := range split {
		switch v {
		case "..":
			if len(res) == 0 {
				continue
			}
			res = res[:len(res)-1]
		case ".", "":
		default:
			res = append(res, v)
		}
	}
	return "/" + strings.Join(res, "/")
}
