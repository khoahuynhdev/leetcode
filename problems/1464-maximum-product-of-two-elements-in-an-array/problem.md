---
number: "1464"
frontend_id: "1464"
title: "Maximum Product of Two Elements in an Array"
slug: "maximum-product-of-two-elements-in-an-array"
difficulty: "Easy"
topics:
  - "Array"
  - "Sorting"
  - "Heap (Priority Queue)"
acceptance_rate: 8395.6
is_premium: false
created_at: "2026-07-27T04:11:41.584531+00:00"
fetched_at: "2026-07-27T04:11:41.584531+00:00"
link: "https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/"
date: "2026-07-27"
---

# 1464. Maximum Product of Two Elements in an Array

Given the array of integers `nums`, you will choose two different indices `i` and `j` of that array. _Return the maximum value of_ `(nums[i]-1)*(nums[j]-1)`. 

 

**Example 1:**
    
    
    **Input:** nums = [3,4,5,2]
    **Output:** 12 
    **Explanation:** If you choose the indices i=1 and j=2 (indexed from 0), you will get the maximum value, that is, (nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12. 
    

**Example 2:**
    
    
    **Input:** nums = [1,5,4,5]
    **Output:** 16
    **Explanation:** Choosing the indices i=1 and j=3 (indexed from 0), you will get the maximum value of (5-1)*(5-1) = 16.
    

**Example 3:**
    
    
    **Input:** nums = [3,7]
    **Output:** 12
    

 

**Constraints:**

  * `2 <= nums.length <= 500`
  * `1 <= nums[i] <= 10^3`
