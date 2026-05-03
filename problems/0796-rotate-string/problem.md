---
number: "0796"
frontend_id: "796"
title: "Rotate String"
slug: "rotate-string"
difficulty: "Easy"
topics:
  - "String"
  - "String Matching"
acceptance_rate: 6560.0
is_premium: false
created_at: "2026-05-03T04:26:36.630548+00:00"
fetched_at: "2026-05-03T04:26:36.630548+00:00"
link: "https://leetcode.com/problems/rotate-string/"
date: "2026-05-03"
---

# 0796. Rotate String

Given two strings `s` and `goal`, return `true` _if and only if_ `s` _can become_ `goal` _after some number of**shifts** on_ `s`.

A **shift** on `s` consists of moving the leftmost character of `s` to the rightmost position.

  * For example, if `s = "abcde"`, then it will be `"bcdea"` after one shift.



 

**Example 1:**
    
    
    **Input:** s = "abcde", goal = "cdeab"
    **Output:** true
    

**Example 2:**
    
    
    **Input:** s = "abcde", goal = "abced"
    **Output:** false
    

 

**Constraints:**

  * `1 <= s.length, goal.length <= 100`
  * `s` and `goal` consist of lowercase English letters.
