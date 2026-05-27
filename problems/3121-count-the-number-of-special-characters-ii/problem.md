---
number: "3121"
frontend_id: "3121"
title: "Count the Number of Special Characters II"
slug: "count-the-number-of-special-characters-ii"
difficulty: "Medium"
topics:
  - "Hash Table"
  - "String"
acceptance_rate: 5097.8
is_premium: false
created_at: "2026-05-27T04:54:00.724668+00:00"
fetched_at: "2026-05-27T04:54:00.724668+00:00"
link: "https://leetcode.com/problems/count-the-number-of-special-characters-ii/"
date: "2026-05-27"
---

# 3121. Count the Number of Special Characters II

You are given a string `word`. A letter `c` is called **special** if it appears **both** in lowercase and uppercase in `word`, and **every** lowercase occurrence of `c` appears before the **first** uppercase occurrence of `c`.

Return the number of __**special** letters __ in __`word`.

 

**Example 1:**

**Input:** word = "aaAbcBC"

**Output:** 3

**Explanation:**

The special characters are `'a'`, `'b'`, and `'c'`.

**Example 2:**

**Input:** word = "abc"

**Output:** 0

**Explanation:**

There are no special characters in `word`.

**Example 3:**

**Input:** word = "AbBCab"

**Output:** 0

**Explanation:**

There are no special characters in `word`.

 

**Constraints:**

  * `1 <= word.length <= 2 * 105`
  * `word` consists of only lowercase and uppercase English letters.
