# 3761. Minimum Absolute Distance Between Mirror Pairs

[LeetCode Link](https://leetcode.com/problems/minimum-absolute-distance-between-mirror-pairs/)

Difficulty: Medium
Topics: Array, Hash Table, Math
Acceptance Rate: 50.5%

## Hints

### Hint 1

The obvious double loop is `O(n^2)`, which is too slow for `n` up to `10^5`. Think about how to turn the pair condition `reverse(nums[i]) == nums[j]` into a fast lookup instead of a search. What data structure is good at "have I seen value X before, and if so, where?"

### Hint 2

Walk the array left to right. At index `j`, you need to know whether any earlier index `i < j` had `reverse(nums[i]) == nums[j]`, and if so, *which* such `i` makes `j - i` smallest. Since you're moving `j` forward, the closest valid `i` is simply the *most recent* one you have seen — so you only need to remember the latest index per value, not all of them.

### Hint 3

Keep a hash map keyed by the reversed value. Concretely: `lastSeen[reverse(nums[i])] = i`, always overwriting with the newest index. Then at each new position `j`, look up `lastSeen[nums[j]]` directly — if the key exists, `j - lastSeen[nums[j]]` is the best possible distance involving `j` as the right endpoint. Track the running minimum across all `j`. One pass, done.

## Approach

Define `reverse(x)` as the integer formed by reversing the digits of `x`; for `x > 0` this is a short loop that repeatedly peels off `x % 10` into an accumulator and divides `x` by `10`. Leading zeros in the reversal fall away automatically because we are building an integer (e.g. `reverse(120) = 21`).

Now, for each valid pair `(i, j)` with `i < j`, the condition `reverse(nums[i]) == nums[j]` is asymmetric: `i` contributes a *reversed* value, `j` contributes its *own* value. So as we sweep `j` from left to right, we want to ask "has any earlier index `i` produced a reverse equal to `nums[j]`?" A hash map answers this in `O(1)` expected time.

The algorithm:

1. Initialize `lastSeen := map[int]int{}` and `best := -1`.
2. For each `j` from `0` to `n-1`:
   - If `lastSeen[nums[j]]` exists as index `i`, then `(i, j)` is a mirror pair with distance `j - i`. Update `best` if this is smaller (or if `best == -1`).
   - Update `lastSeen[reverse(nums[j])] = j`, so later indices can pair against this one.
3. Return `best`.

Why "most recent" is enough: as `j` grows, `j - i` shrinks as `i` grows. So among all prior indices whose reversed value equals `nums[j]`, the largest (most recent) `i` gives the smallest distance. Overwriting the map entry on every occurrence keeps exactly that "most recent" index.

Walk through Example 1, `nums = [12, 21, 45, 33, 54]`:

- `j=0`, `nums[j]=12`: map empty, no match. Store `lastSeen[reverse(12)] = lastSeen[21] = 0`.
- `j=1`, `nums[j]=21`: `lastSeen[21] = 0` — match! Distance `1 - 0 = 1`, `best = 1`. Store `lastSeen[12] = 1`.
- `j=2`, `nums[j]=45`: no match. Store `lastSeen[54] = 2`.
- `j=3`, `nums[j]=33`: no match. Store `lastSeen[33] = 3` (reverse of `33` is `33`).
- `j=4`, `nums[j]=54`: `lastSeen[54] = 2` — match! Distance `4 - 2 = 2`, `best` stays `1`. Store `lastSeen[45] = 4`.

Result: `1`.

## Complexity Analysis

Time Complexity: `O(n * d)` where `n = len(nums)` and `d` is the number of digits of the largest value (at most 10 for `nums[i] <= 10^9`). Effectively `O(n)`.

Space Complexity: `O(n)` for the hash map, which holds at most `n` distinct reversed values.

## Edge Cases

- **Single element (`n == 1`)**: no valid pair `i < j` exists — return `-1`. The loop naturally handles this because the map is empty when `j = 0` and there is no `j = 1`.
- **No mirror pair anywhere** (e.g. `[21, 120]`): `best` stays `-1` for the whole scan. Note the asymmetry — `reverse(21) = 12 != 120`, so order matters.
- **Palindromic numbers** (e.g. `[11, 11]` or `[33, 33]`): `reverse(x) == x`, so consecutive equal palindromes are themselves mirror pairs. The map logic handles this uniformly.
- **Numbers with trailing zeros** (e.g. `[120, 21]`): reversing drops the leading zero, giving `21`. Integer arithmetic in `reverse` produces this for free — don't be tempted to manipulate strings with padding.
- **Large arrays with many repeats**: overwriting `lastSeen` on each occurrence is essential; without it you'd need to remember old indices you will never use.
- **Minimum value `nums[i] = 1`**: single-digit reverse equals itself, no special case needed.
