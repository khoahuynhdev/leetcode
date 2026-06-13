# 3838. Weighted Word Mapping

[LeetCode Link](https://leetcode.com/problems/weighted-word-mapping/)

Difficulty: Easy
Topics: Array, String, Simulation
Acceptance Rate: 88.1%

## Hints

### Hint 1

This is a pure simulation problem — there is no clever trick or hidden data
structure. Read the rules carefully and translate each rule directly into code.
Think about what information you need for a single word before worrying about
the whole array.

### Hint 2

For one word you need its *weight*: the sum of `weights[i]` for each character,
where `i` is the position of the character in the alphabet. Recall that for a
lowercase byte `c`, the index is `c - 'a'`. Once you have the weight, the
problem tells you to reduce it with `weight % 26`.

### Hint 3

The only part that trips people up is the *reverse alphabetical* mapping:
`0 -> 'z'`, `1 -> 'y'`, ..., `25 -> 'a'`. Notice that the index and the letter
move in opposite directions, and they always sum to 25 (`'z'` is index 25 from
`'a'`). So the mapped letter is simply `'z' - r`, or equivalently
`'a' + (25 - r)`. Build the answer one byte at a time and you are done.

## Approach

We process each word independently and build the result string character by
character.

1. **Compute the weight of a word.** Walk over its characters. For each
   character `c`, its alphabet index is `c - 'a'`, so add `weights[c-'a']` to a
   running sum. The constraints (word length ≤ 10, weight ≤ 100) mean this sum
   easily fits in an `int`.

2. **Reduce modulo 26.** Let `r = weight % 26`. This gives a value in the range
   `[0, 25]`.

3. **Map to a letter in reverse order.** The rule is `0 -> 'z'`, `1 -> 'y'`,
   ..., `25 -> 'a'`. Because index `r` and the target letter sum to the value of
   `'z'` (index 25 above `'a'`), the mapped byte is `'z' - r`. For example
   `r = 8` gives `'z' - 8 = 'r'`, matching the first example.

4. **Concatenate.** Append each mapped byte to a `[]byte` (or
   `strings.Builder`) in word order, then convert to a string at the end.

Using a byte slice that we append to and convert once avoids repeated string
allocations, keeping the work proportional to the total number of characters.

Worked example for `words = ["abcd"]` with the third example's weights
`[7,5,3,4,...]`: weight = `7 + 5 + 3 + 4 = 19`, `19 % 26 = 19`, and
`'z' - 19 = 'g'`, so the answer is `"g"`. ✔

## Complexity Analysis

Time Complexity: O(N · L) where N is the number of words and L is the maximum
word length — we touch every character of every word exactly once. (Given the
constraints this is at most 100 · 10 = 1000 operations.)
Space Complexity: O(N) for the output string of one character per word, ignoring
the input. The per-word computation uses O(1) extra space.

## Edge Cases

- **Single-character words.** Words like `"a"` or `"c"` are valid; the weight is
  just `weights[c-'a']`. Example 2 covers this.
- **Single word in the array.** With `words.length == 1` the output is a single
  character (Example 3). Make sure the loop and the builder handle a length-one
  result.
- **Weight divisible by 26 (`r == 0`).** This maps to `'z'`, the start of the
  reverse alphabet. It is easy to accidentally produce a wrong byte here, so it
  is worth a dedicated test.
- **Maximum-weight words.** The largest single word weight is 10 · 100 = 1000,
  which is well within `int` range — no overflow concern, but good to confirm
  the modulo arithmetic is correct for large sums.
