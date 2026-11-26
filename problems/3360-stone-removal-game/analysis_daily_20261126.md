# 3360. Stone Removal Game

[LeetCode Link](https://leetcode.com/problems/stone-removal-game/)

Difficulty: Easy
Topics: Math, Simulation
Acceptance Rate: 51.0%

## Hints

### Hint 1

Despite sounding like a game-theory problem, neither player has any choice to make — every move is forced. Think about whether you even need strategy at all.

### Hint 2

Since every move is predetermined (10, then 9, then 8, ...), you can just play the game out. Track how many stones remain, how many must be removed next, and whose turn it is.

### Hint 3

The game ends the moment the current player needs to remove more stones than remain. That player loses — so the winner is whoever made the *last successful* move. Because the removal amount only shrinks from 10, there are at most 10 moves, so simulation is effectively O(1).

## Approach

1. Start with `take = 10` and `aliceTurn = true`.
2. While `n >= take`, the current player can move:
   - subtract `take` from `n`,
   - decrement `take` (the next player removes one fewer),
   - flip `aliceTurn`.
3. When the loop stops, the player whose turn it is cannot move and loses. So Alice wins exactly when it is **not** her turn: return `!aliceTurn`.

Example, `n = 27`:
- Alice removes 10 → 17 left.
- Bob removes 9 → 8 left.
- Alice removes 8 → 0 left.
- Bob must remove 7 but 0 remain → Bob loses, Alice wins.

A math-flavored alternative: Alice wins iff the number of completed moves is odd. You could compute the cumulative sums 10, 19, 27, 34, ... and find the largest prefix that fits in `n`, but the simulation is just as fast and much clearer.

## Complexity Analysis

Time Complexity: O(1) — at most 10 iterations, since the removal amount drops from 10 and the total 10+9+...+1 = 55 already exceeds the max `n` of 50.
Space Complexity: O(1)

## Edge Cases

- **n < 10**: Alice cannot make her very first move and loses immediately (e.g., `n = 1`, `n = 9`).
- **n exactly equals a cumulative sum** (10, 19, 27, ...): the pile is emptied exactly; the *next* player can't move. Make sure you use `n >= take` (not `>`), otherwise you'd wrongly say the mover failed.
- **Pile emptied to 0**: removing 0 is never required, so a 0-stone pile always means the next player loses.
- **Maximum n = 50**: verifies the loop terminates correctly after several moves (10+9+8+7+6+5+4 = 49, Bob can't take 3 → Alice wins).
