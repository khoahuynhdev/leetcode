# 3756. Concatenate Non-Zero Digits and Multiply by Sum II

[LeetCode Link](https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-ii/)

Difficulty: Medium
Topics: Math, String, Prefix Sum

Acceptance Rate: 31.3%

## Hints

### Hint 1

There can be up to `10^5` queries over a string of length up to `10^5`. Rebuilding the number for every query would be `O(m)` per query and `O(m * q)` overall, which is far too slow. Whenever you see "answer many range queries fast," think about precomputing prefix information so each query becomes `O(1)`. The tricky part here is that the quantity you care about — a *concatenated* integer — does not obviously decompose into a simple prefix sum. Try to express it as one anyway.

### Hint 2

The answer is `x * sum`, where `sum` is just the sum of the non-zero digits in the range. That half is easy: keep a prefix sum of digit values and subtract. The hard half is `x`, the integer formed by gluing the non-zero digits together. Think about what each digit contributes to the final value of `x`. A non-zero digit `d` sitting `k` non-zero-digits before the end of the range contributes `d * 10^k`. Notice that `k` depends only on how many non-zero digits come *after* it — which is itself a prefix-countable quantity.

### Hint 3

Let `nz[i]` be the count of non-zero digits in `s[0..i-1]`. For a query `[l, r]`, a non-zero digit at position `j` contributes `d_j * 10^(nz[r+1] - nz[j+1])`. Factor out the range-dependent exponent:

```
x = 10^(nz[r+1]) * sum_{j in [l, r], s[j] != 0}  d_j * 10^(-nz[j+1])
```

The inner sum no longer depends on `r`, so it *is* a prefix sum. Precompute `pref[i] = sum of d_j * inverse10^(nz[j+1])` for all non-zero `j < i` under the modulus (using the modular inverse of 10). Each query then costs a single subtraction and one multiplication by `10^(nz[r+1])`.

## Approach

We work entirely modulo `p = 10^9 + 7`. Two independent quantities are needed per query: the concatenated integer `x` and the digit sum `sum`.

**Digit sum.** Zeros contribute nothing to the sum, so `sum` is simply the sum of the non-zero digits in the range. Maintain a prefix array `dsum[i]` = sum of non-zero digit values in `s[0..i-1]`. Then `sum = dsum[r+1] - dsum[l]`.

**The concatenated integer `x`.** When we glue the non-zero digits `d_{j1} d_{j2} ... d_{jc}` of the range together (in original order), the value is

```
x = d_{j1} * 10^(c-1) + d_{j2} * 10^(c-2) + ... + d_{jc} * 10^0
```

where `c` is how many non-zero digits the range has. The exponent on a given digit equals the number of non-zero digits that follow it inside the range.

Define `nz[i]` = number of non-zero digits in `s[0..i-1]`. For a non-zero digit at position `j`, the number of non-zero digits after it within `[l, r]` is `nz[r+1] - nz[j+1]`. Hence

```
x = sum_j d_j * 10^(nz[r+1] - nz[j+1])
  = 10^(nz[r+1]) * sum_j ( d_j * inv10^(nz[j+1]) )
```

where `inv10` is the modular inverse of 10. The bracketed term is range-independent, so we precompute the prefix sum

```
pref[i] = sum over non-zero positions j < i of  d_j * inv10^(nz[j+1])   (mod p)
```

For a query, `x = 10^(nz[r+1]) * (pref[r+1] - pref[l])  (mod p)`.

**Putting it together.** Precompute `pow10[]`, `invPow10[]`, `nz[]`, `pref[]`, and `dsum[]` in a single `O(m)` pass. Each query is answered in `O(1)`:

1. `cnt = nz[r+1] - nz[l]`. If `cnt == 0`, there are no non-zero digits, so `x = 0` and the answer is `0`.
2. Otherwise `x = pow10[nz[r+1]] * ((pref[r+1] - pref[l]) mod p)` and `sum = dsum[r+1] - dsum[l]`.
3. Answer is `x * sum mod p`.

*Worked example* (`s = "10203004"`, query `[0, 7]`): the non-zero digits are `1, 2, 3, 4` at positions `0, 2, 4, 7`. Their exponents (non-zero digits after them) are `3, 2, 1, 0`, giving `1000 + 200 + 30 + 4 = 1234`. The digit sum is `10`, so the answer is `12340`, matching the expected output.

## Complexity Analysis

Time Complexity: O(m + q), where `m = len(s)` and `q = len(queries)`. Precomputation is a single linear pass (the one modular exponentiation for `inv10` is `O(log p)`); each query is answered in constant time.
Space Complexity: O(m) for the prefix and power tables.

## Edge Cases

- **No non-zero digits in the range** (e.g. `"000"`): `x = 0`, so the answer must be `0` — guard on `cnt == 0` to avoid multiplying by an empty prefix difference.
- **Trailing zeros** (e.g. `"300"` -> `x = 3`): zeros are simply skipped; the exponent bookkeeping via `nz` handles this automatically because zeros never increment the non-zero counter.
- **Single-character range** (`l == r`): must work whether that character is zero (answer `0`) or non-zero (answer `d * d`).
- **Modular subtraction underflow**: `pref[r+1] - pref[l]` can be negative under the modulus; add `p` before taking the final modulus.
- **Large results**: `x` can be an integer with up to `10^5` digits, so it must be kept modulo `p` throughout — never build the actual number. Also confirm that `x * sum` stays within `int64` (both factors are below `p`, so the product fits).
