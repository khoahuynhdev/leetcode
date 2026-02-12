# 3713. Longest Balanced Substring I

[LeetCode Link](https://leetcode.com/problems/longest-balanced-substring-i/)

Difficulty: Medium
Topics: Hash Table, String, Counting, Enumeration
Acceptance Rate: 57.9%

## Hints

### Hint 1

Think about what "balanced" means in terms of character frequencies. What data structure naturally tracks how often each character appears? Consider how you might check the balanced condition efficiently.

### Hint 2

Since the string length is at most 1000, an O(n^2) approach is feasible. Try enumerating substrings by fixing a left endpoint and expanding rightward, maintaining a frequency map as you go. At each step, check whether all frequencies are equal.

### Hint 3

The key insight for checking "balanced" efficiently: you don't need to compare every pair of frequencies. Instead, track the maximum frequency and the number of distinct characters. A substring is balanced if and only if `max_frequency * distinct_count == substring_length`. This lets you check in O(1) per expansion.

## Approach

1. Use two nested loops to enumerate all substrings. The outer loop fixes the left endpoint `i`, and the inner loop expands the right endpoint `j`.
2. Maintain a frequency array of size 26 (one per lowercase letter) that tracks counts for the current window `s[i..j]`.
3. As you extend `j`, increment the count for `s[j]`. Also maintain:
   - `distinct`: the number of characters with a non-zero count.
   - `maxFreq`: the maximum frequency among all characters.
4. A substring is balanced when every distinct character has the same frequency. This is equivalent to checking `maxFreq * distinct == j - i + 1` (the total character count equals the number of distinct characters times the max frequency — which is only possible if all frequencies are equal to `maxFreq`).
5. Whenever the substring is balanced, update the answer with the current length `j - i + 1`.
6. When moving the left pointer (outer loop increments `i`), reset the frequency array, `distinct`, and `maxFreq`.

### Why the O(1) balanced check works

If there are `d` distinct characters and the maximum frequency is `f`, then the total length is at least `d * f` (each distinct character appears at most `f` times). Equality `d * f == length` forces every character to appear exactly `f` times, which is the definition of balanced.

## Complexity Analysis

Time Complexity: O(n^2) — two nested loops over the string, with O(1) work per step.
Space Complexity: O(1) — a fixed-size frequency array of 26 elements.

## Edge Cases

- **Single character string**: `"a"` — the entire string is balanced (length 1).
- **All same characters**: `"aaaa"` — any substring is balanced since there's only one distinct character. Answer is the full length.
- **All distinct characters**: `"abcde"` — the full string is balanced (each character appears once). Answer is the full length.
- **No balanced substring longer than 1**: This cannot happen since every single character is trivially balanced, so the answer is always at least 1.
- **Multiple balanced substrings of the same max length**: Return the length, not the substring itself, so ties don't matter.
