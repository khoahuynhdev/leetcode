# 1096. Brace Expansion II

[LeetCode Link](https://leetcode.com/problems/brace-expansion-ii/)

Difficulty: Hard
Topics: Hash Table, String, Backtracking, Stack, Breadth-First Search, Sorting
Acceptance Rate: 69.1%

## Hints

### Hint 1

The problem describes a *grammar* with nested braces. Whenever the input is a nested expression with operators, think of how you would evaluate an arithmetic expression like `2*(3+4)`. What plays the role of `+` here, and what plays the role of `*`?

### Hint 2

Union (`,`) acts like `+` and concatenation (placing things side by side) acts like `*`. Concatenation binds tighter than union, just like multiplication binds tighter than addition. A recursive descent parser, or a two-stack evaluator, handles this precedence cleanly. Use sets so duplicates disappear on their own.

### Hint 3

Split the grammar into three levels:
- `expr := term (',' term)*`: take the **union** of the terms.
- `term := factor factor*`: take the **cartesian concatenation** of the factors.
- `factor := letter | '{' expr '}'`: a letter is a singleton set, and a brace group recurses back into `expr`.

Keep one shared index into the string. Each function consumes exactly the characters it owns and returns a set of words.

## Approach

Handle the expression the same way you would handle an arithmetic expression, using **recursive descent** with a shared cursor `pos`:

1. **parseExpr** calls `parseTerm` for its first set. While the next character is `,`, it consumes the comma, parses another term, and adds that term's words into the result set.
2. **parseTerm** starts with the set `{""}`, which is the identity for concatenation. While the next character can start a factor (a letter or `{`), it parses the factor and replaces the result with `{a + b | a ∈ result, b ∈ factor}`. It stops at `,`, `}`, or the end of the input, because those characters belong to an enclosing level.
3. **parseFactor** handles two cases. On `{`, it consumes the brace, recurses into `parseExpr`, and then consumes the matching `}`. Otherwise it reads a run of letters and returns that word as a singleton set. Reading the whole run at once is a small optimization and gives the same result as concatenating one letter at a time.
4. Take the final set, copy it into a slice, and sort it.

**Example:** `{a,b}{c,{d,e}}`
- `parseTerm` reads factor `{a,b}`, which gives `{a, b}`.
- Next it reads factor `{c,{d,e}}`. Inside, `parseExpr` takes the union of `{c}` and `{d, e}`, giving `{c, d, e}`.
- The product is `{ac, ad, ae, bc, bd, be}`, and after sorting that is the answer.

This works because every level of the grammar maps directly to a set operation, and a hash set removes duplicates as they are produced. Example 2 shows why that matters: `a` and `z` each appear more than once there.

## Complexity Analysis

Time Complexity: O(W · L + W log W · L), where W is the number of words produced (including intermediate products) and L is the maximum word length. Parsing visits each character once. Most of the cost is building the concatenation products and then sorting the final words. With `expression.length <= 60`, W is small in practice.
Space Complexity: O(W · L) for the intermediate and final sets, plus O(D) recursion depth, where D is the brace nesting depth.

## Edge Cases

- **Single letter or plain word (`"a"`, `"abc"`)**: there are no braces at all. The term loop must still collect the letters correctly.
- **Duplicates from union (`{{a,b},{b,c}}`)**: the same word can come from several branches. A set removes them.
- **Duplicates from concatenation (`{a,ab}{c,bc}`)**: different pairs can produce the same string (`a+bc` = `ab+c`). Deduplication has to happen after the product, not only after unions.
- **Deep nesting (`{{{a}}}`)**: the recursion must consume every `{` and its matching `}` correctly.
- **Stopping at `,` and `}`**: `parseTerm` must stop at these characters without consuming them. Otherwise the enclosing level loses track of its structure.
- **Sorting**: the output must be sorted lexicographically, so `"a" < "ab" < "ac"`. Set iteration order in Go is random, so always sort at the end.
