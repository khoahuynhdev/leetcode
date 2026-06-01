# 2144. Minimum Cost of Buying Candies With Discount

[LeetCode Link](https://leetcode.com/problems/minimum-cost-of-buying-candies-with-discount/)

Difficulty: Easy
Topics: Array, Greedy, Sorting
Acceptance Rate: 66.3%

## Hints

### Hint 1

Think about which candy you would *want* to get for free. The free candy must be less than or equal to the cheaper of the two you pay for, so a free candy can never be the most expensive one in its group of three. What ordering of the array makes it easy to reason about "the cheapest in a group of three"?

### Hint 2

After sorting, consider processing candies from the most expensive to the cheapest. If you pick the two most expensive remaining candies as the "paid" pair, which candy can you legally claim for free? The structure of the problem strongly suggests grouping the sorted list into chunks of three.

### Hint 3

Sort the array in descending order. Iterate through it and, for every group of three consecutive candies, pay for the first two and skip the third. The third candy in each group is guaranteed to be the cheapest of that triple, and by greedy exchange, no other pairing can save more money. Any leftover 1 or 2 candies at the end must all be paid for in full.

## Approach

The discount lets us get one candy free for every two we buy, with the constraint that the free candy's cost cannot exceed the minimum cost of the two purchased candies. To minimize the total bill we want the candies we skip paying for to be as expensive as possible — but the rule caps each freebie at the smaller of its purchase pair, so the best we can ever do is make the freebie equal to (or smaller than) the cheaper purchased candy.

The clean greedy is:

1. Sort `cost` in descending order.
2. Walk through the sorted array. For indices `0, 1, 2`, pay for `0` and `1`, get `2` free. For indices `3, 4, 5`, pay for `3` and `4`, get `5` free. And so on — every index whose position is `2 mod 3` is free.
3. Sum the prices of all indices where `i % 3 != 2`.

Why this is optimal: in the sorted-descending order, the third candy of any group of three is the smallest of those three, so making it the freebie is always legal. And we've maximized the value of the freebie because each freebie is the third-largest unprocessed candy at the time we encounter it — no rearrangement could promote a larger candy into a free slot without violating the "free ≤ min(paid pair)" rule.

Example with `cost = [6,5,7,9,2,2]`:

- Sort descending: `[9, 7, 6, 5, 2, 2]`.
- Group 1: pay 9 + 7, get 6 free.
- Group 2: pay 5 + 2, get 2 free.
- Total = 9 + 7 + 5 + 2 = 23.

## Complexity Analysis

Time Complexity: O(n log n) — dominated by the sort.
Space Complexity: O(1) extra space if sorting in place (Go's `sort.Sort` on a slice is in place, ignoring sort's internal recursion stack of O(log n)).

## Edge Cases

- **`len(cost) == 1`**: only one candy, no pair possible, so pay its full cost.
- **`len(cost) == 2`**: two candies form a pair but there's no third candy to take free; pay for both.
- **All equal costs** (e.g., `[5,5,5]`): the freebie equals the paid candies, which is allowed since the constraint is `≤`, not `<`.
- **Length not divisible by 3** (e.g., 4 or 5 candies): the trailing 1 or 2 candies after the last full triple must all be paid for; the greedy iteration naturally handles this because `i % 3 != 2` for those indices.
- **Minimum constraint values** (length 1, cost 1): the answer is just the single cost.
