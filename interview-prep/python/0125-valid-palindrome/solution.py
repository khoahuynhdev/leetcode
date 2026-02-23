"""
125. Valid Palindrome (Easy)
https://leetcode.com/problems/valid-palindrome/

A phrase is a palindrome if, after converting all uppercase letters into
lowercase and removing all non-alphanumeric characters, it reads the same
forward and backward.

Given a string s, return true if it is a palindrome, or false otherwise.

Pythonic goals:
- Use str.isalnum() and str.lower() for filtering
- Try the two-pointer approach (more interview-like)
- Also try the one-liner with slicing: compare filtered == filtered[::-1]
- Know the tradeoff: slicing is O(n) space, two-pointer is O(1) space
"""


def is_palindrome(s: str) -> bool:
    pass
