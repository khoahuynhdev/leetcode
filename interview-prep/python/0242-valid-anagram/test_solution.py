import pytest
from solution import is_anagram


@pytest.mark.parametrize(
    "s, t, expected",
    [
        # Basic anagrams
        ("anagram", "nagaram", True),
        ("rat", "car", False),
        # Empty strings
        ("", "", True),
        # Single character
        ("a", "a", True),
        ("a", "b", False),
        # Different lengths
        ("ab", "a", False),
        ("a", "ab", False),
        # Repeated characters
        ("aacc", "ccac", False),
        ("aabb", "bbaa", True),
        # Unicode (follow-up question on LC)
        ("café", "éfac", True),
    ],
)
def test_is_anagram(s, t, expected):
    assert is_anagram(s, t) == expected
