---
number: "3300"
frontend_id: "3300"
title: "Minimum Element After Replacement With Digit Sum"
slug: "minimum-element-after-replacement-with-digit-sum"
difficulty: "Easy"
topics:
  - "Array"
  - "Math"
acceptance_rate: 8709.0
is_premium: false
created_at: "2026-05-29T04:51:25.405449+00:00"
fetched_at: "2026-05-29T04:51:25.405449+00:00"
link: "https://leetcode.com/problems/minimum-element-after-replacement-with-digit-sum/"
date: "2026-05-29"
---

# 3300. Minimum Element After Replacement With Digit Sum

You are given an integer array `nums`.

You replace each element in `nums` with the **sum** of its digits.

Return the **minimum** element in `nums` after all replacements.

 

**Example 1:**

**Input:** nums = [10,12,13,14]

**Output:** 1

**Explanation:**

`nums` becomes `[1, 3, 4, 5]` after all replacements, with minimum element 1.

**Example 2:**

**Input:** nums = [1,2,3,4]

**Output:** 1

**Explanation:**

`nums` becomes `[1, 2, 3, 4]` after all replacements, with minimum element 1.

**Example 3:**

**Input:** nums = [999,19,199]

**Output:** 10

**Explanation:**

`nums` becomes `[27, 10, 19]` after all replacements, with minimum element 10.

 

**Constraints:**

  * `1 <= nums.length <= 100`
  * `1 <= nums[i] <= 104`
