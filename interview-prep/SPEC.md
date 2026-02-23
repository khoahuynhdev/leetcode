# Interview Prep Spec: Python & Bash for Backend/DevOps

## Context

Preparing for a whiteboard interview for a Backend Engineer / DevOps role. The interview involves solving DSA problems in Python and demonstrating Bash scripting competency (text processing and automation). Timeline is 1-2 weeks.

Current skill level: intermediate in both Python and Bash, with strong Go/LeetCode background (107 problems solved). The main gap is translating algorithmic thinking into idiomatic Python and filling in stdlib knowledge that Go doesn't have equivalents for.

## Goals

The prep has two tracks that should be worked on in parallel throughout the 1-2 week window. Track 1 focuses on Python fluency for whiteboard DSA. Track 2 focuses on practical Bash scripting. Each track is broken into topic blocks that can be done in any order, though the suggested ordering builds on earlier material.

---

## Track 1: Python for Whiteboard DSA

### Block 1.1 — Pythonic Idioms (Foundation)

This block covers the Python-specific patterns that make whiteboard code concise and readable. Coming from Go, the biggest shift is that Python rewards brevity and expressiveness — interviewers expect you to use these idioms naturally rather than writing C-style loops.

Topics to internalize:

- List comprehensions and generator expressions: `[x*2 for x in arr if x > 0]`, `sum(x for x in arr)`. Know when to use each (generators for large/lazy iteration, list comps when you need indexing).
- Tuple unpacking and multiple assignment: `a, b = b, a`, `for i, v in enumerate(arr)`, `for k, v in d.items()`.
- Slicing: `arr[::-1]` for reverse, `arr[i:j]`, `s[:k]`. Slices are O(k) not O(1) — know the cost.
- Truthiness: empty collections, `0`, `None`, `""` are falsy. Use `if not arr:` instead of `if len(arr) == 0:`.
- String operations: `"".join(list)` for building strings (not `+=` in a loop), f-strings for formatting, `s.split()`, `s.strip()`.
- `defaultdict`, `Counter`, `deque` from collections — these come up constantly.
- Walrus operator `:=` for inline assignment in conditions/loops (Python 3.8+).
- Star unpacking: `first, *rest = arr`, `*init, last = arr`.

#### Exercise 1.1: Pythonic Rewrites

Take 5 of your existing Go solutions from this repo (suggested: Two Sum, Valid Anagram, Group Anagrams, Contains Duplicate, Valid Palindrome) and rewrite them in Python with the most idiomatic style you can. Focus on using comprehensions, `Counter`, tuple unpacking, and slicing where appropriate. Put solutions in `interview-prep/python/` using the same directory convention.

### Block 1.2 — Standard Library for Interviews

These are the stdlib modules that replace what you'd manually implement in Go. Knowing them saves significant whiteboard time and signals Python fluency.

Key modules and their interview use cases:

- `collections.Counter` — frequency counting, anagram checks, top-k problems. Know `most_common(k)`, arithmetic operations between Counters (`c1 - c2`, `c1 & c2`).
- `collections.defaultdict` — graph adjacency lists, grouping. `defaultdict(list)`, `defaultdict(int)`.
- `collections.deque` — BFS queues, sliding window. `appendleft()`, `popleft()` are O(1). Also useful as a bounded buffer with `maxlen`.
- `heapq` — min-heap operations. `heappush`, `heappop`, `nlargest`, `nsmallest`, `heapify`. Python only has min-heap; negate values for max-heap.
- `bisect` — binary search on sorted lists. `bisect_left`, `bisect_right`, `insort`.
- `itertools` — `combinations`, `permutations`, `product` for brute force / backtracking. `chain` for flattening. `groupby` for consecutive grouping (input must be sorted).
- `functools.lru_cache` — memoization decorator for recursive DP. Use `@lru_cache(maxsize=None)` or `@cache` (3.9+). Arguments must be hashable (convert lists to tuples).
- `sorted()` with `key=lambda` — custom sorting. `sorted(arr, key=lambda x: (-x[1], x[0]))` for multi-key.
- `math.inf`, `float('inf')` — sentinel values for min/max tracking.
- `zip()` — parallel iteration. `zip(*matrix)` for transpose.

#### Exercise 1.2: Stdlib Drills

Solve these problems in Python, deliberately using the stdlib tools above. Each problem maps to specific modules:

1. Top K Frequent Elements (LC 347) — `Counter.most_common` or `heapq.nlargest`
2. Kth Largest Element in a Stream (LC 703) — `heapq`
3. Merge K Sorted Lists (LC 23) — `heapq` with tuples
4. Find K Closest Elements (LC 658) — `bisect` + sliding window or sorted with key
5. Group Anagrams (LC 49) — `defaultdict` with tuple key
6. Climbing Stairs (LC 70) — `@lru_cache` for memoized recursion

