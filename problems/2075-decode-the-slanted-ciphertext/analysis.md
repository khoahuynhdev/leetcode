# 2075. Decode the Slanted Ciphertext

[LeetCode Link](https://leetcode.com/problems/decode-the-slanted-ciphertext/)

Difficulty: Medium
Topics: String, Simulation
Acceptance Rate: 55.1%

## Hints

### Hint 1

Think about what the encoding process looks like as a 2D matrix. If you can reconstruct the matrix, you can reverse the encoding by reading it in the right order.

### Hint 2

The encoded text was formed by reading the matrix row by row. So if you know the number of columns, you can rebuild the matrix. The number of columns is simply `len(encodedText) / rows`. Once the matrix is rebuilt, you need to read the diagonal stripes to recover the original text.

### Hint 3

Each diagonal stripe starts at row 0 and a different column (0, 1, 2, ...). From each starting column, you walk down the diagonal: `(0, col), (1, col+1), (2, col+2), ...` until you leave the matrix bounds. Concatenate the characters from all diagonals in order, then trim trailing spaces.

## Approach

1. **Compute dimensions**: Calculate `cols = len(encodedText) / rows`. The encoded text fills exactly `rows * cols` characters.

2. **Reconstruct the matrix**: The encoded text is the matrix read row by row, so character at index `i` maps to position `(i / cols, i % cols)` in the matrix. We don't actually need to build a 2D array -- we can index directly into the encoded string.

3. **Read diagonals**: The original text was placed along diagonals. Each diagonal starts at `(0, col)` for `col = 0, 1, ..., cols-1`. From each start, walk down-right: increment both row and column by 1 until either exceeds the matrix bounds. Append each character.

4. **Trim trailing spaces**: The problem states the original text has no trailing spaces, so trim them from the result.

The key insight is that you don't need to simulate the full encoding process -- just understand the index mapping. Given a matrix position `(r, c)`, the character is at `encodedText[r * cols + c]`.

## Complexity Analysis

Time Complexity: O(n) where n is the length of encodedText. Each character is visited exactly once during the diagonal reading.

Space Complexity: O(n) for storing the result. The matrix reconstruction is done via index calculation, so no extra 2D array is needed.

## Edge Cases

- **Empty encoded text**: Return an empty string immediately.
- **Single row** (`rows = 1`): The diagonal reading is just left-to-right, so the result is the encoded text with trailing spaces trimmed.
- **Trailing spaces in reconstruction**: The original text has no trailing spaces, but intermediate diagonal characters may be spaces. Only trim from the end.
