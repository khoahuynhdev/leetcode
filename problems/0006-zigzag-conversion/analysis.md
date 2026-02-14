# 0006. Zigzag Conversion

[LeetCode Link](https://leetcode.com/problems/zigzag-conversion/)

Difficulty: Medium
Topics: String
Acceptance Rate: 53.4%

## Hints

### Hint 1

Think about how characters are distributed across rows. Instead of trying to build the zigzag pattern visually, consider collecting all characters that belong to each row separately.

### Hint 2

The zigzag pattern has a repeating cycle. Characters move down through rows 0, 1, 2, ..., numRows-1, then move back up through numRows-2, numRows-3, ..., 1, then down again. Track which direction you're moving (down or up) as you iterate through the string.

### Hint 3

Create an array of strings (or string builders) - one for each row. As you iterate through the input string, append each character to the appropriate row based on the current direction. When you hit the top or bottom row, reverse the direction.

## Approach

The key insight is that we don't need to actually construct the 2D zigzag pattern. Instead, we can simulate the zigzag movement and collect characters for each row.

**Algorithm:**
1. Handle edge cases: if numRows is 1 or the string length is less than numRows, return the original string
2. Create a slice of strings (one for each row) to collect characters
3. Initialize two variables:
   - `currentRow` to track which row we're currently on (starts at 0)
   - `goingDown` to track whether we're moving down or up (starts as false, will flip to true)
4. Iterate through each character in the string:
   - Append the character to the string for `currentRow`
   - If we're at the top row (0) or bottom row (numRows-1), flip the direction
   - Move to the next row: increment if going down, decrement if going up
5. Concatenate all row strings together and return

**Example walkthrough with "PAYPALISHIRING", numRows = 3:**
- Row 0: P (↓), A (↑), H (↑), N (↑)
- Row 1: A (↓), P (↓↑), L (↓), S (↓↑), I (↓), I (↓↑), G (↓)
- Row 2: Y (↓), I (↑), R (↑)
- Result: "PAHNAPLSIIGYIR"

The cycle length is `2 * numRows - 2` (down numRows steps, up numRows-2 steps).

## Complexity Analysis

Time Complexity: O(n) where n is the length of the string. We iterate through the string once, and each character is appended to a row exactly once.

Space Complexity: O(n) for storing the characters in row strings. The final result also requires O(n) space.

## Edge Cases

1. **numRows = 1**: No zigzag pattern possible, return the original string
2. **String length ≤ numRows**: The zigzag won't complete a full cycle, but the algorithm handles this naturally
3. **Single character string**: Should return as-is
4. **String shorter than numRows**: Characters only go down, never zigzag back up
