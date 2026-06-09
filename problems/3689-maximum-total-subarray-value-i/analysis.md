# 3689. Maximum Total Subarray Value I

[LeetCode Link](https://leetcode.com/problems/maximum-total-subarray-value-i/)

Difficulty: Medium
Topics: Array, Greedy
Acceptance Rate: 67.9%

## Hints

### Hint 1

Notice the unusual freedom in this problem: subarrays may overlap **and** you can pick the exact same subarray more than once. Whenever a problem lets you reuse the same choice repeatedly, ask yourself: "What is the single best choice, and what stops me from just taking it every time?"

### Hint 2

Think about what bounds the **value** of a subarray, `max - min`. For any subarray `nums[l..r]`, its maximum can never exceed the maximum of the whole array, and its minimum can never go below the minimum of the whole array. So how large can a single subarray's value possibly get, and which subarray achieves that bound?

### Hint 3

The whole array `nums[0..n-1]` always contains the global maximum and the global minimum, so its value is exactly `max(nums) - min(nums)` — and no subarray can beat that. Since you must pick exactly `k` subarrays and you are allowed to repeat, the optimal move is to pick that single best subarray all `k` times. The answer is simply `k * (max(nums) - min(nums))`.

## Approach

This problem looks intimidating because of the "choose `k` subarrays" framing, but the constraints hide a one-line greedy result.

Step 1 — Bound a single subarray's value. For any subarray, `max(subarray) <= max(nums)` and `min(subarray) >= min(nums)`. Therefore:

```
value(subarray) = max(subarray) - min(subarray) <= max(nums) - min(nums)
```

Step 2 — Show the bound is attainable. The full array is itself a valid subarray, and it contains both the global max and the global min, so its value equals exactly `max(nums) - min(nums)`. No subarray can do better.

Step 3 — Use repetition. We must choose exactly `k` subarrays, overlaps are allowed, and duplicates are allowed. Since every chosen subarray contributes at most `max(nums) - min(nums)`, the best total is achieved by choosing the full array (or any subarray spanning both extremes) all `k` times:

```
answer = k * (max(nums) - min(nums))
```

Let's verify with the examples:

- `nums = [1,3,2], k = 2`: max = 3, min = 1, diff = 2, answer = `2 * 2 = 4`. ✅
- `nums = [4,2,5,1], k = 3`: max = 5, min = 1, diff = 4, answer = `3 * 4 = 12`. ✅

So the algorithm is: scan once to find the global maximum and minimum, then multiply their difference by `k`. The hard part is the realization, not the coding — be honest with yourself about whether you saw *why* the whole array dominates every other choice.

A subtle but important detail: `k` can be up to `10^5` and the difference up to `10^9`, so the product can reach `10^14`. That overflows 32-bit integers, so make sure the result is computed in 64-bit arithmetic.

## Complexity Analysis

Time Complexity: O(n) — a single pass to find the minimum and maximum.
Space Complexity: O(1) — only a couple of scalar variables.

## Edge Cases

- **Single element array (`n == 1`)**: max equals min, so the difference is 0 and the answer is 0 regardless of `k`. The greedy formula handles this naturally.
- **All equal elements**: every subarray has value 0, so the total is 0. Again handled automatically.
- **Large `k` and large values**: the product `k * (max - min)` can exceed the 32-bit integer range, so the computation must use 64-bit integers to avoid overflow.
- **`nums[i] = 0`**: zeros are valid; the minimum may legitimately be 0, which is fine.
