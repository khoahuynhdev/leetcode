# 3105. Longest Strictly Increasing or Strictly Decreasing Subarray

[LeetCode Link](https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/)

Difficulty: Easy
Topics: Array
Acceptance Rate: 66.1%

## Hints

### Hint 1

A subarray is contiguous. That means whether a run continues depends only on how each element compares to the one right before it. Can you solve it in one pass, from left to right?

### Hint 2

Keep track of the length of the current "run" as you scan. You need two runs: one for strictly increasing and one for strictly decreasing. Each one either grows or resets depending on how `nums[i]` compares to `nums[i-1]`.

### Hint 3

At each index `i`:
- If `nums[i] > nums[i-1]`: the increasing run grows by 1, and the decreasing run goes back to 1.
- If `nums[i] < nums[i-1]`: the decreasing run grows by 1, and the increasing run goes back to 1.
- If they are equal: both runs go back to 1 (equal values break *strict* monotonicity).

The answer is the largest run length you see during the scan.

## Approach

This is a classic **run-length scan**. A strictly monotonic subarray that ends at index `i` can only extend a strictly monotonic subarray that ends at `i-1`, so we only need the run lengths that end at the previous index.

1. Set `inc = 1`, `dec = 1`, `best = 1`. A single element is always both strictly increasing and strictly decreasing.
2. For each `i` from 1 to `n-1`:
   - `nums[i] > nums[i-1]` → `inc++`, `dec = 1`
   - `nums[i] < nums[i-1]` → `dec++`, `inc = 1`
   - equal → `inc = 1`, `dec = 1`
   - `best = max(best, inc, dec)`
3. Return `best`.

**Example:** `nums = [1,4,3,3,2]`

| i | pair  | inc | dec | best |
|---|-------|-----|-----|------|
| 0 | -     | 1   | 1   | 1    |
| 1 | 1→4 ↑ | 2   | 1   | 2    |
| 2 | 4→3 ↓ | 1   | 2   | 2    |
| 3 | 3→3 = | 1   | 1   | 2    |
| 4 | 3→2 ↓ | 1   | 2   | 2    |

Answer: `2`.

This works because every maximal strictly monotonic subarray ends somewhere. When we reach its last index, the matching counter holds exactly its length.

Since `n <= 50`, an O(n²) brute force would also pass. The one-pass scan is still the cleaner solution, and the same idea shows up in harder problems such as "Longest Turbulent Subarray".

## Complexity Analysis

Time Complexity: O(n), one pass over the array.
Space Complexity: O(1), only a few counters.

## Edge Cases

- **Single element** (`[5]`): the answer is 1. Starting the counters at 1 handles this without a special case.
- **All equal** (`[3,3,3,3]`): equal neighbors reset both runs, so the answer stays 1. Using `>=` or `<=` by mistake would give the wrong answer here.
- **Whole array monotonic** (`[3,2,1]` or `[1,2,3,4]`): the best run ends at the last index, so update `best` inside the loop, not only when a run resets.
- **Direction changes** (`[1,3,2,4]`): when the direction flips, the run in the other direction must restart at 2 (the previous element plus the current one). Doing `dec = 1` and then `dec++` on the following step handles this naturally.
