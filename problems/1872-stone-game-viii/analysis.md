# 1872. Stone Game VIII

[LeetCode Link](https://leetcode.com/problems/stone-game-viii/)

Difficulty: Hard
Topics: Array, Math, Dynamic Programming, Minimax, Prefix Sum, Game Theory, Zero-Sum Game
Acceptance Rate: 58.5%

## Hints

### Hint 1

The rules look messy — stones get removed, merged, and pushed back onto the left — but notice what the row actually looks like after any sequence of moves. The merged stone on the left is always the sum of some prefix of the *original* array, and everything to its right is untouched original stones. That means the whole game state can be described by a single number. Find that number, and think about what a state-per-index DP would look like. Prefix sums will be your friend.

### Hint 2

Let `p[i] = stones[0] + ... + stones[i]`. If the player to move takes stones up through original index `j`, they score exactly `p[j]`, regardless of how many merges happened before — the merged left stone already carries all of the earlier values. So a move is just "pick an index `j`, score `p[j]`, hand the opponent a game that starts at `j+1`". This is a zero-sum minimax game, so define your DP as *score difference from the perspective of whoever is about to move* and the min-player disappears: the opponent's best difference simply gets subtracted.

### Hint 3

Define `dp[i]` = the best achievable difference (current player minus opponent) when the current player is allowed to end their take at any index `j >= i`. There are only two things the current player can do at `i`: end the take exactly at `i`, scoring `p[i]` and then facing `-dp[i+1]`, or skip past `i` and settle for the best option available from `i+1`, which is `dp[i+1]` itself. So

```
dp[n-1] = p[n-1]
dp[i]   = max(p[i] - dp[i+1], dp[i+1])
```

That "or just reuse `dp[i+1]`" branch is what collapses the inner loop over all `j` into O(1) work. The answer is `dp[1]`, because Alice must remove at least `x > 1` stones on her first move — she cannot stop at index 0. Since each `dp[i]` only looks at `dp[i+1]`, you never need the array at all: a single rolling variable and a running prefix sum suffice.

## Approach

**Step 1: Collapse the state.**

The intimidating part of this problem is the "place a new stone equal to the sum on the left" rule. But that new stone is never anything other than a prefix sum of the original array. Suppose the last completed move ended at original index `j`. The row is then exactly:

```
[ p[j], stones[j+1], stones[j+2], ..., stones[n-1] ]
```

where `p[j] = stones[0] + ... + stones[j]`. So one integer `j` fully describes the position — there is no hidden history. This is the key modeling step, and it turns an exponential-looking game into a linear scan.

**Step 2: Score of a move.**

From state `j`, the player to move must take at least 2 stones from the row: the merged stone `p[j]` plus at least one of the original stones after it. If they take through original index `k` (with `k >= j+1`), their score is

```
p[j] + stones[j+1] + ... + stones[k] = p[k]
```

The merged stone already contains everything before it, so the score telescopes to a single prefix sum. Nice.

**Step 3: Zero-sum minimax as one recurrence.**

Because Alice maximizes `(Alice - Bob)` and Bob minimizes it, this is a zero-sum game, and we can write a single function from the mover's point of view. Let

> `dp[i]` = the best value of (mover's score − opponent's score) over the rest of the game, given the mover may finish their take at any index `j` with `i <= j <= n-1`.

Naively `dp[i] = max over j >= i of (p[j] - dp[j+1])`, which is O(n²). But look at the max: the set of choices available at `i` is exactly the set available at `i+1`, plus the single new option `j = i`. So

```
dp[i] = max( p[i] - dp[i+1],   // stop the take right here at index i
             dp[i+1] )         // don't stop here; take the best from further right
```

with base case `dp[n-1] = p[n-1]` (at the last index there is nothing further right; the mover is forced to sweep everything).

**Step 4: The answer.**

Alice moves first and must choose `x > 1`, so her take must extend at least to index 1. She cannot stop at index 0. Hence the answer is `dp[1]`, not `dp[0]`.

**Step 5: Drop the array.**

`dp[i]` depends only on `dp[i+1]`, so iterate `i` from `n-2` down to `1` carrying one variable. The prefix sums can be built up front in one pass (or maintained as a suffix-running total while scanning backwards). The implementation here precomputes prefix sums in place-free fashion and then runs the rolling recurrence.

**Worked example: `stones = [-1, 2, -3, 4, -5]`**

Prefix sums: `p = [-1, 1, -2, 2, -3]`.

| i | p[i] | p[i] - dp[i+1] | dp[i+1] | dp[i] |
|---|------|----------------|---------|-------|
| 4 | -3   | — (base)       | —       | **-3** |
| 3 | 2    | 2 - (-3) = 5   | -3      | **5**  |
| 2 | -2   | -2 - 5 = -7    | 5       | **5**  |
| 1 | 1    | 1 - 5 = -4     | 5       | **5**  |

Answer `dp[1] = 5`, matching the expected output. And the table tells the story: Alice's winning line is to stop at index 3 (take the first 4 stones for a score of `p[3] = 2`), leaving Bob forced into `p[4] = -3`. Difference `2 - (-3) = 5`.

## Complexity Analysis

Time Complexity: O(n) — one pass to build prefix sums and one backward pass for the DP.
Space Complexity: O(1) extra if prefix sums are computed in place / on the fly (O(n) if you materialize a separate prefix array; the implementation reuses a single prefix slice, and the DP itself needs only one rolling variable).

## Edge Cases

- **`n == 2`** (the minimum allowed). Alice has exactly one legal move: take both stones. The loop body never executes and the answer is the base case `dp[1] = p[1] = stones[0] + stones[1]`. Example 3, `[-10,-12] -> -22`, is precisely this case, and it's a good reminder that the answer can be very negative — Alice is not guaranteed a win.
- **Answer is `dp[1]`, never `dp[0]`.** Off-by-one here is the single most common bug. `x > 1` means Alice must consume at least two stones, so index 0 is not a legal stopping point. Returning `dp[0]` would let Alice "pass", and it silently produces wrong answers on inputs like `[-10,-12]`.
- **All-negative arrays.** Every move hurts the mover, so the interesting behavior is wanting to move as little as possible — but passing isn't allowed, so someone must eat the loss. For `[-1,-2,-3]` the answer is `3`: Alice takes two stones (`-3`), which forces Bob to sweep to `-6`. The `dp[i+1]` branch of the recurrence is what captures "I'd rather defer than stop here".
- **All-positive arrays.** Here the greedy instinct "take everything" is often right, and the DP confirms it: for `[7,-6,5,10,5,-2,-6]` (Example 2) the answer is `13`, the full total. Worth checking that your code doesn't over-think this case.
- **Zeros / a total of zero.** `[0,0,0,0]` should yield `0`. Watch that a "max" initialized to `0` instead of the true base case doesn't accidentally look correct here while being wrong elsewhere — initialize from `p[n-1]`, not from zero or from a sentinel.
- **Overflow.** With `n` up to `10^5` and `|stones[i]| <= 10^4`, prefix sums reach `10^9`, which overflows a 32-bit accumulator in some languages. Go's `int` is 64-bit on all mainstream platforms, so this is safe here, but it's a real trap in C++/Java if you reach for `int` casually.
- **Large `n` (10^5).** Any O(n²) DP — including the natural first formulation `dp[i] = max over all j >= i` — will time out. The collapse to two candidates per index is not an optimization, it's the intended solution.
