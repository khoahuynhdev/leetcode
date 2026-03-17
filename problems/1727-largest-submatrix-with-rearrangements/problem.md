---
number: "1727"
frontend_id: "1727"
title: "Largest Submatrix With Rearrangements"
slug: "largest-submatrix-with-rearrangements"
difficulty: "Medium"
topics:
  - "Array"
  - "Greedy"
  - "Sorting"
  - "Matrix"
acceptance_rate: 7611.1
is_premium: false
created_at: "2026-03-17T03:17:29.130014+00:00"
fetched_at: "2026-03-17T03:17:29.130014+00:00"
link: "https://leetcode.com/problems/largest-submatrix-with-rearrangements/"
date: "2026-03-17"
---

# 1727. Largest Submatrix With Rearrangements

You are given a binary matrix `matrix` of size `m x n`, and you are allowed to rearrange the **columns** of the `matrix` in any order.

Return _the area of the largest submatrix within_`matrix` _where**every** element of the submatrix is _`1` _after reordering the columns optimally._

 

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/12/29/screenshot-2020-12-30-at-40536-pm.png)
    
    
    **Input:** matrix = [[0,0,1],[1,1,1],[1,0,1]]
    **Output:** 4
    **Explanation:** You can rearrange the columns as shown above.
    The largest submatrix of 1s, in bold, has an area of 4.
    

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/12/29/screenshot-2020-12-30-at-40852-pm.png)
    
    
    **Input:** matrix = [[1,0,1,0,1]]
    **Output:** 3
    **Explanation:** You can rearrange the columns as shown above.
    The largest submatrix of 1s, in bold, has an area of 3.
    

**Example 3:**
    
    
    **Input:** matrix = [[1,1,0],[1,0,1]]
    **Output:** 2
    **Explanation:** Notice that you must rearrange entire columns, and there is no way to make a submatrix of 1s larger than an area of 2.
    

 

**Constraints:**

  * `m == matrix.length`
  * `n == matrix[i].length`
  * `1 <= m * n <= 105`
  * `matrix[i][j]` is either `0` or `1`.
