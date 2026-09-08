---
number: "3870"
frontend_id: "3870"
title: "Count Commas in Range"
slug: "count-commas-in-range"
difficulty: "Easy"
topics:
  - "Math"
acceptance_rate: 7346.8
is_premium: false
created_at: "2026-09-08T04:57:47.921771+00:00"
fetched_at: "2026-09-08T04:57:47.921771+00:00"
link: "https://leetcode.com/problems/count-commas-in-range/"
date: "2026-09-08"
---

# 3870. Count Commas in Range

You are given an integer `n`.

Return the **total** number of commas used when writing all integers from `[1, n]` (inclusive) in **standard** number formatting.

In **standard** formatting:

  * A comma is inserted after **every three** digits from the right.
  * Numbers with **fewer** than 4 digits contain no commas.



 

**Example 1:**

**Input:** n = 1002

**Output:** 3

**Explanation:**

The numbers `"1,000"`, `"1,001"`, and `"1,002"` each contain one comma, giving a total of 3.

**Example 2:**

**Input:** n = 998

**Output:** 0

**Explanation:**

All numbers from 1 to 998 have fewer than four digits. Therefore, no commas are used.

 

**Constraints:**

  * `1 <= n <= 105`
