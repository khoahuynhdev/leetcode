# 1732. Find the Highest Altitude

[LeetCode Link](https://leetcode.com/problems/find-the-highest-altitude/)

Difficulty: Easy
Topics: Array, Prefix Sum
Acceptance Rate: 84.1%

## Hints

### Hint 1

The array you are given is not the altitudes themselves — it is the *change* in altitude between consecutive points. Think about what data structure pattern lets you reconstruct cumulative values from a list of deltas. The topic tag "Prefix Sum" is a strong nudge.

### Hint 2

You start at altitude `0`. Each point's altitude is the running total of all the gains seen so far. So instead of building the whole altitude array, ask yourself: do you actually need to store every altitude, or can you track just two things as you sweep left to right?

### Hint 3

Maintain a single running sum that you update with each `gain[i]`, and a separate variable holding the maximum altitude seen. Initialize the maximum to `0` (the starting point itself is a valid altitude). After adding each gain, compare the running sum against the maximum. One pass, no extra array needed.

## Approach

The altitudes form a prefix sum of the `gain` array, anchored at a starting altitude of `0`.

- The biker begins at point `0` with altitude `0`.
- The altitude at point `i + 1` equals `gain[0] + gain[1] + ... + gain[i]`.

So the sequence of altitudes is exactly the running (prefix) sum of `gain`, with `0` prepended for the start.

We want the highest altitude among all `n + 1` points. The naive approach builds the full altitude array and scans it, but we can fold both steps into one pass:

1. Keep a `current` running sum, initialized to `0` (the starting altitude).
2. Keep a `highest` value, also initialized to `0`, because the starting point is a candidate answer.
3. For each `g` in `gain`, add `g` to `current`, then update `highest = max(highest, current)`.
4. Return `highest`.

**Walkthrough with `gain = [-5, 1, 5, 0, -7]`:**

| step | g  | current | highest |
|------|----|---------|---------|
| start| —  | 0       | 0       |
| 1    | -5 | -5      | 0       |
| 2    | 1  | -4      | 0       |
| 3    | 5  | 1       | 1       |
| 4    | 0  | 1       | 1       |
| 5    | -7 | -6      | 1       |

The reconstructed altitudes are `[0, -5, -4, 1, 1, -6]`, and the highest is `1`. ✅

This is the prefix-sum pattern in its purest form: you never need to materialize the full prefix array when you only care about an aggregate (here, the maximum) over it.

## Complexity Analysis

Time Complexity: O(n) — a single pass over the `gain` array.
Space Complexity: O(1) — only two scalar variables are kept, regardless of input size.

## Edge Cases

- **Starting point is the highest:** When every gain is negative (e.g. `[-4, -3, -2, -1, 4, 3, 2]` from Example 2), the altitude never rises above the starting `0`. Initializing `highest` to `0` ensures we return `0` rather than a negative running sum.
- **Single-element input (`n == 1`):** The constraints guarantee at least one gain, so the array is never empty, but a length-1 array still has two points (start and one move). The single-pass loop handles this naturally.
- **All positive gains:** The altitude is strictly increasing, so the answer is the final running sum — the loop captures it on the last iteration.
- **Mixed gains with a peak in the middle:** The maximum may occur at an interior point and then drop (as in Example 1). Tracking the running maximum, not just the final sum, is essential.
