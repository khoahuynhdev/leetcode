# 3867. Sum of GCD of Formed Pairs

[LeetCode Link](https://leetcode.com/problems/sum-of-gcd-of-formed-pairs/)

Difficulty: Medium
Topics: Array, Math, Two Pointers, Sorting, Simulation, Number Theory
Acceptance Rate: 69.6%

## Hints

### Hint 1

The problem reads like a chain of separate steps: build one array, sort it, pair
things up, then reduce. Don't try to be clever about merging the steps — this is a
simulation problem. Handle each stage faithfully and the pieces fall into place. The
only questions are: how do I build `prefixGcd` efficiently, and how do I pair the
sorted elements without an expensive nested loop?

### Hint 2

For `prefixGcd`, notice that `mxi` is just a running (prefix) maximum. You can compute
the whole array in one pass while carrying the max seen so far. For the pairing step,
"smallest unpaired with largest unpaired" after sorting is the textbook signal for the
**two-pointers** technique: one pointer at the front, one at the back, walking inward.

### Hint 3

Once `prefixGcd` is sorted, set `left = 0` and `right = n - 1`. While `left < right`,
add `gcd(prefixGcd[left], prefixGcd[right])` to the answer, then move `left++` and
`right--`. The loop condition `left < right` handles the odd-length case for free: the
single middle element is exactly where the two pointers meet, so it is never paired and
never counted. No special-casing required.

## Approach

The task is a direct simulation with two efficiency-sensitive stages.

**Stage 1 — Build `prefixGcd`.** For each index `i`, `mxi` is the maximum of
`nums[0..i]`, which is a prefix maximum. Keep a variable `mx` that we update to
`max(mx, nums[i])` as we scan left to right, and set
`prefixGcd[i] = gcd(nums[i], mx)`. This is a single `O(n)` pass (each `gcd` is
`O(log V)`), avoiding any `O(n^2)` recomputation of the max.

For example, with `nums = [3, 6, 2, 8]`:

- `i=0`: `mx=3`, `gcd(3,3)=3`
- `i=1`: `mx=6`, `gcd(6,6)=6`
- `i=2`: `mx=6`, `gcd(2,6)=2`
- `i=3`: `mx=8`, `gcd(8,8)=8`

So `prefixGcd = [3, 6, 2, 8]`.

**Stage 2 — Sort.** Sort `prefixGcd` in non-decreasing order: `[2, 3, 6, 8]`.

**Stage 3 — Pair with two pointers.** The rule "pair the smallest unpaired with the
largest unpaired, repeatedly" is exactly a two-pointer sweep from both ends of the
sorted array. Start `left = 0`, `right = n - 1`:

- `left=0, right=3`: `gcd(2, 8) = 2`
- `left=1, right=2`: `gcd(3, 6) = 3`
- `left=2, right=1`: `left < right` fails, stop.

Sum `= 2 + 3 = 5`, matching the expected output.

When `n` is odd, the two pointers eventually satisfy `left == right` on the middle
element, and the `left < right` guard stops the loop before it is used — so the middle
element is naturally ignored, as the problem requires.

Computing `gcd` uses the Euclidean algorithm. Since values can be up to `10^9`, the sum
of GCDs across up to ~`5 * 10^4` pairs can exceed a 32-bit range in the worst case
(each GCD up to `10^9`), so accumulate the answer in a 64-bit integer to be safe.

## Complexity Analysis

Time Complexity: O(n log n) — the sort dominates; building `prefixGcd` is `O(n log V)`
and the two-pointer pass is `O(n log V)`, where `V` is the maximum value.
Space Complexity: O(n) — for the `prefixGcd` array (or O(1) auxiliary beyond it,
ignoring the sort's internal usage).

## Edge Cases

- **Single element (`n == 1`):** `prefixGcd` has one element; no pair can be formed and
  the answer is `0`. The two-pointer loop never executes since `left < right` is false
  immediately.
- **Odd length:** The middle element after sorting must be skipped. The `left < right`
  condition handles this automatically — no explicit middle-index arithmetic needed.
- **All equal values:** e.g. `[5, 5, 5, 5]` gives `prefixGcd = [5, 5, 5, 5]`, and each
  pair contributes `gcd(5, 5) = 5`. A good sanity check that pairing/summing works.
- **Large values (up to `10^9`) and many pairs:** Accumulate the sum in a 64-bit
  integer to avoid overflow, since the total can exceed the 32-bit range.
- **Strictly increasing input:** `mxi == nums[i]` at every index, so
  `prefixGcd[i] = gcd(nums[i], nums[i]) = nums[i]`; a useful case to confirm the prefix
  max is tracked correctly.
