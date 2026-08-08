# 3302. Find the Lexicographically Smallest Valid Sequence

[LeetCode Link](https://leetcode.com/problems/find-the-lexicographically-smallest-valid-sequence/)

Difficulty: Medium
Topics: Two Pointers, String, Dynamic Programming, Greedy
Acceptance Rate: 27.2%

## Hints

### Hint 1

Strip away the "almost equal" clause for a moment. If zero changes were allowed,
this would be the classic *is `word2` a subsequence of `word1`* problem, and the
greedy two-pointer scan (always grab the earliest character that matches) would
already give you the lexicographically smallest index array.

So the real question is: how do you extend that greedy to a budget of exactly one
mismatch, without letting the mismatch be spent in a spot that dooms you later?

### Hint 2

The danger with a greedy is committing to a choice you cannot undo. Spending your
one change at index `i` is only safe if everything left of `word2` can still be
matched **exactly** in `word1[i+1:]`.

That is a question about a *suffix*, and it does not depend on anything you did
earlier. Which means you can precompute it once, before the forward scan ever
starts, and then answer "is spending the change here safe?" in O(1).

### Hint 3

Define `suffixMatch[i]` = the length of the longest **suffix** of `word2` that is
a subsequence of `word1[i:]`. Build it right-to-left in one pass: keep a counter
`k`, and at each `i` bump it when `word1[i] == word2[m-1-k]`.

Now scan forward with `i` over `word1` and `j` over `word2`:

- If `word1[i] == word2[j]`, take `i`. A free match is never worse than paying
  for the same index with your change.
- Otherwise, if the change is unused and `suffixMatch[i+1] >= m-1-j`, spend it on
  `i`. That inequality is exactly "the remaining `m-1-j` characters of `word2`
  still fit exactly into `word1[i+1:]`", so the pick can always be completed.
- Otherwise `i` is unusable — skip it.

Every accepted index is the smallest one that keeps a completion reachable, which
is precisely the definition of lexicographically smallest.

## Approach

Let `n = len(word1)`, `m = len(word2)`.

**Step 1 — the suffix table.** We want `suffixMatch[i]`: how many trailing
characters of `word2` we can still match exactly, as a subsequence, using only
`word1[i:]`. Set `suffixMatch[n] = 0` and walk `i` from `n-1` down to `0`:

```
k = suffixMatch[i+1]
if k < m && word1[i] == word2[m-1-k] { k++ }
suffixMatch[i] = k
```

This is the ordinary subsequence-matching greedy, just run backwards. Because the
greedy match is optimal, `suffixMatch[i] >= t` is a correct and exact test for
"the last `t` characters of `word2` are a subsequence of `word1[i:]`".

**Step 2 — the forward greedy.** Keep `j` (next position of `word2` to satisfy)
and a flag `changed`. For `i = 0 .. n-1`:

- `word1[i] == word2[j]` → append `i`, `j++`.
- else if `!changed && suffixMatch[i+1] >= m-1-j` → append `i`, set
  `changed = true`, `j++`.
- else → skip `i`.

If the scan ends with `j < m`, no valid sequence exists; return an empty array.

**Why the greedy is safe.** The invariant is that the state `(i, j, changed)` is
always *completable* whenever the original input was. Three cases:

1. Taking a free exact match at `i` dominates every alternative: any other
   solution from this state uses an index `>= i` for `word2[j]`, and taking `i`
   leaves both a larger remaining suffix of `word1` and an untouched change
   budget. So it is never wrong.
2. Spending the change at `i` is guarded by the suffix test, so by construction
   the rest can be completed exactly. And index `i` is smaller than any later
   index, so grabbing it is lexicographically optimal.
3. When we skip `i`, index `i` is genuinely unusable for `word2[j]`: either the
   change is already spent and the characters differ, or the change is available
   but the suffix test fails — meaning *any* solution that used `i` here would
   have to burn the change and then complete exactly, which we just proved
   impossible. Every valid solution therefore skips `i` too.

**Worked example** — `word1 = "bacdc"`, `word2 = "abc"`:

`suffixMatch = [2, 1, 1, 1, 1, 0]` (indices `0..5`).

- `i=0`: `'b' != 'a'`. Change would need `suffixMatch[1] = 1 >= 2` — false. Skip.
- `i=1`: `'a' == 'a'`. Take `1`, `j=1`.
- `i=2`: `'c' != 'b'`. Change needs `suffixMatch[3] = 1 >= 1` — true. Take `2`,
  `changed = true`, `j=2`.
- `i=3`: `'d' != 'c'`, change spent. Skip.
- `i=4`: `'c' == 'c'`. Take `4`.

Result `[1, 2, 4]`, matching the expected output.

A subtler case is `word1 = "abcde"`, `word2 = "ace"`. The naive "match exactly,
then patch" instinct produces `[0, 2, 4]`, but the greedy above returns
`[0, 1, 4]`: it pays the change at index `1` (turning `'b'` into `'c'`) precisely
because the suffix table confirms `"e"` still fits in `word1[2:]`. Grabbing the
smaller index `1` beats the free-but-later index `2`.

## Complexity Analysis

Time Complexity: O(n + m) — one backward pass to build the table, one forward
pass to build the answer, both with O(1) work per character. With `n` up to
`3 * 10^5` this is comfortably fast.

Space Complexity: O(n) for the `suffixMatch` table, plus O(m) for the output
array (which is required anyway).

## Edge Cases

- **No valid sequence at all** (`word1 = "aaaaaa"`, `word2 = "aaabc"`): the
  forward scan runs off the end of `word1` with `j < m`. Return an empty array,
  not a partially filled one.
- **Two or more mismatches needed** (`word1 = "ab"`, `word2 = "cd"`): the suffix
  guard rejects the change at every index, so nothing is ever appended. Also
  returns empty.
- **The change must be saved for the very last character** (`word1 = "abz"`,
  `word2 = "abc"`): the greedy naturally holds the budget because it prefers free
  exact matches, and `suffixMatch[n] = 0 >= 0` lets it fire at the final index.
- **The change is best spent immediately at index 0** (`word1 = "vbcca"`,
  `word2 = "abc"`): the suffix test authorizes burning it on the first character,
  which is what makes `[0,1,2]` beat any later-starting sequence.
- **`m == 1` with no matching character anywhere** (`word1 = "zz"`,
  `word2 = "a"`): `m-1-j` is `0`, so the guard is trivially satisfied and the
  answer is `[0]`. Do not accidentally require a real match.
- **Off-by-one in the guard.** `suffixMatch[i+1] >= m-1-j` reads `index i+1`, so
  the table must have length `n+1` with a `0` sentinel — otherwise spending the
  change on the last character of `word1` panics.
- **Returning `nil` vs. an empty slice.** Return `[]int{}` so the result compares
  equal to an empty array under `reflect.DeepEqual` and serializes as `[]`.
- **Never consuming the change** (`word1 = "abcd"`, `word2 = "abc"`): a plain
  exact subsequence is a perfectly valid answer, since "at most one" includes
  zero changes.
