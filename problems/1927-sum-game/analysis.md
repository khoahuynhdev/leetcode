# 1927. Sum Game

[LeetCode Link](https://leetcode.com/problems/sum-game/)

Difficulty: Medium
Topics: Math, String, Greedy, Game Theory
Acceptance Rate: 51.9%

## Hints

### Hint 1

Nothing about this problem actually requires simulating the game. The string can be
100,000 characters long, so any search over move orders is hopeless. That is a strong
signal that the answer is decided by a small number of aggregate quantities you can
compute in one pass — think "what summary of the board fully determines the winner?"

Start by asking which numbers even matter. The individual digits do not; only the sum
of the known digits in each half, and the count of `'?'` in each half.

### Hint 2

Reduce the board. Let `d` be (sum of known digits in the left half) minus (sum of known
digits in the right half), and let `m` be (number of `'?'` in the right half) minus
(number of `'?'` in the left half).

A `'?'` on the left and a `'?'` on the right can be treated as a matched pair: whatever
one player writes into one of them, the opponent can immediately write the same digit
into its partner and restore the balance. So matched pairs are neutral — they cancel.
What is left is `d` plus a block of `|m|` unmatched `'?'` all sitting on the same side.

Also think about parity: who makes the last move? If the total number of `'?'` is odd,
Alice moves last, and the player who moves last on an otherwise-decided position has an
enormous advantage.

### Hint 3

Two facts finish it.

First, parity: the total count of `'?'` has the same parity as `m`. If it is odd, Alice
moves last and can always break a tie (she has 10 digits to choose from and at most one
value would keep the sums equal), so Alice wins.

Second, the value of the unmatched block. Suppose `m` is even and there are `m = 2k`
unmatched `'?'` on the right. Bob's strategy is to answer every Alice move inside that
block with `9 - x`, forcing each pair to contribute exactly `9`. So Bob can guarantee the
block totals exactly `9k`. Symmetrically, Alice can force it away from `9k`: if she wants
the block bigger she opens with `9` (Bob can only add `0..9`, so each of her pairs is
worth at least `9` and she can push strictly above), and if she wants it smaller she
opens with `0`. Therefore the block is pinned to exactly `9k` **only** when neither
player wants to move it, i.e. when `9k` is precisely the value that makes the halves
equal.

Bob wins if and only if `m` is even and `d == 9 * m / 2`. Alice wins in every other case.

## Approach

Scan the string once. Walk index `i` from `0` to `n-1`, treating the left half
(`i < n/2`) with a `+` sign and the right half with a `-` sign:

- If `num[i]` is a digit, add it to `sumDiff` with that sign — so `sumDiff` ends up as
  `leftKnownSum - rightKnownSum`, the quantity called `d` above.
- If `num[i]` is `'?'`, add `1` to `qDiff` with the *opposite* sign — so `qDiff` ends up
  as `rightQuestionCount - leftQuestionCount`, the quantity called `m` above.

Then apply the rule directly:

1. If `qDiff` is odd, return `true` (Alice wins). Note `qDiff = q2 - q1` and
   `q1 + q2` always have the same parity, so checking `qDiff` is the same as checking
   whether the total number of `'?'` is odd. With an odd count Alice places the final
   digit; at that moment at most one of the ten digits would equalize the halves, so she
   simply picks a different one.
2. Otherwise `qDiff` is even, and Bob wins exactly when `d` equals `9 * qDiff / 2`.
   Written without division (which keeps the negative-`qDiff` case honest), Bob wins iff
   `2 * sumDiff == 9 * qDiff`. Return the negation, since we report whether *Alice* wins.

Why the rule is right, in one paragraph: pair off `'?'`s across the two halves. Each such
pair is neutral, because the opponent mirrors the digit into the partner slot and the
difference returns to where it was. That leaves `|m|` question marks concentrated on one
side. If `m` is odd, the parity argument above hands the game to Alice. If `m = 2k` is
even, that block behaves like a tug-of-war over its total: Bob's mirror strategy `x -> 9-x`
locks it at `9k`, and Alice's greedy openings (`9` to inflate, `0` to deflate) let her
push it off `9k` in either direction. So the halves end up equal precisely when the target
`d` already equals the locked value `9k`.

Worked example, `num = "?3295???"`. Left half `?329`: known sum `3+2+9 = 14`, one `'?'`.
Right half `5???`: known sum `5`, three `'?'`. So `d = 14 - 5 = 9` and `m = 3 - 1 = 2`.
`m` is even, `k = 1`, and `9k = 9 == d`. Bob wins, so the answer is `false` — matching the
problem's example.

Contrast with `num = "25??"`: `d = 7 - 0 = 7`, `m = 2 - 0 = 2`, `9k = 9 != 7`. Alice wins,
answer `true`. She writes `9` into one `'?'`, and the right half is already too big to fix.

## Complexity Analysis

Time Complexity: O(n) — a single pass over the string, with O(1) work per character.
Space Complexity: O(1) — only two integer accumulators, independent of input size.

## Edge Cases

- **No `'?'` at all.** The game is over before it starts; the answer is just whether the
  two known sums differ. The formula covers this for free: `qDiff = 0`, so the check
  collapses to `sumDiff != 0`.
- **Odd total number of `'?'`.** Handled by the parity branch and easy to get wrong if you
  only compare `2*sumDiff` against `9*qDiff`. With `qDiff` odd, `9*qDiff` is odd while
  `2*sumDiff` is even, so equality can never hold and the code would still return `true` —
  but relying on that accident is fragile. Check parity explicitly.
- **Unmatched `'?'`s on the *left* half (`qDiff < 0`).** The rule is symmetric, but
  integer division of a negative `qDiff` by `2` is a place bugs hide. Comparing
  `2*sumDiff == 9*qDiff` sidesteps division entirely. Test this direction explicitly,
  e.g. `"??09"` (Bob wins) versus `"??00"` (Alice wins).
- **Balanced `'?'`s across halves (`qDiff == 0`).** Then Bob wins iff the known sums
  already match. `"??"` is Bob's win (he mirrors Alice's digit); `"0?9?"` is Alice's win
  because the known digits are already `9` apart and the mirroring cancels the `'?'`s.
- **Minimum length `n == 2`.** One character per half. Cases like `"?0"` (odd `'?'` count,
  Alice wins) exercise the parity branch at the smallest possible size.
- **Maximum length (10^5).** Only matters as a reminder that anything beyond a single
  linear pass is too slow — no recursion, no memoization over game states.
