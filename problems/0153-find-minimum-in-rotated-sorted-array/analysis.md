# 0153. Find Minimum in Rotated Sorted Array

[LeetCode Link](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)

Difficulty: Medium
Topics: Array, Binary Search
Acceptance Rate: 54.2%

## Hints

### Hint 1

The problem explicitly requires `O(log n)` time, which is a strong signal: a linear scan would be trivial but is not allowed. Whenever you see "sorted array" plus a logarithmic time requirement, your first instinct should be **binary search**. The trick here is that the array is not fully sorted — it's been rotated — so you'll need to adapt the standard binary-search template rather than apply it directly.

### Hint 2

Picture the rotated array visually. If you draw the values on a graph, it looks like two ascending segments with a sharp drop between them. The minimum value sits exactly at the start of the second segment (the "pivot" point). At each step of binary search, ask yourself: *given `nums[mid]`, can I tell whether the pivot lies to my left or to my right?* You only need one comparison to decide, and that decision narrows the search space by half.

### Hint 3

Compare `nums[mid]` to `nums[right]` (the rightmost element of the current window). If `nums[mid] > nums[right]`, then the pivot (and therefore the minimum) must be **strictly to the right** of `mid`, because a rotated section is "wrapping around." Otherwise (`nums[mid] <= nums[right]`), the pivot is at `mid` or to its left, so you can include `mid` in the next search window. Using `right` as the anchor (rather than `left`) is what makes this work cleanly — comparing to `nums[left]` runs into ambiguity when the array isn't rotated at all.

## Approach

We use a modified binary search that converges on the index of the minimum element.

Maintain two pointers `left = 0` and `right = n-1`. While `left < right`:

1. Compute `mid = left + (right-left)/2`.
2. If `nums[mid] > nums[right]`, the smallest value must be in the right half, *excluding* `mid` itself (since `nums[mid]` is bigger than at least one later element). Set `left = mid + 1`.
3. Otherwise, `nums[mid] <= nums[right]`, meaning the right half from `mid` onward is sorted. The minimum is at `mid` or somewhere to its left. Set `right = mid` (keep `mid` in the window).

When the loop exits, `left == right`, and that index holds the minimum.

**Why this works:** the invariant is that the minimum element is always inside `[left, right]`. Each branch preserves the invariant while halving the window. Comparing against `nums[right]` (not `nums[left]`) is critical because in a non-rotated array (e.g., `[1,2,3,4,5]`), every `nums[mid]` would be `>= nums[left]`, which gives no information. But `nums[mid] <= nums[right]` cleanly tells us the entire right half is sorted, regardless of rotation.

**Example walkthrough** with `nums = [4,5,6,7,0,1,2]`:

- `left=0, right=6`, `mid=3`, `nums[3]=7 > nums[6]=2` → `left=4`
- `left=4, right=6`, `mid=5`, `nums[5]=1 <= nums[6]=2` → `right=5`
- `left=4, right=5`, `mid=4`, `nums[4]=0 <= nums[5]=1` → `right=4`
- Loop exits with `left=right=4`. Answer: `nums[4] = 0`. ✓

## Complexity Analysis

Time Complexity: O(log n) — the search window halves each iteration.
Space Complexity: O(1) — only two index pointers.

## Edge Cases

- **Array of length 1**: The loop never executes (`left == right` immediately), and we return the only element. Important to verify the loop condition uses `<` rather than `<=`.
- **Not rotated at all** (e.g., `[11,13,15,17]`, which is a "full" rotation of `n`): `nums[mid] <= nums[right]` is always true, so we keep shrinking `right` toward `0` and correctly return `nums[0]`.
- **Rotated by exactly one position** (e.g., `[2,1]`): `mid=0`, `nums[0]=2 > nums[1]=1`, so `left=1`. Return `nums[1] = 1`. ✓
- **Minimum at the very end is impossible** because the problem guarantees a rotation between `1` and `n` times — the array is never "rotated by 0." Still, the algorithm handles the `n`-rotation case (which looks identical to no rotation) correctly.
- **Negative numbers**: Constraints allow `-5000 <= nums[i] <= 5000`. The comparisons are purely relational, so signs don't matter.
