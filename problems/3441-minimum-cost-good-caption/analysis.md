# 3441. Minimum Cost Good Caption

[LeetCode Link](https://leetcode.com/problems/minimum-cost-good-caption/)

Difficulty: Hard
Topics: String, Dynamic Programming
Acceptance Rate: 20.5%

## Hints

### Hint 1

Think about partitioning the string into contiguous segments, where each segment will be assigned a single character. What constraint must each segment satisfy for the caption to be "good"?

### Hint 2

Consider a DP where you track the current position, the character assigned to the current segment, and how many consecutive positions have used that character so far. Since we only care whether the group length has reached 3 or not, you can cap the consecutive count at 3 to keep the state space manageable.

### Hint 3

Define `dp[i][c][k]` as the minimum cost to process positions `0..i` where position `i` is assigned character `c` and the current run of `c` has length `min(k, 3)`. Transitions either extend the current run (same character, increment k) or start a new character (only allowed if k >= 3). To recover the lexicographically smallest answer among all minimum-cost solutions, reconstruct the string by greedily choosing the smallest character at each position that preserves optimality.

## Approach

1. **State definition**: Let `dp[i][c][k]` be the minimum number of operations to process `caption[0..i]` such that:
   - Position `i` is assigned character `c` (0-25 for 'a'-'z')
   - The current consecutive run of `c` has length `k` (we track k = 1, 2, or 3+)

2. **Base case**: `dp[0][c][1] = |caption[0] - c|` for each character `c`.

3. **Transitions from dp[i][c][k]**:
   - **Continue same character**: `dp[i+1][c][min(k+1, 3)]` = min of itself and `dp[i][c][k] + |caption[i+1] - c|`
   - **Start new character** (only if k >= 3): `dp[i+1][c'][1]` = min of itself and `dp[i][c][3] + |caption[i+1] - c'|`

4. **Answer**: The minimum cost is `min over all c of dp[n-1][c][3]`. If no state with k=3 is reachable at position n-1, return `""`.

5. **Lexicographic reconstruction**: After computing the DP, reconstruct the answer from left to right. At each position, choose the smallest character `c` and appropriate `k` that allows reaching the overall minimum cost. This requires also knowing the "suffix" optimal costs, which can be computed via a backward pass or by storing parent pointers and carefully breaking ties.

The approach used in the solution computes a forward DP, then reconstructs greedily by iterating characters in order ('a' to 'z') and checking if choosing that character still allows achieving the global optimum for the suffix.

## Complexity Analysis

Time Complexity: O(n * 26 * 3 * 26) = O(n * 26^2) which simplifies to O(n). The constant factor is 26 * 3 states per position with 26 possible transitions, giving roughly 2028 operations per position.

Space Complexity: O(n * 26 * 3) for the DP table, which is O(n).

## Edge Cases

- **Length < 3**: Impossible to form any group of 3, so return `""`.
- **All same characters**: Already a good caption, return as-is with 0 cost.
- **Length exactly 3**: Must assign one character to all three positions; the optimal character is the median (or smallest among tied medians for lexicographic order).
- **Large input (n = 50000)**: The DP must be efficient; O(n * 26^2) with small constants is fast enough.
