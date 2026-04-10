# 3740. Minimum Distance Between Three Equal Elements I

[LeetCode Link](https://leetcode.com/problems/minimum-distance-between-three-equal-elements-i/)

Difficulty: Easy
Topics: Array, Hash Table
Acceptance Rate: 64.7%

## Hints

### Hint 1

Think about what data structure lets you efficiently track where each value appears in the array. You need to find three indices with the same value — how would you collect those indices?

### Hint 2

Use a hash map to group indices by their value. For each value that appears at least three times, you need to pick three indices that minimize the distance. Think about what the distance formula simplifies to when you know the ordering of i, j, k.

### Hint 3

The critical insight: for ordered indices i < j < k, the distance formula `abs(i-j) + abs(j-k) + abs(k-i)` simplifies to `2 * (k - i)`. The middle index j doesn't affect the distance at all! So you just need to minimize the span between the first and third element of any triple of equal values. Checking consecutive windows of size 3 among sorted indices is sufficient.

## Approach

1. **Group indices by value**: Iterate through `nums` and build a hash map where each key is a value from the array and each entry is the list of indices where that value appears (naturally in sorted order since we scan left to right).

2. **Simplify the distance**: For any triple (i, j, k) with i < j < k, the distance is:
   - `abs(i-j) + abs(j-k) + abs(k-i) = (j-i) + (k-j) + (k-i) = 2*(k-i)`
   - The middle element cancels out, so only the gap between the first and last index matters.

3. **Find the minimum**: For each value with at least 3 occurrences, slide a window of size 3 across its sorted index list. Compute `2 * (indices[i+2] - indices[i])` for each window and track the global minimum.

4. If no value has 3+ occurrences, return -1.

## Complexity Analysis

Time Complexity: O(n) — one pass to build the index map, then iterating over all indices across all groups is at most n total.

Space Complexity: O(n) — for the hash map storing indices.

## Edge Cases

- **Fewer than 3 elements**: The array is too short to form any good tuple — return -1.
- **No value appears 3 times**: Multiple distinct values but none with 3+ occurrences — return -1.
- **All elements identical**: Every triple is valid; the answer comes from the first three indices, giving `2 * 2 = 4` for a length-3 array.
- **Multiple values with 3+ occurrences**: Must check all of them and take the global minimum.
