package main

// Greedy / game theory, single pass.
//
// Let d = (sum of known digits in the left half) - (sum of known digits in the right
// half), and m = (count of '?' in the right half) - (count of '?' in the left half).
//
// A '?' on the left pairs with a '?' on the right: whatever digit one player writes,
// the opponent mirrors it into the partner slot, so matched pairs never change d.
// What remains is |m| unmatched '?' on one side.
//
//   - If m is odd (equivalently, the total number of '?' is odd), Alice places the last
//     digit and can always avoid the single value that would equalize the halves.
//   - If m = 2k is even, Bob can lock the unmatched block at exactly 9k by answering
//     every x with 9-x, and Alice can force it off 9k by opening with 9 (to inflate) or
//     0 (to deflate). So Bob wins iff d == 9k.
//
// The comparison is written as 2*d == 9*m to avoid dividing a possibly negative m.
func sumGame(num string) bool {
	n := len(num)
	sumDiff, qDiff := 0, 0

	for i := 0; i < n; i++ {
		if num[i] == '?' {
			// qDiff counts right-half '?' minus left-half '?'.
			if i < n/2 {
				qDiff--
			} else {
				qDiff++
			}
			continue
		}
		// sumDiff counts left-half digits minus right-half digits.
		digit := int(num[i] - '0')
		if i < n/2 {
			sumDiff += digit
		} else {
			sumDiff -= digit
		}
	}

	// Odd number of '?': Alice moves last and always breaks the tie.
	if qDiff%2 != 0 {
		return true
	}

	// Even: Bob wins exactly when the known gap equals the value he can lock in.
	return 2*sumDiff != 9*qDiff
}
