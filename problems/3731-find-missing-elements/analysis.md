# 3731. Find Missing Elements

[LeetCode Link](https://leetcode.com/problems/find-missing-elements/)

Difficulty: Easy
Topics: Array, Hash Table, Sorting
Acceptance Rate: 85.1%

## Hints

### Hint 1

The problem never tells you what the original range was — but it does promise that the smallest and largest values of that range survived in `nums`. So before you can hunt for anything missing, you first need to recover the boundaries of the range. What single pass over the array gives you those?

### Hint 2

Once you know the boundaries, the question becomes a membership test: for every integer between the low and high boundary, is it present in `nums` or not? Membership tests in O(1) are exactly what a hash set (or, since the values are small and bounded, a plain boolean array) is for.

### Hint 3

The output must be sorted, but you never need to call a sort function. If you iterate the range from the smallest value up to the largest and append each value that fails the membership test, the results come out in increasing order for free. That turns the whole problem into two linear scans and no sorting at all.

## Approach

The solution has three short phases, and each one is a single loop.

**Phase 1 — recover the range.** Scan `nums` once, tracking the minimum and maximum. The problem guarantees both endpoints of the original range are still present, so `[min, max]` *is* the original range. In Example 1, `nums = [1,4,2,5]` gives `min = 1`, `max = 5`, so the original array was `[1,2,3,4,5]`.

**Phase 2 — record what is present.** Allocate a boolean slice `seen` of length `max - min + 1`, where index `i` stands for the value `min + i`. Scan `nums` again and set `seen[v-min] = true` for each value `v`. Using the offset `min` keeps the slice exactly as large as the range instead of as large as the maximum value, which matters when the range starts far from zero (e.g. `[7,8,6,9]` only needs 4 slots, not 10). A `map[int]bool` would work identically; the boolean slice is just faster and avoids hashing.

**Phase 3 — collect the gaps in order.** Walk `i` from `0` to `max - min` and append `min + i` to the result whenever `seen[i]` is false. Because we walk the range in increasing order, the output is sorted by construction — no `sort.Ints` needed.

Tracing Example 3, `nums = [5,1]`: phase 1 gives `min = 1`, `max = 5`. Phase 2 marks indices 0 and 4 of a 5-element `seen` slice. Phase 3 walks indices 0..4 and emits `1+1 = 2`, `1+2 = 3`, `1+3 = 4` for the unmarked indices 1, 2, 3, producing `[2,3,4]`.

Why does this work? The problem states the input contains *unique* integers, so there are no duplicates to confuse the membership marking, and every value in `nums` necessarily lies inside `[min, max]` — meaning no index goes out of bounds. Any integer in that closed interval that is not flagged in `seen` must be one that "went missing."

An alternative worth knowing: sort `nums`, then walk adjacent pairs and emit everything strictly between `nums[i]` and `nums[i+1]`. It reaches the same answer and uses O(1) extra space beyond the output, but it costs O(n log n) and mutates the caller's array unless you copy first. The counting approach is both simpler and asymptotically faster here.

## Complexity Analysis

Time Complexity: O(n + R), where `n = len(nums)` and `R = max - min + 1` is the size of the original range. Given the constraints (`n <= 100`, values in `[1, 100]`), both terms are bounded by 100, so this is effectively linear and trivially fast.

Space Complexity: O(R) for the `seen` slice, plus O(k) for the output where `k` is the number of missing integers. Excluding the required output, the auxiliary space is O(R).

## Edge Cases

- **Nothing is missing** (Example 2, `[7,8,6,9]`): every slot in `seen` is marked, so the loop appends nothing. The function must return an **empty** slice rather than `nil` — `reflect.DeepEqual(nil, []int{})` is `false` in Go, so initializing the result with `make([]int, 0)` avoids a surprising test failure even though both marshal to `[]` in JSON.
- **Unsorted input** (Examples 1 and 3): the input is explicitly not guaranteed sorted, so you cannot assume `nums[0]` is the minimum or `nums[len-1]` is the maximum. Deriving `min`/`max` by scanning handles this.
- **Only the two endpoints present** (Example 3, `[5,1]`): the array has the legal minimum length of 2 and everything strictly between the endpoints is missing. This is the case that breaks any solution which assumes at most one gap.
- **Two adjacent values** (`[3,4]`): `max - min + 1 == 2` and both slots are marked, so the answer is empty. A good check that the loop bounds are inclusive on both ends and produce no off-by-one output.
- **Range not starting at 1** (`[50,53]`): the `min` offset in the `seen` index is what keeps this correct and memory-efficient; indexing by raw value would still work under these constraints but wastes space and breaks the moment values can be negative.
- **Widest possible range** (`[1,100]`): 98 missing values, the largest output this problem can produce. Nothing special is needed, but it is a good sanity check that the result slice grows correctly.
