# 1665. Minimum Initial Energy to Finish Tasks

[LeetCode Link](https://leetcode.com/problems/minimum-initial-energy-to-finish-tasks/)

Difficulty: Hard
Topics: Array, Greedy, Sorting
Acceptance Rate: 66.6%

## Hints

### Hint 1

The tasks can be done in any order, so the real question is: *what order* makes the starting requirement smallest? Whenever the order of operations is a free choice and you want to minimize/maximize some bottleneck, think **greedy with a custom sort** — and try to discover the comparison rule with an exchange argument on two adjacent tasks.

### Hint 2

Fix any order `t_1, t_2, ..., t_n`. The total `actual` energy you spend is the same no matter the order. So the only thing the order affects is whether you ever fail the "you need `minimum_i` energy *before* starting task `i`" check. Write down, for each task, the inequality the starting energy `E` must satisfy, in terms of the `actual` costs of the tasks that come before it. The answer is the maximum of those lower bounds — now you want the order that makes that maximum as small as possible.

### Hint 3

Compare two adjacent tasks `a` then `b` versus `b` then `a`. Doing `a` first is at least as good when

```
max(min_a, actual_a + min_b) <= max(min_b, actual_b + min_a)
```

Working through the cases (and using `actual <= minimum`), this reduces to: do `a` before `b` when `min_a - actual_a >= min_b - actual_b`. So **sort by `minimum - actual` in descending order** — handle the tasks that leave the least "slack" relative to their cost while you still have lots of energy in reserve.

## Approach

For a fixed processing order `t_1, ..., t_n`, let `S_k = actual_1 + actual_2 + ... + actual_{k-1}` be the energy already spent before task `k`. To legally start task `k` we need

```
E - S_k >= minimum_k     i.e.     E >= minimum_k + S_k
```

So the minimum feasible starting energy for that order is `max_k (minimum_k + S_k)`. The sum of all `actual` values is constant, so reordering only shuffles which `minimum_k` gets paired with which prefix sum `S_k`.

To minimize that maximum, sort the tasks by `minimum_i - actual_i` in **descending** order (proved via the adjacent-swap exchange argument in Hint 3). Then sweep once:

- keep a running total `spent` of `actual` costs seen so far (the prefix sum `S_k`),
- for each task, the bound it imposes is `spent + minimum`; track the largest such bound,
- then add this task's `actual` to `spent`.

The largest bound encountered is the answer.

**Walkthrough on Example 1:** `tasks = [[1,2],[2,4],[4,8]]`. The `minimum - actual` values are `1, 2, 4`, so the descending order is `[4,8], [2,4], [1,2]`.

| step | task   | spent before | bound = spent + minimum |
|------|--------|--------------|--------------------------|
| 1    | [4,8]  | 0            | 0 + 8 = **8**            |
| 2    | [2,4]  | 4            | 4 + 4 = 8                |
| 3    | [1,2]  | 6            | 6 + 2 = 8                |

Maximum bound = `8`, which matches the expected output. Starting with `8`: `8 → (do [4,8]) → 4 → (do [2,4]) → 2 → (do [1,2]) → 1`. Starting with `7` fails immediately because you cannot begin `[4,8]`.

This is a genuinely tricky "discover the sort key" greedy — it is rated Hard for good reason — but once the exchange argument clicks, the implementation is short and clean.

## Complexity Analysis

Time Complexity: O(n log n) — dominated by sorting the `n` tasks; the sweep is O(n).
Space Complexity: O(1) extra if the input may be sorted in place (Go's `sort.Slice` sorts in place; it uses O(log n) stack for recursion). O(n) if you copy the input first to avoid mutating it.

## Edge Cases

- **Single task `[[a, m]]`** — the answer is just `m` (you need `m` to start, and you finish with `m - a >= 0` left). Make sure the sweep handles a one-element slice.
- **`actual_i == minimum_i` for some/all tasks** — allowed by the constraints (`actual <= minimum`); the slack `minimum - actual` is then `0`. The greedy still works; ties in the sort key can be broken arbitrarily.
- **All tasks identical** — e.g. `[[2,2],[2,2],[2,2]]`. The sort is a no-op and the answer is `actual*(n-1) + minimum` (here `6`); a useful sanity check that the prefix-sum accumulation is correct.
- **Order genuinely matters** — e.g. `[[1,2],[1,5]]` must be done as `[1,5]` then `[1,2]` to need only `5`; the other order needs `6`. This is the case that fails if you forget to sort or sort by the wrong key.
- **Large input / large values** — up to `10^5` tasks with values up to `10^4`, so the answer can be on the order of `10^9`. Fine for Go's 64-bit `int`, but worth noting if porting to a 32-bit `int` language.
