# 3514. Number of Unique XOR Triplets II

[LeetCode Link](https://leetcode.com/problems/number-of-unique-xor-triplets-ii/)

Difficulty: Medium
Topics: Array, Math, Bit Manipulation, Enumeration
Acceptance Rate: 40.2%

## Hints

### Hint 1

The brute-force idea is to iterate over every triplet `(i, j, k)` with `i <= j <= k`, XOR the three values, and drop the results into a set. With `n` up to 1500 that is on the order of `n^3 / 6 ≈ 5 * 10^8` operations — too slow. Before optimizing, ask yourself: what is special about the *values* here, and does the ordering `i <= j <= k` actually restrict which XOR results are reachable?

### Hint 2

Because indices may repeat (`i <= j <= k` allows `i == j == k`), the ordering constraint adds nothing: you can reach every value `a ^ b ^ c` where `a`, `b`, `c` are chosen from the array with repetition. XOR is commutative and associative, so only the *set of distinct values* matters, not positions or counts. Reduce `nums` to its distinct values first.

### Hint 3

Every `nums[i]` is at most 1500, so any XOR of these values stays below 2048 (11 bits). That means the number of *reachable* XOR values is tiny — bounded by 2048 — no matter how large `n` is. Split the work in two stages: first compute the set of all pairwise XORs `P = { x ^ y }`, then XOR that set against the distinct values once more to get all triplet XORs. Each stage is bounded by `|distinct| * 2048`, which is comfortably fast.

## Approach

Let `D` be the set of distinct values in `nums`.

1. **Reduce to distinct values.** Duplicates never create a new XOR result, so collapse `nums` to `D`. Since `1 <= nums[i] <= 1500`, `|D| <= 1500`.

2. **Build the pairwise-XOR set `P`.** For every ordered pair `(x, y)` with `x, y in D`, record `x ^ y`. This automatically includes `0` (from `x ^ x`) and every single value `x` is representable later. Because each value is under 2048, `P` lives in a boolean array of size 2048. This costs `O(|D|^2)` time.

3. **Build the triplet-XOR set `T`.** For every `p in P` and every `d in D`, record `p ^ d`. This yields every `x ^ y ^ z` for `x, y, z in D`, which is exactly the set of reachable triplet values. This costs `O(|P| * |D|) = O(2048 * |D|)`.

4. **Count** the number of `true` entries in `T`.

Why the two-stage split works: a triplet XOR is `(x ^ y) ^ z`. Stage 2 collects all pairwise results into `P`, and stage 3 combines each of them with a third distinct value. Because we already store `0 in P`, values reachable with "fewer distinct picks" are also covered — e.g. `p = 0` gives back the single elements `d`, and `p = x ^ y` with `z` reproducing one of them gives back a single value too.

**Walkthrough on `nums = [1, 3]`:** `D = {1, 3}`. Pairwise XORs: `1^1 = 0`, `1^3 = 2`, `3^3 = 0`, so `P = {0, 2}`. Triplet XORs: `0^1 = 1`, `0^3 = 3`, `2^1 = 3`, `2^3 = 1`, so `T = {1, 3}`. Count is `2`. ✅

## Complexity Analysis

Let `V = 2048` be the value bound (next power of two above the max XOR) and `d = |D| <= 1500`.

Time Complexity: `O(d^2 + V * d)` — the pairwise stage dominates at roughly `d^2` (about `2.25 * 10^6`), and the triplet stage is `V * d`. Both are well within limits.
Space Complexity: `O(V + d)` — two boolean arrays of size `V` for `P` and `T`, plus the distinct-value list.

## Edge Cases

- **Single element (`n == 1`):** the only triplet is `a ^ a ^ a = a`, so the answer is `1`. The distinct set is `{a}`, `P = {0}`, and `T = {a}` — handled naturally.
- **All identical elements:** collapses to a single distinct value, same as the `n == 1` case; answer is `1`.
- **Two elements:** `P = {0, a^b}` and `T = {a, b}`, giving `2` distinct triplet values (matches Example 1). Confirms the ordering constraint imposes no real restriction.
- **Values that XOR to `0`:** `0` is always reachable via `a ^ a ^ a`... but note `0` only appears in `T` if some `d` cancels a `p`; for a single distinct value the smallest result is `a` itself, so make sure the sizing (2048) covers the largest possible XOR, not just the largest input.
