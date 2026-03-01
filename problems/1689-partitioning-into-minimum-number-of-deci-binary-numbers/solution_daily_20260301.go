package main

// Greedy: the minimum number of deci-binary numbers needed equals the
// maximum digit in n, since each deci-binary number contributes at most
// 1 to each digit position.
func minPartitions(n string) int {
	maxDigit := 0
	for _, ch := range n {
		d := int(ch - '0')
		if d > maxDigit {
			maxDigit = d
		}
		if maxDigit == 9 {
			break
		}
	}
	return maxDigit
}
