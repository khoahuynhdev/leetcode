---
number: "3718"
frontend_id: "3718"
title: "Smallest Missing Multiple of K"
slug: "smallest-missing-multiple-of-k"
difficulty: "Easy"
topics:
  - "Array"
  - "Hash Table"
acceptance_rate: 6447.0
is_premium: false
created_at: "2026-08-25T01:58:20.704602+00:00"
fetched_at: "2026-08-25T01:58:20.704602+00:00"
link: "https://leetcode.com/problems/smallest-missing-multiple-of-k/"
date: "2026-08-25"
---

# 3718. Smallest Missing Multiple of K

Given an integer array `nums` and an integer `k`, return the **smallest positive multiple** of `k` that is **missing** from `nums`.

A **multiple** of `k` is any positive integer divisible by `k`.

 

**Example 1:**

**Input:** nums = [8,2,3,4,6], k = 2

**Output:** 10

**Explanation:**

The multiples of `k = 2` are 2, 4, 6, 8, 10, 12... and the smallest multiple missing from `nums` is 10.

**Example 2:**

**Input:** nums = [1,4,7,10,15], k = 5

**Output:** 5

**Explanation:**

The multiples of `k = 5` are 5, 10, 15, 20... and the smallest multiple missing from `nums` is 5.

 

**Constraints:**

  * `1 <= nums.length <= 100`
  * `1 <= nums[i] <= 100`
  * `1 <= k <= 100`
