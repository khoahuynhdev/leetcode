---
number: "1510"
frontend_id: "1510"
title: "Stone Game IV"
slug: "stone-game-iv"
difficulty: "Hard"
topics:
  - "Math"
  - "Dynamic Programming"
  - "Minimax"
  - "Game Theory"
  - "Nim Game"
  - "Sprague–Grundy Theorem"
  - "Zero-Sum Game"
acceptance_rate: 6052.4
is_premium: false
created_at: "2026-08-10T02:40:01.992585+00:00"
fetched_at: "2026-08-10T02:40:01.992585+00:00"
link: "https://leetcode.com/problems/stone-game-iv/"
date: "2026-08-10"
---

# 1510. Stone Game IV

Alice and Bob take turns playing a game, with Alice starting first.

Initially, there are `n` stones in a pile. On each player's turn, that player makes a _move_ consisting of removing **any** non-zero **square number** of stones in the pile.

Also, if a player cannot make a move, he/she loses the game.

Given a positive integer `n`, return `true` if and only if Alice wins the game otherwise return `false`, assuming both players play optimally.

 

**Example 1:**
    
    
    **Input:** n = 1
    **Output:** true
    **Explanation:** Alice can remove 1 stone winning the game because Bob doesn't have any moves.

**Example 2:**
    
    
    **Input:** n = 2
    **Output:** false
    **Explanation:** Alice can only remove 1 stone, after that Bob removes the last one winning the game (2 -> 1 -> 0).
    

**Example 3:**
    
    
    **Input:** n = 4
    **Output:** true
    **Explanation:** n is already a perfect square, Alice can win with one move, removing 4 stones (4 -> 0).
    

 

**Constraints:**

  * `1 <= n <= 105`
