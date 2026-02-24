"""
242. Valid Anagram (Easy)
https://leetcode.com/problems/valid-anagram/

Given two strings s and t, return true if t is an anagram of s, and false
otherwise. An anagram uses all the original letters exactly once.

Pythonic goals:
- Use collections.Counter for one-liner comparison
- Also try the sorted() approach as a contrast
- Compare brevity vs your Go map-based solution
"""


def is_anagram(s: str, t: str) -> bool:
    pass


if __name__ == "__main__":
    assert is_anagram("anagram", "nagaram") == True
    assert is_anagram("rat", "car") == False
    assert is_anagram("", "") == True
    assert is_anagram("a", "a") == True
    assert is_anagram("a", "b") == False
    assert is_anagram("ab", "a") == False
    assert is_anagram("aacc", "ccac") == False
    assert is_anagram("listen", "silent") == True
    assert is_anagram("hello", "llohe") == True
    assert is_anagram("aabbcc", "abcabc") == True
    print("All tests passed!")
