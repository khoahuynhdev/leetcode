---
number: "1021"
frontend_id: "1021"
title: "Remove Outermost Parentheses"
slug: "remove-outermost-parentheses"
difficulty: "Easy"
topics:
  - "String"
  - "Stack"
  - "Bracket Sequences"
acceptance_rate: 8780.7
is_premium: false
created_at: "2026-10-08T06:25:46.996809+00:00"
fetched_at: "2026-10-08T06:25:46.996809+00:00"
link: "https://leetcode.com/problems/remove-outermost-parentheses/"
date: "2026-10-08"
---

# 1021. Remove Outermost Parentheses

A valid parentheses string is either empty `""`, `"(" + A + ")"`, or `A + B`, where `A` and `B` are valid parentheses strings, and `+` represents string concatenation.

  * For example, `""`, `"()"`, `"(())()"`, and `"(()(()))"` are all valid parentheses strings.



A valid parentheses string `s` is primitive if it is nonempty, and there does not exist a way to split it into `s = A + B`, with `A` and `B` nonempty valid parentheses strings.

Given a valid parentheses string `s`, consider its primitive decomposition: `s = P1 + P2 + ... + Pk`, where `Pi` are primitive valid parentheses strings.

Return `s` _after removing the outermost parentheses of every primitive string in the primitive decomposition of_`s`.

 

**Example 1:**
    
    
    **Input:** s = "(()())(())"
    **Output:** "()()()"
    **Explanation:** 
    The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
    After removing outer parentheses of each part, this is "()()" + "()" = "()()()".
    

**Example 2:**
    
    
    **Input:** s = "(()())(())(()(()))"
    **Output:** "()()()()(())"
    **Explanation:** 
    The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
    After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".
    

**Example 3:**
    
    
    **Input:** s = "()()"
    **Output:** ""
    **Explanation:** 
    The input string is "()()", with primitive decomposition "()" + "()".
    After removing outer parentheses of each part, this is "" + "" = "".
    

 

**Constraints:**

  * `1 <= s.length <= 105`
  * `s[i]` is either `'('` or `')'`.
  * `s` is a valid parentheses string.
