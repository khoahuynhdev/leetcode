---
number: "0301"
frontend_id: "301"
title: "Remove Invalid Parentheses"
slug: "remove-invalid-parentheses"
difficulty: "Hard"
topics:
  - "String"
  - "Backtracking"
  - "Breadth-First Search"
acceptance_rate: 5075.9
is_premium: false
created_at: "2026-10-07T06:16:06.453689+00:00"
fetched_at: "2026-10-07T06:16:06.453689+00:00"
link: "https://leetcode.com/problems/remove-invalid-parentheses/"
date: "2026-10-07"
---

# 0301. Remove Invalid Parentheses

Given a string `s` that contains parentheses and letters, remove the minimum number of invalid parentheses to make the input string valid.

Return _a list of**unique strings** that are valid with the minimum number of removals_. You may return the answer in **any order**.

 

**Example 1:**
    
    
    **Input:** s = "()())()"
    **Output:** ["(())()","()()()"]
    

**Example 2:**
    
    
    **Input:** s = "(a)())()"
    **Output:** ["(a())()","(a)()()"]
    

**Example 3:**
    
    
    **Input:** s = ")("
    **Output:** [""]
    

 

**Constraints:**

  * `1 <= s.length <= 25`
  * `s` consists of lowercase English letters and parentheses `'('` and `')'`.
  * There will be at most `20` parentheses in `s`.
