# 1606. Find Servers That Handled Most Number of Requests

[LeetCode Link](https://leetcode.com/problems/find-servers-that-handled-most-number-of-requests/)

Difficulty: Hard
Topics: Array, Heap (Priority Queue), Simulation, Ordered Set
Acceptance Rate: 53.1%

## Hints

### Hint 1

This is a simulation: go through the requests in arrival order and decide which server takes each one. Scanning all `k` servers for every request is O(n·k), which is too slow when both are up to 10^5. Two questions come up over and over: "which busy servers have finished by now?" and "which free server is the first one at or after `i % k`?" Pick a data structure that answers each of them quickly.

### Hint 2

For "who has finished?", use a **min-heap of busy servers keyed by end time**. Before each request, pop every server whose end time is `<= arrival[i]` and mark it free.

For "first free server at or after `i % k`, wrapping around", an **ordered set** of free server IDs works. Look up the smallest ID `>= i % k`. If there isn't one, take the smallest ID overall (that's the wrap-around). Go has no built-in ordered set, so think about whether a plain min-heap could do the same job.

### Hint 3

A trick lets a plain min-heap replace the ordered set: store each free server under a **virtual index** instead of its real ID. When server `j` becomes free just before request `i`, push the smallest number `v >= i` with `v % k == j`, which is `v = i + ((j - i) % k + k) % k`. Every value in the heap is then `>= i`, and the smallest value is exactly the first free server you reach by scanning `i%k, i%k+1, ...` with wrap-around. Get the real ID back with `v % k`.

## Approach

1. **State**
   - `count[j]`: how many requests server `j` has handled.
   - `busy`: a min-heap of `(endTime, id)`.
   - `available`: a min-heap of virtual indices. At the start it holds `0, 1, ..., k-1`, which are valid virtual indices for `i = 0`.

2. **For each request `i` arriving at time `t`:**
   - **Release:** while the top of `busy` has `endTime <= t`, pop it. Push its virtual index `i + ((id - i) % k + k) % k` onto `available`.
   - **Drop:** if `available` is empty, every server is busy, so skip the request.
   - **Assign:** pop the smallest virtual index `v`. The server is `id = v % k`. Increment `count[id]` and push `(t + load[i], id)` onto `busy`.

3. **Answer:** find the maximum in `count` and return every server ID that has that count.

### Why the virtual index trick works

The scan order for request `i` is `i, i+1, i+2, ...` (taken mod `k`). If each free server sits at a number in `[i, i+k)` that matches its residue mod `k`, then the smallest number is the first free server in that scan order.

The invariant is that every value in `available` is `>= i` at step `i`. To see why it holds:
- Newly pushed values are `>= i` by construction.
- Different servers have different residues mod `k`, so at most one value equals `i`. If it exists, it is the minimum, and it gets popped at step `i`.
- So when we move on to `i + 1`, every value left is `>= i + 1`. A value pushed at an earlier step is still in `[i, i+k)` relative to the current step, so the order stays correct even if we never update it.

### Example walk-through (Example 1: k = 3, arrival = [1,2,3,4,5], load = [5,2,3,3,3])

| i | t | released | available (virtual) | assigned | busy after |
|---|---|----------|---------------------|----------|------------|
| 0 | 1 | -        | {0,1,2}             | 0 → server 0 | (6,0) |
| 1 | 2 | -        | {1,2}               | 1 → server 1 | (4,1),(6,0) |
| 2 | 3 | -        | {2}                 | 2 → server 2 | (4,1),(6,0),(6,2) |
| 3 | 4 | server 1 → v = 3 + ((1-3)%3+3)%3 = 4 | {4} | 4 → server 1 | (6,0),(6,2),(7,1) |
| 4 | 5 | -        | {}                  | dropped  | |

Counts are `[1, 2, 1]`, so the answer is `[1]`.

### Alternative

You can also keep a balanced BST, a TreeSet, or a segment tree over server IDs and query "first free ID `>= i % k`, otherwise the first free ID overall". This has the same complexity but needs much more code in Go.

## Complexity Analysis

Time Complexity: O(n log k), where n is the number of requests. Each request causes at most one push and one pop on each heap, and each heap holds at most k entries. The final pass over `count` is O(k).
Space Complexity: O(k) for the two heaps and the `count` array.

## Edge Cases

- **All servers busy:** the request must be dropped, not queued. This is the case where `available` is empty.
- **Server finishes exactly when a request arrives:** a server is free at time `end` if `end <= arrival`. Use `<=`, not `<`, or you will wrongly mark it busy.
- **Wrap-around:** if no server from `i % k` to `k-1` is free, the search must continue from server 0. The virtual index handles this without any special case.
- **k = 1:** there is only one server, so each request is either handled by server 0 or dropped.
- **More servers than requests:** some servers handle zero requests. They must not appear in the answer unless the maximum count is also zero, which can't happen because the first request is always handled.
- **Ties:** several servers can share the maximum count, and all of them must be returned (Example 3).
- **Negative modulo in Go:** `(j - i) % k` can be negative in Go. Add `k` and take `% k` again to get a value in `[0, k)`.
- **Large values:** `arrival[i] + load[i]` can reach 2·10^9. Go's `int` is 64-bit on LeetCode, so it won't overflow.
