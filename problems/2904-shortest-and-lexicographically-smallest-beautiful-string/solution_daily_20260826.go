package main

// 2904. Shortest and Lexicographically Smallest Beautiful String
//
// A shortest beautiful substring must start and end with '1': any leading or
// trailing '0' can be trimmed while keeping exactly k ones, yielding a strictly
// shorter beautiful substring. So every candidate is the span from the i-th '1'
// to the (i+k-1)-th '1'. We collect the indices of all ones, slide a fixed
// window of size k over that index list, and keep the shortest candidate,
// breaking ties with a lexicographic comparison.
func shortestBeautifulSubstring(s string, k int) string {
	ones := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ones = append(ones, i)
		}
	}
	if len(ones) < k {
		return ""
	}

	best := ""
	for i := 0; i+k-1 < len(ones); i++ {
		candidate := s[ones[i] : ones[i+k-1]+1]
		if best == "" || len(candidate) < len(best) ||
			(len(candidate) == len(best) && candidate < best) {
			best = candidate
		}
	}
	return best
}
