# 3736. Minimum Moves to Equal Array Elements III

[LeetCode Link](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-iii/)

Difficulty: Easy
Topics: Array, Math
Acceptance Rate: 86.7%

## Hints

### Hint 1

You can only **increase** values. Think about what that means for the final common value: can any element ever go down to meet the others?

### Hint 2

The final value must be at least as large as every element. Which target value makes the total amount of increasing as small as possible?

### Hint 3

The best target is the **maximum** element. Each element `x` then needs exactly `max - x` moves, so the answer is `sum(max - x)`, which equals `n * max - sum(nums)`.

## Approach

1. Because moves only increase values, no element can decrease. So the final common value `T` must satisfy `T >= max(nums)`.
2. The cost of reaching target `T` is `sum(T - x)` over all elements, which equals `n * T - sum(nums)`. This grows as `T` grows, so the smallest allowed `T`, which is `max(nums)`, gives the minimum.
3. Do one pass to find the max and the sum, then return `n * max - sum`. You could also do two passes and add `max - x` for each element. Both are fine.

Example: `nums = [2,1,3]`. max = 3, sum = 6, n = 3, so the answer is `3*3 - 6 = 3`. That matches `(3-2) + (3-1) + (3-3) = 1 + 2 + 0 = 3`.

Compare this with LeetCode 453 (Minimum Moves to Equal Array Elements). There each move increments n-1 elements, and the trick is to flip it into "decrement one element", so the target becomes the minimum. In this version you increment one element at a time, so the answer comes straight from the max.

## Complexity Analysis

Time Complexity: O(n)
Space Complexity: O(1)

## Edge Cases

- **Single element**: it is already equal to itself, so the answer is 0.
- **All elements already equal**: every `max - x` is 0, so the answer is 0.
- **Max appears multiple times**: elements equal to the max need no moves. The formula handles this without any special case.
- **Large spread** (e.g. one 100 and many 1s): the result is at most 99 * 99, which fits easily in an `int`.
