# 3863. Minimum Operations to Sort a String

[LeetCode Link](https://leetcode.com/problems/minimum-operations-to-sort-a-string/)

Difficulty: Medium
Topics: String
Acceptance Rate: 26.8%

## Hints

### Hint 1

Don't try to simulate the operations or search over them. The answer can only be a handful of values. Ask yourself: what is the largest number of operations you could ever need?

### Hint 2

Any single operation leaves out at least one end of the string: either index `0` or index `n-1` stays put. So focus on the two endpoints. What must be true about `s[0]` (or `s[n-1]`) for one operation to be enough?

### Hint 3

In the final sorted string, `s[0]` is the smallest character and `s[n-1]` is the largest. One operation works exactly when `s[0]` is already the minimum (sort `s[1:]`) or `s[n-1]` is already the maximum (sort `s[:n-1]`). If neither is true, you can usually fix one end with a first sort and finish with a second. The only case that fails is when the minimum appears **only** at the last index and the maximum appears **only** at the first index. That case takes 3.

## Approach

Let `n = len(s)`, `mn = min(s)`, `mx = max(s)`.

1. **Already sorted → 0.** This also covers `n == 1`.
2. **`n == 2` and unsorted → -1.** The only substrings you're allowed to sort are single characters, and sorting those does nothing.
3. **One operation.** One operation never touches index 0 or never touches index n-1.
   - If `s[0] == mn`, sort `s[1:]`. The string is now sorted.
   - If `s[n-1] == mx`, sort `s[:n-1]`. Same idea.
   - If neither holds, one operation can't work. Whichever end it leaves alone is in the wrong place.
4. **Two operations.**
   - If `mn` shows up anywhere in `s[0..n-2]`, sort `s[:n-1]`. That moves `mn` to index 0, and then sorting `s[1:]` finishes the job.
   - If `mx` shows up anywhere in `s[1..n-1]`, do the mirror image: sort `s[1:]`, then `s[:n-1]`.
5. **Three operations.** Otherwise `mn` appears only at `n-1` and `mx` appears only at `0`. Two operations can't work. The first operation leaves one end alone, so either `mn` is still stuck at the last index or `mx` is still stuck at the first index. To fix that, the second operation would need to cover both index 0 and index n-1, which means the whole string, and that's not allowed. Three operations are enough: sort `s[:n-1]` (moves `mx` off index 0), sort `s[1:]` (moves `mx` to the end), then sort `s[:n-1]`.

Example: `"card"`. `mn = 'a'` and `mx = 'r'`. `s[0] = 'c'` isn't the minimum and `s[3] = 'd'` isn't the maximum, so the answer isn't 1. `'a'` sits at index 1, which is inside `s[:n-1]`, so the answer is 2.

Example: `"cba"`. `'a'` appears only at the end and `'c'` appears only at the start, so the answer is 3: `"cba"` → `"bca"` → `"bac"` → `"abc"`.

The tricky part is proving these cases are complete. The idea that unlocks it is that every operation must leave at least one endpoint alone.

## Complexity Analysis

Time Complexity: O(n). There are a few linear scans: the sortedness check, min/max, and the position checks.
Space Complexity: O(1)

## Edge Cases

- **Single character:** it's already sorted, so the answer is 0. Make sure the `n == 2` impossibility check doesn't catch it.
- **Already sorted string (including all characters equal):** the answer is 0. Check this before anything else.
- **Length 2, unsorted (e.g. `"ba"`):** the answer is -1, the only impossible case.
- **Duplicated min or max (e.g. `"zbza"`):** an extra copy of `mx` away from index 0, or of `mn` away from index n-1, drops the answer from 3 to 2. The "only at" condition has to be checked carefully.
- **Strictly reversed strings of length ≥ 3 (e.g. `"cba"`):** this is the worst case, and the answer is 3.
