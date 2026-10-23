# 3644. Maximum K to Sort a Permutation

[LeetCode Link](https://leetcode.com/problems/maximum-k-to-sort-a-permutation/)

Difficulty: Medium
Topics: Array, Bit Manipulation
Acceptance Rate: 44.3%

## Hints

### Hint 1

Elements already at their correct position (`nums[i] == i`) never have to move. Concentrate on the elements that are out of place: every one of them must take part in at least one swap.

### Hint 2

Each swap requires `a AND b == k`. That means every element taking part in a swap must have **all bits of `k` set**. What is the largest number whose bits are a subset of the bits of every misplaced element?

### Hint 3

Take the bitwise AND of all misplaced elements, and call it `M`. Any valid `k` must be a submask of `M`, so `k <= M`. Then show that `k = M` actually works. The element whose value is exactly `M` is your key: every misplaced element contains all the bits of `M`, so ANDing any of them with `M` gives back `M`. That makes `M` a universal "pivot" you can swap through.

## Approach

1. Find all indices `i` where `nums[i] != i`. If there are none, the array is already sorted, so return `0`.
2. Compute `M = AND of nums[i]` over all those misplaced indices.
3. Return `M`.

**Why `k` cannot exceed `M` (upper bound):** every misplaced value `x` has to be swapped at least once, and a swap involving `x` needs `x AND y == k`, so every bit of `k` is set in `x`. That holds for every misplaced value, so `k` is a submask of their AND, which gives `k <= M`.

**Why `k = M` is achievable:** `nums` is a permutation of `[0, n-1]` and `M < n`, so the value `M` is somewhere in the array.
- If `M` is misplaced, it belongs to the set we ANDed over. For any misplaced value `x`, `x AND M == M`, because `M` is a submask of `x`. So the value `M` can swap with any other misplaced element. With one "pivot" that can swap with everybody, we can sort the misplaced elements, the same way you sort with a buffer slot: repeatedly swap the pivot to wherever it lets you place the correct element.
- If `M` is already at index `M`, it still works as a pivot. `M` is a submask of every misplaced value, so swapping it with any misplaced `x` is allowed. Swap it into a cycle, use it to put each element in place, and it finishes back at index `M`. Every swap involving the pivot has AND exactly `M`.

So the answer is simply the AND of all misplaced elements.

**Example:** `nums = [0,3,2,1]`. The misplaced elements are `3` (index 1) and `1` (index 3). `3 AND 1 = 1`, so the answer is `1`.

**Example:** `nums = [3,2,1,0]`. Every element is misplaced. `3 & 2 & 1 & 0 = 0`, so the answer is `0`.

## Complexity Analysis

Time Complexity: O(n), from a single pass over the array.
Space Complexity: O(1)

## Edge Cases

- **Already sorted array:** no element is misplaced. The AND over an empty set would be "all ones", but the problem says to return `0`, so handle this case explicitly.
- **Single element (`n = 1`):** always sorted, so the answer is `0`.
- **Value `0` is misplaced:** the AND collapses to `0`, so only `k = 0` works.
- **Only two elements swapped:** the answer is just the AND of those two values, as in `[0,1,3,2]` giving `3 & 2 = 2`.
- **Initializing the accumulator:** start with `-1` (all bits set) or a flag, not `0`, or else the AND is always `0`.
