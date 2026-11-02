---
number: "0462"
frontend_id: "462"
title: "Minimum Moves to Equal Array Elements II"
slug: "minimum-moves-to-equal-array-elements-ii"
difficulty: "Medium"
topics:
  - "Array"
  - "Math"
  - "Sorting"
acceptance_rate: 6379.9
is_premium: false
created_at: "2026-11-02T06:33:58.676167+00:00"
fetched_at: "2026-11-02T06:33:58.676167+00:00"
link: "https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/"
date: "2026-11-02"
---

# 0462. Minimum Moves to Equal Array Elements II

Given an integer array `nums` of size `n`, return _the minimum number of moves required to make all array elements equal_.

In one move, you can increment or decrement an element of the array by `1`.

Test cases are designed so that the answer will fit in a **32-bit** integer.

 

**Example 1:**
    
    
    **Input:** nums = [1,2,3]
    **Output:** 2
    **Explanation:**
    Only two moves are needed (remember each move increments or decrements one element):
    [_1_ ,2,3]  =>  [2,2,_3_]  =>  [2,2,2]
    

**Example 2:**
    
    
    **Input:** nums = [1,10,2,9]
    **Output:** 16
    

 

**Constraints:**

  * `n == nums.length`
  * `1 <= nums.length <= 105`
  * `-109 <= nums[i] <= 109`
