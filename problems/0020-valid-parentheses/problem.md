---
number: "0020"
frontend_id: "20"
title: "Valid Parentheses"
slug: "valid-parentheses"
difficulty: "Easy"
topics:
  - "String"
  - "Stack"
  - "Bracket Sequences"
acceptance_rate: 4505.6
is_premium: false
created_at: "2026-10-01T06:18:26.535875+00:00"
fetched_at: "2026-10-01T06:18:26.535875+00:00"
link: "https://leetcode.com/problems/valid-parentheses/"
date: "2026-10-01"
---

# 0020. Valid Parentheses

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

  1. Open brackets must be closed by the same type of brackets.
  2. Open brackets must be closed in the correct order.
  3. Every close bracket has a corresponding open bracket of the same type.



 

**Example 1:**

**Input:** s = "()"

**Output:** true

**Example 2:**

**Input:** s = "()[]{}"

**Output:** true

**Example 3:**

**Input:** s = "(]"

**Output:** false

**Example 4:**

**Input:** s = "([])"

**Output:** true

**Example 5:**

**Input:** s = "([)]"

**Output:** false

 

**Constraints:**

  * `1 <= s.length <= 104`
  * `s` consists of parentheses only `'()[]{}'`.
