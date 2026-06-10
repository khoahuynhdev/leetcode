# 3691. Maximum Total Subarray Value II

[LeetCode Link](https://leetcode.com/problems/maximum-total-subarray-value-ii/)

Difficulty: Hard
Topics: Array, Greedy, Segment Tree, Heap (Priority Queue)
Acceptance Rate: 28.9%

This is a genuinely hard problem — the greedy insight is short, but turning it into something
that runs within the `n <= 5·10^4` limit takes a careful counting argument. Don't be discouraged
if the path to an efficient solution isn't obvious; work through the hints in order.

## Hints

### Hint 1

Since every subarray's value `max - min` is non-negative and we want the **maximum** total over
exactly `k` *distinct* subarrays, distinctness never forces us to take a smaller value when a
larger one is available. So the task reduces to a clean one: **sum the `k` largest values** among
all `O(n^2)` subarrays. There are far too many subarrays to list them, so think about how you'd
find the top `k` of a huge implicit multiset *without materializing it*.

### Hint 2

A classic trick for "sum of the top `k` elements of an implicit set" is to **binary search on a
threshold value `T`**. If you can quickly answer "how many subarrays have value `>= T`?", then you
can find the cutoff value where the count first drops below `k`. The monotonicity you need is real:
for a fixed left endpoint, growing the subarray to the right only makes `max - min` larger, so the
count of subarrays with value `>= T` is non-increasing in `T`.

### Hint 3

Counting (and summing) subarrays by their range is easiest via the **complement**: count subarrays
whose range `max - min <= X`. That predicate is *monotone in a sliding window* — if `[l..r]` has
range `<= X`, so does any `[l'..r]` with `l' >= l`. So a two-pointer window where `left` only ever
moves right works, with two **monotonic deques** tracking the window max and min. The extra spark:
maintain not just the max/min, but the **windowed sum of subarray-maxes and subarray-mins ending at
`r`**. Each deque element owns a contiguous block of left endpoints; store that block's length so
you can update the running sums in O(1) as the window grows on the right and shrinks on the left.
That gives you both `count(range <= X)` and `sum(range <= X)` in one O(n) pass.

## Approach

**Reduction.** Because values are non-negative and subarrays are distinct, the answer is simply the
sum of the `k` largest values of `range(l, r) = max(nums[l..r]) - min(nums[l..r])` over all
subarrays.

**One O(n) primitive.** Define `f(X) -> (cnt, sum)` where `cnt` is the number of subarrays with
`range <= X` and `sum` is the total of their ranges. Compute it with a sliding window `[left, r]`:

- Two monotonic deques hold the window's maximum and minimum. Each deque element is a *segment*: an
  index `i` together with `len`, the number of left endpoints `l` for which this element is the max
  (resp. min) of `nums[l..r]`. The front segment of the max deque is the maximum of the whole
  window; the front of the min deque is the minimum.
- Maintain `sumMax = Σ_{l=left..r} max(nums[l..r])` and `sumMin` analogously. When `r` enters,
  pop the back segments it dominates (`<=` for max, `>=` for min), accumulate their lengths, and push
  `r` as a new segment of length `poppedLen + 1`, adjusting `sumMax` / `sumMin` by the segment value
  times length. When `range > X`, advance `left`: drop one unit from the front segment of each deque
  (subtract its value once, decrement its length, pop it if it hits 0).
- After settling the window for this `r`, add `r - left + 1` to `cnt` and `sumMax - sumMin` to `sum`.

A single call with a huge `X` (so the window never shrinks) yields `total = n(n+1)/2` and
`totalSum = Σ range` over *all* subarrays.

**From the primitive to the answer.** Let `geCount(T) = total - f(T-1).cnt` be the number of
subarrays with `range >= T`; it is non-increasing in `T`. Binary search for the largest `T*` with
`geCount(T*) >= k` (search range `[0, max(nums) - min(nums)]`). Then:

- `c1 = geCount(T*+1)` subarrays have `range > T*`; all of them are in the top `k`. Their total is
  `sumGE(T*+1) = totalSum - f(T*).sum`.
- Fill the remaining `k - c1` picks with subarrays of range exactly `T*`, each contributing `T*`.

So `answer = sumGE(T*+1) + T* · (k - c1)`.

**Why it works.** `geCount(T*) >= k > geCount(T*+1)` guarantees enough range-`T*` subarrays exist to
finish the selection, and every subarray strictly above the cutoff is included exactly once.

**Worked example.** `nums = [4,2,5,1], k = 3`. Subarray ranges are
`{0,0,0,0,2,3,4,3,4,4}`; sorted descending the top three are `4,4,4`. Here `total = 10`,
`geCount(4) = 3 >= 3` and `geCount(5) = 0`, so `T* = 4`, `c1 = 0`, `sumGE(5) = 0`, and the answer is
`0 + 4·(3 - 0) = 12`.

## Complexity Analysis

Time Complexity: O(n · log V), where `V = max(nums) - min(nums)`. Each `f(X)` runs in O(n) (every
index is pushed/popped from each deque at most once), and binary search does O(log V) calls plus a
couple of extra evaluations.
Space Complexity: O(n) for the two monotonic deques.

## Edge Cases

- **Single element (`n = 1`).** The only subarray has range `0`; `k` must be `1` and the answer is
  `0`. The window logic and the `T* = 0` branch both produce `0`.
- **All elements equal.** Every range is `0`, so the answer is `0` regardless of `k`. `T* = 0`,
  `c1 = 0`, and `T*·(k - c1) = 0`.
- **`T* = 0` in general.** When fewer than `k` subarrays have a positive range, we take all positive
  ones (`sumGE(1) = totalSum`) and pad with zero-range subarrays; the formula collapses to
  `totalSum`.
- **Large counts overflow `int32`.** The number of subarrays can reach ~`1.25·10^9` and the running
  sums reach ~`10^18`; use 64-bit integers throughout for counts, sums, and the answer.
- **`k` equal to the total number of subarrays.** Then `T* = 0` and the answer is the sum of *all*
  subarray ranges, i.e. `totalSum`.
