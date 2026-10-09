---
number: "1541"
frontend_id: "1541"
title: "Minimum Insertions to Balance a Parentheses String"
slug: "minimum-insertions-to-balance-a-parentheses-string"
difficulty: "Medium"
topics:
  - "String"
  - "Stack"
  - "Greedy"
  - "Bracket Sequences"
acceptance_rate: 5659.4
is_premium: false
created_at: "2026-10-09T06:27:05.476213+00:00"
fetched_at: "2026-10-09T06:27:05.476213+00:00"
link: "https://leetcode.com/problems/minimum-insertions-to-balance-a-parentheses-string/"
date: "2026-10-09"
---

# 1541. Minimum Insertions to Balance a Parentheses String

Given a parentheses string `s` containing only the characters `'('` and `')'`. A parentheses string is **balanced** if:

  * Any left parenthesis `'('` must have a corresponding two consecutive right parenthesis `'))'`.
  * Left parenthesis `'('` must go before the corresponding two consecutive right parenthesis `'))'`.



In other words, we treat `'('` as an opening parenthesis and `'))'` as a closing parenthesis.

  * For example, `"())"`, `"())(())))"` and `"(())())))"` are balanced, `")()"`, `"()))"` and `"(()))"` are not balanced.



You can insert the characters `'('` and `')'` at any position of the string to balance it if needed.

Return _the minimum number of insertions_ needed to make `s` balanced.

 

**Example 1:**
    
    
    **Input:** s = "(()))"
    **Output:** 1
    **Explanation:** The second '(' has two matching '))', but the first '(' has only ')' matching. We need to add one more ')' at the end of the string to be "(())))" which is balanced.
    

**Example 2:**
    
    
    **Input:** s = "())"
    **Output:** 0
    **Explanation:** The string is already balanced.
    

**Example 3:**
    
    
    **Input:** s = "))())("
    **Output:** 3
    **Explanation:** Add '(' to match the first '))', Add '))' to match the last '('.
    

 

**Constraints:**

  * `1 <= s.length <= 105`
  * `s` consists of `'('` and `')'` only.
