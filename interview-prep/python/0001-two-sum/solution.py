"""
1. Two Sum (Easy)
https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the
two numbers such that they add up to target. You may assume that each input
would have exactly one solution, and you may not use the same element twice.

Pythonic goals:
- Use enumerate() for index + value iteration
- Use dict as a hash map (not defaultdict — simple lookup is enough here)
- One-pass approach
"""



def two_sum(nums: list[int], target: int) -> list[int]:
    dtc = {}
    for i, v in enumerate(nums):
        if (target - v) in dtc:
            return [dtc[target - v], i]
        else:
            dtc[v] = i
    return []

"""
The key changes: enumerate(nums) gives you index i and value v. The dict stores {value: index} so you can look up whether the complement target - v has been seen. The in check happens before any bracket access, so no KeyError.
Also notice the dict maps v -> i (not target - v -> i). Storing the complement as the key works too, but storing the actual value is more readable — you look up "have I seen the number I need?" rather than "have I seen someone who needs me?"
"""
