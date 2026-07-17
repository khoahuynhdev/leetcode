# 3312. Sorted GCD Pair Queries

[LeetCode Link](https://leetcode.com/problems/sorted-gcd-pair-queries/)

Difficulty: Hard
Topics: Array, Hash Table, Math, Binary Search, Combinatorics, Counting, Number Theory, Prefix Sum
Acceptance Rate: 34.5%

## Hints

### Hint 1

There can be up to `n * (n - 1) / 2 ≈ 5 * 10^9` pairs, so you can never actually build the `gcdPairs` array. But notice that every GCD value is bounded: `1 <= gcd <= max(nums) <= 5 * 10^4`. That is a much smaller universe than the number of pairs. What can you count over that small range instead of enumerating pairs?

### Hint 2

Instead of asking "what is the gcd of each pair", flip it around: for each candidate value `g`, ask "how many pairs have gcd *exactly* `g`?" If you had that count for every `g`, the sorted `gcdPairs` array is just each `g` repeated that many times — and a prefix sum over `g` lets you answer any query with binary search. This is a classic **counting + number theory** setup.

### Hint 3

Counting pairs with gcd *exactly* `g` directly is hard, but counting pairs where both elements are *divisible by* `g` is easy: if `m` numbers are multiples of `g`, there are `C(m, 2) = m*(m-1)/2` such pairs. That count includes pairs whose gcd is `g`, `2g`, `3g`, ... So process `g` from large to small (or subtract as you go): `exact[g] = pairsDivisibleBy[g] - exact[2g] - exact[3g] - ...`. This inclusion–exclusion sieve is the crux. Then prefix-sum `exact[]` and binary-search each query.

## Approach

**Step 1 — Frequency table.** Let `M = max(nums)`. Build `cnt[v]` = how many elements of `nums` equal `v`, for `1 <= v <= M`.

**Step 2 — Count pairs divisible by each `g`.** For every `g` from `1` to `M`, sum `cnt[g] + cnt[2g] + cnt[3g] + ...` to get `m`, the number of array elements that are multiples of `g`. The number of pairs where *both* elements are multiples of `g` is `pairs[g] = m*(m-1)/2`. This harmonic-style double loop costs `O(M log M)`.

**Step 3 — Inclusion–exclusion to get exact gcd counts.** `pairs[g]` counts every pair whose gcd is a multiple of `g`. To isolate pairs with gcd *exactly* `g`, subtract the exact counts of all larger multiples:

```
exactCount[g] = pairs[g] - exactCount[2g] - exactCount[3g] - ...
```

Process `g` from `M` down to `1` so that `exactCount[2g], exactCount[3g], ...` are already finalized when you compute `exactCount[g]`. (You can also fold Steps 2 and 3 into one downward loop.)

**Step 4 — Prefix sums.** Build `prefix[g] = exactCount[1] + exactCount[2] + ... + exactCount[g]`. `prefix[g]` is the number of pairs whose gcd is `<= g`, i.e. how many entries of the conceptual sorted `gcdPairs` array are `<= g`.

**Step 5 — Answer queries.** For a query index `q` (0-based), the answer is the smallest `g` such that `prefix[g] > q`. Binary search over `g` in `[1, M]` finds it in `O(log M)`.

**Why it works.** The sorted `gcdPairs` array places all `exactCount[1]` ones first, then `exactCount[2]` twos, and so on. The prefix sum tells you exactly where each value's block ends, so locating the `q`-th element is a binary search on those block boundaries — no need to materialize the (possibly billions-long) array.

**Tiny example.** `nums = [2,3,4]`, `M = 4`, `cnt = [_,0,1,1,1]`.
- `pairs[1]`: multiples of 1 → 3 numbers → `C(3,2)=3`.
- `pairs[2]`: multiples of 2 → {2,4} → 2 numbers → `C(2,2)=1`.
- `pairs[3]`: {3} → 1 number → 0. `pairs[4]`: {4} → 0.
- Going down: `exact[4]=0`, `exact[3]=0`, `exact[2]=1-0=1`, `exact[1]=3-1=2`.
- `prefix = [_,2,3,3,3]`. Query 0 → smallest g with prefix>0 → g=1. Query 2 → smallest g with prefix>2 → g=2. Matches `[1,2,2]`.

## Complexity Analysis

Time Complexity: O(n + M log M + Q log M), where `n = len(nums)`, `M = max(nums)`, and `Q = len(queries)`. The `M log M` term comes from the harmonic sieve summing multiples; each query is a binary search.

Space Complexity: O(M) for the frequency, exact-count, and prefix arrays.

## Edge Cases

- **Only one distinct pair** (`n = 2`): `gcdPairs` has a single element; every query must return that element. Covered by the general prefix-sum logic.
- **Duplicate values** (e.g. `nums = [2,2]`): `gcd(2,2) = 2`, and `cnt[2] = 2` yields `C(2,2)=1` pair — duplicates must be counted, not deduplicated.
- **All numbers coprime**: most pairs collapse to gcd `1`, so `exactCount[1]` dominates; the prefix sum still resolves queries correctly.
- **Large query indices**: `queries[i]` can be up to `n*(n-1)/2 ≈ 5*10^9`, which overflows 32-bit integers. Use 64-bit (`int` on 64-bit Go, or `int64`) for query values, `m*(m-1)/2`, and prefix sums.
- **`max(nums)` sizing**: allocate arrays up to `M` inclusive; off-by-one on the upper bound drops the largest possible gcd.
