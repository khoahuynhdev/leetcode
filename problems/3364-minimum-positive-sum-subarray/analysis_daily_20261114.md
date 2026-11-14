# 3364. Minimum Positive Sum Subarray 

[LeetCode Link](https://leetcode.com/problems/minimum-positive-sum-subarray/)

Difficulty: Easy
Topics: Array, Sliding Window, Prefix Sum
Acceptance Rate: 51.7%

## Hints

### Hint 1

Look at the constraints first: `nums.length <= 100`. Checking every subarray whose length is in `[l, r]` is completely affordable. The real question is how to get each subarray's sum quickly, without adding it up from scratch every time.

### Hint 2

Use **prefix sums**. If `prefix[i]` is the sum of the first `i` elements, then the sum of any contiguous range is the difference of two prefix values, which takes O(1). (A fixed-size sliding window for each length `k` works just as well.)

### Hint 3

Loop over every end position `j` and every allowed length `k` in `[l, r]` with `k <= j`. The window sum is `prefix[j] - prefix[j-k]`. Skip it unless it is **strictly greater than 0**, and keep the smallest one you see. If you never find a positive sum, return `-1`.

## Approach

1. Build `prefix` of size `n+1`, where `prefix[0] = 0` and `prefix[i+1] = prefix[i] + nums[i]`.
2. Set `best = -1`, meaning "nothing found yet".
3. For every exclusive end index `j` from `l` to `n`, and every length `k` from `l` to `min(r, j)`:
   - `sum = prefix[j] - prefix[j-k]` is the sum of `nums[j-k .. j-1]`.
   - If `sum > 0` and (`best == -1` or `sum < best`), set `best = sum`.
4. Return `best`.

**Example:** `nums = [3, -2, 1, 4]`, `l = 2`, `r = 3` → `prefix = [0, 3, 1, 2, 6]`.
- j=2, k=2: 1 - 0 = 1 → best = 1
- j=3, k=2: 2 - 3 = -1 (skip); k=3: 2 - 0 = 2
- j=4, k=2: 6 - 1 = 5; k=3: 6 - 3 = 3

So the answer is 1.

This is correct because we check every valid subarray exactly once. Prefix sums make each check O(1), which removes the extra factor you would get from summing each window by hand.

**Why not a classic variable-size sliding window?** Because the array has negative numbers, adding an element can make the sum smaller, so the usual "grow or shrink" rule doesn't hold. A faster O(n log n) method exists: for each `j`, search an ordered set of `prefix[j-r .. j-l]` for the largest value that is strictly less than `prefix[j]`. With n ≤ 100, though, the simple enumeration is the clean choice, and it is what an interviewer expects for an Easy problem. Mention the ordered-set idea as a follow-up.

## Complexity Analysis

Time Complexity: O(n · (r − l + 1)), which is O(n²) in the worst case
Space Complexity: O(n) for the prefix array

## Edge Cases

- **No positive-sum subarray** (all negative, or every window sums to ≤ 0): you must return `-1`, not 0 and not some sentinel like `MaxInt`.
- **Zero sums**: a sum of 0 is *not* positive. Use the strict check `sum > 0`.
- **`l == r`**: only one window size is allowed, so the method reduces to a single fixed-size window scan.
- **`r == n`**: the whole array is a candidate. Check `k <= j` so the index never goes out of bounds.
- **Single element** (`n = 1`): return that element if it is positive, otherwise `-1`.
- **A longer window beats a shorter one**: with negative numbers, a longer subarray can have a smaller positive sum (e.g. `[5, -4, 3]` → `[5, -4]` = 1). Check every allowed length, not just `l`.
