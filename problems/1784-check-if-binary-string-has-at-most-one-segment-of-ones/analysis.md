# 1784. Check if Binary String Has at Most One Segment of Ones

[LeetCode Link](https://leetcode.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/)

Difficulty: Easy
Topics: String
Acceptance Rate: 41.5%

## Hints

### Hint 1

Think about what it means for a binary string **without leading zeros** to have more than one segment of ones. What pattern in the string would indicate that the ones are split into separate groups?

### Hint 2

If the string starts with ones and then has some zeros, the only way to get a second segment of ones is if a '1' appears after a '0'. Focus on detecting that specific transition.

### Hint 3

You don't need to count segments or track state. Since the string has no leading zeros (it always starts with '1'), the string has more than one segment of ones if and only if the substring `"01"` appears anywhere in it. A '0' followed by a '1' means ones were interrupted and restarted.

## Approach

Since the string has no leading zeros, it always starts with '1'. This simplifies the problem significantly.

A single contiguous segment of ones means the string looks like `111...000...0` — some ones followed by some zeros (or all ones). The only way to violate this is if a '0' is followed by a '1', which means the ones were split by zeros.

So the algorithm is simply: check whether the substring `"01"` exists in `s`. If it does, return `false`; otherwise, return `true`.

**Walkthrough with examples:**
- `s = "1001"`: Contains `"01"` at index 2 → return `false` (ones are split).
- `s = "110"`: Does not contain `"01"` → return `true` (ones form one contiguous block).

## Complexity Analysis

Time Complexity: O(n) — we scan the string once looking for the substring "01".
Space Complexity: O(1) — no extra space is used.

## Edge Cases

- **Single character `"1"`**: No zeros at all, so there's trivially one segment of ones. Should return `true`.
- **All ones (e.g., `"1111"`)**: No zeros means one contiguous segment. Should return `true`.
- **Ones followed by zeros (e.g., `"10"`, `"1100"`)**: Only one segment of ones at the start. Should return `true`.
- **Alternating pattern (e.g., `"101"`, `"10101"`)**: Contains `"01"`, so multiple segments. Should return `false`.
