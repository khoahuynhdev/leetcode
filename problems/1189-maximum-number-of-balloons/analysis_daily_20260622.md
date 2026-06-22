# 1189. Maximum Number of Balloons

[LeetCode Link](https://leetcode.com/problems/maximum-number-of-balloons/)

Difficulty: Easy
Topics: Hash Table, String, Counting
Acceptance Rate: 61.4%

## Hints

### Hint 1

The order of the characters in `text` does not matter — you are only allowed to *reuse* each character once to build copies of a fixed target word. When order is irrelevant and you only care about *how many of each item you have*, think about **counting** rather than scanning or two pointers.

### Hint 2

Count how many times each letter appears in `text`. Then compare those counts against the letters needed for one copy of `"balloon"`. The answer is limited by whichever required letter "runs out" first.

### Hint 3

The word `"balloon"` needs the letters `b, a, n` once each and `l, o` **twice** each. So for `l` and `o`, the number of complete copies they can support is `count / 2` (integer division). The maximum number of balloons is the **minimum** over all five required letters of (available count ÷ needed count).

## Approach

This is a classic frequency-counting problem.

1. **Tally `text`.** Walk through `text` once and build a frequency table of all 26 lowercase letters. An array of size 26 (indexed by `c - 'a'`) is the cleanest container, but a map works equally well.

2. **Identify the requirement.** The target word `"balloon"` is composed of:
   - `b` × 1
   - `a` × 1
   - `l` × 2
   - `o` × 2
   - `n` × 1

   Only these five distinct letters matter; every other letter in `text` is irrelevant.

3. **Bottleneck = minimum ratio.** Each complete copy of `"balloon"` consumes the letters above. If you have, say, 5 `l`'s, you can only contribute `5 / 2 = 2` copies from the `l` supply. Apply this for every required letter and take the smallest result — that letter is the bottleneck that caps how many full words you can spell.

   - For `b`, `a`, `n`: copies supported = `count`
   - For `l`, `o`: copies supported = `count / 2`

   Answer = `min(count[b], count[a], count[l]/2, count[o]/2, count[n])`.

**Example:** `text = "loonbalxballpoon"`.
Counts: `b=2, a=2, l=4, o=4, n=2` (the `x`, `p` are ignored).
- `b`: 2, `a`: 2, `n`: 2
- `l`: 4/2 = 2, `o`: 4/2 = 2

The minimum is `2`, so the answer is `2`. ✅

Why it works: spelling each copy is independent and consumes a fixed letter cost, so the total is purely supply-limited by the scarcest required letter relative to its per-word demand.

## Complexity Analysis

Time Complexity: O(n), where n is the length of `text` — a single pass to count, then O(1) work to compare the five required letters.
Space Complexity: O(1) — a fixed-size table of 26 counters regardless of input size.

## Edge Cases

- **No matching letters / missing a required letter** (e.g. `"leetcode"`): if any of `b, a, l, o, n` is absent, its count is 0, the minimum is 0, and the answer is correctly `0`.
- **The double letters `l` and `o`**: the most common mistake is treating every letter as needing one occurrence. Forgetting the integer division by 2 for `l` and `o` produces wrong (too-large) answers.
- **Exactly enough letters for one word** (e.g. `"balloon"` or `"nlaebolko"`): all ratios equal 1, answer is 1 — verifies the boundary.
- **Extra irrelevant characters**: letters not in `"balloon"` should not affect the count; the algorithm naturally ignores them.
- **Short input** (length 1): cannot form a word, returns 0.
