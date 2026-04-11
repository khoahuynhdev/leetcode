# 3741. Minimum Distance Between Three Equal Elements II

[LeetCode Link](https://leetcode.com/problems/minimum-distance-between-three-equal-elements-ii/)

Difficulty: Medium
Topics: Array, Hash Table
Acceptance Rate: 70.3%

## Hints

### Hint 1

Think about how you would group indices that share the same value. What data structure lets you efficiently collect and look up indices by value?

### Hint 2

Once you have the indices grouped by value, notice that the indices within each group are naturally sorted. Now simplify the distance formula: for three sorted indices i < j < k, what does `abs(i-j) + abs(j-k) + abs(k-i)` reduce to?

### Hint 3

The distance simplifies to `2 * (k - i)` — the middle index j cancels out entirely. This means you only need to minimize the span between the first and last index of any triple. For a sorted list of indices, which three should you pick to minimize that span?

## Approach

1. **Simplify the distance formula.** When i < j < k, the absolute values resolve to:
   `(j - i) + (k - j) + (k - i) = 2 * (k - i)`.
   The middle index j does not affect the distance at all — only the outermost two indices matter.

2. **Group indices by value.** Use a hash map where each key is a number from `nums` and the value is a list of indices where that number appears. Because we iterate left to right, each list is already sorted.

3. **Slide a window of size 3.** For each value whose index list has at least 3 entries, slide a window of 3 consecutive indices across the list. The distance for each window is `2 * (list[i+2] - list[i])`. Track the global minimum.

4. **Return the result.** If no value appeared 3+ times, return -1.

**Why consecutive indices?** If you skip an index in a sorted list, the span from first to last only grows. So the tightest triple always uses three consecutive positions in the sorted index list.

**Walk-through with Example 2:** `nums = [1,1,2,3,2,1,2]`
- Value 1 → indices [0, 1, 5] → span = 2*(5-0) = 10
- Value 2 → indices [2, 4, 6] → span = 2*(6-2) = 8
- Minimum = 8

## Complexity Analysis

Time Complexity: O(n) — one pass to build the map, one pass over all index lists (total entries = n).
Space Complexity: O(n) — for the hash map storing all indices.

## Edge Cases

- **Fewer than 3 elements:** The array cannot have a good triple, return -1.
- **No value appears 3+ times:** Return -1.
- **All elements identical:** The first three indices form the optimal triple with distance `2 * 2 = 4`.
- **Multiple values each with 3+ occurrences:** Must check all groups and return the global minimum.
- **Large clusters of the same value at consecutive indices:** The minimum distance is 4 (indices i, i+1, i+2 → 2*2).