### Block 1.3 — OOP Patterns for Interviews

Python OOP comes up in design-flavored whiteboard problems and when implementing data structures. The key differences from Go: Python has real inheritance, dunder methods replace interfaces, and there's no explicit interface keyword.

Topics:

- `__init__`, `__repr__`, `__str__` — constructor and display. Always define `__repr__` for debugging on whiteboard.
- `__eq__`, `__hash__`, `__lt__` — needed if your objects go into sets/dicts or need sorting. If you define `__eq__`, you must also define `__hash__` (or it becomes unhashable).
- `@dataclass` — quick way to get `__init__`, `__repr__`, `__eq__` for free. Use `@dataclass(frozen=True)` for hashable immutable objects.
- Inheritance: `class LRUCache(OrderedDict)` is a legitimate pattern. `super().__init__()`.
- `@property` — computed attributes without getter/setter boilerplate.
- Context managers: `__enter__` / `__exit__` or `@contextmanager` decorator. Less common in interviews but shows depth.

#### Exercise 1.3: OOP Data Structures

Implement these as Python classes:

1. LRU Cache (LC 146) — use `OrderedDict` or implement with dict + doubly linked list
2. Min Stack (LC 155) — class with `push`, `pop`, `top`, `getMin`
3. Implement Trie (LC 208) — nested `defaultdict` or explicit TrieNode class
4. Design Twitter (LC 355) — combines `heapq`, `defaultdict`, OOP design

### Block 1.4 — Arrays, Strings, Trees & Graphs in Python

This block is about practicing the actual DSA categories you'll face, but in Python. Since you already know the algorithms from Go, the focus is on Python-specific implementation patterns.

Array/String patterns in Python:

- Two pointers: use `while l < r` with `l, r = 0, len(arr)-1`.
- Sliding window: track window state with `Counter` or `defaultdict(int)`, adjust counts as window moves.
- Prefix sums: `list(itertools.accumulate(arr))` or manual loop. `accumulate` also takes a `func` argument.
- String building: always use `list.append()` then `"".join()`, never `+=` in a loop (O(n^2)).

Tree/Graph patterns in Python:

- TreeNode: define inline `class TreeNode: def __init__(self, val=0, left=None, right=None)`.
- BFS: `queue = deque([root])`, `while queue: node = queue.popleft()`.
- DFS: recursive (natural in Python) or iterative with `stack = [root]`.
- Graph adjacency: `graph = defaultdict(list)`, build with `graph[u].append(v)`.
- Visited set: `visited = set()`, `visited.add(node)`.
- Level-order trick: `for _ in range(len(queue))` to process one level at a time.

#### Exercise 1.4: Core DSA in Python

Solve these in Python, focusing on clean idiomatic code you'd write on a whiteboard:

Arrays & Strings:

1. 3Sum (LC 15) — two pointers, skip duplicates
2. Longest Substring Without Repeating Characters (LC 3) — sliding window + set or dict
3. Product of Array Except Self (LC 238) — prefix/suffix
4. Container With Most Water (LC 11) — two pointers
5. Minimum Window Substring (LC 76) — sliding window + Counter

Trees & Graphs:

