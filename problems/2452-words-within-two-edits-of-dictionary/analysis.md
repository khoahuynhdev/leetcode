# 2452. Words Within Two Edits of Dictionary

[LeetCode Link](https://leetcode.com/problems/words-within-two-edits-of-dictionary/)

Difficulty: Medium
Topics: Array, String, Trie

Acceptance Rate: 65.8%

## Hints

### Hint 1

The problem says "change any letter to any other letter" — notice that you're only
substituting characters in place, never inserting or deleting. All words have the
same length, so the number of edits between two words is simply the number of
positions where they differ. What well-known string distance does this describe?

### Hint 2

This is exactly the **Hamming distance** between two equal-length strings. For
each query, you want to know: does there exist some dictionary word whose Hamming
distance to the query is at most 2? You don't actually need to compute the full
distance — you can short-circuit as soon as you've seen 3 mismatches.

### Hint 3

Given the small constraints (at most 100 queries × 100 dictionary words × 100
characters = 10^6 comparisons), a direct pairwise scan with early exit is fast
enough and easy to get right. A Trie with a DFS that carries an "edit budget"
can share the work of comparing against many dictionary words simultaneously,
but it is overkill here. Focus on correctness: count mismatches, bail out at 3,
and remember to preserve the original order of queries in the output.

## Approach

Use a straightforward "Hamming-distance with early termination" strategy:

1. For each word `q` in `queries`, iterate over every word `d` in `dictionary`.
2. Walk the two strings index by index, counting how many positions differ.
3. As soon as the mismatch count exceeds 2, stop comparing — this `d` is not a
   match, move on to the next dictionary word.
4. If any `d` yields a final mismatch count of 2 or less, append `q` to the result
   and break out of the inner loop (we only need one matching dictionary word per
   query).
5. Return the collected results. Because we iterate `queries` in order, the output
   naturally preserves the original query order.

Why it works: all words share the same length `n`, so the minimum number of
single-letter edits transforming one into the other is exactly the count of
positions where their letters differ. We're not allowed insertions or deletions,
so this is Hamming distance, not Levenshtein — a key simplification.

Worked example: for `queries = ["word","note","ants","wood"]` and
`dictionary = ["wood","joke","moat"]`:
- `"word"` vs `"wood"`: differ only at index 2 (`r` vs `o`) → 1 edit, accept.
- `"note"` vs `"wood"`: 4 differences, reject. vs `"joke"`: positions 0 and 2 differ
  (`n` vs `j`, `t` vs `k`) → 2 edits, accept.
- `"ants"` vs each dictionary word: 4, 4, 3 differences → reject.
- `"wood"` vs `"wood"`: 0 edits, accept.

## Complexity Analysis

Let `Q = len(queries)`, `D = len(dictionary)`, and `n` be the common word length.

Time Complexity: O(Q · D · n) — each pair is compared in O(n), with early exit
bringing the constant down in practice but not the worst-case bound.

Space Complexity: O(Q · n) for the output in the worst case (every query matches).
The algorithm itself uses O(1) auxiliary space beyond the input and output.

## Edge Cases

- **Word length 1**: any single-character query matches any dictionary word (at
  most 1 edit), so all queries should be returned. The mismatch counter never
  reaches 3, so the algorithm handles this naturally.
- **Exact match (0 edits)**: the query equals some dictionary word. Mismatch
  count is 0, clearly ≤ 2, so it is included.
- **Exactly 2 edits**: the boundary case — must be included. The early-exit guard
  should use strict `>` 2, not `>=` 2.
- **Exactly 3 edits**: must be excluded. Good coverage here protects against
  off-by-one errors in the threshold.
- **Empty result**: no query is close enough to any dictionary word. Return an
  empty slice (not `nil`) for idiomatic Go and cleaner JSON-style equality in
  tests.
- **Duplicate or repeated queries**: each is evaluated independently; duplicates
  in the input produce duplicates in the output, preserving the given order.
