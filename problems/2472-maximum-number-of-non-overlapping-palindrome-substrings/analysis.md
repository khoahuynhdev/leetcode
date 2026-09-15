# 2472. Maximum Number of Non-overlapping Palindrome Substrings

[LeetCode Link](https://leetcode.com/problems/maximum-number-of-non-overlapping-palindrome-substrings/)

Difficulty: Hard
Topics: Two Pointers, String, Dynamic Programming, Greedy
Acceptance Rate: 51.2%

## Hints

### Hint 1

Strip the palindrome part away for a moment and ask what the problem really is: you are handed a collection of intervals `[i, j]` (every palindromic substring of length at least `k`) and asked to pick as many as possible without any two touching. That is the classic **interval scheduling** / **non-overlapping intervals** setup. Every solution to this problem is really "enumerate the candidate intervals efficiently, then schedule them."

### Hint 2

Two sub-problems, solve them separately:

1. **Which intervals exist?** Deciding "is `s[i..j]` a palindrome?" for every pair naively costs O(n) per pair, which is O(n³) overall — too slow for `n = 2000`. There is a standard O(n²) way to get *all* of them at once, and it comes from a tiny recurrence: a string is a palindrome exactly when its two ends match **and** the string inside them is a palindrome. Think about what order you must fill a 2-D table in for that recurrence to be usable.
2. **How do you schedule them?** Define a one-dimensional DP over suffixes: `dp[i]` = the best answer achievable using only `s[i:]`. At index `i` you either skip character `i`, or you commit to some palindrome starting at `i`.

### Hint 3

The insight that makes both halves click:

- For the palindrome table, iterate `i` **downward** (from `n-1` to `0`) and `j` upward from `i`. Then `isPal[i][j] = s[i] == s[j] && (j-i < 2 || isPal[i+1][j-1])` — the inner substring `[i+1, j-1]` has a larger left index, so it was already computed. Lengths 1 and 2 are the base cases folded into `j-i < 2`.
- For the scheduling, `dp[i] = max(dp[i+1], 1 + dp[j+1])` over every `j >= i+k-1` with `isPal[i][j]`. And here is the greedy kicker: among all palindromes starting at `i`, you never need more than the **shortest** valid one. If an optimal solution uses a longer palindrome starting at `i`, swapping in the shorter one keeps the same count and frees up strictly more room on the right — an exchange argument. So the inner loop can `break` at the first match.

## Approach

The solution has two phases.

**Phase 1 — precompute every palindrome in O(n²).**

Build a boolean table `isPal[i][j]` meaning "`s[i..j]` is a palindrome". The recurrence is:

```
isPal[i][j] = (s[i] == s[j]) && (j - i < 2 || isPal[i+1][j-1])
```

`j - i < 2` covers the two base cases: a single character (`i == j`) is always a palindrome, and a two-character block is a palindrome iff its characters match. For anything longer, we peel off the outer pair and defer to the inside.

The dependency is `(i, j) -> (i+1, j-1)`, so we fill with `i` descending and `j` ascending. In the code the table is one flat `[]bool` of size `n*n` indexed as `i*n + j` — for `n = 2000` that is 4 MB in one allocation instead of 2000 separate slices, which is both faster and friendlier to the garbage collector.

**Phase 2 — DP over suffixes.**

Let `dp[i]` be the maximum number of non-overlapping valid substrings selectable from `s[i:]`, with `dp[n] = 0`. Walk `i` from `n-1` down to `0`:

- Option A: don't start a substring at `i`. That gives `dp[i+1]`.
- Option B: start one at `i`. Scan `j` from `i+k-1` rightward for the first `j` with `isPal[i][j]` true. Taking it gives `1 + dp[j+1]`.

`dp[i]` is the max of the two, and the answer is `dp[0]`.

The `break` after the first match in Option B is the greedy step justified in Hint 3: a shorter palindrome starting at `i` is never worse than a longer one, because `dp` is non-increasing in its index (`dp[j+1] >= dp[j'+1]` whenever `j <= j'`), so `1 + dp[j+1]` is already maximal at the smallest feasible `j`.

**Worked example: `s = "abaccdbbd"`, `k = 3`.**

Indices: `a(0) b(1) a(2) c(3) c(4) d(5) b(6) b(7) d(8)`. Palindromes of length ≥ 3 are `"aba"` at `[0,2]` and `"dbbd"` at `[5,8]`.

Filling `dp` right-to-left: `dp[9] = 0`, and `dp[8] = dp[7] = dp[6] = 0` (no length-3 palindrome starts there). At `i = 5`, `j = 8` gives `isPal[5][8] = true` (`"dbbd"`), so `dp[5] = max(dp[6], 1 + dp[9]) = 1`. Indices 4 and 3 start nothing, so `dp[4] = dp[3] = 1`. At `i = 0`, `j = 2` matches (`"aba"`), so `dp[0] = max(dp[1], 1 + dp[3]) = max(1, 2) = 2`. Answer: **2**.

Note how `dp[3] = 1` already carries the `"dbbd"` choice forward — that is exactly the "schedule the rest of the string optimally" subproblem, and it is why the greedy break is safe.

**A note on the alternative.** There is a slicker O(n·k)-ish solution that expands around each of the `2n-1` centers left to right and greedily takes the first palindrome of length ≥ `k` that starts at or after the last chosen endpoint. It uses O(1) extra space and is worth studying once you have the DP down, but the center ordering argument is fiddlier to get right. The DP above is the version to reach for under interview pressure: it is easy to justify out loud and comfortably fast enough for `n = 2000`.

## Complexity Analysis

Time Complexity: O(n²) — the palindrome table is O(n²) to fill, and the DP phase does at most O(n) work per index in the worst case (for example `s = "aaaa...a"` with large `k`), so O(n²) overall. With n = 2000 that's ~4·10⁶ operations.

Space Complexity: O(n²) for the palindrome table, plus O(n) for the `dp` array. The flat `[]bool` layout keeps this at ~4 MB for the maximum input.

## Edge Cases

- **`k == 1`** — every single character is a palindrome of sufficient length, so the answer is always `len(s)`. Make sure the inner loop starts at `j = i + k - 1 = i` and not `i + 1`; an off-by-one here silently halves the answer on strings like `"abc"`.
- **`k > len(s)`** — no substring can be long enough, so the answer is `0`. The loop bound `j <= n-1` makes the inner loop body never execute, and `dp` stays all zeros. Worth guarding explicitly (the code returns early) so the O(n²) table is never even built.
- **Empty string** — outside the stated constraints (`1 <= s.length`), but the early return keeps it from panicking on an index. Cheap insurance.
- **No palindrome of length ≥ k exists** (Example 2, `"adbcda"` with `k = 2`) — the answer is `0`, not `1`. It's easy to write a greedy that accidentally counts a single character.
- **Highly repetitive strings like `"aaaaaaaa"`** — this is the worst case for the DP and the case where the "take the shortest palindrome" greedy matters most. With `k = 2` the answer is `len(s) / 2`; a solution that greedily grabs the *longest* palindrome at each position would return `1`.
- **Even vs. odd length palindromes** — `"aa"` and `"aba"` must both be found. The `j - i < 2` base case handles both; a table built only from odd centers would miss `"cc"` in Example 1.
- **The palindrome that must be split** — e.g. `s = "abbaaa"`, `k = 2`. Taking `"abba"` yields 2, and so does taking `"bb"` then `"aa"`. Cases like this are where an untested greedy tends to go wrong, so keep one in your test table.
