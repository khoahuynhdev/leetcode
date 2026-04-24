# 2833. Furthest Point From Origin

[LeetCode Link](https://leetcode.com/problems/furthest-point-from-origin/)

Difficulty: Easy
Topics: String, Counting
Acceptance Rate: 69.0%

## Hints

### Hint 1

Think about what each character contributes to your final position. `'L'` pulls you one unit in the negative direction, `'R'` pushes you one unit in the positive direction, and `'_'` is a free move you get to choose. The question is really about how to spend those free moves to maximize the absolute displacement.

### Hint 2

You do not actually need to simulate the moves in order — position on a number line is commutative. Just count how many of each character the string contains. Once you have those three counts, ask yourself: if I already have some imbalance between `L` and `R`, what is the optimal thing to do with every `_`?

### Hint 3

The furthest you can get is to pick a direction and commit every `_` to it. If `L > R`, send every `_` left; if `R > L`, send every `_` right; if they are tied, it does not matter. In every case the final distance from the origin is `|count(L) - count(R)| + count(_)`. No simulation needed — one pass to count, then plain arithmetic.

## Approach

1. Scan the string once, tallying the number of `'L'`, `'R'`, and `'_'` characters.
2. The two directional moves cancel out pairwise, leaving a fixed net displacement of `|L - R|` away from the origin (in whichever direction is larger).
3. Every `'_'` is a free move, and since we only care about *distance from origin* (not a specific side), the optimal play is to push all of them in the same direction that already dominates — this strictly increases the absolute displacement.
4. Return `abs(L - R) + underscores`.

**Walkthrough with `moves = "L_RL__R"`:**
- `L` count = 2, `R` count = 2, `_` count = 3.
- `|2 - 2| = 0`, plus 3 free moves. Send all free moves one direction (say left): net position is `-3`, distance is `3`. ✓

**Walkthrough with `moves = "_R__LL_"`:**
- `L` count = 2, `R` count = 1, `_` count = 4.
- `|2 - 1| = 1`, plus 4 free moves all sent left: net position is `-5`, distance is `5`. ✓

## Complexity Analysis

Time Complexity: O(n) — a single linear scan over the input string.
Space Complexity: O(1) — only three integer counters are maintained, independent of input size.

## Edge Cases

- **All underscores** (`"_______"`): no fixed moves to cancel, every free move goes one way, answer equals the string length.
- **All one direction** (`"LLLL"` or `"RRRR"`): no free moves at all, answer equals the count of that direction.
- **Perfectly balanced L and R with no underscores** (`"LR"`, `"LLRR"`): `|L - R| = 0` and no free moves, answer is `0`.
- **Single character** (`"L"`, `"R"`, or `"_"`): smallest valid input per constraints; the formula handles it naturally (answer is always `1`).
- **Equal L and R with underscores** (`"LR__"`): fixed moves cancel; the answer is exactly the number of underscores.
