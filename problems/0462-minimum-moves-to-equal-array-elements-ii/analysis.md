# 0462. Minimum Moves to Equal Array Elements II

[LeetCode Link](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/)

Difficulty: Medium
Topics: Array, Math, Sorting
Acceptance Rate: 63.8%

## Hints

### Hint 1

Every element ends up at the same target value `t`, and each move changes one element by 1. So the total cost is a sum of distances: `sum(|nums[i] - t|)`. What value of `t` makes that sum as small as possible? Think about statistics.

### Hint 2

Sort the array and imagine sliding `t` along the number line. When you move `t` one step right, every element to the left of `t` costs one more move, and every element to the right costs one less. When does moving stop helping?

### Hint 3

The sum of absolute deviations is minimized at the **median**. Once sorted, you can skip computing the median explicitly: pair the smallest with the largest, the second smallest with the second largest, and so on. Each pair `(lo, hi)` costs exactly `hi - lo` for any target between them, and the median lies between every pair.

## Approach

1. **Reframe the problem.** Choosing a final value `t`, the answer is `sum(|nums[i] - t|)`. We need the `t` that minimizes this.
2. **Why the median?** Let `L` be the number of elements less than `t` and `R` the number greater. Increasing `t` by 1 changes the cost by `L - R`. The cost decreases while `R > L` and increases once `L > R`, so the minimum is at the point where the counts balance — the median. (For an even-length array, any value between the two middle elements works.)
3. **Two-pointer pairing.** Sort `nums`. Place `i` at the start and `j` at the end. For the pair `(nums[i], nums[j])`, any target in `[nums[i], nums[j]]` costs exactly `nums[j] - nums[i]` moves for those two elements. Since the median lies inside every such nested interval, add `nums[j] - nums[i]`, move both pointers inward, and repeat until they meet. A middle element (odd `n`) is the median itself and contributes 0.

Example: `nums = [1, 10, 2, 9]` → sorted `[1, 2, 9, 10]`.
- Pair `(1, 10)` → 9
- Pair `(2, 9)` → 7
- Total = 16.

An alternative is Quickselect to find the median in average O(n) time, but sorting is simpler and fast enough for `n ≤ 10^5`.

## Complexity Analysis

Time Complexity: O(n log n) for sorting (O(n) average with Quickselect).
Space Complexity: O(1) extra beyond the sort's internal stack (O(log n)).

## Edge Cases

- **Single element:** already equal, answer is 0.
- **All elements equal:** every pair contributes 0.
- **Even length:** there is no single middle element; the pairing approach handles it naturally.
- **Large negative and positive values:** differences like `10^9 - (-10^9)` reach `2 * 10^9`, which overflows 32-bit ints. Go's `int` is 64-bit on LeetCode's platform, so accumulating in `int` is safe.
- **Unsorted input:** the order of `nums` doesn't matter, but you must sort (or select) before pairing.
