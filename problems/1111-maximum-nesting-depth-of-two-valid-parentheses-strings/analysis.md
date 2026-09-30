# 1111. Maximum Nesting Depth of Two Valid Parentheses Strings

[LeetCode Link](https://leetcode.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/)

Difficulty: Medium
Topics: String, Stack, Bracket Sequences
Acceptance Rate: 77.1%

## Hints

### Hint 1

The wordy statement hides a simple question. If the original string has nesting depth `D`, what is the best you could possibly hope for when splitting it into two groups? Think about a lower bound first.

### Hint 2

Track the current nesting depth as you scan the string, just like you would when checking whether parentheses are balanced. Every parenthesis "lives" at some depth level. Try deciding which group a parenthesis goes into based only on its depth level.

### Hint 3

Alternate the levels: send parentheses at odd depths to one group and parentheses at even depths to the other. An opening parenthesis and its matching closing parenthesis always sit at the same depth, so each group stays valid, and each group gets roughly half of the levels — giving `ceil(D / 2)`, which matches the lower bound.

## Approach

**Lower bound.** If `seq` has depth `D`, then at some moment there are `D` open parentheses stacked on top of each other. Those `D` parentheses are split between `A` and `B`, so one of them holds at least `ceil(D / 2)`. So no answer can do better than `ceil(D / 2)`.

**Construction.** Scan the string while maintaining `depth`, the current number of unmatched `(`:

1. On `(`: increment `depth`, then assign this char to group `depth % 2`.
2. On `)`: assign this char to group `depth % 2`, then decrement `depth`.

With this ordering, a `(` and its matching `)` see exactly the same `depth` value, so they land in the same group. Within one group, the parentheses form a properly nested sequence (it is just `seq` with every other level deleted), so both groups are VPS's.

Group 1 gets levels 1, 3, 5, ... and group 0 gets levels 2, 4, 6, .... Each group therefore has at most `ceil(D / 2)` levels, matching the lower bound — so the split is optimal.

**Example:** `seq = "(()())"`

| char | depth used | group |
|------|-----------|-------|
| `(`  | 1 | 1 |
| `(`  | 2 | 0 |
| `)`  | 2 | 0 |
| `(`  | 2 | 0 |
| `)`  | 2 | 0 |
| `)`  | 1 | 1 |

Result `[1,0,0,0,0,1]`: `A = "()()"` (depth 1), `B = "()"` (depth 1). Max depth is 1 = `ceil(2 / 2)`. LeetCode accepts any optimal split, so this is fine even though it differs from the sample output.

No explicit stack is needed. The depth counter is all the "stack" you need.

## Complexity Analysis

Time Complexity: O(n) for one pass over the string.
Space Complexity: O(1) extra space, not counting the O(n) output array.

## Edge Cases

- **Single pair `"()"`**: depth 1. Both chars go to the same group, and the other group is empty (depth 0), which is still a valid VPS.
- **Flat sequence `"()()()"`**: every pair is at depth 1, so everything goes to one group. That is optimal because the answer is `ceil(1/2) = 1`.
- **Deeply nested `"(((())))"`**: this is where alternating levels matters. Depth 4 splits into 2 + 2.
- **Odd depth**: with depth 3 one group gets 2 levels and the other gets 1. You can't do better than `ceil(3/2) = 2`.
- **Order of increment/decrement**: if you increment after assigning `(` (or decrement before assigning `)`), matching parentheses may end up in different groups. That is the most common bug here.
