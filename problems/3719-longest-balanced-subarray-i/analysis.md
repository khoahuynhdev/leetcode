# 3719. Longest Balanced Subarray I

[LeetCode Link](https://leetcode.com/problems/longest-balanced-subarray-i/)

Difficulty: Medium
Topics: Array, Hash Table, Divide and Conquer, Segment Tree, Prefix Sum
Acceptance Rate: 56.7%

## Hints

### Hint 1

Think about how you can efficiently track the count of distinct even and odd numbers as you expand a window through the array. A brute force approach checking all possible subarrays would work, but can you optimize it?

### Hint 2

For each possible subarray, you need to count distinct even and odd numbers. Hash sets are perfect for tracking distinct elements. Consider iterating through all possible starting positions, and for each start, expand to the right while maintaining two sets.

### Hint 3

The key insight is that with the constraint of n ≤ 1500, an O(n²) solution is acceptable. For each starting index, expand rightward while maintaining sets of distinct even and odd numbers. Check balance at each step and update the maximum length when counts match.

## Approach

The solution uses a nested loop approach with hash sets to track distinct elements:

1. **Outer loop**: Iterate through each possible starting position (i) in the array
2. **Inner loop**: For each start position, expand rightward (j) to explore all subarrays starting at i
3. **Track distinct elements**: Use two sets to maintain distinct even and odd numbers in the current subarray
4. **Check balance**: After adding each element, check if the number of distinct evens equals distinct odds
5. **Update maximum**: When balanced, update the maximum length if current subarray is longer

For example, with nums = [2,5,4,3]:
- Start at index 0 (value 2): Add to evens set {2}
- Expand to index 1 (value 5): evens {2}, odds {5} - counts equal (1,1), length = 2
- Expand to index 2 (value 4): evens {2,4}, odds {5} - not balanced
- Expand to index 3 (value 3): evens {2,4}, odds {5,3} - balanced (2,2), length = 4

The algorithm continues this process for all starting positions and returns the maximum balanced subarray length found.

## Complexity Analysis

Time Complexity: O(n²)
- Outer loop runs n times
- Inner loop runs up to n times for each outer iteration
- Set operations (add, size check) are O(1) on average

Space Complexity: O(n)
- Two sets that could each contain up to n/2 distinct elements in the worst case
- The sets are reset for each new starting position

## Edge Cases

1. **Single element array**: Cannot be balanced (need at least one even and one odd), should return appropriate minimum
2. **All even or all odd numbers**: No balanced subarray possible if all numbers have the same parity
3. **Duplicate numbers**: The problem asks for distinct counts, so [2,2,3] has 1 distinct even and 1 distinct odd
4. **Entire array is balanced**: The maximum length could be the entire array length
5. **No balanced subarray exists**: Return 0 or handle appropriately based on problem constraints
