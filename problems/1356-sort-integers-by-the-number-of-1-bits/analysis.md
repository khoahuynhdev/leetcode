# 1356. Sort Integers by The Number of 1 Bits

[LeetCode Link](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/)

Difficulty: Easy
Topics: Array, Bit Manipulation, Sorting, Counting
Acceptance Rate: 79.5%

## Hints

### Hint 1

This problem combines two fundamental concepts: **sorting with a custom comparator** and **counting bits**. Think about how you can define an ordering that considers two properties of each number at once.

### Hint 2

You can use a custom sort where the primary key is the **popcount** (number of set bits) and the secondary key is the **value itself**. Go's `sort.Slice` lets you define exactly this kind of comparator. For counting bits, consider `math/bits.OnesCount`.

### Hint 3

The comparator logic is straightforward: for two elements `a` and `b`, first compare `OnesCount(a)` vs `OnesCount(b)`. If they're equal, break the tie by comparing `a` vs `b` directly. That's the entire algorithm — the rest is just plugging it into a sort call.

## Approach

1. **Custom sort with two-level comparison**: We sort the input array in-place using `sort.Slice` with a comparator function.
2. **Primary key — bit count**: For each pair of elements, count the number of `1` bits in their binary representations using `bits.OnesCount`.
3. **Secondary key — value**: When two numbers have the same bit count, the smaller number comes first.

Walking through Example 1 with `arr = [0,1,2,3,4,5,6,7,8]`:
- `0` → 0 bits, `1,2,4,8` → 1 bit each, `3,5,6` → 2 bits each, `7` → 3 bits
- Group by bit count, sort within each group by value: `[0, 1,2,4,8, 3,5,6, 7]`

This is essentially a **stable sort by popcount** with value as tiebreaker, which the comparator handles naturally.

## Complexity Analysis

Time Complexity: O(n log n) — dominated by the sort. Each comparison is O(1) since `OnesCount` runs in constant time for bounded integers.

Space Complexity: O(log n) — for the sort's internal stack. No auxiliary data structures are allocated.

## Edge Cases

- **Single element**: Array of length 1 is already sorted; the sort is a no-op.
- **All same bit count** (e.g., all powers of 2): Falls back to sorting by value, producing ascending order.
- **Zero in the array**: `0` has 0 set bits, so it always appears first.
- **All identical elements**: Already sorted; comparator returns `false` for both directions, preserving order.
- **Maximum constraint values**: `arr[i]` up to 10^4 (14 bits) — `OnesCount` handles this trivially.
