# 3702. Longest Subsequence With Non-Zero Bitwise XOR

[LeetCode Link](https://leetcode.com/problems/longest-subsequence-with-non-zero-bitwise-xor/)

Difficulty: Medium
Topics: Array, Bit Manipulation
Acceptance Rate: 39.3%

## Hints

### Hint 1

The word "subsequence" makes this look like a selection or DP problem, but resist that pull. Ask yourself the greedy question first: what is the *best case* answer, and what would have to go wrong to prevent it? The whole array is a valid subsequence — start by checking whether it already works.

### Hint 2

XOR has two properties that do all the heavy lifting here: `x ^ x == 0` and `x ^ 0 == x`. Together they mean that if you take the XOR of the entire array and then remove one element `v` from the subsequence, the new XOR is the old XOR combined with `v` again. So removing elements is not a search — it's an algebraic adjustment you can reason about directly.

### Hint 3

Let `X` be the XOR of all `n` elements. If `X != 0`, the answer is `n` — you're already done. If `X == 0`, then for *any* subset you keep, `XOR(kept) == XOR(removed)` (because the two halves must XOR back to `0`). So dropping exactly one element `v` leaves a subsequence whose XOR is `v`. Pick any non-zero `v` and you get `n - 1`. The only way to fail is if there is no non-zero element at all.

## Approach

The problem collapses to three cases, decided by a single pass over the array.

**Step 1 — compute the total XOR.** Let `X = nums[0] ^ nums[1] ^ ... ^ nums[n-1]`.

**Step 2 — if `X != 0`, return `n`.** The entire array is itself a subsequence, and its XOR is non-zero. No longer subsequence exists, so this is optimal. Example 2 is exactly this case: `2 ^ 3 ^ 4 = 5 != 0`, so the answer is `3`.

**Step 3 — if `X == 0`, try dropping one element.** Suppose we keep a sub-multiset `K` and remove the rest, `R`. Every element appears in exactly one of the two, so `XOR(K) ^ XOR(R) = X = 0`, which means `XOR(K) == XOR(R)`. That is a useful restatement: *the XOR of what you keep equals the XOR of what you throw away.*

Now remove a single element `v`. Then `R = {v}` and `XOR(K) = v`. So the kept subsequence has non-zero XOR precisely when `v != 0`. Since a length-`n` answer is impossible (that would need `X != 0`), `n - 1` is the best achievable length, and it is achievable as long as at least one non-zero element exists.

Example 1 illustrates it: `nums = [1,2,3]`, `X = 1 ^ 2 ^ 3 = 0`. Drop the `1` and the remaining `[2,3]` XORs to `1` — length `2 = n - 1`. Equivalently, drop `2` to get `[1,3]` with XOR `2`, or drop `3` to get `[1,2]` with XOR `3`. Any non-zero element works, which is why the problem says "one longest subsequence" rather than "the".

**Step 4 — if `X == 0` and every element is zero, return `0`.** Any subsequence of all-zeros XORs to `0`, including the empty one, so no valid subsequence exists.

Note that step 3 and step 4 are complementary: when `X == 0`, either some element is non-zero (answer `n - 1`) or all are zero (answer `0`). Both checks fold into the same single loop that computes `X`, so the implementation is one pass with two accumulators.

This is a Medium mostly because of the misdirection — the subsequence framing invites an exponential search or a DP over XOR values, and the acceptance rate reflects how many people go down that road. Once you see the "keep XOR equals removed XOR" identity, the code is about five lines. That gap between apparent and actual difficulty is the lesson worth keeping.

## Complexity Analysis

Time Complexity: O(n) — a single pass to accumulate the running XOR and to detect whether any non-zero element exists.
Space Complexity: O(1) — only a couple of scalar accumulators, independent of input size.

## Edge Cases

- **All elements are zero** (`[0]`, `[0,0,0]`): total XOR is `0` and no element can be dropped to fix it. Must return `0`, not `n - 1`. This is the single case that returns `0` and the one most likely to be missed.
- **Total XOR is already non-zero** (`[2,3,4]`): return `n` directly. Do not reflexively subtract one.
- **Zeros mixed with non-zeros** (`[0,0,5]`): the zeros do not affect the XOR at all, so `X = 5 != 0` and the answer is the full length `3`. Zeros are only fatal when they are *all* you have.
- **Single element** (`n == 1`): `[0]` returns `0`, any `[v]` with `v != 0` returns `1`. Both fall out of the general logic — no special-casing needed, but worth testing.
- **Duplicate pairs that cancel** (`[7,7]`): `X = 0` and a non-zero element exists, so the answer is `1`. A good check that you are not accidentally deduplicating.
- **Large input** (`n` up to `10^5`): any solution that enumerates subsets or builds a DP table over XOR values (up to `2^30` for `nums[i] <= 10^9`) will time out or blow memory. The O(n) pass is not just elegant, it is required.
