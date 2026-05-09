# 1914. Cyclically Rotating a Grid

[LeetCode Link](https://leetcode.com/problems/cyclically-rotating-a-grid/)

Difficulty: Medium
Topics: Array, Matrix, Simulation
Acceptance Rate: 58.4%

## Hints

### Hint 1

The matrix decomposes into independent concentric "rings" (layers). Each ring rotates on its own — rotating one layer never touches another. So the problem reduces to: how do you rotate **a single ring** by `k` steps efficiently?

### Hint 2

A ring is just a closed loop of cells. If you can flatten that loop into a 1-D array, rotation is trivial. Think about how to walk the perimeter in a fixed direction (say, clockwise), what `k` really does to a 1-D array of length `L`, and why the naive "rotate one step at a time, k times" is wasteful when `k` can be up to `10^9`.

### Hint 3

Rotation by `k` steps on an array of length `L` is the same as rotation by `k % L` steps — that is the key to surviving the `k <= 10^9` constraint. Counter-clockwise rotation by `s` on a ring whose perimeter is laid out **clockwise** is equivalent to a **left shift by `s`**: the element that ends up at clockwise position `i` is the one that was originally at clockwise position `(i + s) % L`. After that, write the shifted array back along the same clockwise traversal you used to read it.

## Approach

The matrix is built of `min(m, n) / 2` concentric rings. Layer `L` (0-indexed from the outside) occupies the cells whose row is in `[L, m-1-L]` and column is in `[L, n-1-L]`, but only on the boundary of that sub-rectangle. Rings are disjoint and rotate independently, so process them one at a time.

For a single layer `L`:

1. **Flatten the ring clockwise.** Walk:
   - top edge left→right: `(L, L) ... (L, n-1-L)`
   - right edge top→bottom (skip the corner already taken): `(L+1, n-1-L) ... (m-1-L, n-1-L)`
   - bottom edge right→left (skip the corner): `(m-1-L, n-2-L) ... (m-1-L, L)`
   - left edge bottom→top (skip both corners): `(m-2-L, L) ... (L+1, L)`

   Collect the values into a slice `ring` of length `Lsize = 2*((m - 2*L) + (n - 2*L)) - 4`.

2. **Rotate.** Counter-clockwise rotation moves each element into its CCW-neighbor's slot. With the perimeter laid out clockwise, that is exactly a left shift. Compute `s = k % Lsize` and produce `rotated[i] = ring[(i + s) % Lsize]`.

3. **Write back.** Re-walk the same clockwise path and assign `rotated[idx]` cell by cell.

**Why it works.** Each ring is a simple cycle in the rotation permutation, so a `k`-rotation of the matrix is just a `k`-rotation of every ring. Choosing a single direction (clockwise) for both reading and writing avoids index-bookkeeping bugs; the `% Lsize` collapses the huge `k` to at most `Lsize - 1`.

**Worked example** (Example 2): on `4×4` with `k=2`, layer 0's clockwise perimeter is `[1,2,3,4,8,12,16,15,14,13,9,5]` (length 12). Left-shift by 2: `[3,4,8,12,16,15,14,13,9,5,1,2]` — write it back along the same path. Layer 1's perimeter is `[6,7,11,10]`; with `k % 4 = 2` it becomes `[11,10,6,7]`, which slots back into the inner 2×2.

## Complexity Analysis

Time Complexity: O(m * n) — each cell is read once and written once, across all layers combined.
Space Complexity: O(m + n) — the largest single ring's perimeter is `2*(m + n) - 4`; only one layer's buffer is held at a time.

## Edge Cases

- **Smallest grid (`m == n == 2`)**: exactly one layer of 4 cells; the `left edge` walk in step 1 is empty by construction (`m-2-L > L` is false), and the algorithm still works.
- **Very thin grids (e.g., `2 × 50` or `50 × 2`)**: only one layer exists (`min(m,n)/2 = 1`); make sure the perimeter walk correctly produces `2*(m + n) - 4` cells without double-counting corners.
- **`k` is a multiple of the ring's perimeter length**: `k % Lsize == 0` means that ring is unchanged. Different layers can have different perimeter lengths, so some layers may rotate while others stand still — the per-layer modulo handles this naturally.
- **Large `k` (up to `10^9`)**: never simulate `k` single-step rotations — always reduce with `k % Lsize` per layer.
- **Non-square grids where `m != n`**: layers are still well-defined as long as both dimensions are even; the perimeter formula `2*((m-2L) + (n-2L)) - 4` accounts for rectangular rings.
