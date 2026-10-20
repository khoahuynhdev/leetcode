# 3886. Sum of Sortable Integers

[LeetCode Link](https://leetcode.com/problems/sum-of-sortable-integers/)

Difficulty: Hard
Topics: Array, Math, Sorting, Enumeration
Acceptance Rate: 41.1%

## Hints

### Hint 1

`k` has to divide `n`, so there aren't many candidates. How many divisors can a number up to 10^5 have? If that number is small, you can afford an O(n) check for each one.

### Hint 2

Rotating a block moves its elements around but never changes **which** values are in it. So break the question for a fixed `k` into two parts: (a) can each block be sorted on its own by rotating it, and (b) once every block is sorted, is the whole array sorted?

### Hint 3

A sequence is a rotation of a sorted sequence exactly when, read **cyclically** (last element wraps around to the first), it has at most one "descent" `a[i] > a[i+1]`. Once every block passes that test, the array is sorted iff `max(block j) <= min(block j+1)` for every pair of neighbouring blocks.

## Approach

1. Loop over `k = 1..n` and keep only the `k` that divide `n`. A number up to 10^5 has at most 128 divisors.
2. For a fixed `k`, go through the blocks `[0..k-1], [k..2k-1], ...`. For each block:
   - Count cyclic descents: positions `i` in the block where `nums[i] > nums[next]`, with `next` being `i+1`, or the start of the block when `i` is the last index.
   - Track the block's min and max.
   - If there are more than 1 descent, the block can't be rotated into sorted order, so reject `k`.
   - If the block's min is smaller than the previous block's max, the blocks can't line up in order, so reject `k`.
3. Add every `k` that passes to the answer.

**Why the descent test works:** a sorted sequence read cyclically goes up (non-strictly) the whole way and drops at most once, at the wrap from the largest value back to the smallest. Rotating it only changes where that single drop sits. Going the other way, if a cyclic sequence has at most one drop, cutting it right after the drop gives a non-decreasing sequence. When there are zero drops, every element is equal.

**Example:** `nums = [2,1,4,3]`, `n = 4`.
- `k = 1`: the blocks `[2]`, `[1]` are each sorted, but `1 < 2`, so they're out of order. Rejected.
- `k = 2`: `[2,1]` has one cyclic descent (2>1), and `[4,3]` has one. Max of the first block is 2 and min of the second is 3, and 2 ≤ 3. Accepted.
- `k = 4`: `[2,1,4,3]` has descents 2>1, 4>3, and the wrap 3>2, so 3 descents. Rejected.
- Answer: `2`.

## Complexity Analysis

Time Complexity: O(d(n) · n), where d(n) ≤ 128 is the number of divisors of n, so in practice about 1.3·10^7 operations at most.
Space Complexity: O(1) extra space.

## Edge Cases

- **n = 1**: the only divisor is 1, and a single element is always sorted, so the answer is 1.
- **Array already sorted**: every divisor works, so the answer is the sum of all divisors of n.
- **All values equal**: there are zero descents everywhere. The check must count *strict* descents (`>`), or this case will be rejected by mistake.
- **Duplicates across block boundaries** (e.g. `[2,1,2,3]` with `k=2`): use `prevMax <= curMin`, not `<`, because equal values at the boundary are fine.
- **Don't forget the wrap-around comparison**: `[2,1,4,3]` with `k=4` has only two descents inside the block. The wrap `3 > 2` is the third one, and it's what rules this `k` out.
- **Whole array is a rotation of a sorted array** (e.g. `[4,5,6,1,2,3]`): only `k = n` works, which tests the single-block case.
