package solution

func canBuild(word string, wd map[string]int) bool {
	tmp := map[string]int{}
	for k, v := range wd {
		tmp[k] = v
	}
	for _, w := range word {
		tmp[string(w)]--
		if tmp[string(w)] < 0 {
			return false
		}
	}
	return true
}

func countCharacters(words []string, chars string) int {
	wd := map[string]int{}
	ans := 0
	for _, c := range chars {
		wd[string(c)]++
	}
	for _, w := range words {
		ok := true
		for _, c := range w {
			if wd[string(c)] == 0 {
				ok = false
				break
			}
		}
		if !canBuild(w, wd) {
			ok = false
		}
		if ok {
			ans += len(w)
		}
	}
	return ans
}
