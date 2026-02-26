# 0067. Add Binary

[LeetCode Link](https://leetcode.com/problems/add-binary/)

Difficulty: Easy
Topics: Math, String, Bit Manipulation, Simulation
Acceptance Rate: 56.9%

## Hints

### Hint 1

Think about how you add two numbers by hand on paper. You start from the rightmost digit and work your way left, keeping track of a carry. The same principle applies to binary addition.

### Hint 2

Use two pointers starting at the end of each string and move them leftward simultaneously. At each step, sum the two digits plus any carry from the previous step. The current digit of the result is `sum % 2`, and the new carry is `sum / 2`.

### Hint 3

The strings can have different lengths, so your loop should continue as long as there are digits remaining in either string OR there is a carry left to process. Build the result in reverse order and reverse it at the end (or prepend to the result).

## Approach

Simulate binary addition just like manual addition on paper:

1. Initialize two pointers `i` and `j` at the last index of strings `a` and `b` respectively.
2. Initialize a `carry` variable to 0.
3. Loop while `i >= 0` or `j >= 0` or `carry > 0`:
   - Get the digit from `a[i]` (or 0 if `i` is out of bounds).
   - Get the digit from `b[j]` (or 0 if `j` is out of bounds).
   - Compute `sum = digitA + digitB + carry`.
   - Append `sum % 2` to the result.
   - Update `carry = sum / 2`.
   - Decrement `i` and `j`.
4. Since we built the result from least significant to most significant bit, reverse the result string before returning.

**Example walkthrough** with `a = "1010"` and `b = "1011"`:

| Step | digitA | digitB | carry | sum | result digit | new carry |
|------|--------|--------|-------|-----|-------------|-----------|
| 1    | 0      | 1      | 0     | 1   | 1           | 0         |
| 2    | 1      | 1      | 0     | 2   | 0           | 1         |
| 3    | 0      | 0      | 1     | 1   | 1           | 0         |
| 4    | 1      | 1      | 0     | 2   | 0           | 1         |
| 5    | -      | -      | 1     | 1   | 1           | 0         |

Result (reversed): `10101`

## Complexity Analysis

Time Complexity: O(max(m, n)), where m and n are the lengths of strings `a` and `b`. We process each digit exactly once.

Space Complexity: O(max(m, n)) for the result string. The result can be at most one digit longer than the longer input.

## Edge Cases

- **Different lengths**: The strings `a` and `b` can have very different lengths (e.g., `"1"` and `"1111111111"`). The loop must handle indices going out of bounds independently for each string.
- **Final carry**: Adding `"1"` and `"1"` produces `"10"`, which is longer than both inputs. The loop must continue when carry is non-zero even after both strings are exhausted.
- **Single character inputs**: Both strings could be `"0"` or `"1"`. The algorithm should handle these without special cases.
- **All ones**: Strings like `"1111"` + `"1111"` produce a carry cascade resulting in `"11110"`. Ensures carry propagation works correctly through the entire length.
