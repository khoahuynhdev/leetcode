package main

// Approach: Simulate the zigzag pattern by collecting characters for each row
// As we iterate through the string, we track which row we're on and whether
// we're moving down or up. When we hit the top or bottom row, we reverse direction.
// Finally, concatenate all rows to get the result.

func convert(s string, numRows int) string {
	// Edge case: if only 1 row or string is too short, no zigzag needed
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	// Create a slice to hold characters for each row
	rows := make([]string, numRows)
	currentRow := 0
	goingDown := false

	// Iterate through each character and append to appropriate row
	for _, char := range s {
		rows[currentRow] += string(char)

		// Reverse direction when we hit top or bottom row
		if currentRow == 0 || currentRow == numRows-1 {
			goingDown = !goingDown
		}

		// Move to next row
		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	// Concatenate all rows
	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}
