# 3016. Minimum Number of Pushes to Type Word II

[LeetCode Link](https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-ii/)

Difficulty: Medium
Topics: Hash Table, String, Greedy, Sorting, Counting
Acceptance Rate: 80.9%

## Hints

### Hint 1

The order of characters in `word` is completely irrelevant. Typing `"abab"` costs exactly the same as typing `"aabb"` — only *how many times* each letter appears matters. That should immediately push you toward counting: build a frequency table over the 26 lowercase letters and forget the string itself.

### Hint 2

Once you have frequencies, the problem becomes an assignment problem: hand out 26 letters to 8 keys. A letter placed in the *k*-th slot of a key costs *k* pushes every time it is typed, so a letter with frequency `f` in slot `k` contributes `f * k` to the total. You want to pair large frequencies with small slot numbers. That pairing instinct is the classic greedy exchange argument — and it needs the frequencies sorted.

### Hint 3

Here is the insight that makes it collapse into a few lines: there is never a reason to put a second letter on a key before every key has a first letter. Cost slots become available in *tiers of 8* — the eight cheapest slots (first position on each of the 8 keys) all cost 1, the next eight all cost 2, and so on. So sort frequencies in descending order; the letter at index `i` (0-based) lands in a slot costing `i/8 + 1`. The answer is `sum(freq[i] * (i/8 + 1))`. Since there are only 26 letters, at most 4 tiers are ever used and no key ever holds more than 4 letters.

## Approach

**Step 1 — Count.** Scan `word` once and tally occurrences into a fixed array `counts[26]`. The string can be up to 10^5 characters, but the tally is bounded at 26 entries, which is what makes the rest cheap.

**Step 2 — Sort descending.** Sort those 26 counts from largest to smallest. This is a constant-size sort, so it costs nothing asymptotically. Letters that never appear end up as zeros at the tail and contribute nothing regardless of where they are assigned, so they need no special handling.

**Step 3 — Assign in tiers of 8.** Walk the sorted counts. For index `i`, the push cost per occurrence is `i/8 + 1`:

- indices 0–7 → cost 1 (first slot on keys 2 through 9)
- indices 8–15 → cost 2 (second slot on each key)
- indices 16–23 → cost 3
- indices 24–25 → cost 4

Accumulate `counts[i] * (i/8 + 1)`.

**Why the greedy is optimal.** Suppose an optimal mapping puts a letter with frequency `f1` at cost `c1` and a letter with frequency `f2` at cost `c2`, where `f1 > f2` but `c1 > c2`. Swapping them changes the total by `(f1*c2 + f2*c1) - (f1*c1 + f2*c2) = (f1 - f2)(c2 - c1) < 0`, a strict improvement — contradiction. So in any optimal solution the frequencies must be sorted in the same order as the costs. And the multiset of available costs is fixed and forced: to use a slot costing `k` on some key, that key must already hold `k-1` letters, so the cheapest achievable cost multiset is exactly eight 1s, then eight 2s, then eight 3s, then eight 4s. Pairing the sorted frequencies against that sorted cost list is the minimum.

**Worked example.** `word = "aabbccddeeffgghhiiiiii"` gives counts `a..h = 2` each and `i = 6`. Sorted descending: `[6, 2, 2, 2, 2, 2, 2, 2, 2, 0, ...]`. The first eight entries (`6,2,2,2,2,2,2,2`, summing to 20) cost 1 each → 20. The ninth entry (`2`) falls into the second tier at cost 2 → 4. Total **24**, matching the expected output. Notice the greedy naturally put the heavy `i` (frequency 6) on a first slot and pushed one of the cheap `h`s (frequency 2) into the second slot — exactly the mapping the problem's explanation describes.

This is a genuinely approachable Medium once you see the tier structure; the difficulty is almost entirely in trusting that sorting plus integer division is the whole algorithm.

## Complexity Analysis

Time Complexity: O(n), where n is the length of `word` — one pass to count, plus a sort of a fixed 26-element array (O(26 log 26) = O(1)).

Space Complexity: O(1) — a fixed array of 26 counters, independent of input size.

## Edge Cases

- **Single character (`word = "a"`)**: the minimum length allowed by the constraints. Only one nonzero count, index 0, cost 1 → answer 1. Guards against off-by-one errors in the tier formula (`i/8 + 1` must yield 1, not 0, at `i = 0`).
- **Exactly 8 distinct letters**: fills the first tier perfectly, and the answer equals `len(word)`. The 9th index is where the cost first jumps, so this is the boundary worth testing on both sides (8 distinct vs. 9 distinct).
- **All 26 distinct letters**: uses all four tiers, with the last tier only partially filled (indices 24 and 25). Confirms the loop handles a ragged final tier without special-casing.
- **Letters with zero frequency**: 26 slots always exist in the array but most words use fewer. Zeros must be sorted to the end (descending order does this for free) — if you sorted ascending and forgot to reverse, zeros would occupy the cheap slots and inflate the answer badly.
- **Highly skewed frequencies (`"aaaa...a"`)**: one letter with a huge count. Verifies that the accumulator is sized correctly; with n up to 10^5 and a max cost of 4, the total fits comfortably in an `int`, but the reflex of checking overflow bounds is worth keeping.
- **Repeated characters in non-adjacent positions (`"abab"` vs `"aabb"`)**: both must return the same value, reinforcing that position in the string is irrelevant.
