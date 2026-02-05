# 3379. Transformed Array

[LeetCode Link](https://leetcode.com/problems/transformed-array/)

Difficulty: Easy
Topics: Array, Simulation
Acceptance Rate: 62.8%

## Hints

### Hint 1

This is a simulation problem where you need to follow movement rules in a circular array. Focus on how to handle the circular nature - when you go past the end, you wrap to the beginning, and when you go before the start, you wrap to the end.

### Hint 2

For each element in the original array, you need to calculate a new position by moving a certain number of steps. The key insight is using modular arithmetic to handle the circular wrapping. Consider how the modulo operator can help you stay within array bounds.

### Hint 3

The critical insight is handling negative numbers correctly when moving left. When moving left with negative steps, you need to ensure the result is always positive after applying modulo. The formula `(i + steps%n + n) % n` handles both positive and negative movements correctly.

## Approach

The solution involves simulating the transformation for each element:

1. For each index `i` in the original array:
   - If `nums[i] > 0`: Move `nums[i]` steps to the right
   - If `nums[i] < 0`: Move `abs(nums[i])` steps to the left  
   - If `nums[i] == 0`: Keep the same value

2. To handle circular movement, use modular arithmetic:
   - For positive movement: `(i + steps) % n`
   - For negative movement: `(i - steps + n) % n` or equivalently `(i + steps%n + n) % n`

3. The result at index `i` is the value from the original array at the calculated destination index.

Let's trace through Example 1: `nums = [3,-2,1,1]`
- Index 0: Move 3 steps right from 0 → (0+3)%4 = 3 → `nums[3] = 1`
- Index 1: Move 2 steps left from 1 → (1-2+4)%4 = 3 → `nums[3] = 1` 
- Index 2: Move 1 step right from 2 → (2+1)%4 = 3 → `nums[3] = 1`
- Index 3: Move 1 step right from 3 → (3+1)%4 = 0 → `nums[0] = 3`

Result: `[1,1,1,3]`

## Complexity Analysis

Time Complexity: O(n) where n is the length of the array. We iterate through each element once.
Space Complexity: O(1) extra space (not counting the output array which is required).

## Edge Cases

- **Zero values**: When `nums[i] == 0`, simply copy the value without movement
- **Single element array**: The only valid destination is the same index
- **Large step counts**: Steps larger than array length should wrap around multiple times
- **Negative movements**: Ensure correct wrapping when moving left from early indices