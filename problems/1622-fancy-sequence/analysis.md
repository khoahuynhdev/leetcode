# 1622. Fancy Sequence

[LeetCode Link](https://leetcode.com/problems/fancy-sequence/)

Difficulty: Hard
Topics: Math, Design, Segment Tree
Acceptance Rate: 22.9%

## Hints

### Hint 1

Applying `addAll` and `multAll` to every element in the sequence each time would be O(n) per operation. Think about how you can defer or batch these operations so they only need to be applied when you actually query a value.

### Hint 2

Consider representing the cumulative effect of all `addAll` and `multAll` operations as a single **affine transformation** of the form `f(x) = m*x + a`. How do `addAll(inc)` and `multAll(k)` compose into this form? If you track this globally, can you avoid touching individual elements during bulk operations?

### Hint 3

The critical insight is to store each appended value in a "pre-transformed" (inverse-transformed) form. When you `append(val)`, compute the base value `b` such that applying the current global transform to `b` yields `val`. That is, `b = (val - a) * m^(-1) mod p`. Then `getIndex` simply applies the current global transform to the stored base value. Since the modulus is prime, you can compute the modular inverse using Fermat's little theorem: `m^(-1) = m^(p-2) mod p`.

## Approach

We maintain a global affine transformation `f(x) = mult * x + add (mod 10^9+7)` that represents the cumulative effect of all `addAll` and `multAll` operations performed so far.

**Operations:**

1. **`Append(val)`**: We need to store a base value `b` such that `mult * b + add ≡ val (mod p)`. Solving: `b = (val - add) * mult^(-1) mod p`. We compute the modular inverse of `mult` using fast exponentiation (Fermat's little theorem, since `10^9+7` is prime).

2. **`AddAll(inc)`**: Update the global transform: `add = add + inc`.

3. **`MultAll(m)`**: Update the global transform: `mult = mult * m` and `add = add * m`. This is because composing `g(x) = m*x` with `f(x) = mult*x + add` gives `g(f(x))` ... but actually we want `f` applied first then `g`, so the new transform applied to base values is `m*(mult*x + add) = (m*mult)*x + (m*add)`.

4. **`GetIndex(idx)`**: Return `mult * vals[idx] + add mod p`, or `-1` if the index is out of bounds.

**Example walkthrough:**
- `append(2)`: global is `(1, 0)`, store `b = (2-0)*1 = 2`
- `addAll(3)`: global becomes `(1, 3)`
- `append(7)`: store `b = (7-3)*1 = 4`
- `multAll(2)`: global becomes `(2, 6)`
- `getIndex(0)`: `2*2 + 6 = 10` ✓

## Complexity Analysis

Time Complexity: O(log p) for `append` (due to modular inverse), O(1) for `addAll`, `multAll`, and `getIndex`. Here `p = 10^9+7`.

Space Complexity: O(n) where n is the number of appended elements.

## Edge Cases

- **Index out of bounds**: `getIndex` with an index >= length of the sequence must return -1.
- **No operations before query**: If `getIndex` is called immediately after `append` with no `addAll`/`multAll`, the value should be returned unchanged.
- **Overflow concerns**: All arithmetic must be done modulo `10^9+7`. Since intermediate products of two values up to `10^9` can exceed 32-bit range, 64-bit integers are essential.
- **Multiplying by values that share factors with the modulus**: Since `10^9+7` is prime, every non-zero value has a modular inverse, so this is always safe.
- **Large number of operations**: With up to `10^5` total calls, the O(1) / O(log p) per-operation design ensures we stay well within time limits.
