# Python Cheatsheet: Arrays & Maps for Interviews

## Lists (Go slices equivalent)

### Creation

```python
arr = [1, 2, 3]
zeros = [0] * 10                    # [0, 0, 0, ..., 0]  length 10
matrix = [[0] * cols for _ in range(rows)]  # 2D array — MUST use list comp
# WRONG: [[0]*cols]*rows  — all rows share the same reference!
```

### Indexing & Slicing

```python
arr[0]          # first element
arr[-1]         # last element (Go: arr[len(arr)-1])
arr[-2]         # second to last

arr[1:4]        # elements at index 1,2,3 (start inclusive, end exclusive)
arr[:3]         # first 3 elements
arr[2:]         # everything from index 2 onward
arr[:]          # shallow copy of entire list
arr[::2]        # every 2nd element: [arr[0], arr[2], arr[4], ...]
arr[::-1]       # reversed copy (does NOT modify original)
```

### Adding & Removing

```python
arr.append(x)           # add to end — O(1) amortized (Go: append(arr, x))
arr.pop()               # remove & return last — O(1)
arr.pop(i)              # remove & return at index i — O(n)
arr.insert(i, x)        # insert x at index i — O(n)
arr.extend([4, 5, 6])   # append multiple items — O(k)
arr += [4, 5, 6]        # same as extend

del arr[i]              # delete by index — O(n)
arr.remove(x)           # delete first occurrence of value x — O(n), raises ValueError if missing
```

### Searching

```python
x in arr                # membership check — O(n) for list, O(1) for set
arr.index(x)            # first index of x — O(n), raises ValueError if missing
arr.count(x)            # count occurrences — O(n)
```

### Sorting

```python
arr.sort()                          # in-place, returns None
arr.sort(reverse=True)              # descending in-place
arr.sort(key=lambda x: x[1])       # sort by second element of tuples
arr.sort(key=lambda x: (-x[1], x[0]))  # multi-key: descending by [1], then ascending by [0]

sorted(arr)                         # returns new sorted list, original unchanged
sorted(arr, key=len)                # sort by length
sorted(arr, key=str.lower)          # case-insensitive sort
```

### Comprehensions

```python
[x * 2 for x in arr]                       # map
[x for x in arr if x > 0]                  # filter
[x * 2 for x in arr if x > 0]              # map + filter
[f(x) for sublist in nested for x in sublist]  # flatten (Go: nested loop)
[[row[i] for row in matrix] for i in range(cols)]  # transpose
```

### Unpacking

```python
a, b, c = [1, 2, 3]            # exact unpack
first, *rest = [1, 2, 3, 4]    # first=1, rest=[2,3,4]
*init, last = [1, 2, 3, 4]     # init=[1,2,3], last=4
a, *_, b = [1, 2, 3, 4, 5]     # a=1, b=5, middle discarded
```

### Common Patterns

````python
# Swap (no temp variable needed)
a, b = b, a
arr[i], arr[j] = arr[j], arr[i]

# Enumerate — index + value (Go: for i, v := range arr)
for i, v in enumerate(arr):
    print(i, v)
example:
```python

  def two_sum(nums: list[int], target: int) -> list[int]:
      seen = {}
      for i, v in enumerate(nums):
          if (comp := target - v) in seen:
              return [seen[comp], i]
          seen[v] = i
```

# Zip — parallel iteration over two lists

for a, b in zip(list1, list2):
print(a, b)

# Reverse in-place

arr.reverse() # modifies arr
arr[:] = arr[::-1] # also in-place via slice assignment

# Min/Max with index

min_val = min(arr)
min_idx = arr.index(min(arr))

# or in one pass:

min_idx, min_val = min(enumerate(arr), key=lambda x: x[1])

# Check if sorted

all(arr[i] <= arr[i+1] for i in range(len(arr)-1))

# Flatten one level

flat = [x for sub in nested for x in sub]

# or: list(itertools.chain.from_iterable(nested))

# Chunk into groups of n

chunks = [arr[i:i+n] for i in range(0, len(arr), n)]

````

### Gotchas Coming from Go

