package main

// 3838. Weighted Word Mapping
//
// Pure simulation. For each word we sum the weights of its characters
// (weights[c-'a']), reduce the sum modulo 26, then map the result r to a
// lowercase letter in reverse alphabetical order (0 -> 'z', ..., 25 -> 'a').
// Since index r and its target letter sum to the value of 'z', the mapped byte
// is simply 'z' - r. Results are appended to a byte slice and converted once.
func weightedWordMapping(words []string, weights []int) string {
	result := make([]byte, 0, len(words))
	for _, word := range words {
		weight := 0
		for i := 0; i < len(word); i++ {
			weight += weights[word[i]-'a']
		}
		r := weight % 26
		result = append(result, byte('z'-r))
	}
	return string(result)
}
