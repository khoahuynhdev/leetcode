# 2839. Check if Strings Can be Made Equal With Operations I

[LeetCode Link](https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-i/)

Difficulty: Easy
Topics: String
Acceptance Rate: 51.2%

## Hints

### Hint 1

Think about what positions each character can actually reach through the allowed swaps. Since `j - i = 2`, a swap always moves a character by exactly 2 positions. What does that mean for the reachable positions of each character?

### Hint 2

Characters at even indices (0, 2) can only ever swap with each other, and characters at odd indices (1, 3) can only swap with each other. This splits the string into two independent groups. What condition must hold for each group?

### Hint 3

For two strings to be made equal, the set of characters at even positions in `s1` must match the set at even positions in `s2`, and likewise for odd positions. Since each group has only 2 characters, you just need to check if those 2 characters are the same (possibly swapped).

## Approach

Since the strings have length 4 and swaps are constrained to `j - i = 2`, we observe:

- Index 0 can only swap with index 2 (and vice versa).
- Index 1 can only swap with index 3 (and vice versa).

This means within each string, the characters at even indices {0, 2} can be freely rearranged among themselves, and the characters at odd indices {1, 3} can be freely rearranged among themselves. No operation can move a character from an even index to an odd index or vice versa.

Therefore, for `s1` and `s2` to be made equal:

1. The characters at positions 0 and 2 of `s1` must be the same multiset as positions 0 and 2 of `s2`.
2. The characters at positions 1 and 3 of `s1` must be the same multiset as positions 1 and 3 of `s2`.

For a group of two characters, checking the multiset is equivalent to checking either they're already equal, or they're each other's swap.

## Complexity Analysis

Time Complexity: O(1) — the strings are always length 4, so all operations are constant time.

Space Complexity: O(1) — no extra space beyond a few variables.

## Edge Cases

- **Already equal strings**: No swaps needed, should return `true`.
- **Both groups need swapping**: e.g., `s1 = "abcd"`, `s2 = "cdab"` — both even and odd pairs are swapped.
- **One group matches, the other doesn't**: e.g., `s1 = "abcd"`, `s2 = "abdc"` — even positions match but odd positions have characters that aren't a permutation of each other. Must return `false`.
- **All same characters**: e.g., `s1 = "aaaa"`, `s2 = "aaaa"` — trivially `true`.
