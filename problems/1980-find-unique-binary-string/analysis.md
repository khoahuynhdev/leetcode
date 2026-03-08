# 1980. Find Unique Binary String

[LeetCode Link](https://leetcode.com/problems/find-unique-binary-string/)

Difficulty: Medium
Topics: Array, Hash Table, String, Backtracking
Acceptance Rate: 79.6%

## Hints

### Hint 1

Think about how many possible binary strings of length `n` exist versus how many are in the input. There are `2^n` possible strings but only `n` are given. A missing string is guaranteed to exist — but can you find one without checking all possibilities?

### Hint 2

Consider a diagonal argument: what if you constructed a string that is guaranteed to differ from each input string at a specific position? If your result differs from `nums[i]` at position `i`, it cannot equal any string in the array.

### Hint 3

Cantor's diagonalization: build the result character by character. For position `i`, look at `nums[i][i]` (the diagonal element) and flip it — use `'1'` if it's `'0'`, and `'0'` if it's `'1'`. The resulting string differs from `nums[0]` at index 0, from `nums[1]` at index 1, and so on, so it cannot match any string in the array.

## Approach

This problem has a beautifully elegant solution based on Cantor's diagonal argument, the same technique used to prove that the real numbers are uncountable.

**Algorithm:**

1. Initialize an empty result string.
2. For each index `i` from `0` to `n-1`:
   - Look at `nums[i][i]`, the character on the "diagonal."
   - Append the opposite character to the result (`'0'` becomes `'1'`, `'1'` becomes `'0'`).
3. Return the result.

**Why it works:**

The result string is guaranteed to differ from `nums[i]` at position `i` for every `i`. Since it differs from every string in at least one position, it cannot be equal to any of them.

**Example with `["111","011","001"]`:**
- `i=0`: `nums[0][0] = '1'` → flip to `'0'`
- `i=1`: `nums[1][1] = '1'` → flip to `'0'`
- `i=2`: `nums[2][2] = '1'` → flip to `'0'`
- Result: `"000"`, which indeed is not in the array.

Note: while backtracking and hash set approaches also work (iterate through all possible strings and check membership), the diagonal approach is optimal and requires no extra data structures beyond the result.

## Complexity Analysis

Time Complexity: O(n) — one pass through the array, looking at one character per string.

Space Complexity: O(n) — only the result string of length `n` is constructed.

## Edge Cases

- **n = 1**: Only one string in the input, either `"0"` or `"1"`. The diagonal approach correctly returns the opposite.
- **Diagonal characters are all the same**: e.g., all diagonal entries are `'0'`, the result is all `'1'`s. Still valid.
- **Maximum n = 16**: The diagonal approach handles this trivially since it's O(n).
