---
number: "1288"
frontend_id: "1288"
title: "Remove Covered Intervals"
slug: "remove-covered-intervals"
difficulty: "Medium"
topics:
  - "Array"
  - "Sorting"
acceptance_rate: 5696.8
is_premium: false
created_at: "2026-07-06T04:49:33.821768+00:00"
fetched_at: "2026-07-06T04:49:33.821768+00:00"
link: "https://leetcode.com/problems/remove-covered-intervals/"
date: "2026-07-06"
---

# 1288. Remove Covered Intervals

Given an array `intervals` where `intervals[i] = [li, ri]` represent the interval `[li, ri)`, remove all intervals that are covered by another interval in the list.

The interval `[a, b)` is covered by the interval `[c, d)` if and only if `c <= a` and `b <= d`.

Return _the number of remaining intervals_.

 

**Example 1:**
    
    
    **Input:** intervals = [[1,4],[3,6],[2,8]]
    **Output:** 2
    **Explanation:** Interval [3,6] is covered by [2,8], therefore it is removed.
    

**Example 2:**
    
    
    **Input:** intervals = [[1,4],[2,3]]
    **Output:** 1
    

 

**Constraints:**

  * `1 <= intervals.length <= 1000`
  * `intervals[i].length == 2`
  * `0 <= li < ri <= 105`
  * All the given intervals are **unique**.
