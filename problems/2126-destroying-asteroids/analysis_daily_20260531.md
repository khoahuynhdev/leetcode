# 2126. Destroying Asteroids

[LeetCode Link](https://leetcode.com/problems/destroying-asteroids/)

Difficulty: Medium
Topics: Array, Greedy, Sorting
Acceptance Rate: 55.9%

## Hints

### Hint 1

You are allowed to pick any order. When you have full freedom over ordering and you want to maximize your chances of "surviving" each step, think about what kind of strategy almost always pays off in problems like this: **greedy**. Ask yourself, "If I have to face every asteroid eventually, which one should I face first?"

### Hint 2

A larger asteroid only becomes destroyable after the planet has absorbed enough mass from smaller ones. That suggests a particular ordering of the asteroids before you start simulating the collisions. What does that ordering look like, and what operation on the array gives it to you in one step?

### Hint 3

Sort `asteroids` in non-decreasing order, then walk through them once. If the current planet mass is ever smaller than the asteroid in front of you, you can return `false` immediately — no later ordering can rescue you because every earlier asteroid is already smaller. Otherwise, absorb the asteroid and keep going. Watch for overflow: the cumulative mass can grow well past `10^5`.

## Approach

The key observation is that if you ever survive a collision with asteroid `a`, you only gain mass — you never get weaker. So the optimal strategy is to face the smallest asteroid you can destroy next, growing your mass as much as possible before encountering larger asteroids.

This translates to a clean greedy algorithm:

1. Sort `asteroids` in non-decreasing order.
2. Iterate through the sorted array. For each asteroid `a`:
   - If `mass < a`, return `false`. There is no way to handle `a` later because every remaining asteroid is `>= a` and contributes nothing to overcoming this barrier.
   - Otherwise, set `mass += a` and continue.
3. If we finish the loop, every asteroid was destroyed; return `true`.

**Why it works:** Suppose some non-sorted order succeeds. Consider the first asteroid `a` in the sorted order where the sorted strategy fails. By that point in the sorted strategy, the planet has accumulated the sum of all asteroids strictly smaller than `a`. Any order that destroys `a` must have already absorbed at least enough mass to match `a`, but it cannot have absorbed more than the sum of all asteroids smaller than `a` before reaching `a` (since asteroids `>= a` cannot be tackled before `a`). So the sorted approach has the maximum possible mass available at the moment of facing `a` — if even that is not enough, no order works.

**Example (`mass = 10`, `asteroids = [3, 9, 19, 5, 21]`):**
After sorting: `[3, 5, 9, 19, 21]`. Trace: `10 → 13 → 18 → 27 → 46 → 67`. All survived → `true`.

**Overflow note:** With `asteroids.length` up to `10^5` and each value up to `10^5`, the running mass can reach up to `10^10 + 10^5`, which overflows `int32` but fits comfortably in Go's 64-bit `int`.

## Complexity Analysis

Time Complexity: O(n log n) for sorting, plus O(n) for the linear sweep.
Space Complexity: O(1) extra space (Go's `sort.Ints` sorts in place; ignoring the O(log n) recursion stack from the sort implementation).

## Edge Cases

- **Single asteroid that is too heavy** (e.g., `mass = 1`, `asteroids = [2]`): The very first comparison must fail and return `false`.
- **Single asteroid that exactly equals mass**: The constraint is `>=`, so equal mass destroys the asteroid. Make sure the comparison is `mass < a`, not `mass <= a`.
- **All asteroids smaller than initial mass**: Should always return `true` regardless of order.
- **Large cumulative mass / overflow**: Up to `10^5` asteroids of value `10^5` each pushes the running mass past `int32` range — use `int` (64-bit on standard Go platforms) to be safe.
- **Strictly increasing asteroid sequence requiring exact greedy ordering** (e.g., `mass = 5`, `asteroids = [5, 10, 20]`): Sorted order succeeds; tackling `10` before `5` would fail. This confirms sorting is essential.
