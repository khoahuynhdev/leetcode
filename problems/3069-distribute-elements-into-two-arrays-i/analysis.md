# 3069. Distribute Elements Into Two Arrays I

[LeetCode Link](https://leetcode.com/problems/distribute-elements-into-two-arrays-i/)

Difficulty: Easy
Topics: Array, Simulation
Acceptance Rate: 75.0%

## Hints

### Hint 1

Read the problem statement one more time and ask yourself: is there any decision here that depends on the *future*? Every operation is fully determined by the state you already have. When the rules are that mechanical, you usually do not need a clever algorithm — you need to faithfully reproduce the process the statement describes.

### Hint 2

The rule for operation `i` only ever looks at *one* value from each of the two arrays: the element most recently appended. So the "state" you carry through the loop is much smaller than the arrays themselves. Think about what container gives you cheap access to the most recently added element in O(1).

### Hint 3

A Go slice already gives you O(1) access to its last element via `s[len(s)-1]`, and O(1) amortized appends via `append`. That means you can build `arr1` and `arr2` directly, compare their tails on each step, and never need auxiliary bookkeeping. Seed `arr1` with `nums[0]` and `arr2` with `nums[1]` (the problem is 1-indexed; your code is 0-indexed), loop from index 2, and finally concatenate. The whole thing is one pass.

## Approach

This problem is pure simulation. The trap is not the algorithm — it is the off-by-one risk from the problem being stated in 1-indexed terms while Go is 0-indexed.

Step by step:

1. **Seed the two arrays.** Operation 1 appends `nums[1]` (1-indexed) to `arr1`, and operation 2 appends `nums[2]` to `arr2`. In 0-indexed Go that is `arr1 = [nums[0]]` and `arr2 = [nums[1]]`. The constraint `n >= 3` guarantees both of these indices exist, so no guard is needed — but it is worth knowing *why* it is safe rather than assuming it.

2. **Simulate the remaining operations.** For each `i` from `2` to `n-1`, compare the tails: if `arr1[len(arr1)-1] > arr2[len(arr2)-1]`, append `nums[i]` to `arr1`; otherwise append to `arr2`. Note the tie-breaking direction — the statement says "otherwise", so anything that is *not* strictly greater sends the element to `arr2`. Since `nums` contains distinct integers, an exact tie can never occur, but writing the comparison as strict `>` with an `else` branch matches the statement exactly and stays correct even if the distinctness guarantee were dropped.

3. **Concatenate.** `result = arr1 ++ arr2`. In Go, `append(arr1, arr2...)` does this. One subtlety: that call may reuse and mutate `arr1`'s backing array. That is fine here because `arr1` is a local slice we own and never read again, but in general it is a habit worth being deliberate about.

**Walking through Example 2, `nums = [5,4,3,8]`:**

| Step | arr1 | arr2 | Comparison | Action |
| --- | --- | --- | --- | --- |
| seed | `[5]` | `[4]` | — | — |
| `i=2` (`nums[2]=3`) | `[5,3]` | `[4]` | `5 > 4` ✓ | append to `arr1` |
| `i=3` (`nums[3]=8`) | `[5,3]` | `[4,8]` | `3 > 4` ✗ | append to `arr2` |

Result: `[5,3] ++ [4,8] = [5,3,4,8]`. ✅

Why the simulation is optimal: every element must be placed, and its placement depends on all prior placements, so there is no way to shortcut the sequence. You must touch each element at least once, which makes O(n) a lower bound — and the simulation already achieves it.

A worthwhile micro-optimization: since the two arrays partition `nums`, you can preallocate with `make([]int, 0, len(nums))` for each, avoiding repeated reallocation as the slices grow. With `n <= 50` this is invisible in practice, but it is the kind of detail an interviewer notices favorably.

## Complexity Analysis

Time Complexity: O(n) — a single pass over `nums`, with O(1) work (two tail lookups and one append) per element, plus an O(n) final concatenation.

Space Complexity: O(n) — `arr1` and `arr2` together hold exactly `n` elements. This is required for the output itself, so the auxiliary space beyond the answer is O(1).

## Edge Cases

- **1-indexed statement vs. 0-indexed code.** The single most likely bug. The problem says operations 1 and 2 seed the arrays with `nums[1]` and `nums[2]`; in Go those are `nums[0]` and `nums[1]`, and the loop starts at index `2`, not `3`.
- **Minimum length `n == 3`.** Only one real decision is made after seeding. Test it explicitly to confirm the loop bounds are right — an off-by-one here either skips the last element or panics.
- **All elements go to one array.** If `nums[0] > nums[1]` and every subsequent element keeps `arr1`'s tail on top (e.g. `[100,1,2,3,4]`), `arr2` stays at length 1 for the entire run. Verify your concatenation still produces the right length and that nothing assumes both arrays keep growing.
- **The mirrored case, where `arr2` absorbs everything** (e.g. `[1,100,2,3,4]`). Confirms the `else` branch is exercised just as thoroughly as the `if`.
- **Tail comparison, not max comparison.** It is tempting to misread the rule as comparing the largest elements of each array. It is strictly the *last appended* element. `[7,4,9,2,8]` distinguishes these: after `arr1 = [7,9,2]`, `arr1`'s max is 9 but its tail is 2, so `nums[4]=8` goes to `arr2`.
- **Distinct values mean no ties.** You never hit equality, so no tie-breaking rule is needed. Do not invent one — but do keep the comparison strict so the code still matches the specification's wording.
