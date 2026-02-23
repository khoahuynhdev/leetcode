package main

// hasAllCodes checks whether every binary code of length k appears as a
// substring of s. It uses a rolling bit-hash: as we slide a window of size k
// across s, we maintain the window's integer value with O(1) shift-and-mask
// operations, recording seen codes in a boolean array of size 2^k.
// Time: O(n)  Space: O(2^k)
func hasAllCodes(s string, k int) bool {
	n := len(s)
	if n < k {
		return false
	}

	need := 1 << k
	// Early exit: not enough windows to cover all 2^k codes.
	if n-k+1 < need {
		return false
	}

	seen := make([]bool, need)
	mask := need - 1
	hash := 0

	for i := 0; i < n; i++ {
		hash = (hash<<1 | int(s[i]-'0')) & mask
		if i >= k-1 {
			if !seen[hash] {
				seen[hash] = true
				need--
				if need == 0 {
					return true
				}
			}
		}
	}

	return false
}
