# 3789. Minimum Cost to Acquire Required Items

[LeetCode Link](https://leetcode.com/problems/minimum-cost-to-acquire-required-items/)

Difficulty: Medium
Topics: Math, Greedy
Acceptance Rate: 44.4%

## Hints

### Hint 1

The needs can be as large as 10^9, so you can't simulate buying items one at a time. Pick one quantity that, once fixed, makes every other choice obvious.

### Hint 2

Let `k` be how many type-3 ("both") items you buy. With `k` fixed, the rest is forced: buy `max(0, need1 - k)` type-1 items and `max(0, need2 - k)` type-2 items. Now the total cost is a function of `k` alone. What does that function look like?

### Hint 3

`f(k) = k*costBoth + max(0, need1-k)*cost1 + max(0, need2-k)*cost2` is **piecewise linear** in `k`. Its slope only changes at `k = min(need1, need2)` and `k = max(need1, need2)`, and buying more than `max(need1, need2)` never helps. A piecewise-linear function on an interval has its minimum at an endpoint or a breakpoint, so you only need to check `k ∈ {0, min, max}`.

## Approach

1. Let `lo = min(need1, need2)` and `hi = max(need1, need2)`.
2. Define `f(k)` as above.
3. On each segment, the slope of `f` is:
   - `[0, lo]`: `costBoth - cost1 - cost2`. A both-item replaces one of each single item.
   - `[lo, hi]`: `costBoth - cost(larger-need type)`. Only one requirement is still open.
   - `k > hi`: `costBoth > 0`. You pay for nothing useful, so this never helps.
4. A linear function on a segment is smallest at one of the segment's ends, so the answer is `min(f(0), f(lo), f(hi))`.

The greedy reading is the same thing: use both-items for the overlap if `costBoth < cost1 + cost2`. Then, for the leftover of the larger need, use both-items only if `costBoth < cost` of that single type. Checking all three candidates covers every combination without any case analysis.

**Example:** `cost1=5, cost2=4, costBoth=7, need1=2, need2=5`
- `f(0) = 2*5 + 5*4 = 30`
- `f(2) = 2*7 + 0 + 3*4 = 26`
- `f(5) = 5*7 = 35`
- Answer: **26**. Both-items are cheaper than a pair (7 < 9) but dearer than a single type-2 (7 > 4).

## Complexity Analysis

Time Complexity: O(1)
Space Complexity: O(1)

## Edge Cases

- **Both needs are zero:** every candidate is `k = 0`, so the cost is 0.
- **One need is zero:** the overlap is empty (`lo = 0`), but a both-item may still be cheaper than a single item of the needed type, so `k = hi` still matters.
- **Overflow:** costs up to 10^6 times needs up to 10^9 give about 2·10^15, which overflows 32-bit ints. Use `int64`.
- **Ties (`costBoth == cost1 + cost2`):** any choice gives the same cost. Taking the minimum handles this naturally.
- **Over-buying:** buying more than `max(need1, need2)` both-items only adds cost, so `k` never needs to go past `hi`.
