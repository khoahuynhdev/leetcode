# 1477. Find Two Non-overlapping Sub-arrays Each With Target Sum

[LeetCode Link](https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/)

Difficulty: Medium
Topics: Array, Hash Table, Binary Search, Dynamic Programming, Sliding Window
Acceptance Rate: 40.4%

## Hints

### Hint 1

Read the constraints carefully: `1 <= arr[i] <= 1000`. Every element is strictly positive, which means the prefix sums are strictly increasing. That single fact unlocks a much cheaper tool than a hash map of prefix sums — as the right end of a window moves forward, the left end that keeps the window sum at `target` can only move forward too. What technique does a monotone window enable?

### Hint 2

Start by solving the easier half of the problem: enumerate *every* sub-array whose sum equals `target`. With a two-pointer window you can do this in one pass, and there are at most `n` of them (one per right endpoint). Now the real question is how to pair two of them without overlap. Pairing every candidate with every other candidate is O(n²) — too slow. Think about fixing one of the two sub-arrays and asking a question about "everything before it".

### Hint 3

Fix the *second* sub-array. If it occupies `arr[l..r]`, the first one must live entirely inside `arr[0..l-1]`, and to minimize the total you want the shortest valid sub-array anywhere in that prefix. So precompute (or maintain on the fly) `best[i]` = the length of the shortest valid sub-array contained in `arr[0..i-1]`, with `best[i] = min(best[i-1], length of a window ending exactly at i-1)`. Since `best` is a running prefix minimum, it can be filled in the *same* sweep as the sliding window: when the window `[l..r]` hits `target`, the answer candidate is `(r-l+1) + best[l]`. One pass, no nested loop, no binary search needed.

## Approach

The problem has two independent pieces: *find* the candidate sub-arrays, and *combine* two of them optimally.

**Finding candidates.** Because all elements are positive, a sliding window works. Keep `left`, `right`, and the running `sum` of `arr[left..right]`. Extend `right` by one element at a time and add it to `sum`. While `sum > target`, shrink from the left (`sum -= arr[left]; left++`). Shrinking is safe because dropping a positive element can only decrease the sum, and once the sum drops to or below `target` no further left position for this `right` can reach `target` again. After the shrink loop, `sum == target` exactly when the window `arr[left..right]` is a valid sub-array — and it is the *only* valid one ending at `right`, again because prefix sums are strictly increasing. So the sweep yields at most one candidate per right endpoint, i.e. at most `n` candidates, in O(n) total.

**Combining candidates.** Define `best[i]` as the length of the shortest valid sub-array that fits entirely within the prefix `arr[0..i-1]` (so `best[0]` is "impossible", represented by a sentinel infinity). This array is a running minimum:

```
best[i] = min(best[i-1], len of the window ending at index i-1, if one exists)
```

Now suppose the *second* chosen sub-array is the window `arr[left..right]`. Non-overlap means the first sub-array ends at index `left-1` or earlier — exactly the prefix `arr[0..left-1]`, whose shortest valid sub-array has length `best[left]`. So the best total with this window as the second piece is `(right-left+1) + best[left]`. Because every valid pair has a well-defined "second" sub-array, iterating over all candidates as the second piece covers every pair, and taking the minimum over all of them is the answer.

The key scheduling detail is that `best[left]` only refers to indices strictly less than `left`, and `left <= right`, so by the time we process right endpoint `right` we have already computed `best` up to index `right` — the two computations interleave in a single pass. Concretely, at each step we first copy `best[right+1] = best[right]`, then, if the current window is valid, use `best[left]` for the answer candidate and only afterwards fold the current window's length into `best[right+1]`. Folding *after* reading is what prevents the same window being used as both halves of the pair.

**Worked example** on `arr = [7,3,4,7]`, `target = 7` (`∞` marks "no candidate yet"):

| right | window after shrink | sum | valid? | `best[left]` | candidate | `best[right+1]` |
|-------|--------------------|-----|--------|--------------|-----------|-----------------|
| 0     | `[7]` (left=0)     | 7   | yes    | `best[0]=∞`  | —         | 1 |
| 1     | `[3]` (left=1)     | 3   | no     | —            | —         | 1 |
| 2     | `[3,4]` (left=1)   | 7   | yes    | `best[1]=1`  | 1+2 = 3   | 1 |
| 3     | `[7]` (left=3)     | 7   | yes    | `best[3]=1`  | 1+1 = **2** | 1 |

The minimum candidate is 2, matching the expected output: the sub-arrays `[7]` and `[7]` at the two ends.

If no candidate pair is ever formed, the running answer stays at the sentinel and we return `-1`.

## Complexity Analysis

Time Complexity: O(n) — each of `left` and `right` advances at most `n` times across the whole sweep, so the inner shrink loop is amortized O(1) per step.
Space Complexity: O(n) for the `best` prefix-minimum array. This can be reduced to O(1) extra space by noticing that only `best[left]` and the running minimum are ever read, but keeping the array is clearer and still comfortably within limits.

## Edge Cases

- **No valid sub-array at all, or exactly one** (`arr = [4,3,2,6,2,3,4]`, `target = 6`). The answer must be `-1`, not the length of the single sub-array. Guarding with a sentinel "infinity" and checking `best[left] != inf` before forming a candidate handles both.
- **Two valid sub-arrays that overlap** (e.g. `[1,2,1]` with `target = 3`: `[1,2]` and `[2,1]` share index 1). Reading `best[left]` — strictly the prefix *before* the window — rejects these automatically, and the answer is `-1`.
- **Array shorter than two sub-arrays** (`len(arr) == 1`). The loop runs once, `best[0]` is infinity, and no candidate is ever formed, so `-1` falls out naturally with no special-case code.
- **Adjacent sub-arrays** (`arr = [2,2,2]`, `target = 2`). Touching but not overlapping is legal; the boundary must be `best[left]`, not `best[left-1]`, or you would wrongly exclude a sub-array ending at exactly `left-1`.
- **Order of update vs. read.** If you fold the current window's length into `best` *before* reading `best[left]`, a single sub-array can pair with itself when `left == 0`-ish situations arise. Always read first, then update.
- **`target` larger than the total sum.** The window sum never reaches `target`, the shrink loop can empty the window (`left` passes `right`), and the code must not index out of range — the `sum > target` condition combined with positive elements guarantees `left <= right+1` at all times.
- **Overflow is not a concern here** (`n <= 1e5`, `arr[i] <= 1000` gives a max total of 1e8, and `target <= 1e8` fits in an `int`), but the sentinel must be chosen so that sentinel-plus-length comparisons never wrap; guarding explicitly against the sentinel avoids the question entirely.
