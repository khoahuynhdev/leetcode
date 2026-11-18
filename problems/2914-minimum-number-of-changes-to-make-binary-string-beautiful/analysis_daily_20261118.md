# 2914. Minimum Number of Changes to Make Binary String Beautiful

[LeetCode Link](https://leetcode.com/problems/minimum-number-of-changes-to-make-binary-string-beautiful/)

Difficulty: Medium
Topics: String
Acceptance Rate: 77.6%

## Hints

### Hint 1

Don't try to search over all the ways to partition the string. Ask yourself: what is the *smallest* even-length block that could exist, and can every valid partition be broken down into those?

### Hint 2

Any uniform block of even length (like `0000` or `111111`) can be cut into blocks of length 2 that are still uniform. So you only need to think about the string as fixed, non-overlapping pairs: `s[0..1]`, `s[2..3]`, `s[4..5]`, ...

### Hint 3

Each pair `(s[i], s[i+1])` for even `i` is independent of all the others. If the two characters already match, it costs 0. If they differ, one change fixes it. The answer is just the number of mismatched pairs.

## Approach

**Key observation:** A string is beautiful exactly when every aligned pair `s[2k], s[2k+1]` has two equal characters.

- *If the string is beautiful*, each part of the partition has even length and starts at an even index (since all earlier parts have even length). So each part covers whole aligned pairs, and each of those pairs is uniform.
- *If every aligned pair is uniform*, then partitioning into the pairs themselves is a valid beautiful partition.

So the problem reduces to: make every aligned pair uniform using as few changes as possible. The pairs don't overlap, so they don't affect each other:

- Pair already equal (`00` or `11`): 0 changes.
- Pair different (`01` or `10`): exactly 1 change (flip either character).

Algorithm:
1. Set `count = 0`.
2. For `i = 0, 2, 4, ...` up to `len(s) - 2`: if `s[i] != s[i+1]`, increase `count`.
3. Return `count`.

Example: `s = "1001"` → pairs `10`, `01` → both mismatched → answer 2.

Example: `s = "110100"` → pairs `11`, `01`, `00` → one mismatch → answer 1.

This is a "looks hard, turns out simple" problem. The hard part is spotting that length-2 blocks are enough; the code after that is tiny.

## Complexity Analysis

Time Complexity: O(n), one pass over the string looking at n/2 pairs.
Space Complexity: O(1), just a counter.

## Edge Cases

- **Shortest string (length 2)**: there's only one pair. The answer is 0 or 1.
- **Already beautiful string** (for example `"0000"` or `"1100"`): the answer should be 0.
- **Every pair mismatched** (for example `"101010"`): the answer is `n/2`, which is the maximum possible.
- **Unaligned runs** (for example `"0110"`): the two `1`s are next to each other but sit in different aligned pairs, so you still need 2 changes. This catches solutions that count run lengths instead of looking at aligned pairs.
