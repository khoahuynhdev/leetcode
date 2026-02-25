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
from collections import defaultdict

def group_anagrams(strs: list[str]) -> list[list[str]]:
    groups = defaultdict(list)
    for s in strs:
        groups[tuple(sorted(s))].append(s)
    return list(groups.values())


if __name__ == "__main__":
    def check(result, expected):
        """Order-independent comparison: sort inner lists and outer list."""
        normalize = lambda groups: sorted(sorted(g) for g in groups)
        assert normalize(result) == normalize(expected), f"got {result}, want {expected}"

    check(group_anagrams(["eat", "tea", "tan", "ate", "nat", "bat"]),
          [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]])
    check(group_anagrams([""]), [[""]])
    check(group_anagrams(["a"]), [["a"]])
    check(group_anagrams(["", ""]), [["", ""]])
    check(group_anagrams(["abc", "bca", "cab", "xyz", "zyx"]),
          [["abc", "bca", "cab"], ["xyz", "zyx"]])
    check(group_anagrams(["a", "b", "c"]), [["a"], ["b"], ["c"]])
    check(group_anagrams(["aaa", "aaa", "aaa"]), [["aaa", "aaa", "aaa"]])
    print("All tests passed!")
