# 3727. Maximum Alternating Sum of Squares

[LeetCode Link](https://leetcode.com/problems/maximum-alternating-sum-of-squares/)

Difficulty: Medium
Topics: Array, Greedy, Sorting

Acceptance Rate: 70.9%

## Hints

### Hint 1

You can rearrange the elements however you like. So the order of the input doesn't matter. What matters is which elements end up in the positions that get added and which end up in the positions that get subtracted.

### Hint 2

Since every element gets squared, its sign doesn't matter: `-3` and `3` contribute the same amount. Think about the values `x²` (or `|x|`) instead of `x`. Then count how many positions are added (`+`) and how many are subtracted (`-`) for an array of length `n`.

### Hint 3

There are `ceil(n/2)` added positions (even indices) and `floor(n/2)` subtracted positions (odd indices). To get the largest score, put the `ceil(n/2)` **largest** squares in the added positions and the rest in the subtracted positions. Sort by absolute value and split the sorted list.

## Approach

1. **Ignore signs.** The score only uses `arr[i]²`, so replace every value by its absolute value. Ordering by `|x|` is the same as ordering by `x²`.
2. **Count the slots.** For length `n`, the indices `0, 2, 4, ...` are added and there are `ceil(n/2) = (n+1)/2` of them. The indices `1, 3, 5, ...` are subtracted and there are `n/2` of them.
3. **Greedy split.** The score is `(sum of squares in + slots) - (sum of squares in - slots)`. The total of all squares is fixed, so we want the `+` group to be as large as possible. With a fixed group size of `(n+1)/2`, the largest possible sum comes from taking the `(n+1)/2` largest squares. An exchange argument shows why: if a smaller square sits in a `+` slot while a larger one sits in a `-` slot, swapping them increases the score by `2·(big − small) ≥ 0`.
4. **Compute.** Sort the absolute values in ascending order. The first `n/2` go in the `-` group and the remaining `(n+1)/2` go in the `+` group. Add up their squares with the right signs.

**Example:** `nums = [1,-1,2,-2,3,-3]`
- Absolute values sorted: `[1,1,2,2,3,3]`, `n = 6`, so 3 `-` slots and 3 `+` slots.
- `-` group: `1,1,2` gives `1+1+4 = 6`.
- `+` group: `2,3,3` gives `4+9+9 = 22`.
- Score = `22 − 6 = 16`.

**Overflow note:** each square is at most `(4·10⁴)² = 1.6·10⁹`. With up to `10⁵` elements the sum can reach `1.6·10¹⁴`, so use 64-bit integers (`int64`).

## Complexity Analysis

Time Complexity: O(n log n) for the sort. A counting sort over `|x| ≤ 4·10⁴` would make it O(n + V), but sorting is simpler and fast enough.
Space Complexity: O(n) for the copy of absolute values. You could sort in place instead to get O(1) extra space, but that changes the input.

## Edge Cases

- **Single element:** there is one `+` slot and no `-` slots, so the answer is `x²`. This is also the case when `x` is negative.
- **Mixed signs with equal magnitudes:** `-3` and `3` are the same once squared. If you sort by raw value instead of absolute value, you get the wrong answer.
- **Even vs. odd length:** odd length gives one extra `+` slot. Using `(n+1)/2` handles both cases.
- **All equal magnitudes with even length:** every `+` cancels a `-`, so the answer is `0`.
- **Large values:** sums can go past 32-bit range, so accumulate in `int64`.
