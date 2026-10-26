---
number: "3619"
frontend_id: "3619"
title: "Count Islands With Total Value Divisible by K"
slug: "count-islands-with-total-value-divisible-by-k"
difficulty: "Medium"
topics:
  - "Array"
  - "Depth-First Search"
  - "Breadth-First Search"
  - "Union-Find"
  - "Matrix"
acceptance_rate: 5989.3
is_premium: false
created_at: "2026-10-26T05:06:20.123304+00:00"
fetched_at: "2026-10-26T05:06:20.123304+00:00"
link: "https://leetcode.com/problems/count-islands-with-total-value-divisible-by-k/"
date: "2026-10-26"
---

# 3619. Count Islands With Total Value Divisible by K

You are given an `m x n` matrix `grid` and a positive integer `k`. An **island** is a group of **positive** integers (representing land) that are **4-directionally** connected (horizontally or vertically).

The **total value** of an island is the sum of the values of all cells in the island.

Return the number of islands with a total value **divisible by** `k`.

 

**Example 1:**

![](https://assets.leetcode.com/uploads/2025/03/06/example1griddrawio-1.png)

**Input:** grid = [[0,2,1,0,0],[0,5,0,0,5],[0,0,1,0,0],[0,1,4,7,0],[0,2,0,0,8]], k = 5

**Output:** 2

**Explanation:**

The grid contains four islands. The islands highlighted in blue have a total value that is divisible by 5, while the islands highlighted in red do not.

**Example 2:**

![](https://assets.leetcode.com/uploads/2025/03/06/example2griddrawio.png)

**Input:** grid = [[3,0,3,0], [0,3,0,3], [3,0,3,0]], k = 3

**Output:** 6

**Explanation:**

The grid contains six islands, each with a total value that is divisible by 3.

 

**Constraints:**

  * `m == grid.length`
  * `n == grid[i].length`
  * `1 <= m, n <= 1000`
  * `1 <= m * n <= 105`
  * `0 <= grid[i][j] <= 106`
  * `1 <= k <= 106`
