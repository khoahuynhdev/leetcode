# 3524. Find X Value of Array I

[LeetCode Link](https://leetcode.com/problems/find-x-value-of-array-i/)

Difficulty: Medium
Topics: Array, Math, Dynamic Programming
Acceptance Rate: 50.0%

## Hints

### Hint 1

Removing a prefix and a non-overlapping suffix while keeping the array non-empty leaves exactly one contiguous block behind. So "number of ways to perform the operation" is really "number of non-empty subarrays" — each choice of `(prefix, suffix)` maps one-to-one to a subarray `nums[i..j]`. With `n` up to `10^5` there are about `5 * 10^9` subarrays, so you can never enumerate them; you need to count them in groups. Notice how tiny `k` is (`k <= 5`) — that is the real hint about what the state space should look like.

### Hint 2

You do not care about products, only about products **modulo k**. Since `(a * b) mod k = ((a mod k) * (b mod k)) mod k`, every subarray collapses to one of only `k` possible values. That suggests a sweep over the right endpoint `j`, carrying a small bucket of counts: for each residue `x`, how many subarrays *ending at j* have product `x mod k`? This is a classic "DP over the array where the state is an aggregate of the suffix, not an index" pattern — the same shape as counting subarrays by XOR or by GCD.

### Hint 3

Let `dp[x]` = the number of subarrays ending at the current index `j` whose product is congruent to `x (mod k)`. When you move from `j-1` to `j`, every one of those subarrays gets extended by `nums[j]`, so a subarray counted at residue `x` moves to residue `(x * nums[j]) % k`. Then add one brand-new subarray, the single element `[nums[j]]`, at residue `nums[j] % k`. Accumulate `dp` into the answer after processing each `j`. That is `O(n * k)` time with `O(k)` state — and the counts get large, so accumulate in 64-bit integers.

## Approach

**Step 1 — reframe the operation.** Choosing a prefix to delete and a non-overlapping suffix to delete, with something left over, is the same as choosing the surviving block: a non-empty subarray `nums[i..j]`. Different `(prefix, suffix)` pairs give different blocks and vice versa, so the answer for residue `x` is simply *the number of non-empty subarrays whose product is `x` mod `k`*. This reframing is most of the problem; once you see it, the rest is a standard counting sweep.

**Step 2 — work in the modular world.** Products of up to `10^5` values each as large as `10^9` are astronomically big, but modular arithmetic is multiplicative: reducing each element mod `k` up front and multiplying mod `k` as you go gives the same residue. Because `k <= 5`, there are at most 5 distinct residues, so all subarrays ending at a fixed index fall into at most 5 buckets.

**Step 3 — the rolling DP.** Sweep the right endpoint `j` from left to right and maintain `dp`, an array of length `k` where `dp[x]` counts subarrays ending at `j` with product `≡ x (mod k)`. Transition from `j-1` to `j`:

- Extend: each subarray counted in `dp[x]` becomes a subarray ending at `j` with residue `(x * nums[j]) % k`. So build a fresh `next` array and do `next[(x*v)%k] += dp[x]` for every residue `x`, where `v = nums[j] % k`.
- Start fresh: the length-1 subarray `[nums[j]]` is new at this step, so `next[v] += 1`.
- Then `dp = next` and add every `dp[x]` into `result[x]`.

A fresh `next` array (rather than updating `dp` in place) matters: in-place updates would let a subarray be multiplied by `nums[j]` twice when residues collide.

**Worked example.** `nums = [1,2,3,4,5]`, `k = 3`, residues `[1,2,0,1,2]`:

| j | nums[j] % k | dp after step (residue → count) | running result |
|---|---|---|---|
| 0 | 1 | `{1: 1}` | `[0,1,0]` |
| 1 | 2 | `{2: 2}` (`[1,2]` and `[2]`) | `[0,1,2]` |
| 2 | 0 | `{0: 3}` (everything hits the 0) | `[3,1,2]` |
| 3 | 1 | `{0: 3, 1: 1}` | `[6,2,2]` |
| 4 | 2 | `{0: 3, 2: 2}` | `[9,2,4]` |

Final answer `[9,2,4]`, matching the expected output — and note `9 + 2 + 4 = 15 = 5*6/2`, the total number of subarrays, which is a great self-check.

**Why it is correct.** Every non-empty subarray has exactly one right endpoint, so summing `dp` over all `j` counts each subarray exactly once. The transition is exact because `prod(nums[i..j]) mod k = (prod(nums[i..j-1]) mod k * nums[j] mod k) mod k`, so extending preserves the bucket structure, and the only subarray ending at `j` that is not an extension is `[nums[j]]` itself.

## Complexity Analysis

Time Complexity: O(n * k) — one pass over the array, with O(k) work per element (`k <= 5`, so effectively linear).
Space Complexity: O(k) — two rolling arrays of size `k` plus the size-`k` output (output excluded, the working space is O(1) given the constraint on `k`).

## Edge Cases

- **`k = 1`**: every product is `≡ 0 (mod 1)`, so the answer is the single value `n*(n+1)/2`. The DP handles this automatically since the only residue is 0, but it is the case most likely to overflow a 32-bit counter.
- **Counter overflow**: with `n = 10^5` the total is `5,000,050,000`, well past `2^31`. Counts must be 64-bit (`int64` in Go / `long long` in C++), which is why LeetCode's signature returns a 64-bit array.
- **Single element array**: exactly one way to operate (delete nothing), so the result has a single 1 at `nums[0] % k` and zeros elsewhere.
- **Elements far larger than `k`**: reduce with `nums[j] % k` before multiplying. Skipping the reduction risks overflow on the intermediate product for larger `k` and is wasted work here.
- **A zero residue in the middle**: once a subarray's residue hits 0, every extension stays 0 (since `0 * v ≡ 0`). This is not a special case in code, but it explains why `result[0]` is usually dominant — it is the trap in Example 1's explanation.
- **Residues that are never reachable**: e.g. Example 2 has `result[3] == 0`. The output array must still be full length `k` with explicit zeros, so allocate `make([]int64, k)` rather than building it sparsely.
