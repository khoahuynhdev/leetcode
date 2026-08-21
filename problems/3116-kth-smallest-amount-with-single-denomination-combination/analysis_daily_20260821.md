# 3116. Kth Smallest Amount With Single Denomination Combination

[LeetCode Link](https://leetcode.com/problems/kth-smallest-amount-with-single-denomination-combination/)

Difficulty: Hard
Topics: Array, Math, Binary Search, Bit Manipulation, Combinatorics, Number Theory
Acceptance Rate: 23.7%

## Hints

### Hint 1

The "no mixing denominations" rule is doing a lot of work here. Strip away the coin framing and ask: which integers are actually reachable? Once you can describe the reachable set with a single sentence, notice that `k` can be as large as `2 * 10^9` — far too large to generate values one at a time with a heap or a sieve. When you cannot *enumerate* to the answer but the answer is a number on a line, think about *searching* for it instead.

### Hint 2

Binary search on the answer needs a predicate you can evaluate quickly: given a candidate `x`, how many reachable values are `<= x`? For one coin `c` that count is just `floor(x / c)`. For several coins you cannot add those counts up, because a value like 12 is counted by both coin 2 and coin 6. Overcounting a union of sets is exactly the situation one classical counting technique was invented for — and look at the constraint `coins.length <= 15`, which quietly tells you that touching every subset of coins is affordable.

### Hint 3

Inclusion-exclusion over subsets: a value `<= x` is reachable iff it is a multiple of at least one coin, so

```
count(x) = Σ over non-empty subsets S of  (-1)^(|S|+1) * floor(x / lcm(S))
```

`count` is non-decreasing in `x`, so binary search for the smallest `x` with `count(x) >= k`. Two details make it practical: pick the upper bound `minCoin * k` (the multiples of the smallest coin alone already give you `k` values by then), and note that any subset whose LCM exceeds that bound contributes `0` forever — compute the LCMs once, capped, and throw those subsets away.

## Approach

**Step 1 — characterize the reachable set.** Since you may never combine denominations, an amount is producible exactly when it is a positive multiple of *some* single coin. The answer is the `k`-th smallest element of `M(c₁) ∪ M(c₂) ∪ … ∪ M(cₙ)`, where `M(c)` is the set of positive multiples of `c`. Nothing about summing coins survives — this is a union-of-arithmetic-progressions problem.

**Step 2 — drop redundant coins.** If `c` is a multiple of another coin `b`, then `M(c) ⊆ M(b)` and `c` contributes nothing. Sorting ascending and keeping a coin only when it is not divisible by an already-kept coin collapses `[3,6,9]` to `[3]` and `[2,3,4,6,8]` to `[2,3]`. This is optional for correctness but can shrink the subset enumeration by orders of magnitude.

**Step 3 — count with inclusion-exclusion.** The number of multiples of `c` that are `<= x` is `floor(x / c)`. A value in the intersection `M(a) ∩ M(b)` is precisely a multiple of `lcm(a, b)`, so the union size follows from the inclusion-exclusion principle: add all single-coin counts, subtract all pairwise LCM counts, add back all triples, and so on. With at most 15 coins that is `2^15 - 1 = 32767` terms, each a division.

**Step 4 — binary search the answer.** `count(x)` is monotone non-decreasing, and the goal — the `k`-th smallest reachable value — is the smallest `x` where `count(x)` first reaches `k`. Search on `[1, minCoin * k]`. The upper bound is valid because `minCoin, 2·minCoin, …, k·minCoin` are already `k` distinct reachable values, so `count(minCoin * k) >= k`. The classic `lo < hi` / `hi = mid` / `lo = mid + 1` template lands on that first crossing point.

**Step 5 — precompute the LCMs, and cap them.** Recomputing subset LCMs inside every binary-search iteration is wasteful. Build them once with a bitmask DP: `lcm[mask] = lcm(lcm[mask without lowest set bit], coin[lowest bit])`. LCMs of 15 values up to 25 can grow enormous, so cap the computation at `hi + 1` and check `a/g > limit/b` *before* multiplying — that both prevents `int64` overflow and marks the subset as permanently contributing zero. Store only the surviving `(lcm, sign)` pairs; in practice the survivors are far fewer than 32767.

**Worked example — `coins = [5,2]`, `k = 7`.** No coin divides another, so both stay; `hi = 2 * 7 = 14`. The terms are `+x/2`, `+x/5`, `−x/10`. Then `count(11) = 5 + 2 − 1 = 6 < 7`, while `count(12) = 6 + 2 − 1 = 7 >= 7`, and `count(11) < 7` confirms 12 is the first crossing. The answer is **12**, matching the enumeration `2, 4, 5, 6, 8, 10, 12`.

## Complexity Analysis

Let `n` be the number of coins after reduction (`n <= 15`), `m` the smallest coin (`m <= 25`), and `T <= 2^n` the number of surviving inclusion-exclusion terms.

Time Complexity: O(2^n + T · log(m · k)) — one pass to build all `2^n` subset LCMs, then roughly `log2(25 · 2·10^9) ≈ 36` binary-search iterations, each summing `T` divisions. Worst case about `3.3 · 10^4 + 1.2 · 10^6` cheap operations.

Space Complexity: O(2^n) for the subset-LCM table and the surviving-term list.

## Edge Cases

- **A coin equal to 1.** Every positive integer becomes reachable, so the answer is exactly `k`. Coin reduction collapses the whole array to `[1]`, and the upper bound `1 * k` makes the binary search terminate immediately at `k` — worth a test because a wrong upper bound shows up here first.
- **One coin dividing another (`[3,6,9]`).** The naive "sum the counts" approach overcounts badly. Either the reduction step or inclusion-exclusion must handle it; testing both examples covers this since example 1 is exactly this shape.
- **Integer overflow.** `minCoin * k` reaches `25 · 2·10^9 = 5·10^10`, well past `int32`, so the return type is `int64` and all arithmetic must be 64-bit. Worse, an uncapped LCM of 15 coprime-ish values overflows `int64` outright — the `a/gcd > limit/b` guard before multiplying is not a micro-optimization, it is a correctness requirement.
- **`k = 1`.** The answer is simply the smallest coin. The binary search range is `[1, minCoin]`, and the loop must not skip past the lower end.
- **Maximum `k` with a large smallest coin (`coins = [25]`, `k = 2·10^9`).** The answer is `5·10^10`; the search must actually reach the top of its range rather than assuming the answer fits in 32 bits.
- **Unsorted input.** The reduction step relies on ascending order to guarantee that a removed coin is a multiple of a *kept* smaller one, so sort before filtering.
- **All 15 coins, mutually non-dividing (e.g. `11..25`).** This is the performance worst case: no reduction is possible and the full `2^15` subset sweep runs. The LCM cap is what keeps the per-iteration term count small.

This one earns its Hard rating mostly through the *combination* of ideas — binary search on the answer, inclusion-exclusion over bitmasks, and careful overflow handling — rather than any single hard step. If you have seen "count how many values `<= x` satisfy P" before, the shape should feel familiar; the LCM capping is the part most people get wrong on the first attempt.