```python
# Slicing creates a COPY (unlike Go slices which share underlying array)
b = a[1:3]
b[0] = 99       # does NOT affect a

# But assignment without slicing is a reference
b = a
b.append(99)    # DOES affect a — they point to the same list

# list.sort() returns None, not the list
wrong = arr.sort()    # wrong is None
right = sorted(arr)   # right is the sorted list

# Negative indexing is valid (Go would panic)
arr[-1]   # last element, no need for len(arr)-1
```

---

## Dicts (Go map equivalent)

### Creation

```python
d = {"a": 1, "b": 2}
d = dict(a=1, b=2)                         # same thing (keys must be valid identifiers)
d = {x: x**2 for x in range(5)}            # dict comprehension: {0:0, 1:1, 2:4, 3:9, 4:16}
d = dict.fromkeys(["a", "b", "c"], 0)      # all keys with same default: {"a":0, "b":0, "c":0}
```

### Access

```python
d["key"]                # get value — raises KeyError if missing (like Go panic)
d.get("key")            # get value — returns None if missing (no error)
d.get("key", default)   # get value — returns default if missing

# Go equivalent: val, ok := m["key"]
# Python equivalent:
if "key" in d:
    val = d["key"]
# or just use .get()
```

### Adding & Updating

```python
d["key"] = value                # set/overwrite
d.update({"a": 1, "b": 2})     # merge another dict in
d |= {"a": 1, "b": 2}          # merge operator (Python 3.9+)
merged = d1 | d2                # new merged dict (Python 3.9+), d2 wins on conflicts

d.setdefault("key", []).append(x)  # if key missing, set it to [], then append
# This is the manual version of defaultdict(list)
```

### Removing

```python
del d["key"]                    # delete — raises KeyError if missing
val = d.pop("key")              # delete & return — raises KeyError if missing
val = d.pop("key", default)     # delete & return — returns default if missing
d.clear()                       # remove all entries
```

### Iteration

```python
for k in d:                     # iterate keys (Go: for k := range m)
for k, v in d.items():          # iterate key-value pairs (Go: for k, v := range m)
for v in d.values():            # iterate values only
for k in d.keys():              # iterate keys explicitly (same as `for k in d`)

# Sorted iteration
for k in sorted(d):                         # keys sorted
for k, v in sorted(d.items(), key=lambda x: x[1]):  # sorted by value
```

### Membership

```python
"key" in d          # check if key exists — O(1)
"key" not in d      # check if key missing — O(1)
```

### Common Patterns

```python
# Frequency counter (manual)
freq = {}
for x in arr:
    freq[x] = freq.get(x, 0) + 1

# Frequency counter (pythonic)
from collections import Counter
freq = Counter(arr)             # Counter({"a": 3, "b": 2, ...})
freq.most_common(3)             # top 3 by count: [("a", 3), ("b", 2), ...]
freq["missing"]                 # returns 0 (not KeyError — Counter special behavior)

# Grouping (manual)
groups = {}
for item in items:
    key = get_key(item)
    groups.setdefault(key, []).append(item)

# Grouping (pythonic)
from collections import defaultdict
groups = defaultdict(list)
for item in items:
    groups[get_key(item)].append(item)

# Invert a dict (swap keys and values)
inv = {v: k for k, v in d.items()}

# Dict from two lists
d = dict(zip(keys, values))

# Count-based check (e.g., anagram)
Counter(s) == Counter(t)
```

### defaultdict

```python
from collections import defaultdict

dd = defaultdict(int)           # missing keys default to 0
dd = defaultdict(list)          # missing keys default to []
dd = defaultdict(set)           # missing keys default to set()
dd = defaultdict(lambda: "N/A")  # custom default

# Usage — no need to check "if key in dd" before accessing
dd[key] += 1                    # works even if key is new (starts at 0)
dd[key].append(val)             # works even if key is new (starts at [])

# Graph adjacency list
graph = defaultdict(list)
for u, v in edges:
    graph[u].append(v)
    graph[v].append(u)          # undirected
```

### Counter

