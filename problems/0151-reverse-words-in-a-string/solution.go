package solution

import (
	"strings"
)

// intuition: use 2 ptrs
// better to do in the reverse of array

func reverseWords(s string) string {
	wd := []string{}
	l, r := len(s), len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if l == r {
				l, r = i, i
				continue
			}
			wd = append(wd, string(s[l:r]))

			l, r = i, i
		} else {
			l = i
			if l == 0 {
				wd = append(wd, string(s[l:r]))
			}
		}
	}
	return strings.Join(wd, " ")
}
