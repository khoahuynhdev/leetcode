# 3634. Minimum Removals to Balance Array

[LeetCode Link](https://leetcode.com/problems/minimum-removals-to-balance-array/)

Difficulty: Medium
Topics: Array, Sliding Window, Sorting
Acceptance Rate: 39.3%

## Hints

### Hint 1

Think about what it means for an array to be balanced - the relationship between max and min values. If you need to keep certain elements, what property must they satisfy to form a valid balanced subarray?

### Hint 2

Consider sorting the array first. After sorting, if you want to keep a contiguous range of elements, what would be the condition for that range to be balanced? This transforms the problem into finding the largest valid contiguous range.

### Hint 3

Once sorted, for any subarray from index i to j, the minimum is nums[i] and maximum is nums[j]. The condition becomes nums[j] <= k * nums[i]. Use a sliding window approach to find the longest such valid subarray, then the answer is n - (longest valid subarray length).

## Approach

1. **Sort the array**: After sorting, any valid balanced subarray must be contiguous. This is because if we have elements a ≤ b ≤ c and both [a,c] and [b] are valid, then [a,b,c] is also valid.

2. **Sliding window**: Use two pointers to maintain a window where the balance condition holds:
   - For a window from left to right, the condition is: `nums[right] <= k * nums[left]`
   - Expand the right pointer as much as possible
   - When the condition is violated, shrink from the left

3. **Track maximum window size**: The longest valid window gives us the maximum number of elements we can keep. The answer is `n - max_window_size`.

**Example walkthrough** with nums = [1,6,2,9], k = 3:
- After sorting: [1,2,6,9]
- Windows to check: [1], [1,2], [1,2,6], [1,2,6,9]
- [1,2,6]: max=6, min=1, 6 <= 3*1? Yes (6 <= 3 is false, so this breaks)
- Actually: [1,2]: 2 <= 3*1? Yes. [2,6]: 6 <= 3*2? Yes. 
- Best window is [2,6] with length 2, so remove 4-2=2 elements.

## Complexity Analysis

Time Complexity: O(n log n) - dominated by sorting
Space Complexity: O(1) - only using constant extra space (excluding input)

## Edge Cases

- **Single element array**: Always balanced, return 0
- **Already balanced array**: No removals needed, return 0  
- **All elements must be removed except one**: When no two elements can coexist
- **Large k values**: When k is very large, most subarrays might be valid
- **Duplicate elements**: Handle cases where min equals max in a range