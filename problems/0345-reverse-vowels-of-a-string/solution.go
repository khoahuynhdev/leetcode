package solution

func isVowel(s string) bool {
	return s == "a" || s == "e" || s == "i" || s == "o" || s == "u" || s == "A" || s == "E" || s == "I" || s == "O" || s == "U"
}

// Intuition: 2 pointers, this one care about positions
// time complexity: O(n)
// space complexity: O(n)
func reverseVowels(s string) string {
	str := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		if isVowel(string(str[l])) && isVowel(string(str[r])) {
			str[l], str[r] = str[r], str[l]
			l++
			r--
			continue
		}
		if isVowel(string(str[l])) && !isVowel(string(str[r])) {
			r--
			continue
		}
		if !isVowel(string(str[l])) && isVowel(string(str[r])) {
			l++
			continue
		}
		l++
		r--
	}
	return string(str)
}
