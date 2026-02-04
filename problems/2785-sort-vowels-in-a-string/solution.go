package solution

func sortVowels(s string) string {
	vls := []string{"A", "E", "I", "O", "U", "a", "e", "i", "o", "u"}
	isVowel := make(map[string]bool)
	for _, v := range vls {
		isVowel[v] = true
	}
	vowels := make(map[string]int)
	pos := []int{}
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if isVowel[char] {
			vowels[char]++
			pos = append(pos, i)
		}
	}
	// fmt.Println(pos)
	// fmt.Println(vowels)
	str := []byte(s)
	for i := 0; i < len(vls); i++ {
		v := vls[i]
		for vowels[v] > 0 {
			p := pos[0]
			str[p] = byte(v[0])
			pos = pos[1:]
			vowels[v]--
		}
	}
	return string(str)
}
