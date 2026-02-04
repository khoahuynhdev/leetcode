package solution

func halvesAreAlike(s string) bool {
	hlf := len(s) / 2
	cnt := 0
	for i := 0; i < len(s); i++ {
		switch rune(s[i]) {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			if i < hlf {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return cnt == 0
}
