# 0761. Special Binary String

[LeetCode Link](https://leetcode.com/problems/special-binary-string/)

Difficulty: Hard
Topics: String, Divide and Conquer, Sorting
Acceptance Rate: 68.2%

## Hints

### Hint 1

Think about what special binary strings look like structurally. If you map `1` to `(` and `0` to `)`, what does the definition remind you of? Consider how valid parentheses sequences can be decomposed.

### Hint 2

A special binary string is equivalent to a valid (balanced) parentheses sequence. Any valid parentheses sequence can be broken into top-level groups — substrings where the running balance first returns to zero. Each of these groups is itself a special binary string wrapped in an outer pair `1...0`. Use recursion to process the inside of each group, then think about how to arrange the groups for the largest result.

### Hint 3

After recursively making each top-level special substring as large as possible internally, sort all the top-level substrings in descending lexicographic order and concatenate them. This works because swapping adjacent special substrings is exactly the operation needed to reorder them, and any permutation of adjacent elements can be reached through a sequence of adjacent swaps.

## Approach

Map the problem to balanced parentheses: treat `1` as `(` and `0` as `)`. A special binary string is exactly a balanced parentheses expression.

**Algorithm (recursive):**

1. Scan the string and find all top-level special substrings. A top-level substring is one where the running balance (increment on `1`, decrement on `0`) returns to zero.
2. For each top-level substring, it must start with `1` and end with `0`. Strip the outer `1` and `0`, and recursively apply the same algorithm to the inner portion.
3. After recursing, wrap each result back with `1` and `0` to restore the outer layer.
4. Sort all the processed top-level substrings in **descending** lexicographic order.
5. Concatenate them to form the answer.

**Why sorting in descending order works:**

Since we can swap any two adjacent special substrings, we can achieve any permutation of the top-level substrings. To maximize lexicographic order, we simply place the largest substring first. Descending sort achieves this.

**Example walkthrough with `"11011000"`:**

- Top-level substrings: the entire string `"11011000"` is one top-level group (balance returns to 0 only at the end).
- Strip outer `1` and `0` → inner = `"101100"`.
- Recurse on `"101100"`: top-level groups are `"10"` and `"1100"`.
  - `"10"` → strip → `""` → recurse → `""` → wrap → `"10"`.
  - `"1100"` → strip → `"10"` → recurse → `"10"` → wrap → `"1100"`.
- Sort descending: `["1100", "10"]` → concatenate → `"110010"`.
- Wrap back: `"1" + "110010" + "0"` = `"11100100"`.

## Complexity Analysis

Time Complexity: O(n^2) — In each recursive call, we scan the string to find top-level groups and sort them. The total work across all levels of recursion is bounded by O(n^2) in the worst case (each level processes O(n) characters and there can be O(n) levels of nesting).

Space Complexity: O(n^2) — Due to recursive calls and string construction at each level. The recursion depth can be O(n) and each level may create strings of length O(n).

## Edge Cases

- **Minimal special string `"10"`**: Already the largest possible; should return unchanged.
- **Already maximal string**: If the string is already in descending order at every recursive level, the algorithm returns it unchanged.
- **Deeply nested string like `"1111...0000"`**: Only one top-level group at each level of recursion; no sorting needed, just recurse into the interior.
- **Flat structure like `"101010"`**: Multiple top-level groups at the same level; sorting them descending produces `"101010"` (all are `"10"`, so order doesn't matter).
- **Single pair `"10"`**: Base-like case with minimal input length.
