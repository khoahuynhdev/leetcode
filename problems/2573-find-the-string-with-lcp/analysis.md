# 2573. Find the String with LCP

[LeetCode Link](https://leetcode.com/problems/find-the-string-with-lcp/)

Difficulty: Hard
Topics: Array, String, Dynamic Programming, Greedy, Union-Find, Matrix
Acceptance Rate: 41.0%

## Hints

### Hint 1

Think about what the LCP matrix tells you about character equality. If `lcp[i][j] > 0`, what does that say about the characters at positions `i` and `j`? This relationship can help you group positions together.

### Hint 2

Use a greedy strategy to assign characters. Start with 'a' at the first unassigned position, and propagate: any position `j` where `lcp[i][j] > 0` must share the same character as position `i`. This is essentially a Union-Find / grouping problem. After building the string greedily, you need to verify it actually produces the given LCP matrix.

### Hint 3

The verification step is the critical insight. After constructing a candidate string, recompute the LCP matrix using dynamic programming from the bottom-right corner: if `word[i] == word[j]`, then `lcp[i][j] = lcp[i+1][j+1] + 1` (with boundary handling). If the recomputed matrix doesn't match the input, no valid string exists. This DP is the same recurrence that defines LCP for suffixes.

## Approach

The algorithm has two phases: **construction** and **verification**.

**Phase 1 — Greedy Construction:**

Iterate through positions `0` to `n-1`. For each unassigned position `i`, assign the smallest available character (starting from `'a'`). Then scan all positions `j > i`: if `lcp[i][j] > 0`, position `j` must have the same character as `i` (since a positive LCP means the suffixes starting at `i` and `j` share at least their first character). If `j` was already assigned a different character, or if we exhaust all 26 letters, return `""`.

**Phase 2 — Verification:**

The greedy phase only enforces a necessary condition (character equality from `lcp > 0`). It doesn't verify the exact LCP values. We recompute the full LCP matrix from the constructed string using the DP recurrence:

- If `word[i] == word[j]`: `lcp[i][j] = lcp[i+1][j+1] + 1` (or `1` if at boundary)
- If `word[i] != word[j]`: `lcp[i][j] = 0`

Process from bottom-right to top-left. If the recomputed matrix matches the input exactly, return the string; otherwise return `""`.

**Example walkthrough with `lcp = [[4,0,2,0],[0,3,0,1],[2,0,2,0],[0,1,0,1]]`:**

- Position 0: unassigned → assign `'a'`. Check lcp[0][j]: lcp[0][1]=0, lcp[0][2]=2>0 → assign `'a'` to position 2. lcp[0][3]=0.
- Position 1: unassigned → assign `'b'`. Check lcp[1][j]: lcp[1][2]=0, lcp[1][3]=1>0 → assign `'b'` to position 3.
- Positions 2,3 already assigned. Result: `"abab"`.
- Verify: recomputed LCP matches input. Return `"abab"`.

## Complexity Analysis

Time Complexity: O(n²) — both the greedy scan and the verification DP iterate over all pairs (i, j).

Space Complexity: O(n²) — for the verification LCP matrix. The string itself is O(n).

## Edge Cases

- **Diagonal mismatch:** `lcp[i][i]` must equal `n - i` (a suffix's LCP with itself is its own length). If violated, return `""`.
- **Asymmetry:** `lcp[i][j]` must equal `lcp[j][i]`. The verification step catches this.
- **Value too large:** `lcp[i][j] > n - max(i,j)` is impossible. The verification catches this too.
- **More than 26 groups:** If the matrix implies more than 26 distinct character groups, we run out of letters. Return `""`.
- **n = 1:** Single character. `lcp` must be `[[1]]`, answer is `"a"`.
