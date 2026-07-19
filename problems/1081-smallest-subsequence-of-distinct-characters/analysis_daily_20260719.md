# 1081. Smallest Subsequence of Distinct Characters

[LeetCode Link](https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/)

Difficulty: Medium
Topics: String, Stack, Greedy, Monotonic Stack
Acceptance Rate: 64.9%

## Hints

### Hint 1

You need to keep exactly one copy of each distinct character while making the result as small as possible in dictionary order. Think about building the answer one character at a time and being able to "undo" a choice you made earlier if a better opportunity comes along. Which data structure lets you look at and remove your most recent decision cheaply?

### Hint 2

A stack is the natural fit. As you scan left to right, you're deciding whether the character you're holding should stay before or after the current one. To make the answer smaller, you'd like a larger earlier character to come *after* a smaller one — so consider popping a character off the stack when the current character is smaller. But you can only safely discard a character if you'll see it again later in the string.

### Hint 3

Precompute the **last occurrence index** of every character. Maintain a stack that stays as close to increasing as possible. For each character `c` at index `i`:
- Skip it entirely if it's already in the stack (each distinct char appears once).
- Otherwise, while the stack's top is greater than `c` **and** that top character appears again later (`lastIndex[top] > i`), pop it — you can afford to place it later where it helps.
- Push `c`.

The "appears again later" guard is the critical piece: it guarantees you never remove a character you can't get back, so the final stack still contains every distinct character exactly once.

## Approach

This is the classic monotonic-stack greedy problem (identical to LeetCode 316, *Remove Duplicate Letters*).

We want the lexicographically smallest subsequence containing each distinct character exactly once. The greedy intuition: whenever we can move a larger character to a later position without losing it, we should, because putting a smaller character earlier makes the whole string smaller in dictionary order.

**Steps:**

1. **Record last positions.** Compute `last[ch]` = the highest index at which `ch` occurs. This tells us, at any point `i`, whether a character will reappear after `i`.
2. **Scan with a stack.** Keep a `stack` (the answer being built) and an `inStack` set marking which characters are currently on the stack.
3. For each character `c` at index `i`:
   - If `c` is already in the stack, it's already placed optimally — skip it.
   - Otherwise, while the top of the stack is lexicographically greater than `c` and the top character still occurs later in the string (`last[top] > i`), pop it and unmark it. This shoves that larger character rightward to a place where it can only help.
   - Push `c` and mark it in `inStack`.
4. The stack, read bottom to top, is the answer.

**Walkthrough on `"cbacdcbc"`:**
- last = {c:7, b:6, a:2, d:4}
- `c`(0): stack `[c]`
- `b`(1): b < c and last[c]=7>1 → pop c; stack `[b]`
- `a`(2): a < b and last[b]=6>2 → pop b; stack `[a]`
- `c`(3): c > a, push → `[a, c]`
- `d`(4): d > c, push → `[a, c, d]`
- `c`(5): already in stack → skip
- `b`(6): top d > b but last[d]=4 < 6 → can't pop d; push b → `[a, c, d, b]`
- `c`(7): already in stack → skip

Result: `"acdb"`. ✔

## Complexity Analysis

Time Complexity: O(n) — each character is pushed and popped at most once, and the `last`/`inStack` lookups are O(1) over a fixed 26-letter alphabet.
Space Complexity: O(1) auxiliary (at most 26 characters on the stack and in the tracking arrays), or O(k) where k is the number of distinct characters, which is bounded by the alphabet size.

## Edge Cases

- **Single character** (`"a"`): nothing to pop; result is just `"a"`. Confirms the loop handles minimal input.
- **All identical characters** (`"aaaa"`): after the first `a` is pushed, every later `a` is skipped via `inStack`; result `"a"`. Verifies the duplicate-skip path.
- **Already strictly increasing** (`"abc"`): no pops ever happen; result equals the input. Verifies we never remove characters that shouldn't move.
- **Strictly decreasing with unique letters** (`"cba"`): no character reappears, so `last[top] > i` is always false and nothing is popped; result stays `"cba"`. This guards against over-popping — a character with no future occurrence must never be discarded.
- **Repeated blocks** (`"bcabc"`): earlier larger characters get popped once smaller ones arrive and are still recoverable, producing `"abc"`.
