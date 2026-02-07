---ANALYSIS_START---
# 1653. Minimum Deletions to Make String Balanced

[LeetCode Link](https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/)

Difficulty: Medium
Topics: String, Dynamic Programming, Stack
Acceptance Rate: 65.8%

## Hints

### Hint 1

Think about what it means for a string to be "balanced" - all 'a's must come before all 'b's. At each position in the string, you have a choice about which characters to keep and which to delete. Can you track the state of your deletions as you scan through the string?

### Hint 2

Consider using dynamic programming where you track the minimum deletions needed to reach each position. At each position, you need to decide whether to keep or delete the current character. If you keep it, what conditions must be satisfied based on what you've seen so far?

### Hint 3

The key insight is tracking how many 'b's you've encountered so far. When you see an 'a', you have two choices: delete this 'a' (cost +1), or delete all previous 'b's to maintain balance. The minimum of these two options gives you the optimal choice at each step.

## Approach

We can solve this problem using dynamic programming with a single pass through the string.

**Core Idea**: A string is balanced when all 'a's appear before all 'b's. As we scan left to right, we need to maintain this invariant by making minimum deletions.

**Algorithm**:
1. Maintain two variables:
   - `deletions`: minimum deletions needed to make the string balanced up to current position
   - `bCount`: number of 'b's we've kept in our balanced string so far

2. For each character:
   - If it's 'b': We can keep it safely (it comes after all kept 'a's). Increment `bCount`.
   - If it's 'a': We have two choices:
     - Delete this 'a' (cost: `deletions + 1`)
     - Delete all previous 'b's to maintain balance (cost: `bCount`)
     - Take the minimum of these two options
     - If we chose to delete previous 'b's, reset `bCount` to 0

3. Return the final `deletions` count.

**Example walkthrough** with "aababbab":
- Position 0 ('a'): deletions=0, bCount=0 (keep the 'a')
- Position 1 ('a'): deletions=0, bCount=0 (keep the 'a')
- Position 2 ('b'): deletions=0, bCount=1 (keep the 'b')
- Position 3 ('a'): We've seen 1 'b'. Delete this 'a' (cost 1) or delete the 'b' (cost 1). deletions=1
- Position 4 ('b'): deletions=1, bCount=1
- Position 5 ('b'): deletions=1, bCount=2
- Position 6 ('a'): Delete this 'a' (cost 2) or delete 2 'b's (cost 2). deletions=2
- Position 7 ('b'): deletions=2, bCount=1
- Result: 2 deletions

## Complexity Analysis

Time Complexity: O(n) where n is the length of the string. We make a single pass through the string.

Space Complexity: O(1) as we only use two variables regardless of input size.

## Edge Cases

- **All 'a's**: No deletions needed, string is already balanced
- **All 'b's**: No deletions needed, string is already balanced
- **Single character**: Always balanced, return 0
- **Already balanced** (e.g., "aaabbb"): Return 0
- **Completely reversed** (e.g., "bbbaaa"): Must delete all of one type
---ANALYSIS_END---
