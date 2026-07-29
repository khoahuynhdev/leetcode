# 3518. Smallest Palindromic Rearrangement II

[LeetCode Link](https://leetcode.com/problems/smallest-palindromic-rearrangement-ii/)

Difficulty: Hard
Topics: Hash Table, Math, String, Combinatorics, Counting
Acceptance Rate: 24.3%

This one is genuinely hard — it stacks three separate ideas on top of each
other. But each idea is standard on its own, and once you see the reduction in
Hint 1 the problem shrinks to a much more familiar shape. Take the hints one at
a time.

## Hints

### Hint 1

`n` can be 10,000 and `k` can be 1,000,000, so enumerating palindromes is
hopeless — you need to *construct* the answer directly. Before thinking about
"k-th", ask a simpler question: how much freedom does a palindrome actually
give you? If you fix the left half of a palindrome, how many ways are there to
fill in the rest? What does that tell you about which part of the string you
really have to search over?

### Hint 2

You're now looking for the k-th lexicographically smallest **distinct
permutation of a multiset**. That's the classic "k-th permutation sequence"
pattern: decide the answer one position at a time, left to right. At each
position, walk the candidate characters in alphabetical order and ask "how many
complete arrangements start with the prefix I've built so far, followed by this
character?" If that number is at least `k`, commit to the character; otherwise
subtract it from `k` and move on. The counting function you need is the
multinomial coefficient `r! / (c₁! · c₂! · … · c₂₆!)`.

### Hint 3

The counts are astronomically large — `5000!` has thousands of digits — so
computing them exactly is a non-starter. The saving insight is that you never
need the exact value: `k ≤ 10⁶`, and every use of the count is either a
comparison `count ≥ k` or a subtraction `k -= count` that only happens when
`count < k`. So **clamp every count at 10⁶**. Compute the multinomial as a
product of binomials (`C(used, cnt[c])` as you add each character group) and
bail out the instant the running value crosses the cap. Because
`C(2i, i) > 10⁶` by `i = 12`, each clamped binomial finishes in ~25 iterations
and everything stays comfortably inside `int64`.

## Approach

**Step 1 — reduce to half the string.**

Every character in `s` appears an even number of times except at most one (`s`
is given as a palindrome, so this is guaranteed). Let `half[c] = cnt[c] / 2`
and let `mid` be the single odd character, if any. Any palindromic rearrangement
of `s` has the form `H + mid + reverse(H)` where `H` is some permutation of the
multiset `half`, and every such `H` yields a distinct palindrome.

Crucially, **ordering is preserved**: two palindromes of equal length whose
halves first differ at index `i` also first differ at index `i`, in the same
direction. So the k-th smallest palindrome corresponds to the k-th smallest
half. This turns "k-th palindromic permutation of `s`" into "k-th distinct
permutation of a multiset of size `n/2`" — which also handles the "identical
rearrangements count once" note for free, since we work with the multiset of
distinct strings, not with index permutations.

**Step 2 — count, with a cap.**

`countPermutations(cnt)` returns the number of distinct arrangements of the
remaining multiset, clamped to `permLimit = 10⁶`. Rather than `r!` divided by
factorials, it builds the multinomial as a product of binomials: process
characters `a…z`, keep a running `used` total, and multiply in
`C(used, cnt[c])` — "choose which of the slots used so far belong to this
character." Each binomial is itself computed incrementally as
`res = res * (n-r+i) / i`, whose value after step `i` is exactly `C(n-r+i, i)`,
so every division is exact and the sequence is non-decreasing — which is what
makes the early exit at the cap sound.

**Step 3 — greedy construction.**

Check `countPermutations(half) < k` up front; if so there aren't `k` distinct
palindromes and the answer is `""`. Otherwise build the half one position at a
time: for `c` from `a` to `z` with `half[c] > 0`, decrement `half[c]`, count the
arrangements of what's left, and either commit to `c` (if `ways ≥ k`) or restore
and do `k -= ways`. Note that the subtraction branch only runs when
`ways < k ≤ 10⁶`, so the clamped value is always the exact value there — the cap
never corrupts the arithmetic.

One important shortcut: as soon as `k == 1`, the smallest remaining arrangement
is just the leftover characters in sorted order, so dump them and stop. Without
it, a long string would keep re-counting for thousands of positions.

Finally return `H + mid + reverse(H)`.

**Walking through `s = "abba", k = 2`:** counts are `a:2, b:2`, so `half` is
`{a:1, b:1}`, `mid = ""`, `m = 2`. Total arrangements = 2 ≥ 2, so an answer
exists. Position 0: try `'a'` — remaining `{b:1}` has 1 arrangement, and
`1 < 2`, so `k` becomes 1 and we restore. Try `'b'` — remaining `{a:1}` has 1
arrangement and `1 ≥ 1`, so commit `'b'`. Position 1: `k == 1`, so append the
sorted remainder `"a"`. Half is `"ba"`, and the answer is `"ba" + "" + "ab"` =
`"baab"`.

## Complexity Analysis

Let `n = len(s)`, `σ = 26`, and `K = 10⁶` the clamp.

Time Complexity: O(n + n·σ²·log K) in the worst case. Counting the characters is
O(n). The greedy loop covers `n/2` positions, tries up to `σ` characters at
each, and each count is O(σ · log K) — the `σ` character groups each run a
clamped binomial that terminates within ~log₂K ≈ 25 iterations. That bound is
very loose in practice: counts blow past the cap after a handful of groups, and
the `k == 1` shortcut usually ends the loop early. Measured worst case at
`n = 10⁴, k = 10⁶` is well under a millisecond.

Space Complexity: O(n) for the output string, with only O(σ) = O(1) auxiliary
space (two 26-element count arrays).

## Edge Cases

- **Fewer than `k` distinct palindromes** (`"aa", k = 2`) — must return `""`,
  not a wrong string or a panic. Check the total up front; without it the greedy
  loop runs out of characters and silently emits garbage.
- **A string with only one distinct character** (`"aaaa"`) — exactly one
  palindrome exists, so only `k = 1` is valid. Also exercises `C(n, 0) = 1` and
  the multinomial reducing to 1.
- **Odd length** (`"bacab"`) — the lone odd-count character is pinned to the
  middle and must not participate in the permutation search; forgetting to
  re-insert it produces a string of the wrong length.
- **`n = 1`** (`"a"`) — the half is empty, so the greedy loop never runs and the
  count of arrangements of the empty multiset must come out as 1, not 0.
- **`k = 1`** — the answer is just the sorted half mirrored; the fast path must
  produce this without ever calling the counting routine.
- **`k` exactly equal to the total** — the boundary between a valid answer and
  `""`. Off-by-one here (`>` vs `≥` in `ways ≥ k`, or `≤` vs `<` in the total
  check) is the most common bug in this problem.
- **Counts that overflow** — the whole point of the cap. `n = 10⁴` with two
  distinct letters gives a multinomial with over a thousand digits; any
  unclamped `int64` or `float64` computation silently produces nonsense.
- **Multiple odd counts** — impossible per the constraints, but a cheap guard
  returning `""` makes the function safe to reuse on arbitrary input.
