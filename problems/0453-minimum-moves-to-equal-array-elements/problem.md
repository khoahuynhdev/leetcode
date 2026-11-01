---
number: "0453"
frontend_id: "453"
title: "Minimum Moves to Equal Array Elements"
slug: "minimum-moves-to-equal-array-elements"
difficulty: "Medium"
topics:
  - "Array"
  - "Math"
acceptance_rate: 6058.5
is_premium: false
created_at: "2026-11-01T06:33:44.578855+00:00"
fetched_at: "2026-11-01T06:33:44.578855+00:00"
link: "https://leetcode.com/problems/minimum-moves-to-equal-array-elements/"
date: "2026-11-01"
---

# 0453. Minimum Moves to Equal Array Elements

Given an integer array `nums` of size `n`, return _the minimum number of moves required to make all array elements equal_.

In one move, you can increment `n - 1` elements of the array by `1`.

 

**Example 1:**
    
    
    **Input:** nums = [1,2,3]
    **Output:** 3
    **Explanation:** Only three moves are needed (remember each move increments two elements):
    [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
    

**Example 2:**
    
    
    **Input:** nums = [1,1,1]
    **Output:** 0
    

 

**Constraints:**

  * `n == nums.length`
  * `1 <= nums.length <= 105`
  * `-109 <= nums[i] <= 109`
  * The answer is guaranteed to fit in a **32-bit** integer.
