# 3041. Maximize Consecutive Elements in an Array After Modification

[LeetCode Link](https://leetcode.com/problems/maximize-consecutive-elements-in-an-array-after-modification/)

Difficulty: Hard
Topics: Array, Dynamic Programming, Sorting
Acceptance Rate: 50.2%

## Hints

### Hint 1

The original indices don't matter at all — only the multiset of values does. What does sorting buy you when you're looking for runs of consecutive values?

### Hint 2

Think of a DP indexed by **value**, not by position: "what is the longest consecutive run that ends exactly at value `v`?" Each element has only two possible final values, `x` or `x + 1`, so each element can extend a run in only two ways.

### Hint 3

Process elements in sorted order and keep `dp[v]` in a hash map. For an element `x`:

- turn it into `x + 1`: `dp[x+1] = dp[x] + 1`
- keep it as `x`: `dp[x] = dp[x-1] + 1`

The order of these two updates is the whole trick — update `x + 1` **first**, so that it reads the value of `dp[x]` from *before* this element was used.

## Approach

1. Sort `nums`. After sorting, when we handle `x`, every element that could end up as `x - 1` or `x` (i.e. the original values `x - 2`, `x - 1`, `x`) has already been processed.
2. Maintain `dp[v]` = the length of the longest consecutive run of selected values ending exactly at `v`, using only elements processed so far.
3. For each `x`:
   - **Use it as `x + 1`.** The best run ending at `x + 1` is the best run ending at `x` plus this element: `dp[x+1] = dp[x] + 1`. We must read the *old* `dp[x]` because this element can't be both `x` and `x + 1`.
   - **Use it as `x`.** `dp[x] = dp[x-1] + 1`.
   - Track the maximum over all `dp` values written.

Why is assignment (not `max`) safe? Because values are processed in non-decreasing order, `dp[x]` only grows as more copies of smaller-or-equal values arrive, so the newest computation is always at least as large as the previous one.

**Example:** `nums = [2,1,5,1,1]` → sorted `[1,1,1,2,5]`

| x | dp[x+1] = dp[x]+1 | dp[x] = dp[x-1]+1 |
|---|-------------------|-------------------|
| 1 | dp[2] = 0+1 = 1   | dp[1] = 0+1 = 1   |
| 1 | dp[2] = 1+1 = 2   | dp[1] = 1         |
| 1 | dp[2] = 1+1 = 2   | dp[1] = 1         |
| 2 | dp[3] = 2+1 = 3   | dp[2] = 1+1 = 2   |
| 5 | dp[6] = 1         | dp[5] = 1         |

Maximum is `dp[3] = 3` (values 1, 2, 3 from elements 1, 1→2, 2→3).

Don't worry if this took a while to see — the "DP over values, update in the right order" idea is a genuinely tricky pattern, and it's what makes this a Hard.

## Complexity Analysis

Time Complexity: O(n log n) — dominated by sorting; the DP pass is O(n) with hash map operations.
Space Complexity: O(n) — the sorted copy plus at most 2n map entries.

## Edge Cases

- **Single element:** the answer is always at least 1.
- **Many duplicates:** `[3,3,3,3]` can only fill values 3 and 4, so the answer is 2 — the update order prevents counting the same element for both slots.
- **Gap of exactly one:** `[1,1,3]` → 1, 2 (from 1), 3 — an increment can bridge a gap of one.
- **Gap of two or more:** `[1,4]` cannot be bridged; the answer is 1.
- **Shifting a whole block:** `[1,3,3,4]` → 1→2, 3, 3→4, 4→5 gives 2,3,4,5 = 4; most elements are bumped, while one copy of 3 stays put.
- **Large values (up to 10^6):** using a map (or an array of size max+2) keeps `x + 1` in range without overflow concerns.
