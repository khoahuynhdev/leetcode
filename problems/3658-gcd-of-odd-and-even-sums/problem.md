---
number: "3658"
frontend_id: "3658"
title: "GCD of Odd and Even Sums"
slug: "gcd-of-odd-and-even-sums"
difficulty: "Easy"
topics:
  - "Math"
  - "Number Theory"
acceptance_rate: 8639.4
is_premium: false
created_at: "2026-07-15T03:43:29.834554+00:00"
fetched_at: "2026-07-15T03:43:29.834554+00:00"
link: "https://leetcode.com/problems/gcd-of-odd-and-even-sums/"
date: "2026-07-15"
---

# 3658. GCD of Odd and Even Sums

You are given an integer `n`. Your task is to compute the **GCD** (greatest common divisor) of two values:

  * `sumOdd`: the sum of the smallest `n` positive odd numbers.

  * `sumEven`: the sum of the smallest `n` positive even numbers.




Return the GCD of `sumOdd` and `sumEven`.

 

**Example 1:**

**Input:** n = 4

**Output:** 4

**Explanation:**

  * Sum of the first 4 odd numbers `sumOdd = 1 + 3 + 5 + 7 = 16`
  * Sum of the first 4 even numbers `sumEven = 2 + 4 + 6 + 8 = 20`



Hence, `GCD(sumOdd, sumEven) = GCD(16, 20) = 4`.

**Example 2:**

**Input:** n = 5

**Output:** 5

**Explanation:**

  * Sum of the first 5 odd numbers `sumOdd = 1 + 3 + 5 + 7 + 9 = 25`
  * Sum of the first 5 even numbers `sumEven = 2 + 4 + 6 + 8 + 10 = 30`



Hence, `GCD(sumOdd, sumEven) = GCD(25, 30) = 5`.

 

**Constraints:**

  * `1 <= n <= 10​​​​​​​00`
