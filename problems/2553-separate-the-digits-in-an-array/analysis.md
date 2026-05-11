# 2553. Separate the Digits in an Array

[LeetCode Link](https://leetcode.com/problems/separate-the-digits-in-an-array/)

Difficulty: Easy
Topics: Array, Simulation
Acceptance Rate: 82.5%

## Hints

### Hint 1

This is a pure simulation problem. For each number you visit, you need to emit its digits one by one into the output array — preserving both the order of the numbers and the order of digits within each number. Think about the most natural way to walk through the digits of an integer.

### Hint 2

There are two common ways to extract digits from an integer:
1. Convert the integer to a string and iterate over each character.
2. Use repeated division and modulo by 10. Note that this gives the digits in **reverse** order (least significant first), so you'll need a buffer that you flip before appending.

Either approach is acceptable. The math approach avoids string conversion overhead but needs care to preserve digit order.

### Hint 3

If you go with the math approach, the cleanest pattern is: in an inner loop, repeatedly do `digit = n % 10; n /= 10` and push to a temporary slice. Once `n == 0`, walk that slice **backwards** to copy digits into the answer in left-to-right order. Since constraints guarantee `nums[i] >= 1`, you never have to handle a literal `0` integer that would skip the loop entirely.

## Approach

For each integer `n` in `nums`, we need to push its digits onto the answer in their natural left-to-right order.

Using the math approach:

1. Initialize `result` as an empty slice.
2. For each `n` in `nums`:
   - Create a small temporary slice `digits`.
   - While `n > 0`, append `n % 10` to `digits` and divide `n` by 10. This collects digits in reverse order.
   - Iterate `digits` from the last index to the first, appending each to `result`.
3. Return `result`.

Walkthrough on `nums = [13, 25, 83, 77]`:
- `13`: digits collected as `[3, 1]`, reversed into result → `[1, 3]`.
- `25`: digits collected as `[5, 2]`, reversed into result → `[1, 3, 2, 5]`.
- `83`: digits collected as `[3, 8]`, reversed into result → `[1, 3, 2, 5, 8, 3]`.
- `77`: digits collected as `[7, 7]`, reversed into result → `[1, 3, 2, 5, 8, 3, 7, 7]`.

We can also pre-allocate the result with a reasonable capacity (e.g., `len(nums) * 6`, since `nums[i] <= 10^5` has at most 6 digits) to avoid repeated slice growth.

## Complexity Analysis

Time Complexity: O(N · D) where N is `len(nums)` and D is the maximum number of digits per element (at most 6 given the constraints). Effectively linear in the total number of output digits.
Space Complexity: O(N · D) for the output array. Auxiliary space per integer is O(D) for the temporary digit buffer.

## Edge Cases

- **Single-digit values**: Numbers like `7` or `9` collapse to a single push; the inner loop runs once.
- **Largest allowed value (10^5 = 100000)**: 6 digits including a leading `1` and trailing zeros — confirms the math loop handles internal zeros correctly (since we stop when the whole number is 0, not when a single digit is 0).
- **All identical numbers (e.g., [77, 77, 77])**: Confirms ordering is preserved across repeated entries.
- **Single-element input**: Smallest possible input — make sure you don't assume `len(nums) > 1`.
- **Mix of 1-digit and multi-digit numbers**: Output length is not simply `len(nums)`; you must build it dynamically.
