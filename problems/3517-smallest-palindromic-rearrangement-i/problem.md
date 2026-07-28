---
number: "3517"
frontend_id: "3517"
title: "Smallest Palindromic Rearrangement I"
slug: "smallest-palindromic-rearrangement-i"
difficulty: "Medium"
topics:
  - "String"
  - "Sorting"
  - "Counting Sort"
acceptance_rate: 6759.3
is_premium: false
created_at: "2026-07-28T03:44:45.988206+00:00"
fetched_at: "2026-07-28T03:44:45.988206+00:00"
link: "https://leetcode.com/problems/smallest-palindromic-rearrangement-i/"
date: "2026-07-28"
---

# 3517. Smallest Palindromic Rearrangement I

You are given a **palindromic** string `s`.

Return the **lexicographically smallest** palindromic permutation of `s`.

 

**Example 1:**

**Input:** s = "z"

**Output:** "z"

**Explanation:**

A string of only one character is already the lexicographically smallest palindrome.

**Example 2:**

**Input:** s = "babab"

**Output:** "abbba"

**Explanation:**

Rearranging `"babab"` -> `"abbba"` gives the smallest lexicographic palindrome.

**Example 3:**

**Input:** s = "daccad"

**Output:** "acddca"

**Explanation:**

Rearranging `"daccad"` -> `"acddca"` gives the smallest lexicographic palindrome.

 

**Constraints:**

  * `1 <= s.length <= 105`
  * `s` consists of lowercase English letters.
  * `s` is guaranteed to be palindromic.
