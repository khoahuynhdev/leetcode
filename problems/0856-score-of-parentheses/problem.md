---
number: "0856"
frontend_id: "856"
title: "Score of Parentheses"
slug: "score-of-parentheses"
difficulty: "Medium"
topics:
  - "String"
  - "Stack"
  - "Bracket Sequences"
acceptance_rate: 6394.2
is_premium: false
created_at: "2026-10-05T06:00:39.850199+00:00"
fetched_at: "2026-10-05T06:00:39.850199+00:00"
link: "https://leetcode.com/problems/score-of-parentheses/"
date: "2026-10-05"
---

# 0856. Score of Parentheses

Given a balanced parentheses string `s`, return _the**score** of the string_.

The **score** of a balanced parentheses string is based on the following rule:

  * `"()"` has score `1`.
  * `AB` has score `A + B`, where `A` and `B` are balanced parentheses strings.
  * `(A)` has score `2 * A`, where `A` is a balanced parentheses string.



 

**Example 1:**
    
    
    **Input:** s = "()"
    **Output:** 1
    

**Example 2:**
    
    
    **Input:** s = "(())"
    **Output:** 2
    

**Example 3:**
    
    
    **Input:** s = "()()"
    **Output:** 2
    

 

**Constraints:**

  * `2 <= s.length <= 50`
  * `s` consists of only `'('` and `')'`.
  * `s` is a balanced parentheses string.
