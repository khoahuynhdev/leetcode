"""
1. Two Sum (Easy)
https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the
two numbers such that they add up to target. You may assume that each input
would have exactly one solution, and you may not use the same element twice.

Pythonic goals:
- Use enumerate() for index + value iteration
- Use dict as a hash map (not defaultdict — simple lookup is enough here)
- One-pass approach with walrus operator := for clean complement check
"""


def two_sum(nums: list[int], target: int) -> list[int]:
    seen = {}
    for i, v in enumerate(nums):
        if (complement := target - v) in seen:
            return [seen[complement], i]
        seen[v] = i
    return []


if __name__ == "__main__":
    assert two_sum([2, 7, 11, 15], 9) == [0, 1]
    assert two_sum([3, 2, 4], 6) == [1, 2]
    assert two_sum([3, 3], 6) == [0, 1]
    assert two_sum([-1, -2, -3, -4, -5], -8) == [2, 4]
    assert two_sum([0, 4, 3, 0], 0) == [0, 3]
    assert two_sum([1, 5, 8, 3], 4) == [0, 3]
    assert two_sum([1000000, 500000, -1500000], -1000000) == [1, 2]
    print("All tests passed!")
