# 3666. Minimum Operations to Equalize Binary String

[LeetCode Link](https://leetcode.com/problems/minimum-operations-to-equalize-binary-string/)

Difficulty: Hard
Topics: Math, String, Breadth-First Search, Union-Find, Ordered Set
Acceptance Rate: 32.7%

## Hints

### Hint 1

Don't think about the operations step-by-step. Instead, think about the **total number of times** each position gets flipped across all operations. What parity must each position's flip count have?

### Hint 2

If you perform `m` operations, the total number of individual flips is `m * k`. Each zero needs an odd flip count, each one needs an even flip count, and the sum of all flip counts must equal `m * k`. This creates a **parity constraint**: `m * k` must have the same parity as the number of zeros `z`. When does this fail?

### Hint 3

Beyond parity, there is a **capacity constraint**. With `m` operations, each position can be flipped at most `m` times. The maximum achievable total flips (respecting parities) depends on whether `m` is even or odd. For even `m`: the capacity constraint becomes `m * (n - k) >= z`. For odd `m`: it becomes `m * (n - k) >= n - z`. Try both parities of `m` and take the minimum valid value.

## Approach

Let `n = len(s)` and `z = count of '0' in s`.

**Core observation:** The final state of position `i` depends only on whether its total flip count `f_i` across all operations is odd or even, not on the order. So we need:
- `f_i` is odd for each zero (to flip it to '1')
- `f_i` is even for each one (to keep it as '1')
- `sum(f_i) = m * k` (each of `m` operations flips exactly `k` positions)
- `0 <= f_i <= m` (position `i` appears in at most `m` operations)

**Parity constraint:** Since the sum of odd numbers is odd when there's an odd count of them, `sum(f_i) ≡ z (mod 2)`. Thus `m * k ≡ z (mod 2)`. If `k` is even, then `m * k` is always even, so `z` must be even — otherwise it's impossible.

**Capacity constraint:** The maximum achievable sum with correct parities depends on `m`'s parity:
- Even `m`: zeros can have at most `m-1` flips (largest odd ≤ m), ones can have `m`. Max sum = `n*m - z`. So `m*k ≤ n*m - z`, giving `m*(n-k) ≥ z`.
- Odd `m`: zeros can have at most `m` flips, ones can have `m-1`. Max sum = `n*m - (n-z)`. So `m*(n-k) ≥ n - z`.

**Algorithm:**
1. If `z = 0`, return 0.
2. If `k` is even and `z` is odd, return -1.
3. If `k = n`: return 1 if `z = n` (flip everything once), else -1 (can't have different flip parities).
4. For `k < n`, try both parities of `m`:
   - Compute the lower bound from `m*k ≥ z` (i.e., `m ≥ ⌈z/k⌉`).
   - Compute the lower bound from the capacity constraint.
   - Adjust `m` upward if needed to match the required parity.
5. Return the minimum valid `m` across both parities.

**Why decomposition always exists:** Given valid `f_i` values (correct sum, parities, and each ≤ m), we can always decompose them into `m` subsets of size `k` by greedily assigning positions with the highest remaining counts to each operation.

## Complexity Analysis

Time Complexity: O(n) — a single pass to count zeros, then O(1) math.

Space Complexity: O(1) — only a few integer variables.

## Edge Cases

- **All ones (z = 0):** No operations needed, return 0.
- **k = n:** Every operation flips all positions, so all positions always share the same flip parity. Only works if all are zeros (return 1) or all are ones (return 0).
- **k even, z odd:** Impossible because `m * k` is always even but must match `z`'s odd parity. Return -1.
- **k = n - 1:** Each operation leaves exactly one position unflipped. The capacity constraint can force large answers (up to ~n operations).
- **z = n (all zeros):** Always solvable when `k ≤ n`. The minimum operations depend on `k`'s relationship to `n`.
