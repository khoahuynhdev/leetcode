# 0453. Minimum Moves to Equal Array Elements

[LeetCode Link](https://leetcode.com/problems/minimum-moves-to-equal-array-elements/)

Difficulty: Medium
Topics: Array, Math
Acceptance Rate: 60.6%

## Hints

### Hint 1

Simulating the moves will be far too slow, since values can be up to 10^9 apart. This is a math problem in disguise. Try to describe what a single move does to the *differences* between elements, not to the elements themselves.

### Hint 2

Incrementing `n - 1` elements by 1 changes the *relative* layout of the array in exactly the same way as some much simpler operation on a single element. What is that operation?

### Hint 3

Adding 1 to every element except one is the same, relatively speaking, as **subtracting 1 from that one element**. Only the differences matter for "all elements equal", so the question becomes: how many single-element decrements does it take to make everything equal? The cheapest target is the minimum, so the answer is `sum(nums) - n * min(nums)`.

## Approach

**Reframe the move.** A move adds 1 to `n - 1` elements. Picture it as adding 1 to *every* element and then subtracting 1 from the one you left out. Adding 1 to everything doesn't change whether the elements are equal; it just shifts the whole array up. So for our purposes a move is the same as **decrementing one element by 1**.

**Count the decrements.** Now we only want the fewest single-element decrements that make all elements equal. Decrements can only lower values, so the final common value can be at most `min(nums)`. Going any lower costs extra moves, so the target is exactly `min(nums)`. Each element `x` needs `x - min` decrements:

```
moves = Σ (nums[i] - min) = sum(nums) - n * min(nums)
```

**Example:** `nums = [1, 2, 3]`, min = 1
- (1-1) + (2-1) + (3-1) = 0 + 1 + 2 = **3** ✔

**Implementation detail:** compute `Σ (x - min)` straight from the differences rather than computing `sum` and `n * min` separately. Each difference is at most 2·10^9, and the problem promises the total fits in 32 bits, so this keeps the intermediate values small. (Go's `int` is 64-bit on LeetCode anyway, but the difference form doesn't rely on that.)

The algorithm is two passes: one to find the minimum, one to add up the differences.

This one is labelled Medium, but the code is tiny. All the difficulty is in spotting the reframe. If you didn't see it on your own, don't worry; "an operation on n-1 elements equals the opposite operation on 1 element" is a classic trick and worth remembering.

## Complexity Analysis

Time Complexity: O(n), with one pass for the minimum and one for the sum of differences.
Space Complexity: O(1), since only a few scalar variables are used.

## Edge Cases

- **Single element (`n = 1`)**: already equal, so the answer is 0. The formula gives `x - x = 0` naturally.
- **All elements equal**: every difference is 0, so the answer is 0.
- **Negative numbers**: the formula only uses differences, so negatives work the same way (e.g. `[-1, 0, 1]` → 3).
- **Large spread (values near ±10^9)**: one difference can be about 2·10^9, which overflows a 32-bit int. Use a 64-bit int for the running total.
- **Minimum not at the front**: make sure the minimum is found over the whole array before you sum the differences.
