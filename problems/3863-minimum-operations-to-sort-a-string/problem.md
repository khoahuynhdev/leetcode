---
number: "3863"
frontend_id: "3863"
title: "Minimum Operations to Sort a String"
slug: "minimum-operations-to-sort-a-string"
difficulty: "Medium"
topics:
  - "String"
acceptance_rate: 2683.4
is_premium: false
created_at: "2026-10-21T05:06:03.693403+00:00"
fetched_at: "2026-10-21T05:06:03.693403+00:00"
link: "https://leetcode.com/problems/minimum-operations-to-sort-a-string/"
date: "2026-10-21"
---

# 3863. Minimum Operations to Sort a String

You are given a string `s` consisting of lowercase English letters.

In one operation, you can select any **substring** of `s` that is **not** the entire string and **sort** it in **non-descending alphabetical** order.

Return the **minimum** number of operations required to make `s` sorted in **non-descending** order. If it is not possible, return -1.

 

**Example 1:**

**Input:** s = "dog"

**Output:** 1

**Explanation:** ​​​​​​​

  * Sort substring `"og"` to `"go"`.
  * Now, `s = "dgo"`, which is sorted in ascending order. Thus, the answer is 1.



**Example 2:**

**Input:** s = "card"

**Output:** 2

**Explanation:**

  * Sort substring `"car"` to `"acr"`, so `s = "acrd"`.
  * Sort substring `"rd"` to `"dr"`, making `s = "acdr"`, which is sorted in ascending order. Thus, the answer is 2.



**Example 3:**

**Input:** s = "gf"

**Output:** -1

**Explanation:**

  * It is impossible to sort `s` under the given constraints. Thus, the answer is -1.



 

**Constraints:**

  * `1 <= s.length <= 105`
  * `s` consists of only lowercase English letters.
