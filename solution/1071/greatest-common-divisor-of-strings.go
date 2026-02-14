package solution

import "strings"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcdOfStrings(str1 string, str2 string) string {
	var sm string
	var lg string
	// l1,l2 := len(str1), len(str2)
	if len(str1) < len(str2) {
		sm = str1
	} else {
		sm = str2
	}
	if len(str1) < len(str2) {
		lg = str2
	} else {
		lg = str1
	}
	r := 0
	for i := 0; i < len(sm)-1; i++ {
		if !strings.Contains(lg, sm[0:r+1]) {
			return ""
		}
		r++
	}
	var str string
	if len(lg)%len(sm) == 0 {
		str = sm[0 : r+1]
	} else {
		str = sm[0:gcd(len(lg), len(sm))]
	}
	var tmp string = str

	for len(tmp) < len(lg) {
		tmp += str
	}
	if tmp == lg {
		return str
	}

	return ""
}
