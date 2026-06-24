# 3700. Number of ZigZag Arrays II

[LeetCode Link](https://leetcode.com/problems/number-of-zigzag-arrays-ii/)

Difficulty: Hard
Topics: Math, Dynamic Programming
Acceptance Rate: 61.5%

## Hints

### Hint 1

Forget the literal definition for a moment and picture the *shape* of a valid array. "No two adjacent equal" plus "no three consecutive strictly increasing or decreasing" means the direction between neighbors must flip every step: up, down, up, down, ... (or down, up, down, up, ...). The values themselves are bounded by a small range `m = r - l + 1 <= 75`, while the length `n` can be huge. That mismatch — tiny value range, enormous length — is the whole puzzle. Think about a state that depends only on the *last value* and the *last direction*, not on the position.

### Hint 2

A natural DP tracks, for each value `v` in `[1, m]`, how many valid sequences end at `v`. But you need to know whether the next step must go up or down, so split it into two states: `up[v]` = sequences ending at `v` whose last step went **up** (came from a smaller value), and `down[v]` = sequences ending at `v` whose last step went **down**. The transition to a new last value `w` is a simple prefix/suffix sum:
- `new_up[w]   = sum of down[v] for all v < w`
- `new_down[w] = sum of up[v] for all v > w`

This is linear in the state. Linear-recurrence plus a length up to `10^9` should make one technique jump out.

### Hint 3

The transition is a fixed linear map on a vector of size `2m`, applied `n - 2` times (starting from the length-2 base case `up[w] = w - 1`, `down[w] = m - w`). Applying the same linear map `n - 2` times is exactly **matrix exponentiation**: build the `2m × 2m` transition matrix once, raise it to the power `n - 2` in `O(log n)` matrix multiplications, multiply by the base vector, and sum every entry (mod `10^9 + 7`). With `2m <= 150`, each multiply is `~150^3` and there are only `~30` of them.

## Approach

Let `m = r - l + 1`. Because every value must lie in `[l, r]` and the constraints are translation-invariant, we can relabel the values as `1, 2, ..., m`; the count does not depend on the actual offset `l`.

A ZigZag array strictly alternates direction. Define two families of states indexed by the last value `v`:

- `up[v]`: number of valid prefixes that end at `v` where the final step increased (previous element `< v`). The next step must therefore decrease.
- `down[v]`: number of valid prefixes that end at `v` where the final step decreased (previous element `> v`). The next step must therefore increase.

**Base case (length 2).** A length-2 array is just a pair `(a, b)` with `a != b`.
- `up[w]` counts pairs ending at `w` with `a < w`, i.e. `w - 1` choices.
- `down[w]` counts pairs ending at `w` with `a > w`, i.e. `m - w` choices.

**Transition (add one element `w`).** To extend a prefix by a new last value `w`:
- An **up** step into `w` requires the previous step to have been **down**, and the previous value to be smaller than `w`:
  `new_up[w] = sum of down[v] for v < w`.
- A **down** step into `w` requires the previous step to have been **up**, and the previous value to be larger than `w`:
  `new_down[w] = sum of up[v] for v > w`.

Both are linear combinations of the current state, so the whole update is a single matrix-vector product. Stack the `2m` quantities `up[1..m]` and `down[1..m]` into one vector `x`, and let `T` be the `2m × 2m` matrix with:
- `T[up_w][down_v] = 1` for every `v < w`,
- `T[down_w][up_v] = 1` for every `v > w`.

Then one step is `x' = T · x`. Going from length 2 to length `n` applies `T` exactly `n - 2` times, so the answer vector is `T^(n-2) · base`, and the result is the sum of all its entries modulo `10^9 + 7`.

We compute `T^(n-2)` by fast exponentiation: square the matrix and conditionally multiply per bit of the exponent, which needs `O(log n)` matrix multiplications.

**Walkthrough (Example 2: `n = 3, l = 1, r = 3`, so `m = 3`).**
Base (length 2): `up = [0, 1, 2]`, `down = [2, 1, 0]`.
One transition (to reach length 3):
- `new_up = [0, down[1], down[1]+down[2]] = [0, 2, 3]`
- `new_down = [up[2]+up[3], up[3], 0] = [3, 2, 0]`

Total `= 0 + 2 + 3 + 3 + 2 + 0 = 10`, matching the expected output.

## Complexity Analysis

Let `m = r - l + 1` (at most 75), so the matrix dimension is `2m` (at most 150).

Time Complexity: O((2m)^3 · log n) — `O(log n)` matrix multiplications, each costing `O((2m)^3)`.
Space Complexity: O((2m)^2) — storage for the transition matrices.

## Edge Cases

- **Smallest length `n = 3`.** The constraints guarantee `n >= 3`, so `n - 2 >= 1` and we always apply at least one transition; no special handling for very short arrays is needed, but it is worth confirming the base case (length 2) is correct since everything builds on it.
- **Minimum range `r = l + 1` (`m = 2`).** With only two values the array is forced to alternate `a, b, a, b, ...`, giving exactly 2 arrays for any `n`. A good sanity check that the recurrence and modulo logic behave for tiny ranges.
- **Huge `n` (up to `10^9`).** A naive per-element DP would be `O(n · m)` and time out; the matrix-power trick is what keeps it `O(log n)`. Make sure the exponent `n - 2` is handled without overflow and that the identity matrix is used as the starting accumulator.
- **Modular arithmetic.** All intermediate products can be as large as `~150 · (10^9)^2`, which overflows 32-bit but fits comfortably in 64-bit; reduce modulo `10^9 + 7` after each multiply-add to keep values bounded.
