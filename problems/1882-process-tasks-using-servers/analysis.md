# 1882. Process Tasks Using Servers

[LeetCode Link](https://leetcode.com/problems/process-tasks-using-servers/)

Difficulty: Medium
Topics: Array, Heap (Priority Queue)
Acceptance Rate: 45.6%

## Hints

### Hint 1

This is a simulation problem. Over and over you need "the best free server" and "the server that will be free soonest". Which data structure gives you the min of a changing set quickly?

### Hint 2

Use **two** priority queues. One holds idle servers ordered by `(weight, index)`. The other holds busy servers ordered by the time they become free. Servers move back and forth between the two heaps.

### Hint 3

Don't simulate every second. Keep a `time` variable. For task `j`, set `time = max(time, j)`. If no server is idle, **jump** `time` ahead to the earliest free time in the busy heap. Then move every busy server with `freeAt <= time` into the idle heap, and give the task to the top of the idle heap. Tasks are always handed out in arrival order, so the "queue" is just the loop index.

## Approach

1. Put every server into `free`, a min-heap keyed by `(weight, index)`.
2. Keep `busy`, a min-heap keyed by `(freeAt, weight, index)`, which starts empty.
3. Keep a running clock `time = 0`. For each task `j` with duration `d`:
   - `time = max(time, j)`. The task can't start before it arrives. If earlier tasks were delayed, `time` may already be past `j`, and this task waits behind them.
   - If `free` is empty, nothing can happen until the next release, so set `time = busy.top.freeAt`.
   - Pop every server in `busy` with `freeAt <= time` and push it into `free`. Releasing them all at once matters: if several servers come free at the same moment, the choice between them has to follow weight and index, not release order.
   - Pop the top of `free`. That server is the answer for task `j`. Set its `freeAt = time + d` and push it into `busy`.

**Why the clock never goes backwards:** tasks leave the queue in order. Task `j+1` can't start before task `j`, so `time` only moves forward and one pass is enough.

**Walkthrough (Example 1):** servers = `[3,3,2]`, tasks = `[1,2,3,2,1,2]`.
- t=0: free = {2,0,1}. Task 0 → server 2 (free at 1).
- t=1: server 2 is released. Task 1 → server 2 (free at 3).
- t=2: Task 2 → server 0 (free at 5).
- t=3: server 2 is released. Task 3 → server 2 (free at 5).
- t=4: Task 4 → server 1 (free at 5).
- t=5: all three are released. Task 5 → server 2.

Answer: `[2,2,0,2,1,2]`.

## Complexity Analysis

Time Complexity: O((n + m) log n). Building the heap is O(n). Each task does one pop and one push on each heap, and each release is paid for by an earlier push.
Space Complexity: O(n + m): O(n) for the two heaps and O(m) for the answer.

## Edge Cases

- **All servers busy when a task arrives:** the clock has to jump to the next release time instead of using `j`. Forgetting this is the most common bug.
- **Several servers free at the same time:** release all of them before choosing, so the tie is broken by weight and then index.
- **Tasks piling up (delayed clock):** after a jump, `time` can be larger than later arrival indices. Use `max(time, j)` and never reset `time` back to `j`.
- **Equal weights:** the lower index must win, both in the idle heap and among servers released together.
- **Single server:** every task goes to server 0, but the timing logic still has to handle long waits.
- **Large values:** `freeAt` can reach about 2·10⁵ + 2·10⁵·2·10⁵ ≈ 4·10¹⁰. That fits in Go's 64-bit `int`, but it would overflow a 32-bit integer in other languages.
