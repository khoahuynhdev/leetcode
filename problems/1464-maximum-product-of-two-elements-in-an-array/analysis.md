# 1464. Maximum Product of Two Elements in an Array

[LeetCode Link](https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/)

Difficulty: Easy
Topics: Array, Sorting, Heap (Priority Queue)
Acceptance Rate: 84.0%

## Hints

### Hint 1

The brute force is obvious: try every pair `(i, j)` and keep the best product. That is `O(n^2)`, which actually passes for `n <= 500` — but the problem is really asking whether you can characterize *which* pair wins without testing them all. Think about what property the winning pair must have, rather than about a fancy data structure.

### Hint 2

Look at the constraint `1 <= nums[i] <= 10^3`. Every term `nums[i] - 1` is therefore non-negative — there are no negative factors to worry about, so a large negative times another large negative can never sneak in and beat the obvious candidates. With all factors `>= 0`, making each factor as big as possible independently maximizes the product. That reduces the problem to a selection question: which two elements do you want?

### Hint 3

You want the two largest values in the array, and "two largest" means *by position*, not by distinct value — if `nums = [1,5,4,5]` the answer uses both fives, because `i` and `j` only need to be different indices, not different values. So track a running largest and second-largest in a single pass: when a new element beats the current largest, the old largest slides down into second place. Answer is `(largest-1) * (second-1)`.

## Approach

The key observation is a monotonicity argument. Define `f(x) = x - 1`. Under the constraint `1 <= nums[i] <= 10^3`, we have `f(x) >= 0` for every element. The product of two non-negative numbers is maximized by maximizing each factor independently — there is no sign trickery, unlike problems where two large negatives can multiply into a large positive. Therefore the optimal pair is simply the two largest elements of the array (counted with multiplicity, since duplicate values sitting at different indices are a legal choice).

That gives three natural implementations:

1. **Sort** — sort ascending, take the last two elements. `O(n log n)` time, simple to write, but does more work than needed.
2. **Heap** — push everything into a max-heap and pop twice. Also `O(n log n)` to build (or `O(n)` with heapify), and overkill here.
3. **Single linear scan** — keep two variables, `largest` and `second`, and update them as you walk the array. This is the optimal `O(n)` / `O(1)` version and is what the solution implements.

The scan works like this. Initialize both `largest` and `second` to `0`. Since `nums[i] >= 1`, every real element has `nums[i] - 1 >= 0`, so `0` is a safe sentinel that no valid element can be beaten by. For each `v`:

- If `v > largest`, then `v` becomes the new maximum and the *previous* maximum demotes to second place: `second = largest; largest = v`.
- Else if `v > second`, then `v` is not the biggest but does beat the current runner-up: `second = v`.
- Otherwise `v` changes nothing.

The demotion step in the first branch is what makes the scan correct — forgetting it is the classic bug, and it silently breaks inputs like `[1,2,3]` where the maximum keeps being replaced. Note also that both comparisons use strict `>`, which is exactly what lets duplicates work: for `nums = [1,5,4,5]`, the first `5` sets `largest = 5`; the second `5` fails `5 > 5` in the first branch but passes `5 > 4` in the second, so `second` becomes `5` and the answer is `(5-1)*(5-1) = 16`.

Walking Example 1, `nums = [3,4,5,2]`:

| v | largest | second |
|---|---------|--------|
| — | 0       | 0      |
| 3 | 3       | 0      |
| 4 | 4       | 3      |
| 5 | 5       | 4      |
| 2 | 5       | 4      |

Result: `(5-1) * (4-1) = 4 * 3 = 12`. ✓

Because `nums.length >= 2` is guaranteed, both `largest` and `second` are always populated by real elements when the loop finishes, so no post-loop validation is required.

## Complexity Analysis

Time Complexity: O(n) — one pass over the array, constant work per element. (Sorting or heap-based variants are O(n log n).)
Space Complexity: O(1) — only the two tracking variables, regardless of input size.

## Edge Cases

- **Exactly two elements** (`[3,7]`) — the minimum allowed length. Both elements must be used; the scan naturally sets `largest = 7`, `second = 3` and returns `12`. Any solution that special-cases "skip index `i`" logic must not exclude one of only two elements.
- **Duplicate maximum** (`[1,5,4,5]`) — `i != j` constrains *indices*, not *values*, so the same value appearing twice is a legal pair. This is why the comparisons must be strict `>` with an else-if fallthrough rather than deduplicating values.
- **All elements identical** (`[2,2,2,2]`) — same reasoning as above; the answer is `(2-1)*(2-1) = 1`, not `0` and not a "no valid pair" error.
- **Minimum values** (`[1,1]`) — every factor becomes `0`, so the answer is `0`. Worth checking because a sentinel initialized to something other than `0`, or a `-1` sentinel combined with sloppy comparisons, can produce a negative result here.
- **Maximum values** (`[1000, 1000]`) — the largest possible answer is `999 * 999 = 998001`, comfortably inside `int32`, so there is no overflow risk in Go's `int`. Good to confirm rather than assume when a problem asks for a product.
- **Maximum already first, then descending** (`[9, 1, 1]`) — exercises the demotion branch and the "changes nothing" branch. If you forget `second = largest` when a new maximum arrives, this class of input is where it shows up.
