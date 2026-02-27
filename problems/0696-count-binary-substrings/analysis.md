# 0696. Count Binary Substrings

[LeetCode Link](https://leetcode.com/problems/count-binary-substrings/)

Difficulty: Easy
Topics: Two Pointers, String
Acceptance Rate: 67.0%

## Hints

### Hint 1

Think about what makes a valid substring. The 0s and 1s must be grouped — so the structure is always a block of one character followed by a block of the other. What if you focused on the boundaries between groups of consecutive characters?

### Hint 2

Try counting the length of each consecutive group of identical characters. For two adjacent groups (e.g., a block of 0s followed by a block of 1s), how many valid substrings can be formed from just those two groups?

### Hint 3

For any two adjacent groups with lengths `prev` and `curr`, the number of valid substrings spanning that boundary is `min(prev, curr)`. For example, groups "000" and "11" contribute `min(3, 2) = 2` valid substrings: "0011" is not valid (unequal counts), but "01" and "0011" — wait, that's wrong. Actually "01" and "0011" both have equal counts only for "01" (1 each). The valid ones are "01" and "0011"? No: "0011" has 2 zeros and 2 ones, and "01" has 1 each. Both are valid. So from groups of length 3 and 2, we get min(3,2) = 2 substrings. You don't even need to store all the group lengths — just track the previous group's length as you scan.

## Approach

The key observation is that every valid substring consists of a block of one character immediately followed by a block of the other character, with equal counts. For example, "0011", "01", "1100", "10" are all valid.

**Algorithm:**

1. Scan the string left to right, tracking the length of the current group of consecutive identical characters (`curr`) and the length of the previous group (`prev`).
2. Whenever the character changes (a group boundary), add `min(prev, curr)` to the result. This is because from two adjacent groups of lengths `prev` and `curr`, you can form `min(prev, curr)` valid substrings of varying sizes centered on that boundary.
3. After the transition, set `prev = curr` and reset `curr = 1` to start counting the new group.
4. After the loop ends, add `min(prev, curr)` one final time to account for the last pair of adjacent groups.

**Walkthrough with "00110011":**

- Groups: `[00, 11, 00, 11]` → lengths `[2, 2, 2, 2]`
- Between groups 1 and 2: min(2, 2) = 2 → substrings "01", "0011"
- Between groups 2 and 3: min(2, 2) = 2 → substrings "10", "1100"
- Between groups 3 and 4: min(2, 2) = 2 → substrings "01", "0011"
- Total: 6

## Complexity Analysis

Time Complexity: O(n) — single pass through the string.
Space Complexity: O(1) — only two integer variables (`prev` and `curr`) are used.

## Edge Cases

- **Single character string** (e.g., "0" or "1"): Only one group exists, no adjacent pair, so the answer is 0.
- **All same characters** (e.g., "0000"): Only one group, answer is 0.
- **Alternating characters** (e.g., "010101"): Every adjacent pair of groups has length 1, so the answer equals `len(s) - 1`.
- **Two groups only** (e.g., "000111"): Answer is min(3, 3) = 3.
