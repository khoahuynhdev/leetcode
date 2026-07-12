# 1331. Rank Transform of an Array

[LeetCode Link](https://leetcode.com/problems/rank-transform-of-an-array/)

Difficulty: Easy
Topics: Array, Hash Table, Sorting
Acceptance Rate: 71.4%

## Hints

### Hint 1

The rank of an element only depends on how many *distinct* values are smaller than it. Think about what happens if you look at the values in sorted order — the smallest distinct value gets rank 1, the next distinct value gets rank 2, and so on. Which two classic tools let you both order the data and remember a mapping?

### Hint 2

You need to preserve the original positions while assigning ranks based on sorted order. A common trick: work on a *copy* of the array, sort the copy, then build a lookup (hash map) from each value to its rank. Finally, walk the original array and replace each element via the lookup.

### Hint 3

Equal elements must share a rank, and ranks must be "as small as possible" (no gaps). So when you scan the sorted copy, only assign a new rank when you encounter a value you haven't seen before. Storing `value -> rank` in a map means duplicates naturally map to the same rank, and the rank counter increments exactly once per distinct value.

## Approach

The rank of an element is `1 + (number of distinct values strictly smaller than it)`. This gives us a clean, deterministic way to compute every rank.

Step by step:

1. **Copy and sort.** Make a copy of `arr` and sort the copy in ascending order. We sort a copy so the original indices are preserved for the final pass.
2. **Build the rank map.** Iterate over the sorted copy. Maintain a running `rank` starting at 1. For each value, if it is not already a key in the map, insert `value -> rank` and then increment `rank`. Because the array is sorted, all occurrences of a value are contiguous, so the first time we see a value it receives the smallest possible rank, and subsequent duplicates are skipped (they'd map to the same value anyway).
3. **Rebuild the answer.** Create a result slice the same length as `arr`. For each original element, look up its rank in the map and store it at the same index.

Example with `arr = [37,12,28,9,100,56,80,5,12]`:

- Sorted copy: `[5,9,12,12,28,37,56,80,100]`
- Rank map (assigned on first sight): `5->1, 9->2, 12->3, 28->4, 37->5, 56->6, 80->7, 100->8`
- Map back over the original: `[5,3,4,2,8,6,7,1,3]` ✅

The map handles duplicates for free (both `12`s map to rank 3), and sorting guarantees the "as small as possible" requirement since ranks are assigned in increasing value order with no gaps.

## Complexity Analysis

Time Complexity: O(n log n) — dominated by sorting the copy. Building the map and rebuilding the result are each O(n).
Space Complexity: O(n) — the sorted copy, the value→rank map, and the result slice each use O(n) space.

## Edge Cases

- **Empty array (`arr.length == 0`).** The constraints allow a length-0 array. The code should return an empty (non-nil is nice but not required) slice without indexing errors. Sorting an empty copy and looping zero times handles this naturally.
- **Single element.** Always maps to rank 1; a good sanity check that the counter starts at 1.
- **All equal elements (e.g. `[100,100,100]`).** Every element must share the same rank `1`. The map ensures the value is inserted once, so all positions resolve to rank 1.
- **Negative values and full int range (`-10^9 .. 10^9`).** Ranks are based on relative ordering, not magnitude, so negatives work the same as positives — no special handling needed. Values comfortably fit in Go's `int`.
- **Already sorted / reverse sorted input.** Doesn't change correctness; it's just a specific ordering the general algorithm already covers.
