# 3093. Longest Common Suffix Queries

[LeetCode Link](https://leetcode.com/problems/longest-common-suffix-queries/)

Difficulty: Hard
Topics: Array, String, Trie
Acceptance Rate: 43.1%

## Hints

### Hint 1

You need to answer many independent queries against a fixed corpus. A naive scan
for each query is `O(|wordsQuery| * |wordsContainer| * L)`, which is far too slow
given the constraints. Think about which data structure speeds up repeated
string-matching against a fixed dictionary, especially when the matching is
character-by-character from one end of the string.

### Hint 2

The problem talks about *suffixes*, but most well-known string data structures
work naturally on *prefixes*. Is there a simple transformation that converts
"longest common suffix" into "longest common prefix"? Once you do that, a
classic structure makes lookups proportional to the query length.

### Hint 3

Build a trie of every word in `wordsContainer` inserted in *reverse*. At every
trie node, precompute and store the index of the candidate that should win if
the query stops there — i.e. the shortest container word that passes through the
node, breaking ties by the smaller original index. Inserting in input order and
using a strict `<` comparison handles the tiebreak naturally. Then each query is
a single walk down the trie following the reversed query characters, and the
answer is the precomputed best at the deepest node you can reach.

## Approach

Reverse-trie with per-node "best candidate":

1. Define a trie node holding 26 children and two extra fields: `bestIdx` and
   `bestLen`, which describe the optimal container word passing through this
   node so far.
2. Insert each `wordsContainer[i]` into the trie in *reverse order of
   characters*. As you walk through (or create) each node along that path,
   update its `(bestIdx, bestLen)` if the current word is strictly shorter than
   what the node already holds. Because we iterate `i` from `0` upward and use
   a strict `<`, equal-length ties are won by the earlier index automatically.
   The root represents the empty suffix, so it must be updated for every word —
   that handles queries with no common suffix.
3. For each `wordsQuery[i]`, walk the trie following the reversed query, one
   character at a time. Stop as soon as a child is missing. The current node's
   `bestIdx` is the answer.

Why it works:

- A path from the root of length `k` in this reversed trie corresponds exactly
  to a suffix of length `k` shared by every container word whose insertion
  visited that node.
- Walking down with the reversed query keeps us on the deepest node whose
  associated suffix is a suffix of the query — i.e. the longest common suffix.
- The precomputed `(bestIdx, bestLen)` at each node already encodes the
  shortest-then-earliest tiebreak, so the lookup itself is just a walk.

Example trace for `wordsContainer = ["abcdefgh", "poiuygh", "ghghgh"]`,
`query = "acbfgh"` (reversed: `h, g, f, b, c, a`): root → `h` (best: idx 2,
len 6) → `g` (best: idx 2, len 6) → `f` (only "abcdefgh" reaches here, best:
idx 0, len 8) → `b` (idx 0) → `c` (idx 0) → `a` (idx 0). Answer = 0, which
matches the expected output.

## Complexity Analysis

Let `S = sum of |wordsContainer[i]|` and `Q = sum of |wordsQuery[i]|`.

Time Complexity: O(S + Q) — each container character is visited once during
insertion (O(1) work per character because the alphabet is fixed at 26), and
each query character drives one trie step.

Space Complexity: O(26 * S) in the worst case for the trie node array slots,
or O(S) effective node count.

## Edge Cases

- **No shared suffix at all.** When the reversed query immediately falls off
  the trie, the answer must still be a valid index — the root node carries the
  global best (shortest, then earliest), so it covers this case cleanly.
- **Query longer than any container word.** The walk simply stops when a child
  is missing; we return the best at the deepest reachable node.
- **Multiple container words with the same length.** Inserting indices in
  increasing order and using strict `<` for updates ensures the earlier index
  wins ties without any extra logic.
- **Single container word.** Every query returns index 0, since the root and
  every visited node store that lone candidate.
- **Query equals a container word exactly.** The walk descends the full
  reversed-query path; the deepest node's `bestIdx` correctly identifies the
  match (or a shorter container word ending with the same suffix).
