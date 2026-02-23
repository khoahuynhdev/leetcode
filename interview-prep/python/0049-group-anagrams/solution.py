"""
49. Group Anagrams (Medium)
https://leetcode.com/problems/group-anagrams/

Given an array of strings strs, group the anagrams together. You can return
the answer in any order.

Pythonic goals:
- Use defaultdict(list) for grouping
- Use tuple(sorted(s)) as a hashable key (or Counter-based key)
- Compare with your Go [26]byte key approach — Python equivalent would be
  tuple of counts, but sorted() is more idiomatic here
"""


def group_anagrams(strs: list[str]) -> list[list[str]]:
    pass
