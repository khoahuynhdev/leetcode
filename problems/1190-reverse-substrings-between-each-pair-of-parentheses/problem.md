---
number: "1190"
frontend_id: "1190"
title: "Reverse Substrings Between Each Pair of Parentheses"
slug: "reverse-substrings-between-each-pair-of-parentheses"
difficulty: "Medium"
topics:
  - "String"
  - "Stack"
  - "Bracket Sequences"
acceptance_rate: 7290.8
is_premium: false
created_at: "2026-09-27T05:31:01.909324+00:00"
fetched_at: "2026-09-27T05:31:01.909324+00:00"
link: "https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/"
date: "2026-09-27"
---

# 1190. Reverse Substrings Between Each Pair of Parentheses

You are given a string `s` that consists of lower case English letters and brackets.

Reverse the strings in each pair of matching parentheses, starting from the innermost one.

Your result should **not** contain any brackets.

 

**Example 1:**
    
    
    **Input:** s = "(abcd)"
    **Output:** "dcba"
    

**Example 2:**
    
    
    **Input:** s = "(u(love)i)"
    **Output:** "iloveu"
    **Explanation:** The substring "love" is reversed first, then the whole string is reversed.
    

**Example 3:**
    
    
    **Input:** s = "(ed(et(oc))el)"
    **Output:** "leetcode"
    **Explanation:** First, we reverse the substring "oc", then "etco", and finally, the whole string.
    

 

**Constraints:**

  * `1 <= s.length <= 2000`
  * `s` only contains lower case English characters and parentheses.
  * It is guaranteed that all parentheses are balanced.
