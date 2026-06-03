# 3635. Earliest Finish Time for Land and Water Rides II

[LeetCode Link](https://leetcode.com/problems/earliest-finish-time-for-land-and-water-rides-ii/)

Difficulty: Medium
Topics: Array, Two Pointers, Binary Search, Greedy, Sorting

Acceptance Rate: 52.9%

## Hints

### Hint 1

The brute force is to try every pair of (land ride, water ride) in both orders — that is `O(n * m)` combinations. With `n` and `m` up to `5 * 10^4` each, that's `2.5 * 10^9` pairs, far too slow. Ask yourself: do you really need to look at every *pair*, or can you decouple the two categories?

### Hint 2

Consider just one ordering, say **land first, then water**. If you commit to a specific water ride `j`, what is the *best* land ride to pair it with? Notice the finish time only depends on the land ride through a single number: when that land ride finishes. So you don't need to know which land ride you used — you only need the smallest possible value of that number.

### Hint 3

For the "land then water" ordering, the finish time when pairing some land ride with water ride `j` is `max(landFinish, waterStartTime[j]) + waterDuration[j]`. This is *monotonic* in `landFinish`, so the best choice is always the land ride with the **minimum** `landStartTime[i] + landDuration[i]`. Precompute that single minimum once, then sweep every water ride. Do the symmetric thing for "water then land". The two minimums collapse the problem from `O(n*m)` to `O(n+m)`.

## Approach

A valid plan picks exactly one land ride and one water ride and runs them back-to-back in one of two orders. We handle each order independently.

**Order 1 — land ride first, then water ride.**
If the land ride finishes at time `f`, then the water ride `j` can start no earlier than `max(f, waterStartTime[j])` and finishes at `max(f, waterStartTime[j]) + waterDuration[j]`. For a fixed `j`, this expression is non-decreasing in `f`, so among all land rides we always prefer the one with the smallest finish time. Define:

```
minLandFinish = min over i of (landStartTime[i] + landDuration[i])
```

Then the best "land first" result is:

```
min over j of ( max(minLandFinish, waterStartTime[j]) + waterDuration[j] )
```

**Order 2 — water ride first, then land ride.**
By identical reasoning, define:

```
minWaterFinish = min over j of (waterStartTime[j] + waterDuration[j])
```

and the best "water first" result is:

```
min over i of ( max(minWaterFinish, landStartTime[i]) + landDuration[i] )
```

The answer is the minimum of the two orderings.

**Why the decoupling is valid.** The only way the first ride influences the second is through the first ride's finish time, and the second-ride finish time is monotonically non-decreasing in that value. So fixing "use the earliest-finishing ride of the first category" is never worse than any other choice. This greedy collapse is what lets us avoid examining all pairs.

**Walkthrough of Example 1:** `landStartTime = [2,8]`, `landDuration = [4,1]`, `waterStartTime = [6]`, `waterDuration = [3]`.
- Land finishes: `2+4=6`, `8+1=9` → `minLandFinish = 6`.
- Water finishes: `6+3=9` → `minWaterFinish = 9`.
- Land first: for water ride 0, `max(6, 6) + 3 = 9`.
- Water first: for land ride 0, `max(9, 2) + 4 = 13`; for land ride 1, `max(9, 8) + 1 = 10`. Best `= 10`.
- Answer `= min(9, 10) = 9`. ✓

## Complexity Analysis

Time Complexity: O(n + m) — two linear passes to compute the minimums, two linear passes to sweep the other category.
Space Complexity: O(1) — only a handful of scalar accumulators are kept.

## Edge Cases

- **Single ride in each category (`n = m = 1`):** both orderings still get evaluated; the formula handles it without special casing.
- **Second ride opens long before the first finishes:** the `max(...)` clamps the start to the first ride's finish time, so waiting is never modeled as starting early. Example 2 exercises this (water opens at 1 but can't start until land/water predecessor frees up).
- **First ride finishes before the second ride even opens:** then `max(minFinish, startTime) == startTime`, i.e. the tourist waits idle for the second ride to open — the `max` captures this correctly.
- **Large values / overflow:** with values up to `10^5`, a finish time stays well within `int` range, but the reasoning generalizes; using Go's `int` (64-bit on typical platforms) is safe.
- **It is never optimal to pick a slower-finishing first ride** — the monotonic argument guarantees the single minimum suffices, so don't be tempted to keep more candidates than that.
