# 3513. Number of Unique XOR Triplets I

[LeetCode Link](https://leetcode.com/problems/number-of-unique-xor-triplets-i/)

Difficulty: Medium
Topics: Array, Math, Bit Manipulation
Acceptance Rate: 34.9%

## Hints

### Hint 1

The brute force is to enumerate every triplet `(i, j, k)` with `i <= j <= k`, XOR the three values, and collect the results in a set. With `n` up to `10^5` that is on the order of `n^3` combinations — far too many. Before writing any loops, ask yourself: does the *specific structure* of the input (a permutation of `1..n`) let us skip the enumeration entirely? Think in terms of which bit patterns are reachable rather than which index triplets exist.

### Hint 2

The condition `i <= j <= k` allows indices to repeat. That is the key freedom: you can pick the same element twice. Since `x XOR x = 0`, choosing `nums[i] == nums[j]` collapses the triplet to a single element `nums[k]`. So every individual value is trivially reachable, and the real question becomes: what is the full set of values reachable as an XOR of *one or three* distinct elements drawn from `{1, 2, ..., n}`?

### Hint 3

Because the array always contains the numbers `1` and `2` (whenever `n >= 3`), you have enough "building blocks" to flip any bit combination. XOR-ing subsets of `{1, ..., n}` can produce every value from `0` up to the largest number representable with the same number of bits as `n`. That upper bound is `2^b - 1`, where `b` is the number of bits in `n`. So for `n >= 3` the answer is simply `2^b` (the count of values `0 .. 2^b - 1`). Only the tiny cases `n = 1` and `n = 2` behave differently and must be handled by hand.

## Approach

Observe that the constraint `i <= j <= k` permits repeated indices, so any element can appear an even number of times inside a triplet. Two equal picks cancel via XOR (`x XOR x = 0`), which means:

- Picking all three equal gives each single value `v`.
- Picking two equal and one different gives every single value as well.
- Picking three distinct gives genuine 3-way XORs.

So the reachable set is exactly the set of values obtainable by XOR-ing one or three chosen numbers from `{1, 2, ..., n}`.

Now use the small-value building blocks. For `n >= 3`, the numbers `1` and `2` (and more) are present, and XOR combinations of the numbers `1..n` cover the entire range `[0, 2^b - 1]`, where `b` is the bit length of `n`. Intuitively, with distinct powers-of-two-ish contributors you can toggle each of the `b` low bits independently, filling out every value with at most `b` bits. Therefore the number of unique triplet XOR values is `2^b`.

Handle the two degenerate cases explicitly:

- `n == 1`: the only element is `1`, so the only triplet is `1 XOR 1 XOR 1 = 1`. Answer is `1`.
- `n == 2`: elements are `{1, 2}`. Enumerating gives values `{1, 2}`. Answer is `2`.

For everything else, compute the bit length `b` of `n` and return `1 << b`.

**Worked example (`nums = [3,1,2]`, `n = 3`):** `n` in binary is `11`, so `b = 2` and the answer is `2^2 = 4`. Indeed the reachable values are `{0, 1, 2, 3}` — matching the expected output.

## Complexity Analysis

Time Complexity: O(log n) — a single pass over the bits of `n` to find its bit length. (We don't even need to read the array beyond its length.)
Space Complexity: O(1) — only a couple of integer counters.

## Edge Cases

- `n == 1`: Only one element (`1`); the sole triplet XORs to `1`. Returns `1`. Missing this special-case would over-count.
- `n == 2`: Only `1` bit's worth of numbers combined don't fill a full power-of-two range; the true answer is `2`, not `2^b`. Must be handled separately.
- `n == 3`: The first case where the general `2^b` formula kicks in (`b = 2`, answer `4`); a good sanity check that the boundary between special-cased and general logic is correct.
- Powers of two boundaries (e.g. `n = 4`, `n = 8`): the bit length jumps here, so the answer jumps to the next power of two. Verifying these confirms the bit-length computation is off-by-one free.
