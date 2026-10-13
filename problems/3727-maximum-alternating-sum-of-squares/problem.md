---
number: "3727"
frontend_id: "3727"
title: "Maximum Alternating Sum of Squares"
slug: "maximum-alternating-sum-of-squares"
difficulty: "Medium"
topics:
  - "Array"
  - "Greedy"
  - "Sorting"
acceptance_rate: 7093.9
is_premium: false
created_at: "2026-10-13T06:57:00.428067+00:00"
fetched_at: "2026-10-13T06:57:00.428067+00:00"
link: "https://leetcode.com/problems/maximum-alternating-sum-of-squares/"
date: "2026-10-13"
---

# 3727. Maximum Alternating Sum of Squares

You are given an integer array `nums`. You may **rearrange the elements** in any order.

The **alternating score** of an array `arr` is defined as:

  * `score = arr[0]2 - arr[1]2 + arr[2]2 - arr[3]2 + ...`



Return an integer denoting the **maximum possible alternating score** of `nums` after rearranging its elements.

 

**Example 1:**

**Input:** nums = [1,2,3]

**Output:** 12

**Explanation:**

A possible rearrangement for `nums` is `[2,1,3]`, which gives the maximum alternating score among all possible rearrangements.

The alternating score is calculated as:

`score = 22 - 12 + 32 = 4 - 1 + 9 = 12`

**Example 2:**

**Input:** nums = [1,-1,2,-2,3,-3]

**Output:** 16

**Explanation:**

A possible rearrangement for `nums` is `[-3,-1,-2,1,3,2]`, which gives the maximum alternating score among all possible rearrangements.

The alternating score is calculated as:

`score = (-3)2 - (-1)2 + (-2)2 - (1)2 + (3)2 - (2)2 = 9 - 1 + 4 - 1 + 9 - 4 = 16`

 

**Constraints:**

  * `1 <= nums.length <= 105`
  * `-4 * 104 <= nums[i] <= 4 * 104`
