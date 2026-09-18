# 1520. Maximum Number of Non-Overlapping Substrings

[LeetCode Link](https://leetcode.com/problems/maximum-number-of-non-overlapping-substrings/)

Difficulty: Hard
Topics: Hash Table, String, Greedy, Sorting
Acceptance Rate: 50.0%

## Hints

### Hint 1

The string can be up to 10^5 characters, but the alphabet is only 26 letters. That gap is the whole problem: any answer is described by a small set of intervals, not by arbitrary substrings. Start by asking what the *smallest* valid substring containing a given letter looks like, and note that once you turn substrings into `[start, end]` intervals, this becomes a very familiar scheduling-style question.

### Hint 2

For a letter `c`, any valid substring that contains `c` must span at least `[first[c], last[c]]`. But pulling in that range may drag in other letters whose own occurrences reach further out, so the range has to grow. Keep growing until it is closed under "contains every occurrence of every letter inside it". You now have at most 26 candidate intervals — one per letter — and the original problem reduces to picking as many non-overlapping ones as possible.

### Hint 3

Two things make the whole thing click:

1. While expanding `[first[c], last[c]]` to the right, if you ever meet a letter whose *first* occurrence lies to the **left** of your start, then no valid substring can start at `first[c]` — discard this candidate entirely. You don't need to expand leftwards, because that longer interval will be produced by some other letter's candidate anyway.
2. The surviving candidates are **laminar**: any two are either disjoint or one fully contains the other (they can never partially overlap — a partial overlap would violate the closure property). So the classic interval-scheduling greedy of "sort by right endpoint, take every interval that starts after the last taken end" simultaneously maximizes the count *and* minimizes total length: whenever a big interval and a nested small one compete, the small one wins on the right endpoint.

## Approach

**Step 1 — first/last tables.** Scan `s` once and record, for each of the 26 letters, the index of its first and last occurrence. Letters absent from `s` are skipped.

**Step 2 — build one candidate interval per letter.** For a letter `c` present in `s`, set `start = first[c]` and `end = last[c]`, then walk `j` from `start` to `end` (with `end` growing as you go):

- If `first[s[j]] < start`, the interval would have to extend left of `start`; that contradicts starting at `first[c]`, so mark this candidate invalid and stop.
- Otherwise set `end = max(end, last[s[j]])`.

If the walk finishes, `[start, end]` is the **minimal** valid substring whose start is `first[c]`. Because the loop bound `end` only grows and `j` only moves forward, one candidate costs `O(n)` and all 26 cost `O(26n)`.

**Step 3 — greedy selection.** Sort the valid candidates by `end` ascending and sweep: keep a variable `prevEnd` (initially `-1`) and take a candidate whenever `start > prevEnd`, updating `prevEnd = end`. Emit `s[start : end+1]` for each taken candidate.

**Why the greedy is optimal on both criteria.** Candidates cannot partially overlap: if `[a1, b1]` and `[a2, b2]` overlapped with `a1 < a2 <= b1 < b2`, some letter occurring in `[a2, b1]` would have its last occurrence beyond `b1`, so the first interval would have been expanded past `b1` — contradiction. With only nesting or disjointness left, sorting by right endpoint means an inner interval is always considered before the outer one that contains it. Taking the inner one never costs you a substring (anything disjoint from the outer is also disjoint from the inner) and strictly reduces total length, which is exactly the tie-breaking rule the problem asks for.

**Worked example — `s = "adefaddaccc"`.** Candidates: `a → [0,7]` (valid, pulls in `d`), `d → invalid` (hits `a` at index 4 whose first occurrence 0 is left of start 1), `e → [2,2]`, `f → [3,3]`, `c → [8,10]`. Sorted by end: `[2,2], [3,3], [0,7], [8,10]`. The sweep takes `[2,2]` → `"e"`, `[3,3]` → `"f"`, skips `[0,7]` (starts at 0, before `prevEnd = 3`), takes `[8,10]` → `"ccc"`. Answer: `["e", "f", "ccc"]`.

## Complexity Analysis

Time Complexity: O(26 * n) = O(n) — one pass for the first/last tables, at most 26 expansion passes each bounded by `n`, and sorting at most 26 intervals (constant).
Space Complexity: O(1) auxiliary (two 26-entry tables plus at most 26 candidate intervals), excluding the O(n) total size of the returned substrings.

## Edge Cases

- **Single character (`"a"`)** — the shortest legal input; must return `["a"]`, not an empty list.
- **All characters identical (`"aaaa"`)** — exactly one candidate spanning the whole string, so the answer is one substring equal to `s`.
- **All characters distinct (`"abcd"`)** — every letter is its own candidate; the answer is `n` single-character substrings, the maximum possible.
- **Fully nested letters (`"abcba"`)** — candidates are `[0,4]`, `[1,3]`, `[2,2]`, each containing the next; only one can be chosen and it must be the innermost, so the answer is `["c"]`, not `["abcba"]`.
- **Interleaved letters (`"abab"`)** — expanding `a` reaches `b`'s last occurrence and vice versa, so only one candidate survives and the answer is the whole string; a solution that forgets to keep growing `end` inside the loop will wrongly return two substrings.
- **Candidates rejected by the left check (`"adefadda..."`, letter `d`)** — forgetting the `first[s[j]] < start` test produces intervals that are not actually closed, breaking the laminar property the greedy relies on.
- **Ties between a nested and an enclosing interval (`"abbaccd"`)** — both `["abba","cc","d"]` and `["bb","cc","d"]` have three substrings; sorting by `end` is what forces the shorter `"bb"`, satisfying the minimum-total-length requirement.
- **Long input (10^5 characters)** — the expansion loop must reuse the growing `end` rather than restarting, otherwise the per-letter pass degrades and the solution times out.
