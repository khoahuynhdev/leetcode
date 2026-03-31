# 3474. Lexicographically Smallest Generated String

[LeetCode Link](https://leetcode.com/problems/lexicographically-smallest-generated-string/)

Difficulty: Hard
Topics: String, Greedy, String Matching
Acceptance Rate: 42.2%

## Hints

### Hint 1

Think about which positions in the result string are "fixed" (determined by T constraints) and which are "free" (you get to choose). For T positions, you're stamping str2 onto the result. What happens when two T stamps overlap?

### Hint 2

After applying all T constraints, the F constraints say certain windows must NOT match str2. For each F window, if at least one fixed position already differs from str2, the constraint is automatically satisfied. Otherwise, you need at least one free position in the window to differ. Consider which free position gives you the best lexicographic result.

### Hint 3

Process free positions left to right, greedily assigning 'a' (the smallest character). The critical insight: an F constraint only becomes urgent at its **rightmost free position** -- that's the last chance to break the match. At each free position, collect the "forbidden" characters from all F windows whose rightmost free position is here, and pick the smallest character not forbidden.

## Approach

1. **Apply T constraints**: For each index i where str1[i] == 'T', stamp str2 onto word[i..i+m-1]. If two T stamps overlap and disagree on a character, return "".

2. **Analyze F constraints**: For each index i where str1[i] == 'F':
   - If any fixed position in the window already differs from str2, the constraint is automatically satisfied (mark as "broken").
   - Otherwise, find the rightmost free position in the window. If there is none, the window is fully fixed and matches str2 -- impossible, return "".
   - Group unbroken F windows by their rightmost free position.

3. **Greedy fill (left to right)**: For each free position j:
   - Check which unbroken F windows have their rightmost free position at j. For each such window i, str2[j-i] is a "forbidden" character (using it would complete the match).
   - Set word[j] to the smallest character ('a'..'z') not in the forbidden set.
   - After setting word[j], mark any F window containing j as "broken" if word[j] differs from str2 at the corresponding offset.

This works because the rightmost free position is the last opportunity to break an F match. By deferring the constraint to the rightmost position, all earlier positions stay as small as possible.

## Complexity Analysis

Time Complexity: O(n * m) -- applying T constraints is O(n * m), analyzing F windows is O(n * m), and the greedy fill processes each free position in O(m) time for updating broken status.

Space Complexity: O(n + m) -- for the word array, fixed/broken arrays, and the mapping from positions to F window indices.

## Edge Cases

- **Overlapping T conflicts**: Two T stamps that overlap but disagree on a character make the problem impossible.
- **F window fully fixed**: If an F window has no free positions and all fixed positions match str2, it's impossible.
- **All 26 characters forbidden**: If a free position is the rightmost free position for enough F windows that all 26 characters are forbidden, return "".
- **str2 contains characters other than 'a'**: Free positions set to 'a' may automatically satisfy F constraints without any changes, since 'a' != str2[offset].
- **Single character str2**: Each F constraint reduces to a single position that must differ from str2[0].
