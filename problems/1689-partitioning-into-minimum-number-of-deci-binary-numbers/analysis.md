# 1689. Partitioning Into Minimum Number Of Deci-Binary Numbers

[LeetCode Link](https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/)

Difficulty: Medium
Topics: String, Greedy
Acceptance Rate: 89.0%

## Hints

### Hint 1

Think about what a deci-binary number can contribute to each digit position. Since each digit in a deci-binary number is either 0 or 1, what does that tell you about how many deci-binary numbers you need to "build up" a particular digit?

### Hint 2

Consider each digit of `n` independently. If a digit is `d`, you need at least `d` deci-binary numbers that have a `1` in that position. Can you think about the problem as: what is the bottleneck digit that requires the most deci-binary numbers?

### Hint 3

The answer is simply the maximum digit in `n`. Each deci-binary number contributes at most 1 to each digit position, so the digit with the largest value dictates the minimum count. And you can always construct exactly that many deci-binary numbers to cover all positions — just "peel off" one layer at a time by placing a 1 wherever the remaining digit is still positive.

## Approach

The key insight is that each deci-binary number has digits that are either 0 or 1. When we sum multiple deci-binary numbers, each digit position accumulates independently (no carries needed if we think of it digit-by-digit).

For a digit `d` at some position in `n`, we need exactly `d` deci-binary numbers to have a `1` in that position. The remaining deci-binary numbers will have a `0` there.

Since every digit position must be satisfied, the total number of deci-binary numbers we need is determined by the position that demands the most — i.e., the maximum digit in `n`.

**Example walkthrough with n = "32":**
- Digit `3` requires 3 deci-binary numbers with a `1` in the tens place.
- Digit `2` requires 2 deci-binary numbers with a `1` in the ones place.
- Maximum digit is `3`, so we need 3 deci-binary numbers.
- Construction: `10 + 11 + 11 = 32` — the first number has a `1` only in the tens place, while the second and third have `1`s in both places.

**Algorithm:**
1. Iterate through each character in the string `n`.
2. Track the maximum digit value seen.
3. Return the maximum digit value.

## Complexity Analysis

Time Complexity: O(n) — single pass through the string where n is the number of digits.
Space Complexity: O(1) — only a single variable to track the maximum.

## Edge Cases

- **Single digit (e.g., "1"):** The answer is that digit itself. The simplest case.
- **All same digits (e.g., "111"):** The answer equals that repeated digit. Every deci-binary number used contributes a `1` everywhere.
- **Contains a 9:** The answer is 9. This is the theoretical maximum since digits range from 0–9.
- **Very long string:** The algorithm handles strings up to 10^5 characters efficiently in O(n) time.
- **String like "10":** The max digit is 1, so only one deci-binary number is needed — the number itself.
