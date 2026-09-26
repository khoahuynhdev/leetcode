# 1807. Evaluate the Bracket Pairs of a String

[LeetCode Link](https://leetcode.com/problems/evaluate-the-bracket-pairs-of-a-string/)

Difficulty: Medium
Topics: Array, Hash Table, String
Acceptance Rate: 74.3%

## Hints

### Hint 1

You will be asked "what is the value of this key?" many times. Which data structure answers lookup questions in constant time?

### Hint 2

Convert `knowledge` into a hash map first. Then think about walking through `s` exactly once, character by character, building the answer as you go. Brackets are never nested, so you only need to remember whether you are currently "inside" a bracket pair.

### Hint 3

When you see `'('`, record the position right after it. When you see the matching `')'`, the substring between those two positions is the key — look it up in the map and append either the value or `"?"`. Any character outside brackets is copied as-is. Use a string builder so the output is built in linear time instead of repeatedly concatenating strings.

## Approach

1. **Build a lookup table.** Iterate over `knowledge` and store `key -> value` in a hash map. This makes each later lookup O(1) on average (keys are at most 10 characters).
2. **Single left-to-right scan of `s`** with an index `i`:
   - If `s[i] == '('`, find the next `')'` at index `j`. The key is `s[i+1:j]`. Append `m[key]` if it exists, otherwise append `"?"`. Jump `i` to `j + 1`.
   - Otherwise, append `s[i]` to the result and advance `i`.
3. Return the built string.

Because brackets are guaranteed not to nest and every `'('` has a matching `')'`, the first `')'` after a `'('` always closes that pair, so no stack is needed.

**Example:** `s = "(name)is(age)yearsold"`, map = `{name: bob, age: two}`

- `i=0` is `'('` → key `"name"` → append `"bob"`, jump to `i=6`.
- `"is"` copied as-is.
- `'('` at `i=8` → key `"age"` → append `"two"`, jump past `')'`.
- `"yearsold"` copied.
- Result: `"bobistwoyearsold"`.

Each character of `s` is visited exactly once (either copied or read as part of a key), so the scan is linear.

## Complexity Analysis

Time Complexity: O(n + k), where n = len(s) and k = total length of all strings in `knowledge` (building the map + one scan of `s`).
Space Complexity: O(k + n) — the hash map plus the output string.

## Edge Cases

- **Unknown key:** must be replaced with `"?"`, not left as-is or removed entirely.
- **Empty `knowledge`:** every bracket pair becomes `"?"`; make sure a nil/empty map doesn't break lookups.
- **No brackets at all:** the string should be returned unchanged.
- **Same key repeated / adjacent pairs like `"(a)(a)"`:** after handling a pair, resume scanning right after `')'` so the next `'('` is processed correctly.
- **Letters outside brackets that match a key:** e.g. `"aaa"` with key `"a"` — only bracketed text is evaluated.
- **Bracket pair at the very start or end of `s`:** index handling must not go out of bounds.
