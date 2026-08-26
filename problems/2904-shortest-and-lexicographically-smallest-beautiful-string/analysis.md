# 2904. Shortest and Lexicographically Smallest Beautiful String

[LeetCode Link](https://leetcode.com/problems/shortest-and-lexicographically-smallest-beautiful-string/)

Difficulty: Medium
Topics: String, Sliding Window
Acceptance Rate: 43.2%

## Hints

### Hint 1

The constraint `s.length <= 100` is a gift: it tells you that an O(n^2) scan over all candidate substrings is completely acceptable. So the real challenge is not squeezing out the best complexity, it is *enumerating the right candidates* and comparing them correctly. Ask yourself: out of the O(n^2) substrings, which ones could ever be the answer?

### Hint 2

Think about what a *shortest* beautiful substring must look like at its two ends. If a candidate starts with `'0'`, you could chop that `'0'` off and still have exactly `k` ones — a strictly shorter beautiful substring. The same argument applies to a trailing `'0'`. That single observation collapses the search space dramatically, and it points you toward tracking something other than raw indices.

### Hint 3

Record the indices of every `'1'` in an array `ones`. Any minimal beautiful substring is exactly the span from some `ones[i]` to `ones[i+k-1]` inclusive — there are at most `n - k + 1` such windows, and you can generate all of them with one pass. Sweep them, keep the smallest length, and break ties with a plain string comparison (`<` in Go compares byte-by-byte, which is exactly the lexicographic rule the problem describes). Note that among equal-length candidates you genuinely have to compare content: shorter-is-better and smaller-is-better are two separate criteria applied in that order.

## Approach

The key structural insight is that **the shortest beautiful substring must begin and end with `'1'`**. Suppose a beautiful substring `t` starts with `'0'`. Removing that leading `'0'` leaves a substring that still contains exactly `k` ones, so it is also beautiful and is strictly shorter — contradiction with `t` being shortest. Symmetrically for a trailing `'0'`.

That means every shortest candidate is pinned to a pair of `'1'` positions: the position of its first one and the position of its `k`-th one. So:

1. **Collect the ones.** Walk `s` once and push each index `j` with `s[j] == '1'` into a slice `ones`. Let `m = len(ones)`.
2. **Early exit.** If `m < k`, there is no substring anywhere in `s` containing `k` ones, so return `""`. This also covers the "no ones at all" case.
3. **Slide a fixed window over `ones`, not over `s`.** For each `i` in `[0, m-k]`, the window `ones[i..i+k-1]` describes the substring `s[ones[i] : ones[i+k-1]+1]`. It contains exactly `k` ones by construction (all ones strictly between `ones[i]` and `ones[i+k-1]` are precisely `ones[i+1..i+k-2]`), and it is the *tightest* substring whose `k` ones are that particular group.
4. **Two-tier comparison.** Maintain the best answer so far. Replace it when the new candidate is strictly shorter, or when it has the same length but compares lexicographically smaller. Go's `<` on `string` is a byte-wise lexicographic comparison, and since all candidates here have equal length when we tie-break, it matches the problem's definition exactly.

Walking through Example 1, `s = "100011001"`, `k = 3`:

- `ones = [0, 4, 5, 8]`, so `m = 4` and `i` ranges over `{0, 1}`.
- `i = 0`: span `ones[0]=0` to `ones[2]=5` → `s[0:6] = "100011"`, length 6. This is the first candidate, so it becomes the best.
- `i = 1`: span `ones[1]=4` to `ones[3]=8` → `s[4:9] = "11001"`, length 5. Strictly shorter, so it wins.

Answer: `"11001"`. Notice that the seven beautiful substrings the problem enumerates are all just these two windows padded with extra zeros — the padding can only make things longer, never shorter, which is exactly why restricting to `'1'`-anchored windows loses nothing.

One nuance worth internalizing: the tie-break is *not* redundant. Two different windows can produce equal-length candidates with different content, e.g. `s = "1101011"`, `k = 3` yields `"1101"` (from `ones[0..2]`) and `"1011"` (from `ones[2..4]`), both of length 4. The correct answer is `"1011"`, and you only get it if you keep comparing after the length ties.

## Complexity Analysis

Time Complexity: O(n * k) in the worst case — there are at most `n - k + 1` windows, and each one costs O(window length) to slice and compare. Since a window spanning `k` ones has length at most `n`, this is bounded by O(n^2), which is trivial for `n <= 100`. The scan that builds `ones` is O(n).

Space Complexity: O(n) for the `ones` index slice, plus O(n) for the best-candidate string. Go string slicing does not copy the backing array, so the candidates themselves are free; only the retained `best` matters.

## Edge Cases

- **Fewer than `k` ones in `s` (including zero ones).** Examples: `s = "000", k = 1`, or `s = "0110", k = 3`. The window loop must not run at all — guard with `m < k` (or equivalently let the loop bound `m-k` go negative) and return `""`.
- **`k` equals the total number of ones.** There is exactly one window, `ones[0..m-1]`. The loop must still execute once; an off-by-one in the bound (`i < m-k` instead of `i <= m-k`) silently returns `""` here.
- **`k = 1`.** Every single `'1'` is its own window of length 1, and the answer is always `"1"` whenever `s` contains a one. A good sanity check that your window arithmetic handles `k-1 = 0` correctly.
- **Ties on length.** As shown with `"1101011"`, equal-length candidates require the lexicographic comparison. Dropping the tie-break passes both of the first two examples, so this bug hides easily.
- **Leading/trailing zeros in `s`.** Strings like `"001101"` must not have their zeros absorbed into the answer. Anchoring on `'1'` positions handles this automatically — a good reason to prefer this formulation over a raw two-pointer sweep with a running counter, where it is easy to forget to shrink the left edge past zeros.
- **All ones, e.g. `"11111"`.** Every window has length exactly `k`; the first one found is already optimal, and no tie-break can improve on it. Cheap to verify, and it confirms your loop bounds.
