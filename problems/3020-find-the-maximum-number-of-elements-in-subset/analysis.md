# 3020. Find the Maximum Number of Elements in Subset

[LeetCode Link](https://leetcode.com/problems/find-the-maximum-number-of-elements-in-subset/)

Difficulty: Medium
Topics: Array, Hash Table, Enumeration
Acceptance Rate: 31.1%

## Hints

### Hint 1

The pattern `[x, x^2, x^4, ..., x^k, ..., x^4, x^2, x]` is a palindrome. Every value except the single peak appears an even number of times (exactly twice in the chain), so what really matters is *how many copies of each value you have available*. Think about which data structure lets you look up "how many of value `v` do I have?" instantly.

### Hint 2

You don't need to consider arbitrary subsets. Fix a *base* `x`, and the rest of the chain is forced: `x`, then `x^2`, then `(x^2)^2 = x^4`, and so on — each step is the square of the previous value. So enumerate each distinct value as a potential base, walk the squaring chain upward, and count how far you can go.

### Hint 3

Walk up the chain `x -> x^2 -> x^4 -> ...` as long as the current value appears **at least twice** (it sits on both the left and right halves of the palindrome). When you finally reach a value that appears fewer than twice:
- if it appears **exactly once**, it can serve as the unique peak — add 1 to your count;
- if it appears **zero** times, the chain stops one step short, so the last fully-paired value must instead become the (single) peak — subtract 1.

The value `1` is special: `1^2 = 1`, so the whole chain collapses to a run of `1`s. The answer for base `1` is simply the largest **odd** number `<= count(1)`.

## Approach

Because the arrangement is a palindrome built by repeated squaring, the subset is completely determined once you choose the base value `x`. Every value strictly below the peak appears twice (once on each side), and the peak appears exactly once. That gives a clean counting strategy.

1. **Count occurrences.** Build a hash map `freq` from value to how many times it appears in `nums`.

2. **Handle `x = 1` separately.** Since `1^2 = 1`, any chain rooted at `1` is just a block of `1`s of odd length (`[1]`, `[1,1,1]`, `[1,1,1,1,1]`, ...). With `c = freq[1]` ones available, the best odd length is `c` if `c` is odd, otherwise `c - 1`.

3. **Enumerate every other distinct base `x > 1`.** Starting at `cur = x`, repeatedly:
   - while `freq[cur] >= 2`, add `2` to the running length (this value contributes a left copy and a right copy), then move up: `cur = cur * cur`.
   - When the loop stops, inspect `cur`: if `freq[cur] == 1`, this value is the peak, so add `1`; otherwise (`freq[cur] == 0`) the chain ran out, so subtract `1` to convert the previously-paired top value into a single peak.
   - Track the maximum length over all bases.

Note that even a base whose count is `1` yields a valid length of `1` (the trivial subset `[x]`), so the answer is always at least `1`.

**Why squaring never overflows:** values are at most `10^9`, so `cur * cur` is at most `10^18`, which fits comfortably in a 64-bit integer. Once `cur` exceeds `10^9` it cannot appear in `nums`, so `freq[cur]` is `0` and the loop stops before squaring again.

**Walkthrough (Example 1, `nums = [5,4,1,2,2]`):** `freq = {5:1, 4:1, 1:1, 2:2}`.
- Base `2`: `freq[2] = 2` → length `2`, go to `4`. `freq[4] = 1 < 2`, stop; it appears once → peak, add `1`. Total `3` (the arrangement `[2,4,2]`).
- Bases `5`, `4` each give length `1`; base `1` gives `1`.
- Maximum is **3**.

## Complexity Analysis

Time Complexity: O(n log M) — there are at most `n` distinct bases, and each squaring chain has length `O(log M)` before the value exceeds the maximum element `M <= 10^9` (in practice the chain is very short). Each step is an O(1) hash-map lookup.
Space Complexity: O(n) for the frequency map.

## Edge Cases

- **The value `1`:** `1^2 = 1` breaks the "values grow" assumption, so it must be counted as a separate odd-length run. Forgetting this either undercounts or loops forever.
- **Even vs. odd count of `1`s:** a palindrome of `1`s must have odd length, so an even count must be reduced by one.
- **A base with no square present (`freq[cur] == 0` after the loop):** the chain stops short and the top paired value becomes a single peak — remember to subtract `1` rather than add.
- **All distinct values / no chain possible:** every base falls back to length `1`, so the answer is `1` (as in Example 2).
- **Large values near `10^9`:** ensure 64-bit arithmetic so `cur * cur` does not overflow.
