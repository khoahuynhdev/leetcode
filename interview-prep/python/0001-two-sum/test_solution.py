import pytest
from solution import two_sum


@pytest.mark.parametrize(
    "nums, target, expected",
    [
        # Basic cases
        ([2, 7, 11, 15], 9, [0, 1]),
        ([3, 2, 4], 6, [1, 2]),
        ([3, 3], 6, [0, 1]),
        # Negative numbers
        ([-1, -2, -3, -4, -5], -8, [2, 4]),
        ([1, -1, 0, 2], 1, [2, 3]),
        # Larger input
        ([1, 5, 3, 7, 8, 2, 4], 9, [0, 4]),
        # Target is zero
        ([-3, 4, 3, 90], 0, [0, 2]),
    ],
)
def test_two_sum(nums, target, expected):
    result = two_sum(nums, target)
    assert sorted(result) == sorted(expected)
