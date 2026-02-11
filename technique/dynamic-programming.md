# Dynamic Programming in Go

## How to Identify DP

DP applies when a problem has two properties: overlapping subproblems (you solve the same smaller problem multiple times) and optimal substructure (the optimal solution to the whole problem is built from optimal solutions to subproblems). In practice, look for phrases like "minimum/maximum", "count the number of ways", "is it possible", or "longest/shortest" — these are strong signals.

## The Hard Part: Defining the State

The state is what you need to remember about the past to make the optimal decision going forward. Finding the right state is the core challenge of DP, and there's no single formula, but here are practical strategies.

### Strategy 1: Think About What Decision You Make at Each Step

At each position/element, ask "what choices do I have?" The state should capture enough information to evaluate those choices. For problem 1653 (Minimum Deletions to Make String Balanced), at each character you either keep it or delete it. To decide, you need to know: how many deletions so far, and how many `b`s are before this position. That's the state — just two numbers.

A good exercise: write out the decision at position `i` in plain English before writing code. If your sentence references some quantity about the past, that quantity is part of your state.

### Strategy 2: Start with the Obvious State, Then Reduce

Begin with the most straightforward (possibly expensive) state and get a working solution. Then look for redundancies. A 2D DP where one dimension can be derived from the other can often collapse to 1D. A 1D DP where you only look at the previous entry can collapse to O(1) variables.

For example, many substring problems start as `dp[i][j]` (answer for substring `s[i..j]`), and some can be reduced to `dp[i]` if you realize `j` is always derived from the iteration.

### Strategy 3: Ask "What Do I Need to Know About the Past?"

Forget about the future. Stand at position `i` and ask: what's the minimum information about `s[0..i-1]` that lets me make the right choice here? Often it's a single counter (like `bCount` in problem 1653) or a small set of values.

If you find yourself needing to know the exact arrangement of past elements, you probably have the wrong framing. DP states should be summary statistics, not full histories.

### Strategy 4: Try Different "What Does dp[i] Mean?" Definitions

The same problem can often be solved with different state definitions, and some are much easier to work with than others. Common framings:

`dp[i]` = answer considering only the first `i` elements. This is the most natural and usually the first thing to try. Transition looks at `dp[i-1]` or `dp[j]` for various `j < i`.

`dp[i]` = answer where element `i` is the last element included. This is useful for subsequence problems (LIS, etc.) where you need to know what you ended with to decide if the next element extends it.

`dp[i]` = answer for the suffix starting at `i`. Sometimes working backwards makes transitions simpler. If the forward direction has awkward dependencies, try reversing.

If one definition leads to a transition that requires scanning the whole array (O(n) per state), try another definition — it might give O(1) transitions.

## Common DP Patterns

### Linear DP

Process elements left to right, one decision per element.

```go
// dp[i] depends on dp[i-1] and possibly dp[i-2], ..., dp[0]
dp := make([]int, n+1)
dp[0] = baseCase
for i := 1; i <= n; i++ {
    dp[i] = transition(dp[i-1], ...)
}
return dp[n]
```

When `dp[i]` only depends on the previous one or two values, you don't need the array at all:

```go
prev := baseCase
for i := 1; i <= n; i++ {
    curr := transition(prev, ...)
    prev = curr
}
return prev
```

Problem 1653 is a good example — `deletions` is effectively `dp[i]` compressed to a single variable, and `bCount` is auxiliary state tracked alongside it.

### Knapsack / Subset

Choose items with constraints (weight, capacity). State is `dp[i][w]` = best value using first `i` items with capacity `w`.

```go
dp := make([][]int, n+1)
for i := range dp {
    dp[i] = make([]int, capacity+1)
}
for i := 1; i <= n; i++ {
    for w := 0; w <= capacity; w++ {
        dp[i][w] = dp[i-1][w] // skip item i
        if w >= weight[i] {
            dp[i][w] = max(dp[i][w], dp[i-1][w-weight[i]]+value[i]) // take item i
        }
    }
}
```

Space optimization: if `dp[i]` only depends on `dp[i-1]`, use a single 1D array and iterate `w` in reverse (for 0/1 knapsack) to avoid using updated values.