1. Binary Tree Level Order Traversal (LC 102) — BFS with deque
2. Validate Binary Search Tree (LC 98) — DFS with bounds
3. Lowest Common Ancestor (LC 236) — recursive DFS
4. Number of Islands (LC 200) — BFS/DFS on grid
5. Course Schedule (LC 207) — topological sort with DFS or BFS (Kahn's)

---

## Track 2: Bash Scripting

### Block 2.1 — Text Processing

These are the core tools for log analysis, data extraction, and file manipulation that DevOps interviews test on.

Essential commands and patterns:

- `grep`: `-E` for extended regex, `-i` ignore case, `-v` invert, `-c` count, `-l` files only, `-r` recursive, `-o` only matching part, `-P` for Perl regex (lookahead/lookbehind).
- `sed`: `s/old/new/g` substitution, `-i` in-place edit, address ranges (`2,5s/...`), delete lines (`/pattern/d`), print specific lines (`-n '5p'`).
- `awk`: field splitting (`$1`, `$2`, `-F,` for CSV), conditions (`$3 > 100`), aggregation (`{sum += $3} END {print sum}`), printf for formatting, associative arrays.
- Pipelines: chaining `|` for multi-step processing. `sort | uniq -c | sort -rn` for frequency analysis. `tee` for splitting output.
- `cut` for simple field extraction (`cut -d, -f2`), `tr` for character translation/deletion, `sort` flags (`-n` numeric, `-k` key, `-t` delimiter, `-u` unique).
- `xargs`: converting stdin to arguments. `-I{}` for placeholder, `-P` for parallel, `-0` for null-delimited.
- Process substitution: `diff <(cmd1) <(cmd2)`, useful for comparing command outputs.
- Here strings and here docs: `cmd <<< "string"`, `cmd <<EOF ... EOF`.

#### Exercise 2.1: Text Processing Challenges

Create these scripts in `interview-prep/bash/`:

1. `log-analyzer.sh` — Given an nginx-style access log, extract: the top 10 IPs by request count, the distribution of HTTP status codes, and the top 5 most requested URLs. Use only `grep`, `awk`, `sort`, `uniq`, `head`.

2. `csv-transform.sh` — Read a CSV file, filter rows where a specified column exceeds a threshold, reformat the output as `key=value` pairs. Handle quoted fields with commas inside them.

3. `find-and-replace.sh` — Recursively find all files matching a pattern in a directory, perform a regex substitution in each file, and report what changed (filename, line number, before/after). Use `find`, `grep`, `sed`.

4. `log-monitor.sh` — Tail a log file and alert (print a highlighted message) whenever a line matching a given pattern appears. Support multiple patterns from a config file.

### Block 2.2 — Automation Scripts

These cover the scripting patterns used in deployment, CI/CD, and operational automation.

Patterns to know:

- Script structure: shebang (`#!/usr/bin/env bash`), `set -euo pipefail` (exit on error, undefined vars, pipe failures), meaningful exit codes.
- Argument parsing: `getopts` for short options, manual parsing with `case` for long options, positional args with `$1`, `$2`, `shift`.
- Functions: `function_name() { ... }`, local variables with `local`, return values via `echo` + command substitution or global vars.
- Error handling: `trap 'cleanup' EXIT` for cleanup on exit, `||` for fallback commands, custom error messages to stderr (`>&2`).
- Conditionals: `[[ ]]` vs `[ ]` (prefer `[[ ]]`), string comparison, numeric comparison (`-eq`, `-lt`), file tests (`-f`, `-d`, `-x`).
- Loops: `for f in *.log`, `while read -r line`, `for ((i=0; i<n; i++))`. Always use `read -r` to avoid backslash interpretation.
- Arrays: `arr=(a b c)`, `${arr[@]}`, `${#arr[@]}`, associative arrays with `declare -A`.
- String manipulation: `${var##*/}` basename, `${var%.*}` remove extension, `${var:-default}` default value, `${var:+alt}` alternate value.
- Subshells and grouping: `()` for subshell (isolated env), `{}` for grouping in current shell.
- Temp files: `mktemp` for safe temp files, `trap` to clean them up.

#### Exercise 2.2: Automation Challenges

1. `deploy-checker.sh` — A pre-deployment validation script that: checks if required environment variables are set, verifies connectivity to a list of hosts (from a config file), validates that a given Docker image exists in a registry (mock the API call with curl to a local file), and produces a go/no-go summary. Use proper error handling, colored output, and exit codes.

2. `backup-rotate.sh` — A backup rotation script that: takes a source directory and backup destination as arguments, creates a timestamped tarball, keeps only the N most recent backups (configurable via flag), supports dry-run mode (`--dry-run`), and logs all actions. Use `getopts` or manual arg parsing, `trap` for cleanup, and proper quoting throughout.

3. `service-health.sh` — A service health checker that: reads a list of service URLs from a config file, checks each endpoint with configurable timeout and retry count, outputs a status table (service name, URL, status, response time), supports JSON or plain-text output format (via `--format` flag), and returns non-zero exit code if any service is down.

---

## Practice Structure

For each topic block, the recommended workflow is:

Read through the topic notes first, then attempt each exercise without looking at references. After attempting, review your solution for idiom improvements. For Python exercises, run the solution against LeetCode to verify correctness. For Bash exercises, test with sample data you create.

The Python exercises should be placed in `interview-prep/python/NNNN-problem-name/solution.py` following the existing repo convention. The Bash exercises go in `interview-prep/bash/script-name.sh`. Each Bash script should include a comment header describing its purpose and usage.

## Suggested Order

Start with Block 1.1 (Pythonic Idioms) since it underpins everything else in Python. Then alternate between Python and Bash blocks to keep things varied. A reasonable sequence would be: 1.1 → 2.1 → 1.2 → 1.4 → 2.2 → 1.3. Adjust based on which areas feel weakest after the first block.

## Success Criteria

By interview day, you should be able to: write a medium-difficulty LeetCode solution in Python in under 20 minutes using idiomatic stdlib tools, write a 50-line Bash script from scratch with proper error handling and argument parsing, and explain your Python/Bash code choices clearly as you write (the whiteboard communication aspect).
