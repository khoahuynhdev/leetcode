---
number: "3702"
frontend_id: "3702"
title: "Longest Subsequence With Non-Zero Bitwise XOR"
slug: "longest-subsequence-with-non-zero-bitwise-xor"
difficulty: "Medium"
topics:
  - "Array"
  - "Bit Manipulation"
acceptance_rate: 3932.2
is_premium: false
created_at: "2026-08-15T01:55:46.660175+00:00"
fetched_at: "2026-08-15T01:55:46.660175+00:00"
link: "https://leetcode.com/problems/longest-subsequence-with-non-zero-bitwise-xor/"
date: "2026-08-15"
---

# 3702. Longest Subsequence With Non-Zero Bitwise XOR

You are given an integer array `nums`.

Return the length of the **longest subsequence** in `nums` whose bitwise **XOR** is **non-zero**. If no such **subsequence** exists, return 0.

 

**Example 1:**

**Input:** nums = [1,2,3]

**Output:** 2

**Explanation:**

One longest subsequence is `[2, 3]`. The bitwise XOR is computed as `2 XOR 3 = 1`, which is non-zero.

**Example 2:**

**Input:** nums = [2,3,4]

**Output:** 3

**Explanation:**

The longest subsequence is `[2, 3, 4]`. The bitwise XOR is computed as `2 XOR 3 XOR 4 = 5`, which is non-zero.

 

**Constraints:**

  * `1 <= nums.length <= 105`
  * `0 <= nums[i] <= 109`
