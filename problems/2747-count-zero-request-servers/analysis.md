# 2747. Count Zero Request Servers

[LeetCode Link](https://leetcode.com/problems/count-zero-request-servers/)

Difficulty: Medium
Topics: Array, Hash Table, Sliding Window, Sorting
Acceptance Rate: 46.7%

## Hints

### Hint 1

Checking every log for every query costs O(logs × queries), which is too slow. All the windows have the same length `x`. What would you get if you handled the queries in time order instead of the order they were given?

### Hint 2

Sort the logs by time. Also sort the query indices by query time, so you can still write each answer back to its original position. Then slide one window `[q - x, q]` across the sorted logs with two pointers that only move forward.

### Hint 3

Keep a count of how many log entries in the current window belong to each server, plus a counter of how many servers have a count above zero. When the right pointer adds a log and that server's count goes 0 → 1, add one to the active counter. When the left pointer removes a log and a count goes 1 → 0, subtract one. The answer for a query is `n - active`.

## Approach

1. **Sort logs by time.** After sorting, the logs inside any time window form one contiguous block.
2. **Sort the queries offline.** Build an index array `order` of `0..len(queries)-1` and sort it by `queries[i]`. Answers are written to `ans[order[k]]`, so the original order is kept.
3. **Two pointers over the logs:**
   - `right` is the first log not yet added. While `logs[right].time <= q`, add it: `cnt[server]++`, and if the count became 1, `active++`.
   - `left` is the first log still in the window. While `logs[left].time < q - x`, remove it: `cnt[server]--`, and if the count became 0, `active--`.
   - Record `ans[idx] = n - active`.
4. Queries are processed in non-decreasing order, so both window ends `q - x` and `q` never move backward. That means `left` and `right` only move forward, and each log is added once and removed at most once.

**Example 1:** logs sorted = `[(1,3),(1,5),(2,6)]`, x = 5.
- q = 10, window [5,10]: add all three logs, then remove (1,3) because 3 < 5. Counts: server 1 → 1, server 2 → 1. active = 2, so the answer is 3 − 2 = **1**.
- q = 11, window [6,11]: remove (1,5) because 5 < 6. Server 1's count drops to 0, so active = 1 and the answer is 3 − 1 = **2**.

The counter only changes on 0↔1 transitions, so each step costs O(1). You never need to scan the hash table to count active servers.

## Complexity Analysis

Time Complexity: O(L log L + Q log Q + n), where L = len(logs) and Q = len(queries). The two sorts dominate, and the sweep itself is linear.
Space Complexity: O(n + Q) for the per-server count array, the sorted query indices, and the answer.

## Edge Cases

- **Several logs for the same server inside one window:** count individual logs, not just whether a server appeared. Otherwise, removing one old log could wrongly mark the server inactive while a newer log is still in the window.
- **Inclusive boundaries:** a log at exactly `q - x` or exactly `q` belongs in the window. Use `time <= q` to add and `time < q - x` to remove.
- **Duplicate or unsorted queries:** queries can come in any order and can repeat. The offline sort with original indices handles both.
- **All logs outside the window:** the answer is `n`.
- **Servers that never appear in the logs:** they always count as zero-request servers. That is why the answer is `n - active`, not something computed from the servers seen in the logs.
