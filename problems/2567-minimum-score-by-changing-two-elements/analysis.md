# 2567. Minimum Score by Changing Two Elements

[LeetCode Link](https://leetcode.com/problems/minimum-score-by-changing-two-elements/)

Difficulty: Medium
Topics: Array, Greedy, Sorting
Acceptance Rate: 59.9%

## Hints

### Hint 1

The score has two parts: a *low* score and a *high* score. Ask yourself: once you are allowed to change two elements to *any* value, how small can you make the low score? Can you make it trivially zero?

### Hint 2

If you change an element to equal some other element, the low score becomes 0 for free. So the problem really reduces to minimizing the *high* score, which is just `max - min`. Sorting the array makes the max and min easy to reason about.

### Hint 3

Changing an element is effectively "removing" it from the range (set it equal to a value already inside the range, which also makes the low score 0). With two removals on a sorted array, you only ever want to remove extreme elements. There are exactly three options: remove the two smallest, remove the two largest, or remove one smallest and one largest.

## Approach

1. **Low score is free.** Whatever two elements we change, we can set them to a value that already exists in the remaining array (e.g., equal to some untouched element). That creates a duplicate, so the minimum absolute difference is 0. Also, a changed element placed inside `[min, max]` of the remaining elements never increases the high score.

2. **Reduce to range minimization.** The answer is therefore the smallest possible `max - min` after "removing" two elements from consideration.

3. **Sort and try the three extremes.** After sorting `a` (length `n`), removing anything other than extremes cannot shrink the range. The candidates are:
   - Remove the two smallest: `a[n-1] - a[2]`
   - Remove the two largest: `a[n-3] - a[0]`
   - Remove the smallest and the largest: `a[n-2] - a[1]`

   Return the minimum of these three.

**Example:** `nums = [1,4,7,8,5]` → sorted `[1,4,5,7,8]`.
- Remove two smallest: `8 - 5 = 3`
- Remove two largest: `5 - 1 = 4`
- Remove one of each: `7 - 4 = 3`

Answer: `3`.

For `n = 3`, every candidate compares an element with itself (e.g., `a[2] - a[2]`), giving 0 — correct, since we can make all three elements equal.

## Complexity Analysis

Time Complexity: O(n log n) for sorting (an O(n) variant exists by tracking the three smallest and three largest values in one pass).
Space Complexity: O(1) extra beyond the sort (O(log n) for the sort's recursion stack).

## Edge Cases

- **n = 3:** Changing two elements to match the third always yields score 0. The formula handles this naturally.
- **All elements equal:** Range is already 0; every candidate gives 0.
- **Duplicates at the extremes:** e.g., `[1,1,1,100]` — removing duplicates still works because we compare positions in the sorted array, not distinct values.
- **Large values (up to 1e9):** Differences fit comfortably in Go's `int` (64-bit), no overflow concerns.
- **Single outlier on both ends:** e.g., `[1,50,51,52,100]` — removing one from each end is optimal, which is why the mixed case must be checked.
