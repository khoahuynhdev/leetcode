# 3568. Minimum Moves to Clean the Classroom

[LeetCode Link](https://leetcode.com/problems/minimum-moves-to-clean-the-classroom/)

Difficulty: Medium
Topics: Array, Hash Table, Bit Manipulation, Breadth-First Search, Matrix
Acceptance Rate: 41.0%

## Hints

### Hint 1

Every move costs exactly one unit, and you want the *fewest* moves. On a grid
with unit-cost edges that screams BFS. But before you write the usual
`visited[r][c]` grid BFS, ask yourself: is "which cell am I on" really enough to
describe the situation you are in? Two students standing on the same cell are
not necessarily equally well off.

### Hint 2

Look hard at the constraint "at most 10 `'L'` cells". Ten of anything is a loud
hint that a subset should be encoded as a bitmask (`2^10 = 1024` possibilities).
So part of your state is *which litter you have already collected*. And notice
that you may legitimately need to walk over a cell you have already visited —
to double back for a second piece of litter, or to detour to an `'R'` — so
per-cell visited marking would be wrong anyway.

### Hint 3

The state is the triple `(row, col, collectedMask)` plus the remaining energy.
The trick is what to do with energy. Do **not** make it a fourth dimension you
mark as visited independently; instead, for each `(cell, mask)` remember only
the *largest* remaining energy you have ever arrived there with. If you show up
at the same `(cell, mask)` again with less-or-equal energy, that arrival is
strictly dominated — anything the weaker version could do next, the stronger one
could already do. Combine that pruning with a level-by-level BFS and the first
time the mask becomes full, the current level count is the answer.

## Approach

**Step 1 — preprocess the grid.** Scan once to find the starting cell `S` and to
assign each `'L'` cell a bit index `0..k-1` (with `k <= 10`). Keep a flat array
`litterID[r*n+c]` holding that index, or `-1` for non-litter cells. If `k == 0`
there is nothing to clean, so return `0` immediately.

**Step 2 — define the state.** A state is `(r, c, mask, energy)`:

- `(r, c)` — where the student stands,
- `mask` — bitset of litter already collected (`mask == (1<<k)-1` means done),
- `energy` — units left before the student is stuck.

`(r, c, mask)` is the part we index on; `energy` is the part we *compare* on.

**Step 3 — BFS in layers.** Start from `{start, mask: 0, energy: energy}`.
Process the whole frontier for move count `d`, collecting the frontier for
`d+1`. For each state in the frontier:

- If `energy == 0`, the student cannot move at all — skip it. (An `'R'` cell
  refills on arrival, so a state parked on `'R'` never has `energy == 0`.)
- Otherwise try all four neighbours. Skip out-of-bounds cells and `'X'`.
- Compute the new energy as `energy - 1`, then apply the destination cell:
  - `'R'` overrides it back to the full capacity `energy` (unconditionally —
    it "restores to full capacity regardless of current level"),
  - `'L'` sets bit `litterID[dest]` in the mask.
- If the new mask is full, return `d + 1`. This is safe precisely because BFS
  explores in nondecreasing move count.

**Step 4 — the dominance pruning.** Maintain `best[cell][mask]`, initialised to
`-1`, meaning "greatest energy seen at this pair". Push the new state only when
`newEnergy > best[cell][mask]`, and update the entry when you do. This is what
keeps the search finite: without it, a loop through an `'R'` cell would let you
wander forever. It is also what makes revisits *possible* when they matter — a
cell you already stood on gets re-explored the moment you come back with a
fuller tank, which is exactly the situation where a detour to a reset area pays
off.

**Worked example — `["LS","RL"]`, `energy = 4`.** Litter `(0,0)` is bit 0 and
`(1,1)` is bit 1; start is `(0,1)` with energy 4.

- Level 0: `((0,1), mask=00, e=4)`.
- Level 1: move left to `(0,0)`, an `'L'` → `mask=01`, `e=3`. (Moving down to
  `(1,1)` gives `mask=10`, `e=3`; BFS keeps both branches alive.)
- Level 2: from `(0,0)` move down to `(1,0)`, an `'R'` → energy resets to 4,
  `mask=01`. Note the tank is *fuller* than it was two moves ago.
- Level 3: from `(1,0)` move right to `(1,1)`, an `'L'` → `mask=11`, which is
  full. Return `3`.

If you had marked `(1,0)` visited without tracking energy, the branch that
reaches it later with a full tank would have been thrown away and you could
report `-1` on grids that are actually solvable. That is the single most common
way this problem is failed.

## Complexity Analysis

Let `m x n` be the grid size, `k <= 10` the litter count, and `E` the energy
capacity.

Time Complexity: O(m * n * 2^k * E) — there are `m * n * 2^k` distinct
`(cell, mask)` pairs, and each can be enqueued at most `O(E)` times because the
recorded best energy must strictly increase on every re-entry. Each dequeue does
constant work over 4 neighbours. With the given limits (`400 * 1024 * 50`) this
is comfortably fast, and in practice the frontier stays far smaller than the
bound.

Space Complexity: O(m * n * 2^k) for the `best` table, plus the BFS frontier,
which is bounded by the same quantity.

## Edge Cases

- **No litter in the grid.** The full mask is `0`, which is already satisfied at
  move zero. Return `0` before entering the BFS, otherwise the loop only checks
  the goal *after* a move and you would return `-1` or an inflated count.
- **Energy runs out on a non-`'R'` cell.** The student is stuck, not dead — the
  state simply has no outgoing edges. Guard with `if energy == 0 { skip }`
  rather than pruning the state when it is created; it might still be the state
  that lands on the final piece of litter.
- **Landing on `'R'` sets energy to the maximum, not `energy + something`.**
  Arriving with 3 of 4 left and stepping onto `'R'` gives 4, never 5.
- **Revisiting cells is mandatory, not optional.** Grids like `["LSL"]` require
  walking back over the start. A plain `visited[r][c]` grid returns `-1` here.
- **Litter reachable but not *all* of it.** Collecting a subset is worthless;
  only the full mask counts. `["L.S","RXL"]` with `energy = 3` reaches either
  piece but never both — return `-1`.
- **Litter fully walled off by `'X'`,** e.g. `["SXL"]`. BFS exhausts the frontier
  and falls through to `-1`.
- **1x1 grid containing only `'S'`.** Handled by the no-litter early return.
- **Tight-but-sufficient energy.** `["S.L"]` with `energy = 2` succeeds in 2
  moves while `energy = 1` fails; make sure you decrement *before* checking
  whether the destination completes the job, not after.

This one is a fair Medium that punishes a reflexive grid BFS. If your first
instinct was `visited[r][c]`, that is the normal path through this problem — the
lesson worth keeping is that the visited key must contain everything that
distinguishes two situations, and here that means the litter mask plus a
dominance rule on energy.
