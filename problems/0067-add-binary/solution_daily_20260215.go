package main

// addBinary adds two binary strings by simulating manual binary addition
// from right to left, tracking carry at each step.
func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	result := []byte{}

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		result = append(result, byte(sum%2)+'0')
		carry = sum / 2
	}

	// Reverse the result
	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}

	return string(result)
}
