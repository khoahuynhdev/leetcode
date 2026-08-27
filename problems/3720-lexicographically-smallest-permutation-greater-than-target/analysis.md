# 3720. Lexicographically Smallest Permutation Greater Than Target

[LeetCode Link](https://leetcode.com/problems/lexicographically-smallest-permutation-greater-than-target/)

Difficulty: Medium
Topics: Hash Table, String, Greedy, Counting, Enumeration
Acceptance Rate: 39.5%

## Hints

### Hint 1

The words "permutation of `s`" are doing a lot of work here: the *order* of the letters in `s` is irrelevant, only *how many of each letter* you have. So the first move is to throw `s` away and keep a frequency table of 26 counts. From there the question becomes: how do I lay 26 buckets of letters onto `n` slots so the result beats `target`?

Also notice `n <= 300`. That is a strong signal that an `O(n * 26)` or even `O(n^2)` scan is intended — you are meant to *enumerate* something over the positions, not over the permutations (there can be astronomically many of those).

### Hint 2

Think about what it *means* for a string `a` of length `n` to be strictly greater than `target` of the same length. There must be some index `i` where `a[0..i-1]` is character-for-character equal to `target[0..i-1]`, and `a[i] > target[i]`. Everything after `i` is then completely free — it can be anything.

So enumerate that "breaking index" `i` over all `n` positions. For a fixed `i`, ask two questions: can I even spell `target[0..i-1]` out of my letter counts? And do I have a leftover letter strictly greater than `target[i]`? If yes to both, the smallest string with that breaking index is easy to build greedily.

### Hint 3

The insight that makes this click is *which* breaking index to prefer. Compare two candidates, one that breaks at `i` and one that breaks at `j > i`. The second one still agrees with `target` at position `i`, while the first one is strictly *above* `target` at position `i`. Since positions are compared left to right, the candidate that breaks later is always the smaller one.

So there is no need to build all `n` candidates and compare them: **scan `i` from the largest feasible value downward and return the first candidate that works.** And "largest feasible value" is bounded by the longest prefix of `target` you can actually spell from your counts — once a prefix is unspellable, no longer prefix is either. Cap that at `n - 1`, because breaking at index `n` would mean the answer equals `target`, which is not *strictly* greater.

## Approach

Count the letters of `s` into a fixed `[26]int` array. The order of `s` never matters again.

**Step 1 — find the longest spellable prefix of `target`.**
Walk `target` left to right, consuming letters from a copy of the counts. Stop at the first character you cannot afford. Call the number of characters consumed `maxPrefix`. For any breaking index `i`, the answer must reproduce `target[0..i-1]` exactly, so `i <= maxPrefix` is a hard requirement.

**Step 2 — pick the starting breaking index.**
Set `start = min(maxPrefix, n-1)`. The `n-1` cap matters: if `s` is an exact anagram of `target`, then `maxPrefix == n`, but breaking at index `n` would produce `target` itself, which fails the *strictly* greater requirement.

**Step 3 — scan `i` downward and build greedily.**
Maintain `avail`, the multiset left over after consuming `target[0..i-1]`. Initialize it for `i = start`, then as `i` decreases by one, hand the letter `target[i-1]` back to `avail`. That keeps the whole scan `O(n)` on bookkeeping instead of recomputing counts at each step.

At each `i`, look for the smallest letter `c` in `avail` with `c > target[i]`. If one exists, the answer is:

```
target[0..i-1]  +  c  +  (all remaining letters sorted ascending)
```

The prefix is forced. Choosing the *smallest* valid `c` minimizes position `i`, which dominates everything after it. And sorting the leftovers ascending is the smallest possible completion of a free suffix. Return immediately — by Hint 3 this is the global answer.

If the loop runs off the end without finding such a `c`, no permutation of `s` beats `target`, so return `""`.

**Worked example: `s = "abc"`, `target = "bba"`.**
Counts are `{a:1, b:1, c:1}`, `n = 3`. Spelling `target`: `'b'` is affordable, then the second `'b'` is not, so `maxPrefix = 1` and `start = min(1, 2) = 1`.
At `i = 1`: the prefix `"b"` is consumed, leaving `{a:1, c:1}`. We need a letter greater than `target[1] = 'b'` — that is `'c'`. Emit `"b" + "c"`, then the leftovers `{a}` sorted give `"a"`. Answer: `"bca"`. ✓

**Worked example: `s = "leet"`, `target = "code"`.**
`target[0] = 'c'` is not in `s` at all, so `maxPrefix = 0` and the only breaking index is `i = 0`. The smallest available letter above `'c'` is `'e'`; the leftovers `{e, l, t}` sorted give `"elt"`. Answer: `"eelt"`. ✓ Note that the answer shares no prefix with `target` here — the algorithm handles that uniformly.

## Complexity Analysis

Time Complexity: O(n * 26 + n) = O(n), treating the alphabet as a constant. Step 1 is a single `O(n)` pass; step 3 visits each of the `n` breaking indices and scans at most 26 letters at each, and the final string is assembled once in `O(n + 26)`.

Space Complexity: O(1) auxiliary beyond the `O(n)` output buffer — just two fixed `[26]int` count arrays.

## Edge Cases

- **`s` is an exact anagram of `target`** (e.g. `s = "abc"`, `target = "abc"`). `maxPrefix` reaches `n`, and without the `min(maxPrefix, n-1)` cap you would "break" at index `n` and return `target` itself, violating *strictly* greater. Here the correct answer is `"acb"`.
- **No answer exists** (e.g. `s = "baba"`, `target = "bbaa"`). Every breaking index fails, and the function must return `""` rather than a best-effort string. This happens exactly when `target` is greater than or equal to the largest permutation of `s`.
- **`target[0]` is not in `s`.** Then `maxPrefix = 0` and only `i = 0` is considered. Fine — but it means you must not assume the answer shares any prefix with `target`.
- **`n = 1`.** `start` collapses to `0`; the answer is the single character of `s` if it exceeds `target[0]`, else `""`.
- **All characters identical** (e.g. `s = "aaa"`). There is only one permutation, so the answer is `s` itself if `s > target`, else `""`. The counting approach handles this without special-casing.
- **Handing the letter back while scanning down.** An easy off-by-one: when moving from `i` to `i-1` you must return `target[i-1]` to `avail`, not `target[i]`. Getting this wrong silently produces a wrong multiset for the suffix.
- **Duplicate letters in `s`.** Because we work with counts rather than a visited/used boolean array, duplicates never cause the same permutation to be considered twice, and `avail[c]--` correctly leaves the other copies available for the suffix.
