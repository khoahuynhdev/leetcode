# 2161. Partition Array According to Given Pivot

[LeetCode Link](https://leetcode.com/problems/partition-array-according-to-given-pivot/)

Difficulty: Medium
Topics: Array, Two Pointers, Simulation
Acceptance Rate: 90.1%

## Hints

### Hint 1

This problem is really a *stable partition*. Notice the three buckets that the
result must contain, in order: elements `< pivot`, then elements `== pivot`,
then elements `> pivot`. Whenever you see "group items into ordered categories
while preserving relative order," think about a single linear scan that decides
which bucket each element belongs to.

### Hint 2

You only need to make decisions based on a comparison with `pivot`: less,
equal, or greater. Since the relative order *within* the "less" group and
*within* the "greater" group must be preserved, you cannot sort the array —
sorting would destroy that original ordering. Instead, walk the array in its
original order and append each element to the correct group.

### Hint 3

The "equal to pivot" elements are all identical, so you never need to store
them — you only need to **count** how many there are. So one pass can build a
"less than" list (in order) and a "greater than" list (in order) and a count of
pivots. The answer is simply: the less-than list, followed by `count` copies of
the pivot, followed by the greater-than list. You can even do this in place with
two pointers writing from the front for the "less" group and from the back for
the "greater" group.

## Approach

We need a stable three-way partition around `pivot`. Sorting is off the table
because it would reorder elements that compare equal in our three-way sense but
differ in value (e.g. `9` and `5` are both `< 10` and must keep their order).

A clean linear approach uses **two pointers writing into the result**:

1. Allocate a result slice of the same length `n`.
2. Keep a write index `left` starting at `0` (the next slot for a "less than"
   element) and a write index `right` starting at `n-1` (the next slot for a
   "greater than" element).
3. Scan `nums` left to right. For each value `< pivot`, place it at
   `result[left]` and increment `left`. This preserves the relative order of the
   smaller elements because we encounter and place them in their original order.
4. Scan `nums` right to left in a second loop. For each value `> pivot`, place it
   at `result[right]` and decrement `right`. Scanning from the right and writing
   from the right keeps the greater elements in their original relative order.
5. Every index strictly between the final `left` and `right` belongs to a pivot
   value. Fill `result[left..=right]` with `pivot`.

Why it works: after the first loop, slots `0..left-1` hold all smaller elements
in order. After the second loop, slots `right+1..n-1` hold all larger elements
in order. The remaining middle slots count exactly the number of elements equal
to the pivot, so we fill them with `pivot`.

Walking through Example 1, `nums = [9,12,5,10,14,3,10]`, `pivot = 10`:

- Less-than pass collects `9, 5, 3` into front slots → `[9,5,3,_,_,_,_]`.
- Greater-than pass (right to left) sees `10(skip),3(skip... no, <),14(>),...]`
  collecting `14, 12` from the back, but written back-to-front so they land as
  `... 12, 14` → `[9,5,3,_,_,12,14]`.
- The middle two slots get the pivot → `[9,5,3,10,10,12,14]`. ✓

An equivalent and perhaps more obviously-correct alternative is the
**three-bucket** method: build a `less` slice, count `equal`, build a `greater`
slice in one pass, then concatenate `less + equalCopies + greater`. Same time and
space complexity; pick whichever reads more clearly to you. The implementation
uses the two-pointer in-place-style write because it avoids extra slice
allocations beyond the single result.

## Complexity Analysis

Time Complexity: O(n) — a constant number of linear passes over `nums`.
Space Complexity: O(n) — for the returned result array. (This is unavoidable
since we must return a rearranged array; no extra space is used beyond the
output.)

## Edge Cases

- **All elements equal to the pivot**: e.g. `[5,5,5]`, `pivot = 5`. The
  less-than and greater-than passes write nothing, and the entire array is filled
  with the pivot. The two write pointers must correctly leave the whole range to
  the middle.
- **No elements equal to the pivot beyond required**: the constraint guarantees
  `pivot` appears at least once, so the middle region is never empty, but the
  code should still behave if the less/greater groups are empty.
- **All elements less than the pivot** (or all greater): one group fills the
  front (or back) and the rest is pivots. Verify pointers don't overlap
  incorrectly.
- **Single element array**: `[x]` with `pivot = x`. Result is just `[x]`.
- **Negative numbers**: values can be negative (down to `-10^6`); comparisons
  must use the raw integer values, not absolute values.
- **Large input (up to 10^5)**: the O(n) approach handles this comfortably;
  anything involving sorting or nested scans risks being slower than necessary.
