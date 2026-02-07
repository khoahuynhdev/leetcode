### The greedy "extend or shrink" logic doesn't work

Your algorithm tries to maintain a running maxKeep count: if the next character is valid (i.e., the pair s[i-1], s[i] is non-decreasing), you increment it; if not, you walk backward decrementing maxValid until you find a compatible character. The problem is that this treats the "longest valid subsequence" as something you can compute by only looking at the immediate predecessor relationship and doing a local backward scan.

Consider "bbaaaaabb" (expected answer: 2). Your code starts with maxKeep = 1. At index 1, b,b is valid so maxKeep = 2. At index 2, b,a is invalid, so the inner loop walks backward. It decrements maxValid for every b it hits, eventually reaching the start. But maxValid can go to 0 or even negative — there's no max(maxValid, 1) floor — and then that corrupted value becomes the new maxKeep going forward. The entire running count gets thrown off by a single bad
transition.

### It only considers one "path" through the string

The real question at each position is: "what's the best way to partition characters seen so far into kept vs. deleted?" That's inherently a choice with multiple branches. Your code makes a single greedy decision at each step — keep extending or walk back — without ever considering that you might skip the current character entirely (delete it) and preserve a better maxKeep from before. A correct approach needs to track both options: "delete this character" vs. "keep this character and pay the cost of removing conflicting ones."

### The backward scan double-counts deletions

When you hit an invalid pair and walk j backward, you decrement maxValid for each step. But those characters at positions j might already have been "not kept" in a previous step. You're subtracting from a running total that doesn't actually represent which specific characters are in your kept subsequence. There's no bookkeeping of what's actually included vs. excluded.

### Why the correct solution works

The working solution.go reframes the problem entirely. Instead of tracking what to keep, it tracks deletions directly. For each 'a' encountered, there are exactly two choices: delete this 'a' (cost deletions + 1) or delete all bs seen so far (cost bCount). Taking the minimum of these two options at every 'a' gives the globally optimal answer, because the decision is self-contained — it doesn't depend on remembering which specific characters were deleted before, just the count.

### Key takeaway

The pattern here is a classic DP insight: instead of trying to greedily build the longest valid subsequence (which requires tracking exactly which
elements you picked), track the minimum cost of violations so far. The state is just one number (deletions) plus one counter (bCount), and the transition at each character is a clean min of two options. Your attempt was essentially trying to solve a DP problem with a greedy approach, and the backward scan was an ad-hoc attempt to recover from greedy mistakes — but that recovery logic can't capture all the cases correctly.
