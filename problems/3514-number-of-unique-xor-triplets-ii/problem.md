---
number: "3514"
frontend_id: "3514"
title: "Number of Unique XOR Triplets II"
slug: "number-of-unique-xor-triplets-ii"
difficulty: "Medium"
topics:
  - "Array"
  - "Math"
  - "Bit Manipulation"
  - "Enumeration"
acceptance_rate: 4023.6
is_premium: false
created_at: "2026-07-24T03:51:53.056616+00:00"
fetched_at: "2026-07-24T03:51:53.056616+00:00"
link: "https://leetcode.com/problems/number-of-unique-xor-triplets-ii/"
date: "2026-07-24"
---

# 3514. Number of Unique XOR Triplets II

You are given an integer array `nums`.

A **XOR triplet** is defined as the XOR of three elements `nums[i] XOR nums[j] XOR nums[k]` where `i <= j <= k`.

Return the number of **unique** XOR triplet values from all possible triplets `(i, j, k)`.

 

**Example 1:**

**Input:** nums = [1,3]

**Output:** 2

**Explanation:**

The possible XOR triplet values are:

  * `(0, 0, 0) -> 1 XOR 1 XOR 1 = 1`
  * `(0, 0, 1) -> 1 XOR 1 XOR 3 = 3`
  * `(0, 1, 1) -> 1 XOR 3 XOR 3 = 1`
  * `(1, 1, 1) -> 3 XOR 3 XOR 3 = 3`



The unique XOR values are `{1, 3}`. Thus, the output is 2.

**Example 2:**

**Input:** nums = [6,7,8,9]

**Output:** 4

**Explanation:**

The possible XOR triplet values are `{6, 7, 8, 9}`. Thus, the output is 4.

 

**Constraints:**

  * `1 <= nums.length <= 1500`
  * `1 <= nums[i] <= 1500`