```go
dp := make([]int, capacity+1)
for i := 0; i < n; i++ {
    for w := capacity; w >= weight[i]; w-- {
        dp[w] = max(dp[w], dp[w-weight[i]]+value[i])
    }
}
```

### Interval DP

State is a range `dp[i][j]` representing the answer for substring/subarray `s[i..j]`. Iterate by increasing length.

```go
n := len(s)
dp := make([][]int, n)
for i := range dp {
    dp[i] = make([]int, n)
}
// base: length 1
for i := 0; i < n; i++ {
    dp[i][i] = baseCaseValue
}
// expand by length
for length := 2; length <= n; length++ {
    for i := 0; i+length-1 < n; i++ {
        j := i + length - 1
        // try splitting at every k in [i, j-1]
        for k := i; k < j; k++ {
            dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+mergeCost)
        }
    }
}
return dp[0][n-1]
```

### Grid DP

Navigate a 2D grid, usually moving right or down.

```go
dp := make([][]int, m)
for i := range dp {
    dp[i] = make([]int, n)
}
dp[0][0] = grid[0][0]
// fill first row and column, then:
for i := 1; i < m; i++ {
    for j := 1; j < n; j++ {
        dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
    }
}
```

### Two-String DP

Compare or align two strings. State is `dp[i][j]` = answer for `s[:i]` and `t[:j]`.

```go
dp := make([][]int, len(s)+1)
for i := range dp {
    dp[i] = make([]int, len(t)+1)
}
// base cases: dp[0][j], dp[i][0]
for i := 1; i <= len(s); i++ {
    for j := 1; j <= len(t); j++ {
        if s[i-1] == t[j-1] {
            dp[i][j] = dp[i-1][j-1] + 1 // match
        } else {
            dp[i][j] = max(dp[i-1][j], dp[i][j-1]) // skip one
        }
    }
}
```

## Tips and Tricks

### When You're Stuck on the State

Try working through a small example by hand. Write down what you're tracking mentally as you make decisions. That mental state is your DP state. If you're tracking too much, the problem might need a different angle (like problem 1653: tracking "which characters are kept" is too much, but tracking "how many deletions so far" is just right).

### Flip the Problem

If "maximize what to keep" leads to a complicated state, try "minimize what to remove" (or vice versa). Problem 1653 is a textbook example — maximizing the longest valid subsequence is harder than tracking minimum deletions.

### Greedy vs DP

If at every step the locally best choice is also globally best, you can use greedy. If you can construct a counterexample where the greedy choice at one step leads to a worse overall answer, you need DP. A quick way to test: try your greedy on 3-4 small inputs and see if it always matches the expected output.

### Top-Down vs Bottom-Up

Top-down (memoized recursion) is often easier to write because the recursion structure matches how you think about the problem. Bottom-up (iterative) is usually faster in practice (no recursion overhead) and easier to space-optimize. Start with whichever feels more natural, then convert if needed.

```go
// top-down with memoization
memo := make(map[[2]int]int)
var solve func(i, state int) int
solve = func(i, state int) int {
    if i == n {
        return 0
    }
    key := [2]int{i, state}
    if v, ok := memo[key]; ok {
        return v
    }
    // transitions...
    memo[key] = result
    return result
}

// bottom-up equivalent
dp := make([][]int, n+1)
for i := n; i >= 0; i-- {
    for state := range possibleStates {
        // same transitions but reading from dp[i+1][...]
    }
}
```

### Space Optimization Checklist

If `dp[i]` only depends on `dp[i-1]`: use two variables or two rows instead of a full array. If `dp[i]` only depends on `dp[i-1]` and `dp[i-2]`: use three variables. If `dp[i][j]` only depends on the previous row: use a 1D array (careful about iteration order).

## Problems Log

Track problems and what DP pattern/state insight they used, so you can spot recurring themes.

| Problem | Pattern | Key State Insight |
|---|---|---|
| 1653. Minimum Deletions to Make String Balanced | Linear DP (O(1) space) | Track `deletions` and `bCount`; at each `'a'`, choose min(delete this a, delete all previous b's). Flipping from "max subsequence to keep" to "min deletions" makes the state trivial. |
