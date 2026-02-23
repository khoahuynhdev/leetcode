import pytest
from solution import is_palindrome


@pytest.mark.parametrize(
    "s, expected",
    [
        # Classic cases from LC
        ("A man, a plan, a canal: Panama", True),
        ("race a car", False),
        (" ", True),
        # Empty string
        ("", True),
        # Single character
        ("a", True),
        # Only non-alphanumeric
        (".,!?", True),
        # Numbers count as alphanumeric
        ("0P", False),
        ("121", True),
        ("12321", True),
        # Mixed case
        ("Aa", True),
        ("Ab", False),
        # Palindrome with numbers and letters
        ("a1b2b1a", True),
        ("a1b2c1a", False),
    ],
)
def test_is_palindrome(s, expected):
    assert is_palindrome(s) == expected
