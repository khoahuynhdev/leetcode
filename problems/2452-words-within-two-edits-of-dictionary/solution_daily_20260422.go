package main

// twoEditWords returns every query that is within Hamming distance 2 of some
// word in dictionary. All words share the same length, so "edits" are just
// in-place character substitutions — the minimum number of edits between two
// equal-length words is the count of mismatched positions. We scan every
// (query, dictionary) pair and short-circuit as soon as mismatches exceed 2.
func twoEditWords(queries []string, dictionary []string) []string {
	result := []string{}
	for _, q := range queries {
		for _, d := range dictionary {
			diff := 0
			for i := 0; i < len(q); i++ {
				if q[i] != d[i] {
					diff++
					if diff > 2 {
						break
					}
				}
			}
			if diff <= 2 {
				result = append(result, q)
				break
			}
		}
	}
	return result
}
