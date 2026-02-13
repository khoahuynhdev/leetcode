---
number: "3714"
frontend_id: "3714"
title: "Longest Balanced Substring II"
slug: "longest-balanced-substring-ii"
difficulty: "Medium"
topics:
  - "Hash Table"
  - "String"
  - "Prefix Sum"
acceptance_rate: 2663.3
is_premium: false
created_at: "2026-02-13T03:26:32.958422+00:00"
fetched_at: "2026-02-13T03:26:32.958422+00:00"
link: "https://leetcode.com/problems/longest-balanced-substring-ii/"
date: "2026-02-13"
---

# 3714. Longest Balanced Substring II

You are given a string `s` consisting only of the characters `'a'`, `'b'`, and `'c'`.

A **substring** of `s` is called **balanced** if all **distinct** characters in the **substring** appear the **same** number of times.

Return the **length of the longest balanced substring** of `s`.

 

**Example 1:**

**Input:** s = "abbac"

**Output:** 4

**Explanation:**

The longest balanced substring is `"abba"` because both distinct characters `'a'` and `'b'` each appear exactly 2 times.

**Example 2:**

**Input:** s = "aabcc"

**Output:** 3

**Explanation:**

The longest balanced substring is `"abc"` because all distinct characters `'a'`, `'b'` and `'c'` each appear exactly 1 time.

**Example 3:**

**Input:** s = "aba"

**Output:** 2

**Explanation:**

One of the longest balanced substrings is `"ab"` because both distinct characters `'a'` and `'b'` each appear exactly 1 time. Another longest balanced substring is `"ba"`.

 

**Constraints:**

  * `1 <= s.length <= 105`
  * `s` contains only the characters `'a'`, `'b'`, and `'c'`.
