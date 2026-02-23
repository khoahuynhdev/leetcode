import pytest
from solution import group_anagrams


def normalize(groups: list[list[str]]) -> list[tuple[str, ...]]:
    """Sort each group and sort the list of groups for order-independent comparison."""
    return sorted(tuple(sorted(g)) for g in groups)


@pytest.mark.parametrize(
    "strs, expected",
    [
        # Standard case
        (
            ["eat", "tea", "tan", "ate", "nat", "bat"],
            [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]],
        ),
        # Single empty string
        ([""], [[""]]),
        # Single element
        (["a"], [["a"]]),
        # All same anagram
        (["abc", "bca", "cab"], [["abc", "bca", "cab"]]),
        # No anagrams (all unique)
        (["abc", "def", "ghi"], [["abc"], ["def"], ["ghi"]]),
        # Mixed lengths
        (["a", "ab", "ba", "abc"], [["a"], ["ab", "ba"], ["abc"]]),
        # Empty input
        ([], []),
    ],
)
def test_group_anagrams(strs, expected):
    assert normalize(group_anagrams(strs)) == normalize(expected)
