# 1358. Number of Substrings Containing All Three Characters

[LeetCode Link](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/)

Difficulty: Medium
Topics: Hash Table, String, Sliding Window
Acceptance Rate: 74.0%

## Hints

### Hint 1

The brute force is to check every substring and test whether it contains `a`, `b`, and `c`. That is O(n²) substrings (and more if you re-scan each one), which is too slow for `n` up to 5×10⁴. Whenever you are counting substrings over a string and the constraint is in the tens of thousands, ask whether a **sliding window** over the string can do the counting in one pass.

### Hint 2

Instead of trying to count valid substrings directly, fix the **left endpoint** and think about all the right endpoints that make a valid substring. Key observation: if the window `s[left..right]` already contains all three characters, then *every* extension of that window further to the right (`right+1`, `right+2`, … up to the end) is also valid. This "suffix is free" property is the heart of the trick.

### Hint 3

Walk a pointer `right` from left to right and maintain the position of the **most recent** `a`, `b`, and `c` seen so far. For each `right`, the substrings ending at `right` that contain all three are exactly those whose left endpoint is at or before `min(lastA, lastB, lastC)`. So `min(lastA, lastB, lastC) + 1` counts how many valid substrings end at this `right`. Sum that over every `right` and you have the answer in a single O(n) pass.

## Approach

There are two clean linear-time framings; both are worth knowing.

**Framing A — "last seen" positions (count substrings ending at each index).**

Keep three variables `last[0]`, `last[1]`, `last[2]` holding the most recent index where `a`, `b`, `c` appeared (initialized to `-1` meaning "not seen yet"). Iterate `right` from `0` to `n-1`:

1. Update `last[s[right]-'a'] = right`.
2. The substrings ending at `right` that contain all three characters are those starting at any index `left` in the range `0 .. min(last[0], last[1], last[2])`. There are exactly `min(last[0], last[1], last[2]) + 1` such starting positions — but only once all three have been seen (otherwise the minimum is `-1` and `min+1 = 0`, contributing nothing).
3. Add `min(last[0], last[1], last[2]) + 1` to the running total.

Because the minimum of the three last-seen positions is the latest point at which the *rarest-recently* character appeared, any left endpoint at or before it still includes all three characters within `[left, right]`.

Worked example on `s = "abcabc"`:

| right | char | last(a,b,c) | min+1 | running total |
|-------|------|-------------|-------|---------------|
| 0     | a    | (0,-1,-1)   | 0     | 0             |
| 1     | b    | (0,1,-1)    | 0     | 0             |
| 2     | c    | (0,1,2)     | 1     | 1             |
| 3     | a    | (3,1,2)     | 2     | 3             |
| 4     | b    | (3,4,2)     | 3     | 6             |
| 5     | c    | (3,4,5)     | 4     | 10            |

The total is **10**, matching the expected output.

**Framing B — shrinking sliding window (count substrings starting at each left).**

Maintain a count of `a`, `b`, `c` inside a window `[left, right]`. Expand `right` one step at a time. Whenever the window contains all three characters, every substring from the current `left` that ends at `right` or beyond is valid, so add `n - right` to the answer, then shrink from the left (decrementing counts and advancing `left`) and repeat the check. This also runs in O(n) because `left` and `right` each advance at most `n` times.

Both framings produce the same answer; Framing A is the most compact to implement, so the provided solution uses it.

## Complexity Analysis

Time Complexity: O(n) — a single pass over the string, with O(1) work per character.
Space Complexity: O(1) — only three integer positions (or three counts) are tracked, independent of input size.

## Edge Cases

- **Minimum-length string that is exactly `"abc"`**: should return `1`. The constraint guarantees length ≥ 3, so the smallest valid input still has exactly one qualifying substring.
- **No valid substring** (e.g. `"aaa"`, `"aabb"`, `"ccca"`): the minimum last-seen position stays `-1` throughout, so the answer is `0`. Initializing the last-seen positions to `-1` is what makes this fall out correctly.
- **A character appears only at the very end** (e.g. `"aaacb"`): valid substrings only begin once all three have appeared; the running count stays `0` until then, yielding `3`.
- **Long runs of a single character** (e.g. `"aaaaabc"`): the `min+1` formula naturally accounts for the many possible left endpoints once `b` and `c` arrive, without any special-casing.
- **Result magnitude**: for `n = 5×10⁴`, the count can approach `n²/2 ≈ 1.25×10⁹`, which exceeds 32-bit `int` range. Go's `int` is 64-bit on common platforms, so using `int` is safe here, but it is worth being mindful of overflow in fixed-width languages.
