# 3551. Minimum Swaps to Sort by Digit Sum

[LeetCode Link](https://leetcode.com/problems/minimum-swaps-to-sort-by-digit-sum/)

Difficulty: Medium
Topics: Array, Hash Table, Sorting
Acceptance Rate: 55.8%

## Hints

### Hint 1

Split the problem into two parts. The first part is working out where each element *should* end up. The second part is counting the fewest swaps needed to move every element there. The first part is only a custom sort.

### Hint 2

After sorting by `(digitSum, value)`, every element has a known target index. Because all values are distinct, the target order is unique. Use a hash map from value to target index. That gives you a permutation: position `i` holds an element that belongs at position `target[i]`.

### Hint 3

Think of the permutation as a graph where `i → target[i]`. It breaks into disjoint **cycles**. A cycle of length `k` needs exactly `k - 1` swaps, since each swap puts one element in its final place and the last swap places two. So the answer is `n - (number of cycles)`.

## Approach

1. **Compute the target order.** Copy `nums` and sort the copy with this comparator: smaller digit sum comes first, and when digit sums tie, the smaller value comes first. Compute each digit sum once before sorting so the comparator doesn't recompute it.
2. **Map values to target positions.** Build `pos[value] = index in sorted array`. Values are distinct, so the mapping is one-to-one.
3. **Count cycles.** Keep a `visited` array. For each unvisited index `i`, follow `i → pos[nums[i]] → pos[nums[that]] → ...` until you return to a visited index, marking each index along the way. Each walk like this covers one cycle.
4. **Answer** = `n - cycles`. A fixed point (an element already in place) is a cycle of length 1 and adds 0 swaps.

**Why `k - 1` swaps per cycle is optimal:** a single swap can raise the number of cycles by at most one. The sorted state has `n` cycles (all fixed points), and we start with `c` cycles, so we need at least `n - c` swaps. Swapping an element straight into its correct spot always splits off one fixed point, so `n - c` swaps is also enough.

**Example:** `nums = [18, 43, 34, 16]`, with digit sums `[9, 7, 7, 7]`.
Sorted order: `[16, 34, 43, 18]`, so `pos = {16:0, 34:1, 43:2, 18:3}`.
Permutation: index 0 (18) → 3, index 3 (16) → 0, which is a cycle of length 2. Index 1 (43) → 2, index 2 (34) → 1, which is another cycle of length 2.
There are 2 cycles, so the answer is `4 - 2 = 2`.

## Complexity Analysis

Time Complexity: O(n log n). Sorting dominates. Each digit sum costs O(log₁₀ V), which is at most 10 digits. The cycle walk is O(n).
Space Complexity: O(n), for the sorted copy, the position map, and the visited array.

## Edge Cases

- **Single element:** the array is already sorted, so the answer is 0.
- **Already sorted:** every index is a fixed point, so the answer is 0 (Example 2).
- **Ties in digit sum:** order tied numbers by their actual value. If you forget the tie-breaker, the target order is wrong (Example 3).
- **Reverse order / one big cycle:** for example, the elements form a single cycle of length `n`, so the answer is `n - 1`.
- **Large values (up to 10⁹):** digit sums stay small (at most 81), but use a hash map for positions. Don't index an array by value.
