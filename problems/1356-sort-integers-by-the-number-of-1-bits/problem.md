---
number: "1356"
frontend_id: "1356"
title: "Sort Integers by The Number of 1 Bits"
slug: "sort-integers-by-the-number-of-1-bits"
difficulty: "Easy"
topics:
  - "Array"
  - "Bit Manipulation"
  - "Sorting"
  - "Counting"
acceptance_rate: 7952.3
is_premium: false
created_at: "2026-02-25T03:20:39.915392+00:00"
fetched_at: "2026-02-25T03:20:39.915392+00:00"
link: "https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/"
date: "2026-02-25"
---

# 1356. Sort Integers by The Number of 1 Bits

You are given an integer array `arr`. Sort the integers in the array in ascending order by the number of `1`'s in their binary representation and in case of two or more integers have the same number of `1`'s you have to sort them in ascending order.

Return _the array after sorting it_.

 

**Example 1:**
    
    
    **Input:** arr = [0,1,2,3,4,5,6,7,8]
    **Output:** [0,1,2,4,8,3,5,6,7]
    **Explantion:** [0] is the only integer with 0 bits.
    [1,2,4,8] all have 1 bit.
    [3,5,6] have 2 bits.
    [7] has 3 bits.
    The sorted array by bits is [0,1,2,4,8,3,5,6,7]
    

**Example 2:**
    
    
    **Input:** arr = [1024,512,256,128,64,32,16,8,4,2,1]
    **Output:** [1,2,4,8,16,32,64,128,256,512,1024]
    **Explantion:** All integers have 1 bit in the binary representation, you should just sort them in ascending order.
    

 

**Constraints:**

  * `1 <= arr.length <= 500`
  * `0 <= arr[i] <= 104`
