---
number: "1614"
frontend_id: "1614"
title: "Maximum Nesting Depth of the Parentheses"
slug: "maximum-nesting-depth-of-the-parentheses"
difficulty: "Easy"
topics:
  - "String"
  - "Stack"
  - "Bracket Sequences"
acceptance_rate: 8541.5
is_premium: false
created_at: "2026-09-28T05:41:26.612719+00:00"
fetched_at: "2026-09-28T05:41:26.612719+00:00"
link: "https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/"
date: "2026-09-28"
---

# 1614. Maximum Nesting Depth of the Parentheses

Given a **valid parentheses string** `s`, return the **nesting depth** of __`s`. The nesting depth is the **maximum** number of nested parentheses.

 

**Example 1:**

**Input:** s = "(1+(2*3)+((8)/4))+1"

**Output:** 3

**Explanation:**

Digit 8 is inside of 3 nested parentheses in the string.

**Example 2:**

**Input:** s = "(1)+((2))+(((3)))"

**Output:** 3

**Explanation:**

Digit 3 is inside of 3 nested parentheses in the string.

**Example 3:**

**Input:** s = "()(())((()()))"

**Output:** 3

 

**Constraints:**

  * `1 <= s.length <= 100`
  * `s` consists of digits `0-9` and characters `'+'`, `'-'`, `'*'`, `'/'`, `'('`, and `')'`.
  * It is guaranteed that parentheses expression `s` is a VPS.
