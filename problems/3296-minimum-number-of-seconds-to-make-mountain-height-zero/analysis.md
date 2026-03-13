# 3296. Minimum Number of Seconds to Make Mountain Height Zero

[LeetCode Link](https://leetcode.com/problems/minimum-number-of-seconds-to-make-mountain-height-zero/)

Difficulty: Medium
Topics: Array, Math, Binary Search, Greedy, Heap (Priority Queue)
Acceptance Rate: 41.1%

## Hints

### Hint 1

Think about the problem from the perspective of the answer. Instead of figuring out how to optimally distribute work among workers, what if you already knew the total time allowed? Could you then determine whether the workers can finish the job?

### Hint 2

Binary search on the answer. For a candidate time T, each worker i can reduce the mountain height by some maximum amount x, where the cumulative cost `workerTimes[i] * (1 + 2 + ... + x) = workerTimes[i] * x*(x+1)/2` must not exceed T. Sum up all workers' contributions and check if it meets or exceeds `mountainHeight`.

### Hint 3

To find each worker's maximum contribution for a given time T, solve the inequality `workerTimes[i] * x * (x + 1) / 2 <= T`. This is a quadratic in x, giving `x = floor((-1 + sqrt(1 + 8*T/workerTimes[i])) / 2)`. The binary search range for T goes from 1 up to `min(workerTimes) * mountainHeight * (mountainHeight + 1) / 2` (worst case: the fastest worker does everything alone).

## Approach

This problem is a classic "binary search on the answer" pattern. The key observation is that if the workers can finish in T seconds, they can also finish in T+1 seconds (monotonicity). This makes binary search applicable.

**Algorithm:**

1. **Binary search on time T.** Set `lo = 1` and `hi = min(workerTimes) * mountainHeight * (mountainHeight + 1) / 2`. This upper bound represents the scenario where only the fastest worker does all the work alone.

2. **Feasibility check.** For a candidate time T, determine how much height each worker can reduce. Worker i with rate `workerTimes[i]` reducing height by x takes `workerTimes[i] * x * (x + 1) / 2` seconds. We need the largest x such that this quantity is at most T. Solving the quadratic inequality gives:
   - `x = floor((-1 + sqrt(1 + 8 * T / workerTimes[i])) / 2)`

3. **Sum contributions.** Add up the maximum x for each worker. If the total is at least `mountainHeight`, then T seconds is enough.

4. **Narrow the search.** Use standard binary search: if T is feasible, try smaller; otherwise try larger.

**Example walkthrough (mountainHeight=4, workerTimes=[2,1,1]):**
- Try T=3: Worker 0 can do floor((-1+sqrt(1+12))/2) = floor(1.30) = 1. Workers 1 and 2 can each do floor((-1+sqrt(1+24))/2) = floor(2.0) = 2. Total = 1+2+2 = 5 >= 4. Feasible.
- Try T=2: Worker 0 can do 1, Workers 1&2 can each do 1. Total = 3 < 4. Not feasible.
- Answer: 3.

## Complexity Analysis

Time Complexity: O(n * log(minW * H * (H+1) / 2)), where n is the number of workers, H is the mountain height, and minW is the minimum worker time. The binary search runs over the time range and each feasibility check iterates over all workers.

Space Complexity: O(1), only constant extra space is used.

## Edge Cases

- **Single worker**: The worker must do all the work alone. The answer is `workerTimes[0] * mountainHeight * (mountainHeight + 1) / 2`.
- **Mountain height of 1**: Each worker can reduce by at least 1 in `workerTimes[i]` seconds, so the answer is `min(workerTimes)`.
- **All workers have the same rate**: Contributions are evenly distributed, simplifying the feasibility check but the algorithm handles it the same way.
- **Large values**: With mountainHeight up to 10^5 and workerTimes[i] up to 10^6, the maximum time can be around 10^6 * 10^5 * 10^5 / 2 = 5 * 10^15, which fits in int64.
