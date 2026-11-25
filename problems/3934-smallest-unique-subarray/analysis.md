# 3934. Smallest Unique Subarray

[LeetCode Link](https://leetcode.com/problems/smallest-unique-subarray/)

Difficulty: Hard
Topics: Array, Hash Table, Binary Search, Rolling Hash, Suffix Array, Hash Function
Acceptance Rate: 55.2%

## Hints

### Hint 1

Checking one fixed length `L` is fairly easy: list every window of length `L` and count how often each one appears. The hard part is that you can't afford to try every `L` one by one when each window costs `O(L)` to compare. Is there a property of `L` that lets you skip most of the lengths?

### Hint 2

Suppose a window of length `L` appears exactly once. What happens when you make it one element longer? The answer is **monotonic** in `L`, so you can **binary search** on the length. To check one length in linear time, you need a way to compare windows in `O(1)`. That's what a **rolling (polynomial) hash** gives you.

### Hint 3

Any occurrence of a longer subarray also contains an occurrence of its shorter prefix. So if `nums[i..i+L-1]` is unique, then `nums[i..i+L]` (or `nums[i-1..i+L-1]` at the right edge) is unique too. That makes "a unique window of length `L` exists" a false…false, true…true predicate. Binary search for the first `true`. For each candidate `L`, use prefix hashes to get every window's hash in `O(1)`, count the hashes in a map, and check whether any count is exactly 1. Use two independent hashes so collisions are effectively impossible.

## Approach

1. **Monotonicity.** Let `P(L)` = "some subarray of length `L` occurs exactly once."
   - `P(n)` is always true, since the whole array occurs once.
   - If `P(L)` is true for a window `w` and `L < n`, extend `w` by one neighbouring element to get `w'`. Every occurrence of `w'` contains an occurrence of `w` at the same offset. `w` occurs only once, so `w'` does too. So `P(L+1)` is true.

   This means we can binary search for the smallest `L` with `P(L)` true, over `[1, n]`.

2. **Checking `P(L)` in O(n) with rolling hashes.**
   - Precompute prefix hashes `pre[i+1] = pre[i] * B + nums[i]` and powers `pow[k] = B^k`.
   - The hash of `nums[i..i+L-1]` is `pre[i+L] - pre[i] * pow[L]`.
   - Put all `n-L+1` window hashes into a map of counts. If any count is 1, `P(L)` holds.
   - To make collisions negligible, use a **double hash**: one hash mod 2^64 (free overflow on `uint64`) and one mod `1e9+7`. The map key is the pair.

3. **Example:** `nums = [1,1,2,2,1]`
   - `L = 3` (first midpoint of `[1,5]`): windows `112, 122, 221` are all unique, so `P(3)` is true and `hi = 3`.
   - `L = 2`: windows `11, 12, 22, 21` are all unique, so `P(2)` is true and `hi = 2`.
   - `L = 1`: `1` appears 3 times and `2` appears 2 times, so `P(1)` is false and `lo = 2`.
   - The answer is **2**.

A suffix array with an LCP array also works. A suffix starting at `i` gives a unique substring of length `max(lcp[prev], lcp[next]) + 1` if that length fits inside the array. Taking the minimum over all suffixes gives the answer in `O(n log n)`. The binary search + hashing version is much easier to write in an interview.

## Complexity Analysis

Time Complexity: O(n log n). That's `log n` binary-search steps, each doing O(n) expected hash-map work.
Space Complexity: O(n) for the prefix hashes, powers, and the count map.

## Edge Cases

- **Single element** (`[7]`): the only subarray is the whole array, so the answer is 1.
- **All elements equal** (`[3,3,3]`): every shorter window repeats, so the answer is `n`. This checks that the upper bound `hi = n` is reachable.
- **All distinct**: the answer is 1. The binary search must still go all the way down to the lower bound.
- **Periodic arrays** (`[1,2,1,2,1,2]`): many windows repeat, which makes a good test of the hashing. Here the answer is 4 (`2121` appears once).
- **Hash collisions**: a single hash could merge two different windows and hide a unique one. A double hash (or a verification step) prevents this.
- **Large values** (up to 1e5): values go straight into the polynomial, so use 64-bit arithmetic and reduce mod properly when subtracting (add `mod` before taking `%`).
