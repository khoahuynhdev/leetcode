# 0856. Score of Parentheses

[LeetCode Link](https://leetcode.com/problems/score-of-parentheses/)

Difficulty: Medium
Topics: String, Stack, Bracket Sequences
Acceptance Rate: 63.9%

## Hints

### Hint 1

The score is defined recursively on nested structure. Which data structure naturally tracks "how deep am I" and "what has been accumulated at each level" as you scan brackets left to right?

### Hint 2

Try a stack where each entry holds the running score of the current nesting level. When you open a bracket, start a fresh level with score 0. When you close it, the inner level is finished. How should its score fold into the level beneath it?

Alternatively, think about depth alone: can you compute the answer just by counting how deep you are at certain moments?

### Hint 3

Only the innermost `()` pairs actually produce points. Every other bracket just doubles whatever is inside it. A `()` sitting at depth `d` (the number of pairs wrapped around it) therefore adds exactly `2^d` to the total. Scan once, track depth, and whenever you see `(` immediately followed by `)`, add `1 << depth`.

## Approach

**Key observation:** Expand the rules. `(A)` = `2 * A` and `AB` = `A + B` means the total is a sum over every core `()` pair, each multiplied by 2 once for every pair enclosing it. For example:

```
( ( ) ( ( ) ) )
```

- The first `()` sits inside 1 pair: contributes `2^1 = 2`.
- The second `()` sits inside 2 pairs: contributes `2^2 = 4`.
- Total = 6. Check with the rules: `(() (()))` = `2 * (1 + 2*1)` = `2 * 3` = 6. ✓

**Algorithm (O(1) space):**

1. Keep `depth = 0` and `score = 0`.
2. For each index `i`:
   - If `s[i] == '('`, increment `depth`.
   - Otherwise decrement `depth`. If the previous character was `(`, this is a core `()` pair, so add `1 << depth` (using depth *after* decrementing, which equals the number of enclosing pairs).
3. Return `score`.

**Stack alternative (also O(n)):** Push 0 when you see `(`. On `)`, pop the inner score `v` and add `max(2*v, 1)` to the new top. Start with a single 0 on the stack for the outermost level. This mirrors the grammar directly and is a great approach to mention first in an interview before optimizing to the depth trick.

Both work because the grammar is linear in its sub-scores. Doubling distributes over addition, so `2 * (A + B) = 2A + 2B`, which lets each core pair be counted on its own.

## Complexity Analysis

Time Complexity: O(n), with a single pass over the string.
Space Complexity: O(1) for the depth-counting approach (O(n) for the stack variant).

## Edge Cases

- **Single pair `"()"`:** The minimal input. The depth after closing is 0, so it contributes `2^0 = 1`.
- **Deep nesting such as `"((((()))))"`:** Only one core pair, at depth 4, gives 16. This checks that the doubling is applied correctly. With length ≤ 50 the maximum depth is 24, so `1 << 24` fits easily in an `int`.
- **Pure concatenation such as `"()()()"`:** Every pair is at depth 0, so the scores add up to 3 with no doubling.
- **Mixed structure such as `"(()(()))"`:** Core pairs sit at different depths. A naive approach that counts only the maximum depth would give the wrong answer here.
- **Closing brackets that are not core pairs:** A `)` that follows another `)` must not add anything. It only reduces depth.
