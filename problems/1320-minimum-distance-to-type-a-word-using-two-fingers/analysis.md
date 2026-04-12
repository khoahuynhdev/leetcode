# 1320. Minimum Distance to Type a Word Using Two Fingers

[LeetCode Link](https://leetcode.com/problems/minimum-distance-to-type-a-word-using-two-fingers/)

Difficulty: Hard
Topics: String, Dynamic Programming

Acceptance Rate: 62.1%

## Hints

### Hint 1

You have to type the letters of `word` in order, but you can assign each letter to either of two fingers. That freedom is exactly the sort of branching decision that Dynamic Programming is good at — think in terms of "state after typing the i-th letter" rather than trying to greedily pick which finger goes where.

### Hint 2

When designing a state, notice something important: after typing `word[i]`, the finger that just typed it is *pinned* to `word[i]`. So you only need to remember where the **other** finger is. That collapses the state from "positions of both fingers" to "position of the idle finger".

### Hint 3

Use `dp[i][o]` = minimum distance to type the first `i+1` letters such that the "other" (non-active) finger sits at letter `o`. For letter `word[i+1]`:

1. The active finger (sitting on `word[i]`) moves to `word[i+1]` — cost `dist(word[i], word[i+1])`, state `o` unchanged.
2. The idle finger at `o` moves to `word[i+1]` — cost `dist(o, word[i+1])`, new "other" position becomes `word[i]`.

Handle the freebie: the second finger's *first* placement costs nothing. Encode this by starting with an extra "unplaced" sentinel state for the other finger whose move-cost is `0`.

## Approach

Map each uppercase letter to a grid coordinate with 6 columns: letter `L` sits at `(L/6, L%6)`. Verify with the examples: `A=0 -> (0,0)`, `B=1 -> (0,1)`, `P=15 -> (2,3)`, `Z=25 -> (4,1)`. The distance between two letters is the Manhattan distance of their coordinates.

**State.** Let `dp[o]` represent the minimum cost to have typed the word up to the current character, where the finger that just typed the current character is at that character, and the *other* finger is at letter index `o ∈ {0..25}`, or `26` meaning "not yet placed" (still free to drop anywhere at no cost).

**Initialization.** After typing `word[0]`, one finger is on `word[0]` (placement is free) and the other is unplaced: `dp[26] = 0`, everything else `= ∞`.

**Transition.** For each next letter with `prev = word[i-1]` and `curr = word[i]`, compute `next[]` from `dp[]`:

- *Same finger moves:* `next[o] = min(next[o], dp[o] + dist(prev, curr))` — the other finger stays put.
- *Other finger moves:* `next[prev] = min(next[prev], dp[o] + moveCost)` where `moveCost = 0` if `o == 26` (free first placement) else `dist(o, curr)`. The finger previously at `prev` becomes the "other" finger.

**Answer.** `min(dp[o])` over all `o` after processing the full word.

**Walkthrough on "CAKE".** After 'C' (free): `dp[26] = 0`. Typing 'A' (from 'C'): keep finger ⇒ cost 2; switch to fresh finger (free) ⇒ cost 0, leaving 'C' as the idle finger. Typing 'K' (from 'A'): using 'C'-finger it's distance 3, total 3, but better to let the 'C'-idle finger type 'K' freshly only when cheaper. Typing 'E' (from 'K'): best path is idle finger 'A' walks to 'E' for 4 (total 4) *or* the K-finger walks to 'E' for 1 while the A-finger stays — giving total `0 + 2 + 0 + 1 = 3`. The DP finds this minimum.

## Complexity Analysis

Time Complexity: O(n · 27) = O(n), where n = `len(word)`. For each character we iterate 27 "other finger" states with O(1) work.
Space Complexity: O(27) = O(1). Two rolling arrays of size 27.

## Edge Cases

- **`n == 2`:** Both fingers can be placed for free, so the answer is always `0`. The DP handles this because after the first character `dp[26] = 0`, and the "switch to fresh finger" transition gives a `0`-cost placement on the second character.
- **Repeated letters (e.g., `"AAA"`, `"HAPPY"`):** Distance to the same letter is `0`, so the DP naturally favors staying put. `"HAPPY"` exercises this via the `P→P` step.
- **Long runs on one finger vs. alternating:** The DP considers both alternatives at every step, so there is no need to pre-commit to a strategy.
- **Distance formula:** Use `|x1-x2| + |y1-y2|` (Manhattan), not Euclidean — an easy bug to introduce.
- **Free initial placement:** If you don't model the "unplaced" sentinel, the first move of the second finger gets incorrectly charged.
