---
title: Submit LeetCode Solution via Script and GitHub Action
status: accepted
created: 2026-02-09
author: khoahd
tags: [leetcode, github-actions, automation, api]
files:
  new:
    - scripts/submit.sh
    - .github/workflows/submit-leetcode.yml
  modified:
    - .gitignore
secrets:
  - LEETCODE_SESSION
  - LEETCODE_CSRF
---

# Spec: Submit LeetCode Solution via Script and GitHub Action

## Overview

A bash script (`scripts/submit.sh`) that submits a Go solution to LeetCode's submission API, plus a GitHub Actions workflow (`submit-leetcode.yml`) that wraps it with `workflow_dispatch`. The script is the single source of truth — the workflow simply calls it with the right environment variables and input.

## Authentication

The script reads two environment variables: `LEETCODE_SESSION` and `LEETCODE_CSRF`. In GitHub Actions these come from repository secrets. Locally, the user exports them in their shell profile or sources a `.env` file.

Both values are extracted manually from a logged-in browser session (Application > Cookies > leetcode.com). They remain valid for approximately 1-2 weeks. When they expire, the user updates the secrets/env vars manually. There is no automated login flow — LeetCode's reCAPTCHA makes that unreliable.

The script validates that both variables are set and non-empty before doing anything else, and exits with a clear error message if either is missing.

## Script Interface

```
scripts/submit.sh <problem> [--file <filename>]
```

The `<problem>` argument accepts two formats. If it looks like a number (e.g. `1382`), the script searches `problems/` for a directory matching that prefix (zero-padded to 4 digits). If it looks like a path (e.g. `problems/1382-balance-a-binary-search-tree/` or just `1382-balance-a-binary-search-tree`), the script uses it directly.

The optional `--file` flag specifies which solution file to submit (e.g. `--file solution_v2.go`). Defaults to `solution.go`.

### Exit codes

- 0: Accepted
- 1: Non-Accepted verdict (Wrong Answer, TLE, Runtime Error, etc.)
- 2: Usage/argument error (bad input, missing directory, missing file)
- 3: Authentication error (missing/expired cookies)
- 4: API error (network failure, unexpected response)

## Submission Flow

The script performs these steps in order:

### Step 1: Resolve the problem directory

Parse the `<problem>` argument. If numeric, glob for `problems/<zero-padded>-*/` and verify exactly one match exists. Resolve to the full directory path. Verify the solution file (default `solution.go` or the `--file` override) exists within it.

### Step 2: Derive the slug

Extract the slug from the directory name by stripping the leading `NNNN-` prefix. For example, `1382-balance-a-binary-search-tree` becomes `balance-a-binary-search-tree`.

### Step 3: Fetch the backend question ID

The submit endpoint requires the internal `questionId`, not the frontend number. The script maintains a local cache file at `.leetcode-cache/problems.json` (gitignored) mapping `slug -> questionId`. If the slug is found in the cache, use it. Otherwise, query the LeetCode GraphQL API:

```graphql
query questionData {
  question(titleSlug: "<slug>") {
    questionId
    questionFrontendId
  }
}
```

This reuses the same `graphql_request` helper pattern from `fetch-problem.sh` (with `Referer` and `Origin` headers, inline slug instead of GraphQL variables, and retry logic). Write the result into the cache file for future use.

### Step 4: Read the solution file

Read the contents of the solution file into a variable. This becomes the `typed_code` payload field.

### Step 5: Submit

`POST https://leetcode.com/problems/<slug>/submit/` with:

```json
{
  "lang": "golang",
  "question_id": "<questionId>",
  "typed_code": "<solution file contents>"
}
```

Required headers:

```
Content-Type: application/json
X-Requested-With: XMLHttpRequest
X-CSRFToken: $LEETCODE_CSRF
Referer: https://leetcode.com/problems/<slug>/
Origin: https://leetcode.com
Cookie: LEETCODE_SESSION=$LEETCODE_SESSION; csrftoken=$LEETCODE_CSRF
```

The response returns `{"submission_id": 123456}`. Extract the `submission_id`.

### Step 6: Poll for results

`GET https://leetcode.com/submissions/detail/<submission_id>/check/` with the same cookie and CSRF headers.

Poll every 2 seconds, up to a maximum of 30 seconds (15 attempts). On each poll, check the `state` field:

- `PENDING` or `STARTED`: continue polling
- `SUCCESS`: the result is ready, proceed to step 7

If polling times out (state never reaches `SUCCESS`), exit with code 4 and a timeout message.

### Step 7: Report results

Parse the final check response. The key fields are `status_msg` (e.g. "Accepted", "Wrong Answer"), `total_correct` / `total_testcases`, `runtime_percentile`, and `memory_percentile`.

Print a summary to stdout:

```
Problem:  1382. Balance a Binary Search Tree
Verdict:  Accepted
Tests:    95/95 passed
Runtime:  faster than 87.50% (12ms)
Memory:   less than 62.30% (8.2MB)
```

For non-Accepted verdicts, also print `last_testcase` (the failing input) and `expected_output` / `code_output` if available in the response, then exit with code 1.

For Accepted verdicts, exit with code 0.

## ID Cache

The cache lives at `.leetcode-cache/problems.json` in the repo root. The `.leetcode-cache/` directory is added to `.gitignore`. The file is a simple JSON object mapping slugs to their backend question IDs:

```json
{
  "two-sum": "1",
  "balance-a-binary-search-tree": "1490"
}
```

The script reads and writes this file atomically (write to a temp file, then `mv`). If the file doesn't exist or is malformed, the script creates a fresh one. No expiration — problem IDs are permanent.

## GitHub Actions Workflow

File: `.github/workflows/submit-leetcode.yml`

```yaml
name: Submit LeetCode Solution

on:
  workflow_dispatch:
    inputs:
      problem_number:
        description: "Problem number (e.g. 1382)"
        required: true
        type: string
      solution_file:
        description: "Solution file to submit (default: solution.go)"
        required: false
        default: "solution.go"
        type: string
```

The workflow has a single job with these steps:

1. Checkout the repository.
2. Run `scripts/submit.sh "${{ inputs.problem_number }}" --file "${{ inputs.solution_file }}"` with `LEETCODE_SESSION` and `LEETCODE_CSRF` injected from repository secrets.

That's it. The script handles everything. The workflow is intentionally thin — it's just an environment wrapper.

### Required repository secrets

- `LEETCODE_SESSION`: the `LEETCODE_SESSION` cookie value
- `LEETCODE_CSRF`: the `csrftoken` cookie value

## File changes summary

New files:

- `scripts/submit.sh` — the submission script (~150-200 lines of bash)
- `.github/workflows/submit-leetcode.yml` — the dispatch workflow (~30 lines)

Modified files:

- `.gitignore` — add `.leetcode-cache/` entry

## Constraints and edge cases

- The script only supports `golang` as the language. There is no need to make it language-configurable since this repo is Go-only.

- If the problem directory contains no solution file (or the `--file` target doesn't exist), the script fails with exit code 2 before making any API calls.

- If the GraphQL query for the question ID returns no results (e.g. the slug is wrong), the script fails with exit code 4 and suggests checking the directory name.

- The script never modifies any solution files or commits anything. It is purely a read-and-submit tool.
