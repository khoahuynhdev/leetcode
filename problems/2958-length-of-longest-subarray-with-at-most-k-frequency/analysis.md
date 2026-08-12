# 2958. Length of Longest Subarray With at Most K Frequency

[LeetCode Link](https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency/)

Difficulty: Medium
Topics: Array, Hash Table, Sliding Window
Acceptance Rate: 57.4%

## Hints

### Hint 1

We are asked for the *longest contiguous* stretch that satisfies a constraint. Before reaching for anything clever, ask the question that unlocks most problems of this shape: **if a subarray is "good", is every subarray inside it also good?** If the answer is yes, the property is monotone, and a two-pointer window over the array becomes viable instead of checking all O(n²) subarrays.

### Hint 2

You need to know, at any moment, how many times each value appears *inside the current window*. Values go up to 10⁹, so an array indexed by value is out — but a hash map from value to count is exactly the right tool. Maintain that map incrementally as the window's right edge advances, rather than recomputing it.

### Hint 3

Here is the insight that makes the whole thing O(n): **only the element you just added can violate the constraint.** Every other count is unchanged, and the window was valid a step ago. So after incrementing the count of `nums[right]`, you only need to check that single count against `k`. If it exceeds `k`, shrink from the left — decrementing counts as you go — until it doesn't. The left pointer never moves backwards, so the total work across the entire run is linear.

## Approach

This is a textbook variable-size sliding window, and the reason it works comes down to one observation: goodness is **downward closed**. If a subarray has every value occurring at most `k` times, then any subarray contained within it also has every value occurring at most `k` times (removing elements can only lower counts). That means for each right endpoint there is a single leftmost position at which the window is still good, and — critically — that position never moves left as `right` moves right. Two pointers sweeping in one direction is enough.

The algorithm:

1. Keep `count`, a `map[int]int` holding the frequency of every value currently inside the window `[left, right]`.
2. For each `right` from `0` to `n-1`:
   - Increment `count[nums[right]]`.
   - While `count[nums[right]] > k`, remove `nums[left]` from the window: decrement `count[nums[left]]` and advance `left`. The loop condition watches only the newly-added value because that is the only count that grew, and the window was valid before this step — so it is the only possible offender.
   - The window `[left, right]` is now good; update the answer with `right - left + 1`.
3. Return the best length seen.

Note the shrink loop is guaranteed to terminate: as `left` walks forward it must eventually pass an earlier occurrence of `nums[right]`, which drops that count back to `k`. In the worst case it stops when `left == right` and the count is 1, which is `≤ k` since `k >= 1`.

Walking through Example 1, `nums = [1,2,3,1,2,3,1,2]`, `k = 2`:

| right | nums[right] | action | window | length |
|---|---|---|---|---|
| 0 | 1 | count{1:1} | [1] | 1 |
| 1 | 2 | count{1:1,2:1} | [1,2] | 2 |
| 2 | 3 | count{1:1,2:1,3:1} | [1,2,3] | 3 |
| 3 | 1 | count[1]=2, still ≤ 2 | [1,2,3,1] | 4 |
| 4 | 2 | count[2]=2, still ≤ 2 | [1,2,3,1,2] | 5 |
| 5 | 3 | count[3]=2, still ≤ 2 | [1,2,3,1,2,3] | **6** |
| 6 | 1 | count[1]=3 > 2 → shrink past index 0 | [2,3,1,2,3,1] | 6 |
| 7 | 2 | count[2]=3 > 2 → shrink past index 1 | [3,1,2,3,1,2] | 6 |

The answer is 6, matching the expected output. Notice how the last two steps illustrate the pattern nicely: the window slides rather than grows, and the answer stops improving.

The hash map is what makes value-range-independence possible. An alternative is to compress values to indices first and use a slice for counts — that trades a preprocessing sort or map build for faster inner-loop access, but the asymptotics are identical and the map version is far shorter. Not worth the complexity here unless you are micro-optimizing.

One honest note on difficulty: the code is short, but the reasoning about *why* checking only `count[nums[right]]` suffices is the part that separates people who memorize the sliding-window template from people who can adapt it. If that step felt like a leap, it is worth re-deriving rather than memorizing.

## Complexity Analysis

Time Complexity: O(n) — each index is added to the window exactly once by `right` and removed at most once by `left`, and every map operation is O(1) on average. Total map operations are bounded by 2n.

Space Complexity: O(n) — in the worst case (all values distinct and the whole array is good) the frequency map holds up to n entries. More precisely it is O(d) where d is the number of distinct values in the largest window.

## Edge Cases

- **Single-element array** (`nums = [7]`, any `k >= 1`): the answer must be 1, never 0. Because `k >= 1` per the constraints, a one-element window is always good, so the answer is never 0 — but code that initializes `best` and forgets to update it inside the loop will return 0 here.
- **All elements identical** (`nums = [5,5,5,5,5,5,5]`, `k = 4`): the answer is capped at `k`, not `n`. This is the case that exercises the shrink loop on every step past index `k-1` and catches an off-by-one where you use `>= k` instead of `> k`.
- **`k = 1`** (`nums = [1,2,1,2,1,2,1,2]`): forces the window to hold only distinct values, reducing the problem to "longest substring without repeating characters". A good sanity check that your shrink condition is not too permissive.
- **`k >= n`, or all values distinct**: the whole array is good and the shrink loop never runs. Confirms the answer can be the full length `n` and that you are not accidentally shrinking when you should not.
- **Large values up to 10⁹**: rules out indexing counts by value directly. A `map[int]int` (or coordinate compression) is required — a fixed-size array would be impossible to allocate.
- **Repeats separated by long gaps** (e.g. `[1,2,3,4,1]` with `k = 1`): the shrink loop may need to advance `left` several positions in one step. Using an `if` instead of a `for` on the shrink is a classic bug that this case exposes.
