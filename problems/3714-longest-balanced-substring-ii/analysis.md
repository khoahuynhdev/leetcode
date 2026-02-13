# 3714. Longest Balanced Substring II

[LeetCode Link](https://leetcode.com/problems/longest-balanced-substring-ii/)

Difficulty: Medium
Topics: Hash Table, String, Prefix Sum
Acceptance Rate: 26.6%

## Hints

### Hint 1

Think about what "balanced" means in terms of character counts. A substring is balanced if every distinct character in it appears the same number of times. Consider how prefix sums of character frequencies might help you detect equal-count substrings efficiently.

### Hint 2

The classic trick for "equal frequency" problems is to track **differences** between character counts rather than absolute counts. If the difference between counts is the same at two positions, the substring between those positions has equal increments for the relevant characters. But be careful — this only handles the case where **all three** characters appear equally. You need to consider substrings that contain only one or two distinct characters separately.

### Hint 3

Break the problem into cases based on which characters appear in the substring:
- **Single character**: Any maximal run of one character is balanced. Just find the longest run.
- **Two characters**: For each pair like (a,b), find maximal contiguous segments that contain only those two characters, then within each segment use a prefix-difference map (`count_x - count_y`) to find the longest sub-segment where both characters appear equally.
- **Three characters**: Use a 2D prefix-difference state `(count_a - count_b, count_a - count_c)`. Store the earliest index for each state; when you see the same state again, the substring between has equal counts of all three.

The answer is the maximum across all cases.

## Approach

We split the problem into three scenarios by how many distinct characters the balanced substring contains:

**Case 1 — One distinct character:** Scan the string and track the length of consecutive runs of the same character. Any such run is balanced.

**Case 2 — Two distinct characters (3 sub-cases for pairs ab, ac, bc):** For each character pair, identify maximal contiguous segments of the string that contain only characters from that pair. Within each segment, use the prefix-difference technique: maintain a running difference `count_x - count_y` and a hash map storing the first index where each difference value was seen. When the same difference appears again, the substring between those indices has equal counts of both characters.

**Case 3 — All three characters:** Maintain running counts of a, b, c. The state is `(count_a - count_b, count_a - count_c)`. Use a hash map from state to the earliest index. When the same state recurs at index j, the substring from the stored index+1 to j has equal counts of all three characters.

The final answer is the maximum length found across all cases.

### Example walkthrough: s = "abbac"

- Prefix counts at each position: a=[1,1,1,2,2], b=[0,1,2,2,2], c=[0,0,0,0,1]
- For the all-three case, states (a-b, a-c): index 0→(1,-1), 1→(0,1), 2→(-1,1), 3→(0,2), 4→(1,1). No repeated states, so no all-three balanced substring.
- For pair (a,b): segment "abba" (indices 0–3) has only a,b. Differences [1,0,-1,0]. Difference 0 first at index 1, recurs at index 3 → length 2. But also difference 0 at the start (before index 0) maps to length at index 3 = 4. So we get length 4.
- Result: 4.

## Complexity Analysis

Time Complexity: O(n) — We make a constant number of linear passes over the string.

Space Complexity: O(n) — For the hash maps storing first-occurrence indices of prefix-difference states.

## Edge Cases

- **Single character string** (e.g., "a"): The entire string is balanced. Length = 1.
- **All same characters** (e.g., "aaaa"): The entire string is balanced.
- **Two characters alternating** (e.g., "ababab"): Longest balanced substring is the full string if both have equal counts, or length n-1 otherwise.
- **No two characters have equal count in any substring longer than 1**: The answer is 1 (each single character is trivially balanced).
- **Entire string is balanced with all three characters**: The answer is the full string length.
