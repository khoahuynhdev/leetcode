# 3414. Maximum Score of Non-overlapping Intervals

[LeetCode Link](https://leetcode.com/problems/maximum-score-of-non-overlapping-intervals/)

Difficulty: Hard
Topics: Array, Binary Search, Dynamic Programming, Sorting
Acceptance Rate: 49.6%

## Hints

### Hint 1

Forget the "lexicographically smallest indices" part for a moment and ask a simpler question: what is the maximum achievable score if you may pick at most 4 non-overlapping intervals? That reduced question is the classic *weighted interval scheduling* problem with one extra dimension (how many intervals you have spent). Weighted interval scheduling always starts the same way: sort the intervals by an endpoint so that "compatible with the one I just took" becomes "somewhere in a suffix".

### Hint 2

Sort by left endpoint and let `dp[i][j]` be the best you can do using only intervals from sorted position `i` onward while picking at most `j` of them. At each position you either skip the interval, landing on `dp[i+1][j]`, or take it, which forbids everything whose left endpoint is `<= r_i`. Because the array is sorted by left endpoint, the allowed intervals form a contiguous suffix, so the jump target is one **binary search** for the first left endpoint strictly greater than `r_i`. With `j` capped at 4, the whole table is only `5n` states.

### Hint 3

The tie-breaking is what makes this Hard, and the key realization is that it can ride along inside the same DP instead of being a separate post-processing step. Store with every state not just the score but also the winning set of original indices, kept **sorted ascending** (at most 4 of them, so it is a tiny fixed-size array). Compare two candidates by score first, then lexicographically by that sorted array.

Why is that greedy-looking comparison safe? Because if you insert the *same* index `x` into two sorted arrays `S1 < S2`, the merged arrays keep that order — provided neither is a prefix of the other. And a prefix can never happen between two *tied* candidates: a prefix means `S1` is a strict subset of `S2`, and since every weight is at least 1, the superset would have a strictly larger score, contradicting the tie. So locally optimal sub-answers compose into a globally optimal one, and no extra reconstruction pass is needed.

## Approach

**Step 1 — sort and precompute jumps.** Pair every interval with its original index, then sort by left endpoint. Extract the sorted left endpoints into their own array and, for each position `i`, binary search for `next[i]` = the first position whose left endpoint is `>= r_i + 1`. Everything in `[i+1, next[i])` has a left endpoint in `[l_i, r_i]` and therefore genuinely overlaps interval `i` (touching endpoints count as overlapping), so skipping straight to `next[i]` after taking `i` loses nothing. Note `next[i] > i` always holds, because `l_i <= r_i < r_i + 1`.

**Step 2 — the DP state.** Define `dp[j][i]` for `j` in `0..4` and `i` in `0..n` as the best selection drawn from sorted positions `i..n-1` using **at most** `j` intervals. A selection is stored as a small struct: the total score, a count, and a fixed `[4]int32` of original indices kept in ascending order. Base cases are all-empty: `dp[0][i]` for every `i` (no budget) and `dp[j][n]` for every `j` (no intervals left), both the zero value.

**Step 3 — the transition.** For `j` from 1 to 4 and `i` from `n-1` down to 0:

- *skip*: `dp[j][i+1]`, unchanged.
- *take*: `dp[j-1][next[i]]` with interval `i` merged in — add `w_i` to the score and insert the original index into the sorted index array.

Keep the better of the two, where "better" means a strictly larger score, or an equal score with a lexicographically smaller sorted index array (shorter wins if one is a prefix of the other, which is the correct rule for "at most 4" even though ties never actually reach it). Because `dp[j-1]` is fully computed before `dp[j]` starts, and within a row `i` only depends on `i+1`, the rows can be filled with two simple loops.

**Step 4 — read off the answer.** `dp[4][0]` is the answer; copy its index array out as an `[]int`. Since all weights are positive and `n >= 1`, it always contains at least one index.

**Walking example 1**, `intervals = [[1,3,2],[4,5,2],[1,5,5],[6,9,3],[6,7,1],[8,9,1]]`. Sorted by left endpoint: `(1,5,w5,idx2)`, `(1,3,w2,idx0)`, `(4,5,w2,idx1)`, `(6,7,w1,idx4)`, `(6,9,w3,idx3)`, `(8,9,w1,idx5)` (ties in the left endpoint may land in either order — it does not matter, since intervals sharing a left endpoint overlap each other anyway and both orders remain reachable through the skip branch). Taking `idx2` jumps past everything ending at or before 5 and lands on the `l = 6` block, whose best single pick is `idx3` with weight 3, giving score 8 as `[2,3]`. The competing chains `{0,1,3}` and `{2,4,5}` both total 7, so no tie-break is needed and the answer is `[2,3]`.

**Walking example 2**, the optimum is `{6,1,3,5}` with weights `5+7+6+3 = 21`, which the DP emits already sorted as `[1,3,5,6]`.

The tie-break machinery does show up on inputs like `[[1,4,5],[5,10,5],[1,10,10]]`: both `{0,1}` and `{2}` score 10, and comparing the sorted arrays `[0,1]` against `[2]` at their first element picks `[0,1]`.

## Complexity Analysis

Time Complexity: O(n log n) — sorting dominates; the `n` binary searches cost `O(n log n)` and the DP is `O(4n)` states with `O(4)` work each (the merge and the comparison both touch at most 4 slots), so `O(n)` overall for the DP.

Space Complexity: O(n) — the sorted copy, the left-endpoint array, the jump table, and the `5 x (n+1)` DP table of fixed-size structs. The constant is larger than a plain score-only DP because each state carries up to 4 indices, but it is still linear.

## Edge Cases

- **Touching endpoints overlap.** `[1,5]` and `[5,10]` are *not* compatible. This is why the jump target searches for the first left endpoint `> r_i` (i.e. `>= r_i + 1`) rather than `>= r_i`. Getting this off by one wrong silently inflates scores.
- **Fewer than 4 intervals usable, or fewer than 4 available.** "At most 4" must be a real option: the `skip` branch plus the empty base cases handle a single interval (`n == 1` returns `[0]`) and inputs where everything mutually overlaps (returns the single heaviest, lexicographically smallest on ties).
- **Ties between selections of different sizes.** A set of 2 intervals and a set of 3 can score the same, and the answer array lengths then differ. The comparison has to be a genuine lexicographic comparison of sorted index arrays, not "prefer more intervals" or "prefer fewer".
- **Ties in weight everywhere.** With all weights equal, every maximum-score choice has the same size and the tie-break does all the work — e.g. five unit intervals of weight 1 each must yield `[0,1,2,3]`, the smallest four indices.
- **Order of the answer.** Indices are reported ascending, not in the order the intervals were chosen. Inserting into a sorted array during the DP keeps this invariant for free; sorting only at the very end would break the correctness of the comparisons made along the way.
- **Sorting destroys original indices.** Carry the original index through the sort, since the output is expressed in original-input terms.
- **Score magnitude.** Up to `4 * 10^9` exceeds 32-bit range. On platforms where `int` is 64-bit this is fine, but scoring in an explicit `int64` documents the intent.
- **Duplicate intervals.** Identical `[l, r, w]` triples at different indices are legitimate, and the tie-break must then pick the smaller index — which falls out of the lexicographic comparison.
