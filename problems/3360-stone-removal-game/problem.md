---
number: "3360"
frontend_id: "3360"
title: "Stone Removal Game"
slug: "stone-removal-game"
difficulty: "Easy"
topics:
  - "Math"
  - "Simulation"
acceptance_rate: 5096.4
is_premium: false
created_at: "2026-11-26T08:43:12.805958+00:00"
fetched_at: "2026-11-26T08:43:12.805958+00:00"
link: "https://leetcode.com/problems/stone-removal-game/"
date: "2026-11-26"
---

# 3360. Stone Removal Game

Alice and Bob are playing a game where they take turns removing stones from a pile, with _Alice going first_.

  * Alice starts by removing **exactly** 10 stones on her first turn.
  * For each subsequent turn, each player removes **exactly** 1 fewer**** stone**** than the previous opponent.



The player who cannot make a move loses the game.

Given a positive integer `n`, return `true` if Alice wins the game and `false` otherwise.

 

**Example 1:**

**Input:** n = 12

**Output:** true

**Explanation:**

  * Alice removes 10 stones on her first turn, leaving 2 stones for Bob.
  * Bob cannot remove 9 stones, so Alice wins.



**Example 2:**

**Input:** n = 1

**Output:** false

**Explanation:**

  * Alice cannot remove 10 stones, so Alice loses.



 

**Constraints:**

  * `1 <= n <= 50`
