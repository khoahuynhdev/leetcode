# 0115. Distinct Subsequences

[LeetCode Link](https://leetcode.com/problems/distinct-subsequences/)

Difficulty: Hard
Topics: String, Dynamic Programming
Acceptance Rate: 52.8%

## Hints

### Hint 1

You are asked to *count* things, not to decide whether something is possible. That is a strong signal: whenever a greedy "match as early as you can" scan would answer a yes/no version of the question, the counting version almost always becomes dynamic programming over two indices — one walking `s`, one walking `t`. Try to describe the answer in terms of "the number of ways to build a prefix of `t` from a prefix of `s`".

### Hint 2

Think about the decision you make at a single character of `s`. Standing at `s[i]`, you either *use* it to match the next needed character of `t`, or you *skip* it and leave it out of the subsequence. Skipping is always allowed; using is only allowed when the characters agree. Two independent choices that both lead to valid completions means the counts **add**. Define `dp[i][j]` = number of distinct subsequences of `s[0..i)` that equal `t[0..j)` and write the recurrence for those two branches.

### Hint 3

The recurrence is:

- `dp[i][j] = dp[i-1][j]` (always: drop `s[i-1]`)
- plus `dp[i-1][j-1]` when `s[i-1] == t[j-1]` (additionally: spend `s[i-1]` on `t[j-1]`)

with base cases `dp[i][0] = 1` (the empty target is matched exactly one way — take nothing) and `dp[0][j] = 0` for `j > 0` (a non-empty target cannot come from an empty source).

The critical insight for the optimal version: row `i` only ever reads row `i-1`, and only at columns `j` and `j-1`. So you can collapse to a single array of length `len(t)+1` — provided you iterate `j` **downward**. Going upward would overwrite `dp[j-1]` with its new value before you read it, silently letting one character of `s` match two positions of `t`.

## Approach

Let `n = len(s)` and `m = len(t)`. Define `dp[i][j]` as the number of distinct subsequences of the first `i` characters of `s` that spell out the first `j` characters of `t`.

**Base cases.** `dp[i][0] = 1` for every `i`: there is exactly one way to produce the empty string, namely by selecting nothing. `dp[0][j] = 0` for every `j > 0`: you cannot produce a non-empty target from an empty source.

**Transition.** Consider the last character of the source prefix, `s[i-1]`. Every valid subsequence either includes it or does not, and those two families are disjoint, so the counts add:

- *Exclude it.* The whole target must already be formed by `s[0..i-1)`, contributing `dp[i-1][j]`.
- *Include it.* This is only meaningful when `s[i-1] == t[j-1]`; then `s[i-1]` is matched against `t[j-1]`, and the remaining target `t[0..j-1)` must be formed by `s[0..i-1)`, contributing `dp[i-1][j-1]`.

So `dp[i][j] = dp[i-1][j] + (s[i-1] == t[j-1] ? dp[i-1][j-1] : 0)`, and the answer is `dp[n][m]`.

**Why this counts distinctly.** The problem asks for distinct *subsequences by index selection* — `"rabbbit"` has three different index sets that spell `"rabbit"`, and all three count. The recurrence enumerates exactly one path per index set, because at each position of `s` the include/exclude choice is made once and the two branches never produce the same selection.

**Rolling the table.** Row `i` depends only on row `i-1` at columns `j` and `j-1`. Keep a single `dp` of length `m+1`, initialized to `dp[0] = 1` and zeros elsewhere. For each character of `s`, sweep `j` from `m` down to `1` and do `dp[j] += dp[j-1]` whenever `s[i-1] == t[j-1]`. Descending order guarantees that when you read `dp[j-1]` it still holds the previous row's value. This drops memory from `O(n*m)` to `O(m)`.

**Worked example.** `s = "babgbag"`, `t = "bag"`. Tracking `dp = [dp_ε, dp_b, dp_ba, dp_bag]`:

| after char | dp |
| --- | --- |
| start | `[1, 0, 0, 0]` |
| `b` | `[1, 1, 0, 0]` |
| `a` | `[1, 1, 1, 0]` |
| `b` | `[1, 2, 1, 0]` |
| `g` | `[1, 2, 1, 1]` |
| `b` | `[1, 3, 1, 1]` |
| `a` | `[1, 3, 4, 1]` |
| `g` | `[1, 3, 4, 5]` |

The answer is `5`, matching the problem statement. Notice how `dp[1]` (ways to spell `"b"`) grows once per `b` seen, and how the final `g` folds all four ways of spelling `"ba"` into the total.

A small but useful pruning: if `m > n`, return `0` immediately — no subsequence of `s` can be longer than `s`.

This is a genuinely hard problem the first time you meet it, and the "iterate backwards" detail trips up nearly everyone. If you can rederive the two-branch recurrence from scratch, you already have the hard part; the rolling array is a mechanical optimization on top.

## Complexity Analysis

Time Complexity: O(n * m), where `n = len(s)` and `m = len(t)` — each of the `n` characters of `s` triggers one sweep of the `m`-length `dp` array, with `O(1)` work per cell. With `n, m <= 1000` that is at most one million operations.

Space Complexity: O(m) for the single rolling array. The naive two-dimensional table would be O(n * m); collapsing rows removes that without changing the arithmetic.

## Edge Cases

- **`t` longer than `s`.** No subsequence of `s` can exceed `s` in length, so the answer is `0`. The DP already returns `0` here, but the explicit early return avoids pointless work.
- **Empty `t`.** Outside the stated constraints, but worth getting right: the answer is `1`, not `0`. The empty string is a subsequence of everything, obtained by selecting nothing. This is exactly the `dp[0] = 1` base case, and getting it wrong zeroes out the entire table.
- **Empty `s` with non-empty `t`.** The loop over `s` never runs, `dp[m]` stays `0`, which is correct.
- **`s == t`.** Exactly one way. A good sanity check that you are not double counting.
- **No shared characters at all** (e.g. `s = "abc"`, `t = "d"`). Every cell past column 0 stays `0`; the answer is `0`.
- **Heavy repetition** (e.g. `s = "aaaa"`, `t = "aa"`). The answer is the binomial coefficient C(4, 2) = 6. This is the case that exposes an ascending inner loop: sweeping `j` upward would let a single `a` match both target positions and inflate the count.
- **Large answers.** The problem guarantees the result fits in a signed 32-bit integer, so Go's `int` (64-bit on the target platforms) never overflows and no modular arithmetic is needed.
- **Case sensitivity.** `s` and `t` are English letters of either case, and `'a'` does not match `'A'`. Compare bytes directly; do not normalize case.
