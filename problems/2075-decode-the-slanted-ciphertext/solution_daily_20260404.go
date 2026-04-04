// Decode the Slanted Ciphertext
// Approach: Reconstruct the matrix dimensions from encodedText and rows,
// then read characters along each diagonal (top-left to bottom-right)
// to recover the original text. Trim trailing spaces.
package main

import "strings"

func decodeCiphertext(encodedText string, rows int) string {
	n := len(encodedText)
	if n == 0 {
		return ""
	}

	cols := n / rows

	var sb strings.Builder
	// Each diagonal starts at (0, col) and walks down-right
	for col := 0; col < cols; col++ {
		for r, c := 0, col; r < rows && c < cols; r, c = r+1, c+1 {
			sb.WriteByte(encodedText[r*cols+c])
		}
	}

	return strings.TrimRight(sb.String(), " ")
}
