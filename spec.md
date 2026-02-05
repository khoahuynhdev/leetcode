# Daily LeetCode Problem Workflow Specification

## Overview

A GitHub Actions workflow that automatically fetches the daily problem from LeetCode, uses Claude (Anthropic API) to generate a dsa-coach style analysis and Go solution, then creates a PR with all artifacts following the repository's naming conventions.

## Trigger Configuration

The workflow supports two trigger methods:

The **scheduled trigger** runs daily at 00:30 UTC (shortly after LeetCode's daily challenge resets at midnight UTC). The **manual trigger** (workflow_dispatch) allows on-demand execution with an optional input parameter for a specific LeetCode problem URL. When a problem URL is provided, the workflow processes that problem instead of the daily challenge.

All dates in branch names and PR titles use UTC timezone.

## Problem Fetching

### LeetCode GraphQL API

The workflow queries LeetCode's GraphQL API endpoint to fetch problem data. For the daily challenge, it queries the `activeDailyCodingChallengeQuestion` field. For manual URL inputs, it extracts the problem slug and queries `question(titleSlug: $slug)`.

Required data to extract: problem number (frontendQuestionId), title, titleSlug, difficulty, topicTags, acceptance rate (acRate), content (problem description with examples), and isPaidOnly flag.

### Error Handling

On API failure, implement exponential backoff with up to 3 retry attempts. Initial delay of 2 seconds, doubling on each retry. If all retries fail, the workflow fails with a clear error message.

### Premium Problems

When `isPaidOnly` is true, skip the workflow and create a GitHub issue with title "Daily Challenge Skipped: Premium Problem - [Problem Title]" and body containing the problem number, title, and date. No PR is created.

## Claude Integration

### API Configuration

Call the Anthropic API directly using the `ANTHROPIC_API_KEY` repository secret. Use the Claude Sonnet model (`claude-sonnet-4-20250514` or latest stable sonnet).

### Prompt Structure

The prompt to Claude should include: the full problem description, constraints, examples, difficulty level, topic tags, and acceptance rate. Instruct Claude to act as a DSA coach and produce output in two parts: a progressive hints analysis (for the analysis.md file) and a complete Go solution with tests.

### Expected Output Format

Claude should return structured output containing: the analysis content (markdown with progressive hints), the solution.go content, and the solution_test.go content. Use a clear delimiter or JSON structure to parse these sections.

## Output Files

### Directory Structure

For new problems: `problems/[NNNN-problem-name]/` where NNNN is zero-padded problem number and name is kebab-case.

For existing problems (duplicates): add files with suffix `_daily_YYYYMMDD` (e.g., `solution_daily_20240115.go`).

### solution.go

Contains the Go solution with `package main`. All required custom types (TreeNode, ListNode, etc.) must be defined inline in the file, not imported from helper package. Include brief comments explaining the approach at the top.

### solution_test.go

Table-driven tests following the repository's testing pattern. Extract ALL examples from the problem description to generate test cases. Use descriptive test names like "example 1", "example 2", etc.

### analysis.md

Structure:

```markdown
# [Problem Number]. [Problem Title]

[LeetCode Link](https://leetcode.com/problems/[slug]/)

**Difficulty:** [Easy/Medium/Hard]
**Topics:** [Tag1], [Tag2], ...
**Acceptance Rate:** [XX.X%]

## Hints

### Hint 1

[First hint - high level pattern recognition]

### Hint 2

[Second hint - more specific approach]

### Hint 3

[Third hint - key insight for optimal solution]

## Approach

[Detailed explanation of the solution approach]

## Complexity Analysis

**Time Complexity:** O(...)
**Space Complexity:** O(...)

## Edge Cases

- [Edge case 1]
- [Edge case 2]
- ...
```

Do not include the full problem statement in analysis.md, only the link.

## Testing

### Test Execution

After generating solution files, run `go test -v ./problems/[NNNN-problem-name]/` to validate the solution against generated tests.

### Test Failure Handling

Create the PR regardless of test results. Include test pass/fail status in the PR description. This is treated as a learning opportunity, not a blocker.

## Pull Request

### Branch Naming

`feat/daily-YYYYMMDD-problem-slug` (e.g., `feat/daily-20240115-two-sum`)

### PR Title

`daily-YYYY-MM-DD-[NNNN-problem-name]` (e.g., `daily-2024-01-15-0001-two-sum`)

### PR Description

```markdown
## Summary

- **Problem:** [Number]. [Title]
- **Difficulty:** [Easy/Medium/Hard]
- **Topics:** [Tag1], [Tag2], ...
- **Link:** https://leetcode.com/problems/[slug]/

## Analysis Summary

[Brief 2-3 sentence summary of the approach and key insights from the analysis]

## Test Results

✅ All tests passed / ❌ X of Y tests failed

---

🤖 Generated with Claude (Anthropic API) using dsa-coach analysis
```

### PR Configuration

Create as a normal PR (not draft), even if tests fail. No labels, assignees, or reviewers added automatically (minimal automation).

## Workflow File Structure

The workflow should be defined in `.github/workflows/daily-leetcode.yml` with:

1. Trigger configuration (schedule + workflow_dispatch with optional problem_url input)
2. Job to fetch problem data from LeetCode API
3. Job to call Anthropic API and generate files
4. Job to run tests and capture results
5. Job to create branch, commit files, and open PR

Use repository secrets for `CLAUDE_CODE_OAUTH_TOKEN` and the default `GITHUB_TOKEN` for PR creation.

## Environment Variables and Secrets

Required secrets:

- `CLAUDE_CODE_OAUTH_TOKEN` - Anthropic API key for Claude access

The workflow uses `GITHUB_TOKEN` (automatically provided) for git operations and PR creation.
