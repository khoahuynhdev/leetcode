"""
217. Contains Duplicate (Easy)
https://leetcode.com/problems/contains-duplicate/

Given an integer array nums, return true if any value appears at least twice
in the array, and return false if every element is distinct.

Pythonic goals:
- Use set() for O(1) lookup — compare len(set(nums)) vs len(nums)
- This is a one-liner in Python; practice writing it that way
- Also consider the early-exit approach with a growing set
"""


def contains_duplicate(nums: list[int]) -> bool:
    pass


if __name__ == "__main__":
    assert contains_duplicate([1, 2, 3, 1]) == True
    assert contains_duplicate([1, 2, 3, 4]) == False
    assert contains_duplicate([1, 1, 1, 3, 3, 4, 3, 2, 4, 2]) == True
    assert contains_duplicate([]) == False
    assert contains_duplicate([1]) == False
    assert contains_duplicate([1, 1]) == True
    assert contains_duplicate([-1, -2, -3, -1]) == True
    assert contains_duplicate([-1, -2, -3]) == False
    assert contains_duplicate([10**9, -(10**9), 10**9]) == True
    assert contains_duplicate(list(range(10000))) == False
    print("All tests passed!")
