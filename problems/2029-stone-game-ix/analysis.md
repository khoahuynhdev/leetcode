# 2029. Stone Game IX

[LeetCode Link](https://leetcode.com/problems/stone-game-ix/)

Difficulty: Medium
Topics: Array, Math, Greedy, Minimax, Counting, Game Theory, Nim Game, Zero-Sum Game
Acceptance Rate: 33.6%

## Hints

### Hint 1

`stones.length` can be `10^5`, so simulating or memoizing over game states is out of reach — the answer has to come from a closed-form rule. Also notice the only thing that ever matters about a stone is its value **mod 3**: two stones with values 4 and 7 are completely interchangeable. That collapses the whole array into three counters. Think about what the game looks like when your entire input is `(c0, c1, c2)`.

### Hint 2

Split the stones into two very different roles:

- Stones with remainder **1** or **2** actually move the running sum around.
- Stones with remainder **0** leave the sum untouched. They are *pass* moves — the only thing they change is whose turn it is.

Try writing out, by hand, which sequences of 1s and 2s keep the prefix sum away from a multiple of 3. You will find there are surprisingly few, and they are essentially forced. Then ask separately what the pass moves do on top of that.

### Hint 3

The critical insight comes in two halves.

**Half one — the forced pattern.** Alice can never open with a remainder-0 stone (the sum would be `0 mod 3` immediately and she'd lose on move one). So she opens with a 1 or a 2, and from there the only non-losing continuation is forced:

- opening with 1: `1, 1, 2, 1, 2, 1, 2, …`
- opening with 2: `2, 2, 1, 2, 1, 2, 1, …`

After the first two moves the residues must strictly alternate. Whoever runs out of the residue they need is the one forced to hit a multiple of 3.

**Half two — parity of the zeros.** Each remainder-0 stone is a free turn-swap. So only `c0 % 2` matters, not `c0` itself:

- **`c0` is even.** The pass moves cancel out and can't change who is stuck. Alice wins exactly when she has at least one stone of *each* nonzero residue — she needs a 1 **and** a 2 to run the pattern to Bob's disadvantage.
- **`c0` is odd.** The extra pass flips the parity, which hands the game to Bob unless one residue class is lopsided enough to overrun him: Alice wins exactly when `|c1 - c2| > 2`.

## Approach

Count the stones by remainder mod 3 into `c0`, `c1`, `c2` — that single pass is the entire "algorithm", and everything after it is a constant-time decision.

**Why the residues are forced.** Let `S` be the running sum mod 3. The mover loses the instant `S` becomes 0. Starting from `S = 0`:

- Playing a 0-stone is always legal *except* on the very first move (when `S = 0` already, so `0 + 0 ≡ 0`). Afterward it keeps `S` unchanged.
- From `S = 1`, playing a 2 loses (`1 + 2 ≡ 0`), so you must play a 1, moving to `S = 2`.
- From `S = 2`, playing a 1 loses (`2 + 1 ≡ 0`), so you must play a 2, moving to `S = 1`.

So once `S` is nonzero, the required residue alternates `1, 2, 1, 2, …` with no choices at all. Alice's only real decision in the whole game is her opening residue (1 or 2), plus *when* to burn a pass move.

**Case `c0` even.** Think of the pass moves as tokens both players can spend to skip a turn. With an even number of them, whatever Alice spends Bob can mirror, and vice versa — they can't shift who ends up stuck. So the outcome is decided purely by the `1/2` pattern.

If Alice opens with 1, the demanded sequence is `1, 1, 2, 1, 2, …`; if she opens with 2, it's `2, 2, 1, 2, 1, …`. In either case, having at least one 1 *and* at least one 2 lets Alice pick the opening whose pattern exhausts Bob first. If she is missing a residue class entirely — say `c2 = 0`, so the pile is only 1s and 0s — then the run of 1s just marches `S` through `1, 2, 0` and Alice is the one who eventually hits 0, or the stones simply run out and Bob wins by default. Hence:

```
c0 even  →  Alice wins  ⟺  c1 ≥ 1 and c2 ≥ 1
```

**Case `c0` odd.** Now there's one unpaired pass move, and that leftover tempo flips the parity in Bob's favor. Alice can only overcome it by having so many more of one residue than the other that Bob runs dry no matter how the pass is spent. Working through the alternation, the surplus needs to exceed 2:

```
c0 odd   →  Alice wins  ⟺  |c1 - c2| > 2
```

**Worked example — `stones = [5,1,2,4,3]` (Example 3).** Residues are `5→2, 1→1, 2→2, 4→1, 3→0`, giving `c0 = 1, c1 = 2, c2 = 2`. Here `c0` is odd, and `|2 - 2| = 0`, which is not greater than 2. So the answer is `false` — Bob wins, matching the problem's trace. Notice the counts are perfectly balanced, so the lone 0-stone is exactly the tempo Bob needs.

**Contrast — `stones = [3,1,1,1,1]`.** Now `c0 = 1, c1 = 4, c2 = 0`, and `|4 - 0| = 4 > 2`, so Alice wins. Concretely: Alice opens with a 1 (`S = 1`). If Bob plays the 3, Alice plays a 1 (`S = 2`) and Bob, holding only 1s, is forced to `S = 0` and loses. If instead Bob plays a 1 (`S = 2`), Alice spends the 3 as her pass, and Bob is again stuck with only 1s from `S = 2`. Either branch, Bob loses — the surplus of 1s is what absorbs the extra pass move.

This is a genuinely tricky problem — the 33.6% acceptance rate is earned, and the `> 2` threshold in the odd case is the part most people get wrong on the first attempt. If the parity argument doesn't feel airtight yet, the fastest way to trust it is to write a small memoized brute force over `(c0, c1, c2, S, turn)` for small counts and diff it against the formula. It matches everywhere.

## Complexity Analysis

Time Complexity: O(n) — one pass over `stones` to bucket the remainders; the decision itself is O(1).
Space Complexity: O(1) — a fixed array of three counters, regardless of input size.

## Edge Cases

- **Single stone divisible by 3, e.g. `[3]`.** Alice is forced to take it, the sum is immediately `0 mod 3`, and she loses. Covered by the odd-`c0` branch: `|0 - 0| = 0`, not `> 2`, so `false`. A solution that special-cases "Alice can't open with a 0" separately usually trips here.
- **Single stone not divisible by 3, e.g. `[2]`.** Alice removes it safely, but then no stones remain — and Bob wins automatically when the pile empties. Easy to get wrong if you assume "didn't lose" means "won".
- **All stones divisible by 3, e.g. `[3,6,9]`.** `c1 = c2 = 0`, so there is no legal opening move at all; Alice loses. Both branches of the formula return `false` here, which is the correct behavior.
- **One residue class entirely missing, e.g. `[1,1,1,1]`.** `c0` is even but `c2 = 0`, so `false`. This is the case that shows why the even branch needs *both* counts to be nonzero rather than just checking the total.
- **Exactly the boundary of the odd case, `|c1 - c2| == 2`.** The condition is strictly greater than 2, so a difference of exactly 2 is a Bob win. Off-by-one here (`>=` instead of `>`) is the single most common bug in this problem.
- **Large values.** `stones[i]` goes up to `10^4`, well within `int`, but remember to bucket by `stones[i] % 3` — never by the raw value, and never by the running total, which could otherwise invite overflow thinking that isn't needed here.
- **Maximum size input (`10^5` stones).** The O(n) counting pass handles this comfortably; any approach that recurses over game states will not.
