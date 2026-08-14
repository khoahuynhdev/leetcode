# 3090. Maximum Length Substring With Two Occurrences

[LeetCode Link](https://leetcode.com/problems/maximum-length-substring-with-two-occurrences/)

Difficulty: Easy
Topics: Hash Table, String, Sliding Window
Acceptance Rate: 68.1%

## Hints

### Hint 1

The answer is a *contiguous* substring, and the property being tested ("every character appears at most twice") is **monotone**: if a window is valid, every window inside it is also valid; if a window is invalid, every window containing it is also invalid. Whenever you see that shape, brute-forcing all O(n²) substrings is leaving something on the table — think about a window with two moving ends instead.

### Hint 2

Grow a window from the left with two indices, `left` and `right`. Push `right` forward one character at a time and maintain a frequency count of everything currently inside the window. The only thing that can break validity when you add a character is *that one character* — no other count changed. So you only ever need to check one condition after each step.

### Hint 3

When adding `s[right]` pushes its count to 3, the window is invalid, and shrinking from the left is the only way to fix it. Keep advancing `left` (decrementing counts as characters leave) until that offending character's count drops back to 2. Because the counts are already maintained, you don't need to rescan the window. Record `right - left + 1` after every restore, and `left` never moves backwards — that's what makes the whole scan linear.

## Approach

This is the canonical **sliding window with a frequency map** pattern, specialized to "at most k occurrences of each character" with `k = 2`.

Maintain:

- `count[26]`: how many times each lowercase letter currently appears in the window `s[left..right]`.
- `left`: the start of the current window.
- `best`: the longest valid window length seen so far.

Then sweep `right` from `0` to `n-1`:

1. **Expand.** Increment `count[s[right]]`. The window `s[left..right]` may now be invalid, but only because of `s[right]` — every other letter's count is unchanged from the previous (valid) window.
2. **Restore.** While `count[s[right]] > 2`, decrement `count[s[left]]` and advance `left`. This loop is guaranteed to terminate at the right moment: `s[right]` occurs at least three times in the window, so walking `left` forward will eventually remove one of those occurrences and bring the count back to exactly 2.
3. **Record.** The window is valid again, so update `best = max(best, right-left+1)`.

Why an array instead of a `map[byte]int`? The alphabet is fixed at 26 lowercase letters, so a plain array gives O(1) lookups with no hashing and no allocation. A map would work identically but is strictly more overhead here.

**Walkthrough on `s = "bcbbbcba"`** (indices `b0 c1 b2 b3 b4 c5 b6 a7`):

| right | char | action | window | length |
|---|---|---|---|---|
| 0 | b | b→1 | `b` | 1 |
| 1 | c | c→1 | `bc` | 2 |
| 2 | b | b→2 | `bcb` | 3 |
| 3 | b | b→3, shrink past index 0 | `cbb` | 3 |
| 4 | b | b→3, shrink past 1 and 2 | `bb` | 2 |
| 5 | c | c→1 | `bbc` | 3 |
| 6 | b | b→3, shrink past 3 | `bcb` | 3 |
| 7 | a | a→1 | `bcba` | **4** |

The answer is 4, matching the expected output. Note that `left` only ever moved forward — across the whole run it advanced 4 times total, which is why the nested `while` doesn't make this quadratic.

**Amortized cost:** the inner shrink loop looks like a nested loop, but `left` is monotonically non-decreasing and bounded by `n`, so the total work across all iterations of the outer loop is O(n).

Honest difficulty note: this is a genuinely easy problem *if* you already recognize the variable-size sliding window template. If you don't, it is an excellent one to use for building that template, because the invalidating condition is as simple as it gets — a single character crossing a single threshold. Generalizing your solution to "at most k occurrences" is a one-character change, and the same skeleton solves problems like [3](https://leetcode.com/problems/longest-substring-without-repeating-characters/) (k = 1) and [1004](https://leetcode.com/problems/max-consecutive-ones-iii/).

## Complexity Analysis

Time Complexity: O(n) — `right` advances n times and `left` advances at most n times in total, so each character is added and removed from the window at most once.
Space Complexity: O(1) — a fixed 26-element frequency array, independent of input length.

## Edge Cases

- **Every character identical (`"aaaa"`)** — the window can never exceed length 2. This exercises the shrink loop firing on essentially every step and confirms `left` chases `right` correctly instead of stalling.
- **Already-valid whole string (`"aabbcc"`, `"abcdef"`)** — the shrink loop never runs, and the answer is the full length `n`. A solution that unconditionally shrinks, or that updates `best` before restoring validity, breaks here.
- **Minimum length input (`n = 2`)** — the constraints guarantee `2 <= s.length`, so the answer is always at least 1 and there is no empty-string case to defend against. Still worth testing that a two-character string returns 2 (`"ab"`, `"aa"`) rather than an off-by-one 1.
- **The optimal window appears at the very end (`"bcbbbcba"`)** — the best answer is only found on the final iteration. Updating `best` only when the window shrinks, instead of on every iteration, would miss it.
- **Update order matters** — `best` must be computed *after* the shrink loop. Measuring the window while it still holds three copies of a character overcounts.
- **Repeating cycles (`"abcabcabc"`)** — each letter appears three times, so no valid window covers the whole string, but a length-6 window exists in several positions. This catches solutions that shrink too aggressively (e.g. resetting `left = right` instead of advancing it one step at a time).
