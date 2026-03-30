# 2840. Check if Strings Can be Made Equal With Operations II

[LeetCode Link](https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-ii/)

Difficulty: Medium
Topics: Hash Table, String, Sorting
Acceptance Rate: 63.8%

## Hints

### Hint 1

Think about what positions you can actually reach from a given index using swaps where `j - i` is even. Can you group indices into categories based on reachability?

### Hint 2

Indices with the same parity (both even or both odd) can swap with each other. This means all characters at even indices can be freely rearranged among even positions, and likewise for odd indices. Consider treating even-indexed and odd-indexed characters as two separate groups.

### Hint 3

If you can freely rearrange characters within each parity group, then two strings are equivalent if and only if the multiset of characters at even indices matches, AND the multiset of characters at odd indices matches. You can check this by sorting or by counting character frequencies.

## Approach

The key insight is understanding what the swap operation allows. You can swap characters at indices `i` and `j` when `j - i` is even. This means both `i` and `j` have the same parity (both even or both odd).

Through a chain of such swaps, any character at an even index can eventually be moved to any other even index, and any character at an odd index can be moved to any other odd index. This is because swaps within a parity group are transitive: if you can swap positions 0 and 2, and positions 2 and 4, you can effectively move a character from position 0 to position 4.

This means we can independently rearrange:
- All characters at even indices (0, 2, 4, ...)
- All characters at odd indices (1, 3, 5, ...)

So `s1` can be transformed into `s2` if and only if:
1. The sorted characters at even indices of `s1` equal the sorted characters at even indices of `s2`
2. The sorted characters at odd indices of `s1` equal the sorted characters at odd indices of `s2`

Instead of sorting, we can use frequency counting (an array of size 26) for each parity group, which is slightly more efficient.

**Walkthrough with Example 1:** `s1 = "abcdba"`, `s2 = "cabdab"`
- Even indices (0,2,4): s1 has `a,c,b` → sorted `a,b,c`; s2 has `c,b,a` → sorted `a,b,c` ✓
- Odd indices (1,3,5): s1 has `b,d,a` → sorted `a,b,d`; s2 has `a,d,b` → sorted `a,b,d` ✓
- Both match → return `true`

## Complexity Analysis

Time Complexity: O(n) — one pass through both strings to count character frequencies for each parity group, plus O(26) to compare frequency arrays.

Space Complexity: O(1) — four fixed-size arrays of 26 integers.

## Edge Cases

- **Single character strings**: Both strings must simply be equal since there are no valid swaps.
- **Already equal strings**: Should return `true` immediately (the frequency counts will naturally match).
- **All same characters**: Any arrangement is reachable, should return `true`.
- **Mismatch only in one parity group**: e.g., even indices match but odd indices don't — must return `false`.
