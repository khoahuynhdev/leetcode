---
number: "1524"
frontend_id: "1524"
title: "Number of Sub-arrays With Odd Sum"
slug: "number-of-sub-arrays-with-odd-sum"
difficulty: "Medium"
topics:
  - "Array"
  - "Math"
  - "Dynamic Programming"
  - "Prefix Sum"
acceptance_rate: 5673.8
is_premium: false
created_at: "2026-11-15T06:55:34.175321+00:00"
fetched_at: "2026-11-15T06:55:34.175321+00:00"
link: "https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/"
date: "2026-11-15"
---

# 1524. Number of Sub-arrays With Odd Sum

Given an array of integers `arr`, return _the number of subarrays with an**odd** sum_.

Since the answer can be very large, return it modulo `109 + 7`.

 

**Example 1:**
    
    
    **Input:** arr = [1,3,5]
    **Output:** 4
    **Explanation:** All subarrays are [[1],[1,3],[1,3,5],[3],[3,5],[5]]
    All sub-arrays sum are [1,4,9,3,8,5].
    Odd sums are [1,9,3,5] so the answer is 4.
    

**Example 2:**
    
    
    **Input:** arr = [2,4,6]
    **Output:** 0
    **Explanation:** All subarrays are [[2],[2,4],[2,4,6],[4],[4,6],[6]]
    All sub-arrays sum are [2,6,12,4,10,6].
    All sub-arrays have even sum and the answer is 0.
    

**Example 3:**
    
    
    **Input:** arr = [1,2,3,4,5,6,7]
    **Output:** 16
    

 

**Constraints:**

  * `1 <= arr.length <= 105`
  * `1 <= arr[i] <= 100`
