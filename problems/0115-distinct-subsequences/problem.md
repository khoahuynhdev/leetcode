---
number: "0115"
frontend_id: "115"
title: "Distinct Subsequences"
slug: "distinct-subsequences"
difficulty: "Hard"
topics:
  - "String"
  - "Dynamic Programming"
acceptance_rate: 5282.2
is_premium: false
created_at: "2026-09-06T04:55:45.586738+00:00"
fetched_at: "2026-09-06T04:55:45.586738+00:00"
link: "https://leetcode.com/problems/distinct-subsequences/"
date: "2026-09-06"
---

# 0115. Distinct Subsequences

Given two strings s and t, return _the number of distinct_ **_subsequences_** _of_ s _which equals_ t.

The test cases are generated so that the answer fits on a 32-bit signed integer.

 

**Example 1:**
    
    
    **Input:** s = "rabbbit", t = "rabbit"
    **Output:** 3
    **Explanation:**
    As shown below, there are 3 ways you can generate "rabbit" from s.
    **_rabb_** b** _it_**
    **_ra_** b** _bbit_**
    **_rab_** b** _bit_**
    

**Example 2:**
    
    
    **Input:** s = "babgbag", t = "bag"
    **Output:** 5
    **Explanation:**
    As shown below, there are 5 ways you can generate "bag" from s.
    **_ba_** b _**g**_ bag
    **_ba_** bgba** _g_**
    _**b**_ abgb** _ag_**
    ba _**b**_ gb _**ag**_
    babg** _bag_**

 

**Constraints:**

  * `1 <= s.length, t.length <= 1000`
  * `s` and `t` consist of English letters.
