import pytest
from solution import contains_duplicate


@pytest.mark.parametrize(
    "nums, expected",
    [
        # Has duplicates
        ([1, 2, 3, 1], True),
        ([1, 1, 1, 3, 3, 4, 3, 2, 4, 2], True),
        # No duplicates
        ([1, 2, 3, 4], False),
        # Empty and single element
        ([], False),
        ([1], False),
        # Two elements
        ([1, 1], True),
        ([1, 2], False),
        # Negative numbers
        ([-1, -2, -3, -1], True),
        ([-1, -2, -3], False),
        # Large values
        ([1000000000, -1000000000, 1000000000], True),
    ],
)
def test_contains_duplicate(nums, expected):
    assert contains_duplicate(nums) == expected
