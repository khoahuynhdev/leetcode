# 0868. Binary Gap

[LeetCode Link](https://leetcode.com/problems/binary-gap/)

Difficulty: Easy
Topics: Bit Manipulation
Acceptance Rate: 67.0%

## Hints

### Hint 1

Think about how you can examine each bit of the number one at a time. What operation lets you inspect the least significant bit and then move to the next one?

### Hint 2

As you iterate through the bits, track the position of the last `1` you saw. When you encounter a new `1`, you can compute the distance from the previous one. Use right-shifting and bitwise AND to walk through the bits.

### Hint 3

You only need two pieces of state: the index of the last `1` bit you encountered and the maximum gap seen so far. Shift `n` right one bit at a time, check if the current bit is `1` with `n & 1`, and update the gap when you find consecutive `1`s. Make sure to only start measuring after you've found the first `1`.

## Approach

Iterate through every bit of `n` from least significant to most significant:

1. Maintain a counter `pos` for the current bit position (starting at 0).
2. Maintain `last` to record the position of the most recently seen `1` bit. Initialize it to `-1` to indicate no `1` has been seen yet.
3. Maintain `maxGap` to track the largest distance found.
4. In each iteration:
   - Check if the current bit is `1` using `n & 1`.
   - If it is `1` and `last != -1`, compute the distance `pos - last` and update `maxGap` if this distance is larger.
   - If it is `1`, update `last = pos`.
   - Right-shift `n` by 1 (`n >>= 1`) and increment `pos`.
5. Continue until `n` becomes 0.
6. Return `maxGap`.

**Example walkthrough with n = 22 (binary `10110`):**

| Step | pos | bit | last | gap | maxGap |
|------|-----|-----|------|-----|--------|
| 1    | 0   | 0   | -1   | -   | 0      |
| 2    | 1   | 1   | -1→1 | -   | 0      |
| 3    | 2   | 1   | 1→2  | 1   | 1      |
| 4    | 3   | 0   | 2    | -   | 1      |
| 5    | 4   | 1   | 2→4  | 2   | 2      |

Result: 2

## Complexity Analysis

Time Complexity: O(log n) — we examine each bit of `n` exactly once, and there are at most floor(log2(n)) + 1 bits.

Space Complexity: O(1) — only a fixed number of integer variables are used.

## Edge Cases

- **Only one `1` bit** (e.g., n = 8 → `1000`): There are no adjacent pairs of `1`s, so the answer is 0.
- **All bits are `1`** (e.g., n = 7 → `111`): Every adjacent pair has distance 1, so the answer is 1.
- **n = 1** (binary `1`): Only a single bit is set; return 0.
- **Two `1`s far apart** (e.g., n = 5 → `101`): The gap spans across zeros; the answer is 2.
- **Large n near the constraint limit** (e.g., n = 10^9): The algorithm handles up to 30 bits efficiently.
