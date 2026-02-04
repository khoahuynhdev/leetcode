package solution

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	p1, p2, pa1, pa2 := 0, 0, 0, 0
	for pa1 < len(word1) && pa2 < len(word2) {
		if word1[pa1][p1] != word2[pa2][p2] {
			return false
		}
		p1++
		if p1 >= len(word1[pa1]) {
			pa1++
			p1 = 0
		}
		p2++
		if p2 >= len(word2[pa2]) {
			pa2++
			p2 = 0
		}
		if (pa2 == len(word2) && pa1 < len(word1)) || (pa1 == len(word1) && pa2 < len(word2)) {
			return false
		}
	}
	return true
}
