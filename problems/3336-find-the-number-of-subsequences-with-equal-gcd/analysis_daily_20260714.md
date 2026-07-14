# 3336. Find the Number of Subsequences With Equal GCD

[LeetCode Link](https://leetcode.com/problems/find-the-number-of-subsequences-with-equal-gcd/)

Difficulty: Hard
Topics: Array, Math, Dynamic Programming, Number Theory
Acceptance Rate: 44.8%

This one is genuinely tricky — it combines counting over exponentially many
subsequence pairs with GCD tracking. Don't be discouraged if the brute-force
"try every split" idea feels hopeless; that feeling is exactly the signal that
pushes you toward the DP formulation below. Take it hint by hint.

## Hints

### Hint 1

There are up to `2^n` subsequences, and you are pairing them, so enumerating
pairs directly is impossible for `n` up to 200. Whenever you must *count*
configurations over an exponential space, think dynamic programming where the
**state captures just enough summary information** to make the count. Ask
yourself: what is the minimal summary of a chosen subsequence that this problem
actually cares about?

### Hint 2

The only property of a subsequence that matters is its GCD. And there is a
second helpful fact: the values are bounded by `1 <= nums[i] <= 200`, so any
GCD is an integer in `[1, 200]`. That small range is your invitation to make the
GCD part of the DP state. Consider a state that tracks the current GCD of
`seq1` and the current GCD of `seq2` simultaneously as you scan the array.

### Hint 3

Process elements one at a time and let `dp[a][b]` be the number of ways to
assign the elements seen so far to three buckets — put it in `seq1`, put it in
`seq2`, or leave it out — so that `seq1` currently has GCD `a` and `seq2` has
GCD `b`. Use `0` as a sentinel meaning "still empty", which is perfect because
`gcd(0, x) = x`. Each new element branches into exactly three transitions. At
the end, sum `dp[g][g]` over all `g >= 1` (both non-empty, GCDs equal). The
disjointness constraint is handled for free: each index lands in at most one of
the two subsequences because a single element can only take one of the three
choices.

## Approach

We build the answer incrementally with a 2-D DP over GCD values.

**State.** Let `dp[a][b]` = the number of ways to have processed some prefix of
`nums`, assigning each processed element to `seq1`, `seq2`, or "unused", such
that the running GCD of `seq1` is `a` and the running GCD of `seq2` is `b`.
Both `a` and `b` range over `0..maxV` where `maxV = max(nums)`, and the value
`0` is a sentinel meaning "this subsequence is still empty".

**Why `0` works so well.** The standard Euclidean GCD satisfies `gcd(0, x) = x`.
So if a subsequence is empty (GCD `0`) and we add the first element `x`, the new
GCD becomes `gcd(0, x) = x` automatically — no special casing needed for the
first element.

**Initialization.** `dp[0][0] = 1`: before processing anything, both
subsequences are empty, and there is exactly one such (empty) assignment.

**Transition.** For each element `num`, we compute a fresh table `ndp` from `dp`.
Every existing state `dp[a][b]` fans out into three choices:

- *Skip* `num`: it contributes to `ndp[a][b]` unchanged. (Implemented by
  copying `dp` into `ndp` up front.)
- *Put `num` in `seq1`*: the GCD of `seq1` becomes `gcd(a, num)`, contributing
  to `ndp[gcd(a, num)][b]`.
- *Put `num` in `seq2`*: the GCD of `seq2` becomes `gcd(b, num)`, contributing
  to `ndp[a][gcd(b, num)]`.

All additions are taken modulo `1e9 + 7`.

**Answer.** After all elements are processed, every valid pair `(seq1, seq2)`
with both non-empty and equal GCD `g` is counted in `dp[g][g]` for some
`g >= 1`. We sum `dp[g][g]` over `g` from `1` to `maxV`. States with `a = 0` or
`b = 0` are excluded because one of the subsequences would be empty, which the
problem forbids.

**Why disjointness is automatic.** Each element chooses exactly one of the three
buckets, so no index can appear in both `seq1` and `seq2`. The "unused" bucket
is what lets us range over *all* subsequence pairs rather than partitions of the
whole array.

**Worked micro-example.** For `nums = [7, 7]`, after processing both elements
`dp[7][7]` counts: put element 0 in `seq1` and element 1 in `seq2`, or the
reverse — that is `2` ordered pairs, matching the fact that `(seq1, seq2)` is an
ordered pair (the problem counts both directions, as its own examples show).

## Complexity Analysis

Let `n = len(nums)` and `M = max(nums)` (at most 200).

Time Complexity: O(n · M²) — for each of the `n` elements we sweep the full
`(M+1) × (M+1)` table, doing O(1) work (two GCD computations plus modular adds)
per cell. The GCD itself is O(log M), so a tight bound is O(n · M² · log M),
which for the given constraints is comfortably fast.

Space Complexity: O(M²) — two `(M+1) × (M+1)` tables (`dp` and `ndp`); we keep
only the current and next layers, not one per element.

## Edge Cases

- **Single element** (`n = 1`): impossible to form two disjoint non-empty
  subsequences, so the answer is `0`. The DP handles this because no state with
  both `a >= 1` and `b >= 1` can ever be reached.
- **All elements equal** (e.g. `[1,1,1,1]`): many valid pairs; a good stress
  test that the counting (and the modulo) is right. The expected answer is `50`.
- **Coprime pair** (e.g. `[2, 3]`): each single-element GCD differs, and there
  are not enough elements to build a common GCD, so the answer is `0`.
- **Ordered pairs**: `(seq1, seq2)` and `(seq2, seq1)` are counted separately —
  the DP naturally produces both because the two "put in seq1" / "put in seq2"
  transitions are symmetric but distinct.
- **Modulo overflow**: the count can be astronomically large, so every addition
  must be reduced modulo `1e9 + 7`. Forgetting this on any transition silently
  corrupts the result once counts exceed the modulus.
- **Empty-subsequence sentinel**: excluding `g = 0` from the final sum is
  essential — states where either subsequence is empty must not be counted.
