package solution

// NOTE: this implementation is good enough
// for better solution,
// first and last index array to store position of the appeared vowels
// after that start counting normally
func countPalindromicSubsequence(s string) int {
	res := make(map[string]struct{})
	left := make(map[string]struct{})
	right := make(map[string]int)
	// Declare an array of 26 lowercase characters
	var alt [26]string
	// Initialize the array with lowercase letters 'a' to 'z'
	for i := 0; i < 26; i++ {
		alt[i] = string(byte('a' + i))
	}

	for _, v := range s {
		right[string(v)]++
	}

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if right[c]--; right[c] <= 0 {
			delete(right, c)
		}

		for _, v := range alt {
			if _, ok := left[v]; ok && right[v] > 0 {
				res[fmt.Sprintf("%s%s", c, v)] = struct{}{}
			}
		}
		left[c] = struct{}{}
	}

	return len(res)
}
