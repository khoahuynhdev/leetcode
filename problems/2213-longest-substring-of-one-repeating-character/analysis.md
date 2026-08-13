# 2213. Longest Substring of One Repeating Character

[LeetCode Link](https://leetcode.com/problems/longest-substring-of-one-repeating-character/)

Difficulty: Hard
Topics: Array, String, Segment Tree, Ordered Set
Acceptance Rate: 43.7%

## Hints

### Hint 1

Start by asking what the brute force costs. After each query you could rescan the
whole string and measure the longest run of equal characters — that is O(n) per
query, so O(n * k) overall, which is 10^10 operations at the constraint limits.
Too slow, but notice *why* it is wasteful: a single-character update can only
disturb the string near the index you touched. Everything far away is unchanged.

That observation is the signature of a whole family of problems: **point updates
plus a global aggregate query**. Which data structure lets you rebuild only the
O(log n) pieces of a decomposition that actually changed?

### Hint 2

Reach for a **segment tree** over the indices of `s`. The hard part is not the
tree mechanics — it is deciding what to store in a node so that a parent can be
computed from its two children in O(1).

Storing only "longest run inside my range" is not enough. Merging `"aab"` and
`"baa"` gives bests of 2 and 2, but the real answer for `"aabbaa"` is 2, while
merging `"aab"` and `"bba"` (bests 2 and 2) gives `"aabbba"` with answer 3. The
children's `best` values are identical in both cases yet the parents differ, so
`best` alone is not a *mergeable* summary. What extra information about the two
ends of a range would let you tell those cases apart?

### Hint 3

Each node stores three numbers for its range:

- `pref` — length of the run of equal characters starting at the range's left end
- `suf`  — length of the run of equal characters ending at the range's right end
- `best` — length of the longest run anywhere inside the range

Merging left child `L` (covering `[lo, mid]`) and right child `R` (covering
`[mid+1, hi]`) is then O(1). The key insight: **the only run in the parent that
is not already contained in a child is the one straddling the boundary**, and it
exists only when `s[mid] == s[mid+1]`. When it does:

- `best = max(L.best, R.best, L.suf + R.pref)`
- `pref` grows past the boundary only if the left child is *entirely* one
  character (`L.pref == len(L)`), in which case `pref = len(L) + R.pref`
- symmetrically, `suf = len(R) + L.suf` when `R.suf == len(R)`

Note that the merge reads `s[mid]` and `s[mid+1]` from the current string, so
keep a mutable `[]byte` copy of `s` alongside the tree and write the query
character into it *before* running the update. The answer to each query is then
just the root's `best`.

## Approach

Build a segment tree over the `n` positions of `s`, stored in a flat array of
size `4n` with the usual `2*idx` / `2*idx+1` child indexing. Alongside it keep
`s` as a mutable `[]byte`, because the merge step needs to compare the two
characters adjacent to a split point.

**Leaves.** A range of one character is a run of length one, so every leaf is
`{pref: 1, suf: 1, best: 1}`. Leaves never change — even when the character at
that index is overwritten, a single character is still a run of length 1. Only
the *internal* nodes above it need recomputing.

**Merge.** For a node covering `[lo, hi]` split at `mid`, with `lenL = mid-lo+1`
and `lenR = hi-mid`:

```
pref = L.pref
suf  = R.suf
best = max(L.best, R.best)

if s[mid] == s[mid+1]:
    best = max(best, L.suf + R.pref)     // the run crossing the boundary
    if L.pref == lenL: pref = lenL + R.pref
    if R.suf  == lenR: suf  = lenR + L.suf
```

The two conditional extensions are the subtle part. `pref` can only reach beyond
the left child if the left child is one solid run of a single character; only
then does the prefix run touch the boundary and continue into the right child.
The same reasoning, mirrored, applies to `suf`.

**Query loop.** For each `i`, set `s[queryIndices[i]] = queryCharacters[i]`, then
walk down to that leaf and re-merge on the way back up. Every node on the
root-to-leaf path — and no other — can have changed, which is exactly the O(log n)
work a segment tree gives you. Record `tree[1].best` as the answer.

**Worked example** — `s = "babacc"`, third query sets index 3 to `'b'`, giving
`"bbbbcc"`. The tree splits into `[0,2] = "bbb"` and `[3,5] = "bcc"`. The left
child is a solid run: `{pref: 3, suf: 3, best: 3}`. The right child is
`{pref: 1, suf: 2, best: 2}`. At the root `mid = 2`, and `s[2] == 'b' == s[3]`,
so the crossing run is `L.suf + R.pref = 3 + 1 = 4`, giving `best = 4`. Since
`L.pref == 3 == lenL`, the prefix also extends: `pref = 3 + 1 = 4`. The answer is
4, matching the expected output.

Be honest with yourself about this one: the segment tree scaffolding is routine,
but inventing the `(pref, suf, best)` triple and proving the merge is exhaustive
is genuinely Hard-tier. It is the same "make the aggregate mergeable" trick that
powers maximum-subarray segment trees and the classic bracket-matching tree — if
you internalize it here, you get those for free.

## Complexity Analysis

Time Complexity: O(n + k log n) — O(n) to build the tree, then O(log n) per
query for the point update and its re-merges, with the answer read from the root
in O(1).
Space Complexity: O(n) — the `4n` tree array plus the mutable byte copy of `s`.

## Edge Cases

- **`n == 1`.** The tree is a single leaf and the root *is* that leaf. Every
  answer is 1 regardless of the query character. Make sure the build and update
  routines terminate correctly when `lo == hi` at the root, without trying to
  index a child.
- **A query that changes nothing** (writing the character already at that index).
  The code must still produce the correct current answer rather than a stale one;
  re-running the update is harmless, but skipping the answer write is not.
- **A query that splits a long run**, e.g. `"aaaaa"` with `'b'` at index 2 → the
  answer *drops* from 5 to 2. Answers are not monotonically increasing, so any
  approach that only tracks a running maximum across queries is wrong.
- **A query that joins two runs across a node boundary**, e.g. `"aabaa"` with
  `'a'` at index 2 → 5. This is precisely the case the `pref`/`suf` extension
  logic exists for; a merge that only considers `L.suf + R.pref` without also
  growing `pref` and `suf` will produce answers that are too small for ancestors
  higher up the tree.
- **The entire string is one character.** Then `pref == suf == best == n` at the
  root, exercising both conditional extensions simultaneously.
- **Leaves are never recomputed.** After overwriting `s[pos]`, the leaf stays
  `{1,1,1}`. Forgetting this is harmless, but forgetting to write the new
  character into the byte array *before* re-merging is not — the merge would
  compare stale characters.