```python
from collections import Counter

c = Counter("aabbc")            # Counter({"a":2, "b":2, "c":1})
c = Counter([1, 1, 2, 3])      # Counter({1:2, 2:1, 3:1})

c.most_common(2)                # [("a", 2), ("b", 2)]
c.total()                       # sum of all counts (Python 3.10+)
list(c.elements())              # ["a", "a", "b", "b", "c"] — expanded back

# Arithmetic between Counters
c1 + c2                         # add counts
c1 - c2                         # subtract (drops zero and negative)
c1 & c2                         # min of each count (intersection)
c1 | c2                         # max of each count (union)
```

### Gotchas Coming from Go

```python
# KeyError is the default for missing keys (Go returns zero value)
d["missing"]        # KeyError!  Use d.get("missing") or defaultdict

# Dicts preserve insertion order (Python 3.7+)
# Go maps iterate in random order; Python dicts are deterministic

# You can't use lists as dict keys (unhashable)
# Use tuples instead:
d[(1, 2)] = "ok"    # tuple key works
d[[1, 2]] = "bad"   # TypeError: unhashable type: 'list'

# Dict comprehension is your friend for transforms
{k: v * 2 for k, v in d.items() if v > 0}

# Deleting during iteration raises RuntimeError
# Copy keys first: for k in list(d.keys()): ...
```

---

## Sets (no Go equivalent — Go uses map[T]bool)

```python
s = {1, 2, 3}                  # literal (NOT {} — that's an empty dict)
s = set()                       # empty set
s = set([1, 2, 2, 3])          # from list, deduplicates: {1, 2, 3}

s.add(x)                       # add element — O(1)
s.remove(x)                    # remove — raises KeyError if missing
s.discard(x)                   # remove — NO error if missing (prefer this)
s.pop()                        # remove & return arbitrary element
x in s                         # membership — O(1)

# Set operations
s1 | s2                         # union
s1 & s2                         # intersection
s1 - s2                         # difference (in s1 but not s2)
s1 ^ s2                         # symmetric difference (in one but not both)
s1 <= s2                        # is s1 a subset of s2?

# Frozen set (immutable — can be used as dict key or in another set)
fs = frozenset([1, 2, 3])
```

---

## Number Base Conversions

### Int to String Representation

```python
# Binary (base 2)
bin(42)             # '0b101010'
f"{42:b}"           # '101010'       — no prefix
f"{42:08b}"         # '00101010'     — zero-padded to 8 bits

# Octal (base 8)
oct(42)             # '0o52'
f"{42:o}"           # '52'

# Hexadecimal (base 16)
hex(42)             # '0x2a'         — lowercase
f"{42:x}"           # '2a'           — lowercase, no prefix
f"{42:X}"           # '2A'           — uppercase
f"{42:08x}"         # '0000002a'     — zero-padded to 8 hex digits

# Arbitrary base (no built-in — roll your own)
def to_base(n: int, base: int) -> str:
    if n == 0:
        return "0"
    digits, neg = [], n < 0
    n = abs(n)
    while n:
        digits.append("0123456789abcdefghijklmnopqrstuvwxyz"[n % base])
        n //= base
    return ("-" if neg else "") + "".join(reversed(digits))

to_base(255, 16)    # 'ff'
to_base(42, 5)      # '132'
```

### String to Int

```python
int('101010', 2)    # 42   — from binary
int('52', 8)        # 42   — from octal
int('2a', 16)       # 42   — from hex
int('132', 5)       # 42   — from base 5

# Auto-detect base from prefix (0b, 0o, 0x)
int('0b101010', 0)  # 42
int('0o52', 0)      # 42
int('0x2a', 0)      # 42
```

### Base-to-Base Conversion

```python
# There's no direct base-to-base function. Go through int as the intermediate:
binary_str = '101010'
hex_str = f"{int(binary_str, 2):x}"     # '2a'  (bin → int → hex)
oct_str = f"{int(binary_str, 2):o}"     # '52'  (bin → int → oct)
```

### Useful Bit Operations

```python
n.bit_length()          # number of bits needed: (42).bit_length() == 6
n.bit_count()           # popcount (number of 1-bits): (42).bit_count() == 3  (Python 3.10+)
bin(n).count('1')       # popcount fallback for older Python
```
