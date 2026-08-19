# 1386. Cinema Seat Allocation

[LeetCode Link](https://leetcode.com/problems/cinema-seat-allocation/)

Difficulty: Medium
Topics: Array, Hash Table, Greedy, Bit Manipulation
Acceptance Rate: 45.9%

## Hints

### Hint 1

Look hard at the constraints before designing anything. `n` can be up to 10^9, but `reservedSeats.length` is at most 10^4. That gap is the whole problem: you can never afford to iterate over rows, but you *can* afford to iterate over reservations. Ask yourself what a row with zero reservations always contributes, and whether you need to look at it at all.

### Hint 2

Group the reservations by row — a hash map from row number to "what is blocked in this row" is the natural container. Only the rows that appear in the input can ever differ from the default. Each such row is completely independent of every other row, and within a row you only ever have to answer a tiny fixed question: which of the three candidate blocks (2-5, 4-7, 6-9) survive?

### Hint 3

Two insights make this click.

First, seats 1 and 10 belong to no block at all, so a reservation there is pure noise — ignore it. That also means a row whose only reservations are at seats 1 and/or 10 behaves exactly like an empty row.

Second, the greedy choice inside a row is forced, so there is nothing to search. Blocks 2-5 and 6-9 are disjoint, while 4-7 overlaps both. So the answer for a row is: 2 if both the left and right block are free; otherwise 1 if *any* of the three blocks is free; otherwise 0. You never need to "prefer" the middle block — it can only ever be worth 1, and it is only relevant when at least one of the outer blocks is already broken.

A 10-bit mask per row (bit `s` set means seat `s` is reserved) turns each block check into a single `mask & blockMask == 0` test.

## Approach

The rows are fully independent, and a row with no reservations in seats 2..9 always seats exactly 2 groups (blocks 2-5 and 6-9). So the plan is: compute the default for all rows, then correct only the rows that actually appear in the input.

**Step 1 — build the masks.** Walk `reservedSeats` once. For an entry `[row, seat]`, skip it entirely if `seat == 1` or `seat == 10`, since those seats appear in none of the three blocks. Otherwise set bit `seat` in `masks[row]`. After this pass, `masks` contains exactly the rows that are genuinely affected; every other row is untouched.

**Step 2 — count the untouched rows.** There are `n - len(masks)` of them, and each seats 2 groups, contributing `2 * (n - len(masks))`. This is the trick that lets `n = 10^9` cost nothing.

**Step 3 — score each affected row.** With the three block masks precomputed:

- left = seats 2,3,4,5 → `0b0000111100`
- middle = seats 4,5,6,7 → `0b0011110000`
- right = seats 6,7,8,9 → `0b1111000000`

For a row mask `m`, a block is usable iff `m & blockMask == 0`. If left and right are both usable, this row gives 2 (they share no seat, so both groups fit simultaneously). Otherwise, if any one of left, middle, or right is usable, the row gives 1. Otherwise it gives 0. Note that if left and right are both free the middle is irrelevant — taking it would cost us a group, which is why the greedy is safe.

**Walking through Example 1** (`n = 3`, reservations `[[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]`):

- Row 1: seats 2, 3, 8 reserved. Left is broken (2, 3), right is broken (8), middle 4-7 is clear → 1 group.
- Row 2: seat 6 reserved. Left 2-5 is clear, middle and right both contain seat 6 → 1 group.
- Row 3: only seats 1 and 10, both discarded, so row 3 never enters the map → it is one of the `n - len(masks) = 3 - 2 = 1` untouched rows → 2 groups.

Total: `2 * 1 + 1 + 1 = 4`. ✅

This is a satisfying "medium" — the algorithm is short, but it rewards you for reading the constraints first and for noticing that seats 1 and 10 are decoys. If it did not click immediately, that is normal; the constraint-driven insight is the skill being trained here.

## Complexity Analysis

Time Complexity: O(R), where R = `len(reservedSeats)`. One pass to build the map plus one pass over its (at most R) entries; each row is scored with a constant number of bit tests. Independent of `n`.

Space Complexity: O(R) for the map, which holds at most one entry per distinct affected row.

## Edge Cases

- **Huge `n` with few reservations.** With `n` up to 10^9, any loop over rows is fatal. Deriving untouched rows arithmetically is mandatory, not an optimization.
- **Reservations only at seats 1 and 10.** These seats are in no block, so such a row still seats 2 groups. Filtering them out *before* inserting into the map is what makes this fall out automatically — if you insert the row anyway, you must remember to still score it as 2, and forgetting is the classic bug here.
- **Result overflow.** The answer can be as large as `2 * 10^9`, which does not fit in a 32-bit signed integer. Go's `int` is 64-bit on the platforms LeetCode uses, so this is safe, but it is worth noticing in languages where it is not.
- **A row where only the middle block survives.** E.g. seats 2 and 9 reserved: left and right are both broken but 4-7 is clear, so the answer is 1, not 0. A solution that only ever checks the two disjoint outer blocks silently undercounts.
- **A row where left and right are both free.** Must yield 2, not 1 — do not stop at the first usable block.
- **A fully blocked row.** E.g. seats 4 and 5 reserved kills all three blocks at once → 0.
- **Duplicate-free but unsorted input.** The input is guaranteed distinct and arrives in arbitrary order; the map-based accumulation handles both without sorting.
