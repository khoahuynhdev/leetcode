# 3721. Longest Balanced Subarray II

[LeetCode Link](https://leetcode.com/problems/longest-balanced-subarray-ii/)

Difficulty: Hard
Topics: Array, Hash Table, Divide and Conquer, Segment Tree, Prefix Sum
Acceptance Rate: 22.2%

## Hints

### Hint 1

Think about what changes when you extend a subarray by one element to the right. For each possible left endpoint, how does the "balance" (distinct evens minus distinct odds) change? Can you express this as a range update problem?

### Hint 2

For a fixed right endpoint, adding a new element affects the distinct count only for left endpoints where this is the first occurrence of that value. Specifically, if the previous occurrence of `nums[r]` was at index `prev`, then for left endpoints in `[prev+1, r]`, this element introduces a new distinct even or odd. This is a range update on a contiguous segment.

### Hint 3

Maintain a virtual array `balance[l]` representing `distinct_even(l, r) - distinct_odd(l, r)` for all possible left endpoints `l` as you sweep `r` from left to right. Each step is a range add of +1 (even) or -1 (odd) on a suffix-like segment. Use a segment tree with lazy propagation to support range add and "find leftmost index where value equals 0" efficiently.

## Approach

We sweep the right endpoint `r` from 0 to n-1, maintaining the balance (distinct even count minus distinct odd count) for every possible left endpoint simultaneously using a segment tree.

**Key observation:** When we process `nums[r]`:
- Let `prev` be the previous position where `nums[r]` appeared (or -1 if it hasn't).
- If `nums[r]` is even, it adds a new distinct even number to all subarrays starting at positions `l` in `[prev+1, r]`. So we do a range add of +1 on `balance[prev+1 .. r]`.
- If `nums[r]` is odd, it adds a new distinct odd number to those same subarrays. So we do a range add of -1 on `balance[prev+1 .. r]`.

After the update, we want the smallest `l` (0 <= l <= r) such that `balance[l] = 0`, because that gives the longest balanced subarray ending at `r` with length `r - l + 1`.

**Segment tree design:**
- Each node covers a range of left-endpoint indices.
- Supports range add with lazy propagation.
- Each node stores: the minimum value in its range, the count of positions achieving that minimum, and the leftmost such position.
- To find the leftmost zero: if the global minimum is 0, walk the tree to find the leftmost position with value 0. If the minimum is not 0, no balanced subarray ends at the current `r`.

This works because after initialization (all balances start at 0), the balance values change by +1/-1 through range adds. We're always looking for positions that have returned to 0.

**Tracking previous occurrences:** Use a hash map from value to its most recent index, updated as we scan.

## Complexity Analysis

Time Complexity: O(n log n) — Each step does one range update and one query on a segment tree of size n.

Space Complexity: O(n) — For the segment tree and the hash map of previous occurrences.

## Edge Cases

- **Single element**: A single-element subarray has either 1 distinct even and 0 distinct odd, or vice versa. It is never balanced (unless the problem considers 0 == 0, but the subarray has one element contributing one distinct count, so balance is +1 or -1). The answer is 0.
- **All same parity**: e.g., all even or all odd numbers. No subarray can be balanced. Answer is 0.
- **All same value**: e.g., `[2, 2, 2]`. Every subarray has 1 distinct even, 0 distinct odd. Answer is 0.
- **Duplicates that don't affect distinct counts**: e.g., `[3, 2, 2, 5, 4]`. The duplicate 2 extends the subarray length without changing the even distinct count, which can help form a longer balanced subarray.
- **Large n (10^5)**: The O(n log n) solution handles this comfortably within time limits.
