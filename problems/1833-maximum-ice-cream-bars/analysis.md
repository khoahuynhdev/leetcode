# 1833. Maximum Ice Cream Bars

[LeetCode Link](https://leetcode.com/problems/maximum-ice-cream-bars/)

Difficulty: Medium
Topics: Array, Greedy, Sorting, Counting Sort
Acceptance Rate: 75.4%

## Hints

### Hint 1

You want to maximize the *count* of items you buy, not which specific ones. When the goal is "as many as possible" and items have no other constraints tying them together, ask yourself: does the order in which I pick them matter, and is there an obviously "best first" choice?

### Hint 2

This is a greedy problem. To fit the most bars into a fixed budget, always reach for the cheapest bar still available. If you sort the costs ascending and spend from the front, you can never do better by skipping a cheap bar in favor of a more expensive one. The only question left is *how* to sort efficiently.

### Hint 3

Notice the constraint: `1 <= costs[i] <= 10^5`. The values live in a small, bounded range, which is the classic signal for **counting sort**. Instead of comparison-sorting in O(n log n), tally how many bars exist at each price into a frequency array, then walk the prices from cheapest to most expensive, buying as many as your coins allow at each price level. This gives a linear-time greedy.

## Approach

The core idea is greedy: to buy the maximum number of bars, always buy the cheapest ones first. Spending a coin on an expensive bar when a cheaper one is available can only reduce the total count, so sorting ascending and buying from the cheapest end is optimal.

Because the problem explicitly asks for **counting sort**, we avoid a comparison sort and instead exploit the bounded value range (`costs[i] <= 10^5`):

1. **Build a frequency table.** Create a `count` array indexed by price. For each cost `c`, increment `count[c]`. This buckets every bar by its price in O(n).

2. **Walk prices from cheapest to most expensive.** For each price `p` from 1 up to the maximum cost:
   - If `count[p]` is 0, skip it.
   - Otherwise, figure out how many bars at this price we can afford: `affordable = coins / p`, capped at the number actually available `count[p]`.
   - Add that many bars to the answer and subtract their total cost (`bought * p`) from `coins`.
   - If after this we cannot afford even one more bar at the *current* price, then we certainly cannot afford anything more expensive either — so we can stop early.

3. **Return the running count** of bars bought.

**Walkthrough on Example 1:** `costs = [1,3,2,4,1]`, `coins = 7`.
- Frequency: price 1 → 2 bars, price 2 → 1, price 3 → 1, price 4 → 1.
- Price 1: can afford `7/1 = 7`, but only 2 exist → buy 2, coins = 5, total = 2.
- Price 2: can afford `5/2 = 2`, only 1 exists → buy 1, coins = 3, total = 3.
- Price 3: can afford `3/3 = 1`, 1 exists → buy 1, coins = 0, total = 4.
- Price 4: can afford `0/4 = 0` → stop. Answer: **4**.

This is a satisfying problem — the greedy insight is intuitive once you see it, and the counting-sort twist is a clean exercise in recognizing when bounded inputs let you beat O(n log n).

## Complexity Analysis

Time Complexity: O(n + M), where `n` is the number of bars and `M` is the maximum cost value (here at most 10^5). Building the frequency table is O(n), and scanning the price buckets is O(M).

Space Complexity: O(M) for the frequency array sized to the maximum cost.

## Edge Cases

- **Cannot afford any bar** (e.g., Example 2, all costs exceed `coins`): the loop never buys anything and returns 0. Make sure the affordability check correctly yields 0 here.
- **Can afford every bar** (Example 3): the budget is large enough that `count[p]` is always the binding limit, so every bar is bought.
- **Single bar** (`n == 1`): either you can afford it (return 1) or you cannot (return 0).
- **Duplicate prices**: many bars share the same price; the frequency table naturally handles this by bucketing, and you must cap purchases at the available count, not just affordability.
- **Empty input** (`len(costs) == 0`): defensively returns 0, since there is nothing to buy. (LeetCode guarantees `n >= 1`, but the code handles it cleanly.)
