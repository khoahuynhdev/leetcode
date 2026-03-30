package main

// Approach: Characters at even indices can be freely rearranged among even positions,
// and characters at odd indices among odd positions (since j-i even means same parity).
// We just need to check that character frequencies match for each parity group.
func checkStrings(s1 string, s2 string) bool {
	var evenCount, oddCount [26]int

	for i := 0; i < len(s1); i++ {
		if i%2 == 0 {
			evenCount[s1[i]-'a']++
			evenCount[s2[i]-'a']--
		} else {
			oddCount[s1[i]-'a']++
			oddCount[s2[i]-'a']--
		}
	}

	for i := 0; i < 26; i++ {
		if evenCount[i] != 0 || oddCount[i] != 0 {
			return false
		}
	}
	return true
}
