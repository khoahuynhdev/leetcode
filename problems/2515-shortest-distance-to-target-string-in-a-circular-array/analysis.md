# 2515. Shortest Distance to Target String in a Circular Array

[LeetCode Link](https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/)

Difficulty: Easy
Topics: Array, String
Acceptance Rate: 54.5%

## Hints

### Hint 1

Think about how a circular array changes the notion of "distance" between two indices. For each candidate index `i` where `words[i] == target`, what are the two possible ways to reach it from `startIndex`?

### Hint 2

You only need to scan the array once. For each index `i` where `words[i] == target`, compute the forward distance and the backward distance, then take the smaller of the two. Track the global minimum across all such indices.

### Hint 3

If the array length is `n`, the forward distance from `startIndex` to `i` is `(i - startIndex + n) % n` and the backward distance is `(startIndex - i + n) % n`. Equivalently, once you have one of them (say `d = (i - startIndex + n) % n`), the other is just `n - d`. The shortest distance for that index is `min(d, n - d)`.

## Approach

Iterate through the array a single time. For every index `i`, check whether `words[i]` equals `target`. If it does, compute:

- `forward = (i - startIndex + n) % n` — the number of right-moves needed to reach `i` from `startIndex` going forward.
- `backward = n - forward` (equivalently `(startIndex - i + n) % n`) — the number of left-moves needed.

The shortest distance to that particular occurrence is `min(forward, backward)`. Maintain a running minimum across all matching indices and return it. If no occurrence of `target` was found, return `-1`.

For example, with `words = ["hello","i","am","leetcode","hello"]`, `target = "hello"`, `startIndex = 1`, `n = 5`:

- `i = 0`: `forward = (0 - 1 + 5) % 5 = 4`, `backward = 1`, min = 1.
- `i = 4`: `forward = (4 - 1 + 5) % 5 = 3`, `backward = 2`, min = 2.

Overall minimum = 1.

This works because moving in a circular array of length `n` between two indices can only happen in two directions, and these two distances always sum to `n`.

## Complexity Analysis

Time Complexity: O(n * L), where `n` is the number of words and `L` is the maximum word length (string equality check is O(L)).
Space Complexity: O(1), only constant extra space is used.

## Edge Cases

- **Target not present**: No index matches; the result must be `-1`.
- **`startIndex` itself is the target**: The forward distance is `0`, which is the minimum possible answer.
- **Multiple occurrences of target**: Must consider every match and take the global minimum.
- **Length-1 array**: `n = 1` means `startIndex = 0` and either the single word equals `target` (distance 0) or not (`-1`).
- **Target at the "wrap-around" boundary**: e.g., target at index `0` while `startIndex = n - 1` — going forward by 1 step is shorter than walking back through the array.
