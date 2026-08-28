# 3734. Lexicographically Smallest Palindromic Permutation Greater Than Target

[LeetCode Link](https://leetcode.com/problems/lexicographically-smallest-palindromic-permutation-greater-than-target/)

Difficulty: Hard
Topics: Two Pointers, String, Enumeration
Acceptance Rate: 48.5%

## Hints

### Hint 1

The search space looks enormous — `n` can be 300, so you can never enumerate permutations. But a palindrome carries far less information than its length suggests. Ask yourself: how many characters do you actually have to *decide* before the whole string is pinned down? Start by counting letters and figuring out when a palindromic permutation exists at all.

### Hint 2

Once you realise the first half (plus a middle character when `n` is odd) determines everything, the problem becomes the classic "smallest string greater than `target`" construction: pick a position `i` where your answer stops matching `target` and becomes strictly larger, then fill the remainder as small as possible. The twist is that the letters you may use at each position come from a shrinking multiset, not from a free alphabet.

### Hint 3

Among all valid answers, the best one is the one that agrees with `target` on the **longest** prefix — if candidate A matches `target` through index `i` and candidate B breaks earlier at `j < i`, then at position `j` candidate A still holds `target[j]` while B holds something bigger, so A is smaller. So enumerate the break position from late to early and take the first that works. The special case is "no break inside the first half at all": then the first half must equal `target[:h]` exactly, which produces exactly **one** candidate string — build it and compare it to `target` directly. That single check silently covers the middle character and the entire mirrored second half, which is where most buggy solutions get lost.

## Approach

Let `n = len(s)` and `h = n / 2`.

**Step 0 — feasibility.** Count each letter of `s`. A palindromic permutation exists iff at most one letter has an odd count. (Parity does the rest of the work for you: if `n` is odd the number of odd counts is odd, so "at most one" means "exactly one"; if `n` is even it means "zero".) Otherwise return `""`.

**Step 1 — collapse the problem to the first half.** Every palindromic permutation has the shape

```
P = H + [mid] + reverse(H)
```

where `H` is any arrangement of the *half multiset* `half[c] = cnt[c] / 2`, and `mid` (only when `n` is odd) is forced to be the unique odd-count letter. So we are no longer choosing among permutations of `s` — we are choosing `H`, and `H` determines `P` completely.

**Step 2 — order candidates by prefix agreement.** We want the smallest `P > target`. Classify candidates by the first index `i` where `P[i] > target[i]` (with `P[:i] == target[:i]`). As Hint 3 argues, larger `i` gives a smaller `P`, so we want the largest feasible `i`.

Note that any `i >= h` forces `H == target[:h]` — the break happens in the middle character or in the mirrored tail, but the first half was already locked to `target`'s first half. And `H == target[:h]` yields exactly one string. So all of those cases collapse into a single candidate:

- Walk `target[0..h-1]` consuming from the half multiset. If it consumes cleanly, the multiset of `target[:h]` *is* the half multiset. Build `cand = target[:h] + mid + reverse(target[:h])` and, if `cand > target`, return it immediately — nothing can beat it.

**Step 3 — break inside the first half.** Otherwise the break index `i` satisfies `i < h`. While walking in Step 2, record `states[i]`: the multiset remaining after consuming `target[:i]`. The walk stops at the first infeasible index, giving `maxPrefix`; any `i > maxPrefix` is impossible because the prefix itself cannot be built.

For `i` from `min(maxPrefix, h-1)` down to `0`:

- Look for the smallest letter `c > target[i]` still present in `states[i]`.
- If found: the first half is `target[:i] + c + (everything left, sorted ascending)`. The sorted tail is optimal because the string already exceeds `target` at index `i`, so the remaining first-half positions are compared only against each other, and they are the earliest undetermined positions in `P`.
- Mirror it (inserting `mid` if `n` is odd) and return.

If no `i` works, return `""`.

**Worked example** — `s = "baba"`, `target = "abba"`. Counts are `a:2, b:2`, no odd letter, `h = 2`, half multiset `{a:1, b:1}`. Step 2: `target[:2] = "ab"` consumes cleanly, so `cand = "ab" + "ba" = "abba"`, which is *not* strictly greater than `target`. Step 3: at `i = 1` the remaining multiset is `{b:1}` and `target[1] = 'b'` — nothing larger. At `i = 0` the remaining multiset is `{a:1, b:1}` and `target[0] = 'a'`, so pick `'b'`; the leftover `{a:1}` sorted gives half `"ba"`, and mirroring gives `"baab"`. ✅

## Complexity Analysis

Time Complexity: O(n · Σ) where Σ = 26 — the prefix walk is O(n), and the backward scan tries at most `h` break positions each doing an O(Σ) alphabet lookup, with the O(n) string construction happening only once on the successful position. Effectively O(n · 26) ≈ O(n).

Space Complexity: O(n · Σ) for the stored per-prefix multisets (`h + 1` arrays of 26 ints), plus O(n) for the output. This can be reduced to O(n + Σ) by walking the prefix backwards and un-consuming letters instead of caching states, but with `n <= 300` the cached version is clearer.

## Edge Cases

- **No palindromic permutation at all** (two or more odd-count letters, e.g. `s = "abc"`): must return `""` before any construction logic runs.
- **`n == 1`**: `h = 0`, the first half is empty and the answer is just the single letter of `s`. Step 2 must still fire (the empty prefix trivially "consumes cleanly") so that `s > target` is checked; Step 3's loop is empty. Getting this wrong returns `""` for `s = "b", target = "a"`.
- **The mirrored candidate equals `target`** (`s = "abba"`, `target = "abba"`): the comparison must be *strict*, so this candidate is rejected and the search falls through to Step 3, yielding `"baab"`.
- **The mirrored candidate is smaller than `target`** (odd `n` where `mid < target[h]`, or a mirrored tail that loses): also rejected, and Step 3 must still run rather than returning `""`.
- **Odd `n` with a forced middle character**: `mid` is never a free choice. If the first half equals `target[:h]`, the middle is whatever letter had the odd count — you cannot bump it upward to manufacture a win.
- **`target` above every palindromic permutation** (`s = "aabb"`, `target = "zzzz"`): every break position fails and the answer is `""`.
- **All letters identical** (`s = "aaaa"`): the half multiset has a single letter, so there is exactly one palindrome; the answer is it-or-nothing.
- **Prefix becomes infeasible early** (`s = "aabb"`, `target = "aaaa"`): `target[:h]` uses `'a'` twice but only one `'a'` lives in the half multiset, so `maxPrefix = 1` and break positions beyond it must not be attempted — indexing `states` past its end is the easy crash here.
