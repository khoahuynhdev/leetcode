package main

// Simulate binary operations by scanning right-to-left with a carry.
// Even bit (0 or 2) → 1 step (divide by 2).
// Odd bit (1) → 2 steps (add 1, then divide by 2) and propagate carry.
func numSteps(s string) int {
	steps := 0
	carry := 0

	for i := len(s) - 1; i >= 1; i-- {
		bit := int(s[i]-'0') + carry
		if bit == 1 { // odd
			steps += 2
			carry = 1
		} else { // bit == 0 or bit == 2, both even
			steps += 1
			if bit == 2 {
				carry = 1
			} else {
				carry = 0
			}
		}
	}

	// If carry overflows the MSB, one more divide step is needed.
	if carry == 1 {
		steps += 1
	}

	return steps
}
