# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

This is a personal LeetCode solutions repository in Go for interview preparation. All 107 solutions are organized in a standardized structure under `problems/` directory:

- `problems/[0001-problem-name]/` - All solutions follow this zero-padded, kebab-case naming convention
  - Example: `problems/0104-maximum-depth-of-binary-tree/`
- Each problem directory contains:
  - `solution.go` - Primary solution implementation
  - `solution_test.go` - Tests (when available)
  - `solution_v2.go` - Alternative approaches (when available)

All solutions use `package main`.

## Development Commands

### Running Tests
```bash
# Run all tests
go test -v ./...

# Run tests for a specific problem
go test -v ./problems/0104-maximum-depth-of-binary-tree

# Run tests in helper package
go test -v ./helper
```

### Building
```bash
# Note: Not all solutions will build successfully as standalone programs
# This is expected - LeetCode solutions are function implementations, not complete programs
# Individual solutions can be tested but may lack imports/types that LeetCode provides

# Build helper utilities
go build -v ./helper
```

### Creating New Solutions

When adding a new LeetCode solution:

1. **Directory structure**: Create a new directory under `problems/[NNNN-problem-name]/`
   - Use 4-digit zero-padded number (e.g., `0001`, `0104`, `1337`)
   - Use kebab-case for problem name (e.g., `two-sum`, `maximum-depth-of-binary-tree`)
   - Example: `problems/0001-two-sum/`

2. **File naming**:
   - `solution.go` - Primary solution
   - `solution_test.go` - Tests
   - `solution_v2.go` - Alternative approaches (optional)

3. **Package declaration**: Always use `package main`

4. **Tests**: Create `solution_test.go` using table-driven tests pattern:

```go
package main

import "testing"

func TestSolution(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected int
    }{
        {"example 1", []int{1, 2, 3}, 6},
        {"edge case", []int{}, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := solution(tt.input)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}
```

## Code Architecture Notes

### Common Patterns

- **Memoization for DP problems**: Many solutions use map-based memoization with closure functions (see `problems/1137-n-th-tribonacci-number/solution.go`)
- **Stack implementations**: Manual stack using slices with index tracking (see `problems/0150-evaluate-reverse-polish-notation/solution.go`)
- **Binary tree problems**: `TreeNode` struct is typically defined inline in each solution
- **Helper utilities**: Reusable algorithms go in `helper/` package (e.g., fast exponentiation)

### Package Organization

```
leetcode/
├── problems/              # All LeetCode solutions (107 problems)
│   ├── 0001-two-sum/
│   ├── 0104-maximum-depth-of-binary-tree/
│   └── ...
├── helper/               # Shared utility functions with tests
├── migrate.go            # Migration script (historical)
├── fix-packages.go       # Package standardization script (historical)
├── main.go              # Placeholder
├── go.mod
├── README.md
└── CLAUDE.md            # This file
```

### Testing Philosophy

Not all solutions have tests yet. This is being improved as part of interview preparation. When tests exist, they follow Go's table-driven testing conventions.
