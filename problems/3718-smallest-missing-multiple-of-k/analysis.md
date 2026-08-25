# 3718. Smallest Missing Multiple of K

[LeetCode Link](https://leetcode.com/problems/smallest-missing-multiple-of-k/)

Difficulty: Easy
Topics: Array, Hash Table
Acceptance Rate: 64.5%

## Hints

### Hint 1

The question "is this value present in `nums`?" is going to be asked over and over. Scanning the array from scratch each time works, but there is a data structure whose entire purpose is answering membership questions in constant time. Reach for it first, and the rest of the problem becomes a simple search.

### Hint 2

You are not searching over the array — you are searching over the *candidates*. The candidates are `k, 2k, 3k, ...` in increasing order. Walk them in that order and return the first one that fails the membership test. Because you generate candidates smallest-first, the first miss is automatically the answer; no comparison or `min` tracking is needed.

### Hint 3

The loop is guaranteed to terminate quickly, and knowing why is the insight that makes this safe to write as an unbounded `for`. Every element satisfies `nums[i] <= 100`, so no value above `100` can ever be "present." That means the candidate `k * (100/k + 1)` exceeds `100` and *must* be missing. So at most `100/k + 1` iterations run — one for `k = 100`, about 100 for `k = 1`. The answer is never larger than 200 under these constraints.

## Approach

The problem asks for the smallest element of the infinite set `{k, 2k, 3k, ...}` that does not appear in `nums`. Two ingredients solve it:

1. **A hash set for membership.** Copy every element of `nums` into a `map[int]struct{}` (Go's idiomatic set — `struct{}` occupies zero bytes). Duplicates in `nums` collapse naturally, which is exactly what we want: we only care *whether* a value appears, not how many times. This costs one linear pass and turns each later lookup into O(1).

2. **An ordered scan of the candidates.** Start at `multiple = k` and keep adding `k`. The moment a candidate is absent from the set, return it.

Because the candidates are generated in strictly increasing order, the first failure *is* the minimum failure — there is nothing smaller left to check. This is the same "generate in sorted order, stop at first success" idea behind problems like *First Missing Positive*, just without the in-place trickery.

Walking through Example 1, `nums = [8,2,3,4,6]`, `k = 2`:

- set = `{8, 2, 3, 4, 6}`
- `multiple = 2` → present, continue
- `multiple = 4` → present, continue
- `multiple = 6` → present, continue
- `multiple = 8` → present, continue
- `multiple = 10` → **absent** → return `10`

Example 2, `nums = [1,4,7,10,15]`, `k = 5`: the very first candidate `5` is absent, so we return immediately. Note that `10` and `15` being in `nums` is irrelevant — we stop at the first gap, we do not need the largest present multiple.

**Why the unbounded loop is safe.** The constraint `nums[i] <= 100` caps what can possibly block us. Once `multiple > 100`, it cannot be in `nums`, so the loop exits. The worst case is `k = 1` with `nums` containing `1..100`, which returns `101` after 101 iterations. If you dislike unbounded loops in an interview, you can equivalently loop `for i := 1; ; i++` with an explicit early exit when `k*i > 100` — but stating the termination argument out loud is usually what the interviewer is listening for.

An alternative without a hash set: sort `nums`, then binary search each candidate. That is `O(n log n)` up front and `O(log n)` per probe — strictly worse, and more code. Another alternative given the tiny bounds is a fixed `[101]bool` array instead of a map, which is faster in practice and worth mentioning as a micro-optimization; the map version is chosen here because it stays correct if the value bound is ever relaxed.

## Complexity Analysis

Time Complexity: O(n + m), where `n = len(nums)` and `m` is the number of candidates tried. Building the set is `O(n)`; the scan runs at most `100/k + 1` times, so `m <= 101`. With the given constraints this is effectively `O(n)`.

Space Complexity: O(n) for the hash set holding the distinct values of `nums`. (Using a `[101]bool` instead makes this `O(1)` under the stated value bound.)

## Edge Cases

- **The answer is `k` itself.** When no multiple of `k` appears in `nums` at all — Example 2 is this case — the loop must return on its very first iteration. A solution that mistakenly starts at `2k`, or that first scans for the largest present multiple, gets this wrong.
- **`k = 1`.** Every positive integer is a multiple, so the problem degenerates into "smallest missing positive integer." This is the worst case for iteration count (up to 101 steps) and a good sanity check that your loop terminates.
- **`k` larger than every element**, e.g. `nums = [1,2,3], k = 7`. No multiple can possibly be present, so the answer is `k`. Confirms you never index into `nums` by value or assume a multiple exists.
- **Duplicates in `nums`**, e.g. `[2,2,2]` with `k = 2` → `4`. A set handles this for free, but a counting or two-pointer approach might double-count.
- **A contiguous run of multiples**, e.g. `[2,4,6,8,10]` with `k = 2` → `12`. Verifies you keep advancing past a long block of present values rather than stopping early.
- **Non-multiples interleaved.** Values in `nums` that are not multiples of `k` (the `3` in Example 1) must be ignored, not treated as blockers. Filtering them out is optional but they must never affect the result.
- **The answer exceeding the value range**, e.g. `nums = [100], k = 100` → `200`. Guards against any implementation that caps its search at `100` and returns a wrong value (or nothing) when the answer lies above the input bound.
