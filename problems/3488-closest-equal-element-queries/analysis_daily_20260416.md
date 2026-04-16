# 3488. Closest Equal Element Queries

[LeetCode Link](https://leetcode.com/problems/closest-equal-element-queries/)

Difficulty: Medium
Topics: Array, Hash Table, Binary Search
Acceptance Rate: 38.2%

## Hints

### Hint 1

Think about *where* the equal elements live. For every query, you only care about the indices in `nums` that share the value at the queried index. What data structure lets you collect "all indices for each distinct value" in a single pass? Once you have that mapping, the problem stops being about the whole array and becomes about a small sorted list of indices per value.

### Hint 2

For a circular array of length `n`, the distance between indices `a` and `b` is `min(|a - b|, n - |a - b|)`. Given that, if you have the *sorted* list of all indices carrying a particular value, which two indices in that list are candidates for the closest one to a queried index `q`? Convince yourself it is only the immediate predecessor and immediate successor of `q` in the cyclically-ordered list — anything further out is strictly farther away in the same direction.

### Hint 3

Precompute, for every index `i`, two things: `nextSame[i]` (the next index going right with the same value, wrapping around) and `prevSame[i]` (the previous index going left with the same value, wrapping around). Because the indices for a given value are already encountered in increasing order when scanning `nums`, you get these essentially for free by bucketing with a hash map. Each query then becomes O(1): compute the forward distance `(nextSame[q] - q + n) % n` and the backward distance `(q - prevSame[q] + n) % n`, then take the minimum. If the bucket for `nums[q]` has size 1, return -1.

## Approach

1. **Bucket indices by value.** Build a `map[int][]int` where the key is a value from `nums` and the value is the list of indices (already sorted because we iterate left-to-right) at which that value occurs.

2. **Precompute circular neighbors.** For each bucket `idxList` of size `k >= 2`, walk through it and for the element at position `m`:
   - `nextSame[idxList[m]] = idxList[(m+1) % k]`
   - `prevSame[idxList[m]] = idxList[(m-1+k) % k]`

   Buckets of size 1 leave their index with the sentinel `-1`, meaning "no twin exists".

3. **Answer each query in O(1).** For query index `q`:
   - If `nextSame[q] == -1`, answer is `-1`.
   - Otherwise, compute `forward = (nextSame[q] - q + n) % n` and `backward = (q - prevSame[q] + n) % n`. The answer is `min(forward, backward)`.

**Why only neighbors matter.** Consider the sorted cyclic list of same-value indices. Starting from `q` and walking forward (right, with wrap-around), the very first same-value index you hit is `nextSame[q]`; any further same-value index is reached by continuing forward, so it is strictly farther in the forward direction. The same argument holds backward. The global minimum circular distance is therefore the smaller of these two one-step distances.

**Walkthrough on example 1.** `nums = [1,3,1,4,1,3,2]`, `n = 7`.

- Buckets: `1 -> [0,2,4]`, `3 -> [1,5]`, `4 -> [3]`, `2 -> [6]`.
- `nextSame / prevSame`: index 0 → (2, 4), index 2 → (4, 0), index 4 → (0, 2), index 1 → (5, 5), index 5 → (1, 1); indices 3 and 6 have no twin.
- Query 0: forward `(2-0+7)%7 = 2`, backward `(0-4+7)%7 = 3` → **2**.
- Query 3: no twin → **-1**.
- Query 5: forward `(1-5+7)%7 = 3`, backward `(5-1+7)%7 = 4` → **3**.

Matches the expected output `[2, -1, 3]`.

## Complexity Analysis

Time Complexity: O(n + q), where `n = len(nums)` and `q = len(queries)`. One pass builds the buckets and neighbor arrays; each query is O(1).
Space Complexity: O(n), for the hash map buckets and the two neighbor arrays.

## Edge Cases

- **Unique values.** Every bucket has size 1, so every query returns `-1` (matches example 2).
- **Exactly two occurrences.** The two indices are each other's `nextSame` and `prevSame`, but the forward and backward circular distances generally differ; the minimum is the correct answer (e.g., value `3` in example 1, query index 5).
- **All elements identical.** Every index's closest twin is the neighbor ±1, yielding answer `1` for every query (except when `n == 1`, where no twin exists).
- **`n == 1`.** The single element has no twin, so any query returns `-1`. The bucket logic handles this naturally because the bucket has size 1.
- **Query index equals the only occurrence of its value.** Same as the `n == 1` case conceptually — no twin, return `-1`.
- **Wrap-around is the shorter path.** The forward or backward distance may cross the array boundary; the `(... + n) % n` formulation handles both cases uniformly without special-casing.
