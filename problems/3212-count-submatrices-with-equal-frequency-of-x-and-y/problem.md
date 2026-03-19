---
number: "3212"
frontend_id: "3212"
title: "Count Submatrices With Equal Frequency of X and Y"
slug: "count-submatrices-with-equal-frequency-of-x-and-y"
difficulty: "Medium"
topics:
  - "Array"
  - "Matrix"
  - "Prefix Sum"
acceptance_rate: 5604.8
is_premium: false
created_at: "2026-03-19T03:25:19.837760+00:00"
fetched_at: "2026-03-19T03:25:19.837760+00:00"
link: "https://leetcode.com/problems/count-submatrices-with-equal-frequency-of-x-and-y/"
date: "2026-03-19"
---

# 3212. Count Submatrices With Equal Frequency of X and Y

Given a 2D character matrix `grid`, where `grid[i][j]` is either `'X'`, `'Y'`, or `'.'`, return the number of submatrices that contain:

  * `grid[0][0]`
  * an **equal** frequency of `'X'` and `'Y'`.
  * **at least** one `'X'`.



 

**Example 1:**

**Input:** grid = [["X","Y","."],["Y",".","."]]

**Output:** 3

**Explanation:**

**![](https://assets.leetcode.com/uploads/2024/06/07/examplems.png)**

**Example 2:**

**Input:** grid = [["X","X"],["X","Y"]]

**Output:** 0

**Explanation:**

No submatrix has an equal frequency of `'X'` and `'Y'`.

**Example 3:**

**Input:** grid = [[".","."],[".","."]]

**Output:** 0

**Explanation:**

No submatrix has at least one `'X'`.

 

**Constraints:**

  * `1 <= grid.length, grid[i].length <= 1000`
  * `grid[i][j]` is either `'X'`, `'Y'`, or `'.'`.
