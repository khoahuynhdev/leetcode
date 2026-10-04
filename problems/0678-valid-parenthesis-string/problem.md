---
number: "0678"
frontend_id: "678"
title: "Valid Parenthesis String"
slug: "valid-parenthesis-string"
difficulty: "Medium"
topics:
  - "String"
  - "Dynamic Programming"
  - "Stack"
  - "Greedy"
  - "Bracket Sequences"
acceptance_rate: 4113.3
is_premium: false
created_at: "2026-10-04T06:09:07.468526+00:00"
fetched_at: "2026-10-04T06:09:07.468526+00:00"
link: "https://leetcode.com/problems/valid-parenthesis-string/"
date: "2026-10-04"
---

# 0678. Valid Parenthesis String

Given a string `s` containing only three types of characters: `'('`, `')'` and `'*'`, return `true` _if_ `s` _is**valid**_.

The following rules define a **valid** string:

  * Any left parenthesis `'('` must have a corresponding right parenthesis `')'`.
  * Any right parenthesis `')'` must have a corresponding left parenthesis `'('`.
  * Left parenthesis `'('` must go before the corresponding right parenthesis `')'`.
  * `'*'` could be treated as a single right parenthesis `')'` or a single left parenthesis `'('` or an empty string `""`.



 

**Example 1:**
    
    
    **Input:** s = "()"
    **Output:** true
    

**Example 2:**
    
    
    **Input:** s = "(*)"
    **Output:** true
    

**Example 3:**
    
    
    **Input:** s = "(*))"
    **Output:** true
    

**Example 4:**
    
    
    **Input:** s = "("
    **Output:** false
    

 

**Constraints:**

  * `1 <= s.length <= 100`
  * `s[i]` is `'('`, `')'` or `'*'`.
