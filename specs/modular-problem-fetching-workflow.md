# Specification: Modular LeetCode Problem Fetching Workflow

**Version:** 1.0
**Date:** 2026-02-12
**Status:** Draft - Ready for Implementation

---

## 1. Executive Summary

### 1.1 Problem Statement

The current `daily-leetcode.yml` workflow fetches LeetCode problem data on-the-fly for every execution and stores it temporarily in `problem.json`. This approach has several limitations:

- Problem data is not persisted in the repository
- Problem metadata is lost after the workflow completes
- No way to review problem details without running the workflow
- Difficult to manually trigger problem fetching without solution generation
- Duplicate API calls when re-running failed workflows

### 1.2 Proposed Solution

Refactor the problem-fetching logic into a reusable, modular architecture:

1. **Create a composite action** (`.github/actions/fetch-leetcode-problem/action.yml`) that:
   - Fetches problem data from LeetCode API
   - Converts HTML content to clean markdown
   - Stores problem metadata in `problem.md` with YAML frontmatter
   - Is idempotent (skips fetch if `problem.md` already exists)
   - Outputs all necessary metadata for downstream workflows

2. **Create supporting scripts**:
   - `scripts/convert-problem-json-to-md.py` - Converts API JSON to markdown format
   - `scripts/parse-problem-md.py` - Extracts metadata from `problem.md` for workflows

3. **Modify `daily-leetcode.yml`** to:
   - Use the new composite action
   - Read from `problem.md` when it exists
   - Maintain backward compatibility with existing logic

4. **Create a manual workflow** (`.github/workflows/fetch-problem.yml`):
   - Allows ad-hoc problem fetching
   - Simple wrapper around the composite action
   - Commits `problem.md` to repository

### 1.3 Benefits

- **Persistence**: Problem data stored in git history
- **Transparency**: Problem details visible in PRs and repository
- **Efficiency**: Avoid redundant API calls for duplicate runs
- **Modularity**: Reusable action for multiple workflows
- **Maintainability**: Clear separation of concerns
- **Auditability**: Track when problems were fetched and if they changed over time

---

## 2. Architecture Design

### 2.1 Component Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                     GitHub Actions Workflows                    │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌────────────────────────┐      ┌─────────────────────────┐    │
│  │ daily-leetcode.yml     │      │ fetch-problem.yml       │    │
│  │ (scheduled/manual)     │      │ (manual only)           │    │
│  └───────────┬────────────┘      └────────────┬────────────┘    │
│              │                                │                 │
│              └────────────┬───────────────────┘                 │
│                           │ uses                                │
│                           ▼                                     │
│              ┌─────────────────────────────┐                    │
│              │  Composite Action           │                    │
│              │  fetch-leetcode-problem     │                    │
│              │  (.github/actions/...)      │                    │
│              └────────────┬────────────────┘                    │
│                           │ calls                               │
│                           ▼                                     │
└───────────────────────────┼─────────────────────────────────────┘
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
        ▼                   ▼                   ▼
┌──────────────┐  ┌──────────────────┐  ┌──────────────────┐
│fetch-problem │  │convert-problem-  │  │parse-problem-    │
│.sh (existing)│  │json-to-md.py     │  │md.py (new)       │
│              │  │(new)             │  │                  │
└──────┬───────┘  └────────┬─────────┘  └────────┬─────────┘
       │                   │                     │
       │ outputs           │ creates             │ reads
       ▼                   ▼                     ▼
┌──────────────┐  ┌──────────────────┐  ┌──────────────────┐
│ problem.json │  │ problem.md       │  │ problem.md       │
│ (temporary)  │─>│ (permanent)      │<─│ (extract data)   │
└──────────────┘  └──────────────────┘  └──────────────────┘
```

### 2.2 Data Flow

#### Flow 1: Daily Challenge (problem.md doesn't exist)

```
1. daily-leetcode.yml runs (scheduled or manual)
2. Composite action invoked
3. fetch-problem.sh fetches from API → problem.json (temp)
4. convert-problem-json-to-md.py converts → problem.md
5. Composite action outputs metadata
6. build-prompt.py reads problem.json (backward compat)
7. Claude generates solution files
8. Git commits problem.md + solution files
9. PR created
```

#### Flow 2: Daily Challenge (problem.md exists)

```
1. daily-leetcode.yml runs
2. Composite action invoked
3. Action detects existing problem.md → skip API call
4. Action outputs metadata from problem.md
5. parse-problem-md.py extracts data → problem.json (temp)
6. build-prompt.py reads problem.json
7. Claude generates solution files (with _daily_YYYYMMDD suffix)
8. Git commits only solution files (problem.md unchanged)
9. PR created
```

#### Flow 3: Manual Problem Fetch

```
1. fetch-problem.yml triggered with problem URL
2. Composite action invoked with force_refetch=false
3. fetch-problem.sh fetches → problem.json
4. convert-problem-json-to-md.py → problem.md
5. Git commits problem.md
6. No PR created (just a commit)
```

### 2.3 File System Structure

```
leetcode/
├── .github/
│   ├── actions/
│   │   └── fetch-leetcode-problem/
│   │       └── action.yml                    # New composite action
│   └── workflows/
│       ├── daily-leetcode.yml                # Modified
│       ├── fetch-problem.yml                 # New manual workflow
│       └── submit-leetcode.yml               # Unchanged
├── problems/
│   └── 0001-two-sum/
│       ├── problem.md                        # New persistent file
│       ├── solution.go
│       ├── solution_test.go
│       └── analysis.md
├── scripts/
│   ├── fetch-problem.sh                      # Unchanged (reused)
│   ├── build-prompt.py                       # Unchanged
│   ├── convert-problem-json-to-md.py         # New
│   ├── parse-problem-md.py                   # New
│   └── prompt-template.md                    # Unchanged
└── specs/
    └── modular-problem-fetching-workflow.md  # This document
```

---

## 3. File Specifications

### 3.1 problem.md Format

#### 3.1.1 Structure

```markdown
---
number: "1"
frontend_id: "1"
title: "Two Sum"
slug: "two-sum"
difficulty: "Easy"
topics:
  - "Array"
  - "Hash Table"
acceptance_rate: 49.2
is_premium: false
created_at: "2026-02-12T00:30:15Z"
fetched_at: "2026-02-12T00:30:15Z"
link: "https://leetcode.com/problems/two-sum/"
date: "2026-02-12"
---

# 1. Two Sum

## Description

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

## Examples

### Example 1:

**Input:** nums = [2,7,11,15], target = 9
**Output:** [0,1]
**Explanation:** Because nums[0] + nums[1] == 9, we return [0, 1].

### Example 2:

**Input:** nums = [3,2,4], target = 6
**Output:** [1,2]

### Example 3:

**Input:** nums = [3,3], target = 6
**Output:** [0,1]

## Constraints

- 2 <= nums.length <= 10^4
- -10^9 <= nums[i] <= 10^9
- -10^9 <= target <= 10^9
- Only one valid answer exists.

## Follow-up

Can you come up with an algorithm that is less than O(n^2) time complexity?
```

#### 3.1.2 Field Definitions

**YAML Frontmatter Fields:**

| Field             | Type                | Description                              | Example                                    |
| ----------------- | ------------------- | ---------------------------------------- | ------------------------------------------ |
| `number`          | string              | Zero-padded problem number (4 digits)    | `"0001"`                                   |
| `frontend_id`     | string              | LeetCode's frontend ID (not zero-padded) | `"1"`                                      |
| `title`           | string              | Problem title                            | `"Two Sum"`                                |
| `slug`            | string              | URL slug                                 | `"two-sum"`                                |
| `difficulty`      | string              | Easy, Medium, or Hard                    | `"Easy"`                                   |
| `topics`          | array[string]       | List of topic tags                       | `["Array", "Hash Table"]`                  |
| `acceptance_rate` | number              | Acceptance percentage                    | `49.2`                                     |
| `is_premium`      | boolean             | Whether problem requires premium         | `false`                                    |
| `created_at`      | string (ISO 8601)   | First time fetched                       | `"2026-02-12T00:30:15Z"`                   |
| `fetched_at`      | string (ISO 8601)   | Last time fetched                        | `"2026-02-12T00:30:15Z"`                   |
| `link`            | string              | Full LeetCode URL                        | `"https://leetcode.com/problems/two-sum/"` |
| `date`            | string (YYYY-MM-DD) | Date from daily challenge or fetch date  | `"2026-02-12"`                             |

**Markdown Body:**

- H1 heading with number and title
- H2 sections: Description, Examples, Constraints, Follow-up (if applicable)
- Clean markdown converted from LeetCode's HTML content
- No inline HTML tags
- Code formatting using backticks
- Proper list formatting

#### 3.1.3 HTML to Markdown Conversion Rules

The `convert-problem-json-to-md.py` script must convert HTML to markdown using these rules:

| HTML                     | Markdown           |
| ------------------------ | ------------------ |
| `<p>text</p>`            | `text\n\n`         |
| `<strong>text</strong>`  | `**text**`         |
| `<em>text</em>`          | `*text*`           |
| `<code>text</code>`      | `` `text` ``       |
| `<pre>code</pre>`        | ` ```\ncode\n``` ` |
| `<ul><li>item</li></ul>` | `- item\n`         |
| `<ol><li>item</li></ol>` | `1. item\n`        |
| `<sup>text</sup>`        | `^text`            |
| `&nbsp;`                 | ` ` (space)        |
| `&lt;`                   | `<`                |
| `&gt;`                   | `>`                |
| `&amp;`                  | `&`                |

**Special handling:**

- Remove all `<div>` and `<span>` tags (keep content)
- Convert `<pre>` blocks to fenced code blocks
- Preserve line breaks in examples
- Clean up excessive whitespace (max 2 consecutive newlines)

---

### 3.2 Composite Action Specification

**File:** `.github/actions/fetch-leetcode-problem/action.yml`

```yaml
name: "Fetch LeetCode Problem"
description: "Fetches a LeetCode problem and stores it as problem.md with metadata"
author: "khoahuynhdev"

inputs:
  problem_url:
    description: "LeetCode problem URL (leave empty for daily challenge)"
    required: false
    default: ""
  force_refetch:
    description: "Force re-fetch even if problem.md exists (updates fetched_at)"
    required: false
    default: "false"

outputs:
  # Essential outputs
  dir_name:
    description: "Problem directory name (e.g., 0001-two-sum)"
    value: ${{ steps.metadata.outputs.dir_name }}
  is_premium:
    description: "Whether the problem is premium-only (true/false)"
    value: ${{ steps.metadata.outputs.is_premium }}
  problem_md_path:
    description: "Path to problem.md file"
    value: ${{ steps.metadata.outputs.problem_md_path }}

  # Additional metadata outputs (for backward compatibility)
  number:
    description: "Zero-padded problem number (4 digits)"
    value: ${{ steps.metadata.outputs.number }}
  frontend_id:
    description: "LeetCode frontend ID (not zero-padded)"
    value: ${{ steps.metadata.outputs.frontend_id }}
  title:
    description: "Problem title"
    value: ${{ steps.metadata.outputs.title }}
  slug:
    description: "Problem slug"
    value: ${{ steps.metadata.outputs.slug }}
  difficulty:
    description: "Problem difficulty (Easy/Medium/Hard)"
    value: ${{ steps.metadata.outputs.difficulty }}
  topics:
    description: "Comma-separated list of topics"
    value: ${{ steps.metadata.outputs.topics }}
  ac_rate:
    description: "Acceptance rate percentage"
    value: ${{ steps.metadata.outputs.ac_rate }}
  date:
    description: "Problem date (YYYY-MM-DD)"
    value: ${{ steps.metadata.outputs.date }}
  date_compact:
    description: "Problem date (YYYYMMDD)"
    value: ${{ steps.metadata.outputs.date_compact }}

  # Status outputs
  already_exists:
    description: "Whether problem.md already existed (true/false)"
    value: ${{ steps.check.outputs.exists }}
  was_fetched:
    description: "Whether API was called (true/false)"
    value: ${{ steps.fetch.outputs.fetched }}

runs:
  using: "composite"
  steps:
    - name: Check if problem.md already exists
      id: check
      shell: bash
      run: |
        # Implementation in Section 4.1

    - name: Fetch problem from LeetCode API
      id: fetch
      if: steps.check.outputs.exists != 'true' || inputs.force_refetch == 'true'
      shell: bash
      run: |
        # Implementation in Section 4.2

    - name: Convert JSON to Markdown
      id: convert
      if: steps.fetch.outputs.fetched == 'true'
      shell: bash
      run: |
        # Implementation in Section 4.3

    - name: Extract and output metadata
      id: metadata
      shell: bash
      run: |
        # Implementation in Section 4.4
```

**Key Design Decisions:**

1. **Idempotency**: By default, skip fetch if `problem.md` exists
2. **Force refetch**: Allow updating existing `problem.md` via input flag
3. **Comprehensive outputs**: Provide all fields needed by `daily-leetcode.yml`
4. **Status tracking**: Output `already_exists` and `was_fetched` for debugging
5. **No git operations**: Action only creates files; caller decides commit strategy

---

### 3.3 Script Specifications

#### 3.3.1 convert-problem-json-to-md.py

**File:** `scripts/convert-problem-json-to-md.py`

**Purpose:** Convert `problem.json` to `problem.md` format

**Usage:**

```bash
python3 scripts/convert-problem-json-to-md.py problem.json problems/0001-two-sum/problem.md
```

**Arguments:**

1. Input JSON file path (default: `problem.json`)
2. Output markdown file path (required)

**Behavior:**

1. Read JSON file
2. Extract metadata fields
3. Convert HTML `content` field to markdown
4. Generate YAML frontmatter
5. Write to output file
6. Exit code 0 on success, 1 on error

**Dependencies:**

- Standard library: `json`, `sys`, `os`, `datetime`
- External: `html2text` (for HTML conversion)
  - Install: `pip3 install html2text`
  - Fallback: Use regex if html2text not available (basic conversion)

**Implementation Notes:**

- Set `created_at` to current timestamp if creating new file
- Preserve `created_at` from existing file if updating (read first)
- Always update `fetched_at` to current timestamp
- Round `acRate` to 1 decimal place
- Zero-pad `questionFrontendId` to 4 digits for `number` field
- Use UTC timezone for all timestamps

#### 3.3.2 parse-problem-md.py

**File:** `scripts/parse-problem-md.py`

**Purpose:** Extract metadata from `problem.md` and output as JSON (for `build-prompt.py`)

**Usage:**

```bash
python3 scripts/parse-problem-md.py problems/0001-two-sum/problem.md > problem.json
```

**Arguments:**

1. Input markdown file path (required)

**Behavior:**

1. Read markdown file
2. Extract YAML frontmatter
3. Parse markdown body
4. Output JSON to stdout matching `fetch-problem.sh` format
5. Exit code 0 on success, 1 on error

**Output Format:**

```json
{
  "date": "2026-02-12",
  "link": "/problems/two-sum/",
  "questionId": "1",
  "questionFrontendId": "1",
  "title": "Two Sum",
  "titleSlug": "two-sum",
  "difficulty": "Easy",
  "isPaidOnly": false,
  "acRate": 0.492,
  "topicTags": ["Array", "Hash Table"],
  "content": "<p>Given an array...</p>"
}
```

**Implementation Notes:**

- Parse YAML frontmatter using `yaml` library or regex fallback
- Convert `acceptance_rate` (49.2) to `acRate` (0.492)
- Convert `is_premium` (boolean) to `isPaidOnly` (boolean)
- Convert markdown body back to HTML for `content` field (or keep as markdown if `build-prompt.py` is updated)
- Handle missing fields gracefully with sensible defaults

**Dependencies:**

- Standard library: `json`, `sys`, `re`
- Optional: `pyyaml` (for robust YAML parsing)
  - Install: `pip3 install pyyaml`
  - Fallback: Use regex to extract frontmatter

---

### 3.4 Manual Fetch Workflow

**File:** `.github/workflows/fetch-problem.yml`

```yaml
name: Fetch LeetCode Problem

on:
  workflow_dispatch:
    inputs:
      problem_url:
        description: "LeetCode problem URL"
        required: true
        type: string
      force_refetch:
        description: "Force re-fetch if problem.md exists"
        required: false
        type: boolean
        default: false

jobs:
  fetch:
    runs-on: ubuntu-latest
    permissions:
      contents: write

    steps:
      - name: Checkout repository
        uses: actions/checkout@v4
        with:
          fetch-depth: 1

      - name: Setup Python
        uses: actions/setup-python@v5
        with:
          python-version: "3.11"

      - name: Install dependencies
        run: |
          pip3 install html2text pyyaml

      - name: Fetch problem
        id: fetch
        uses: ./.github/actions/fetch-leetcode-problem
        with:
          problem_url: ${{ inputs.problem_url }}
          force_refetch: ${{ inputs.force_refetch }}

      - name: Check if premium
        if: steps.fetch.outputs.is_premium == 'true'
        run: |
          echo "::error::Cannot fetch premium problem: ${{ steps.fetch.outputs.title }}"
          exit 1

      - name: Commit problem.md
        if: steps.fetch.outputs.was_fetched == 'true'
        run: |
          git config user.name "github-actions[bot]"
          git config user.email "github-actions[bot]@users.noreply.github.com"
          git add "problems/${{ steps.fetch.outputs.dir_name }}/problem.md"
          git commit -m "feat: add problem metadata for ${{ steps.fetch.outputs.number }}. ${{ steps.fetch.outputs.title }}"
          git push

      - name: Summary
        run: |
          echo "## Problem Fetched Successfully" >> "$GITHUB_STEP_SUMMARY"
          echo "" >> "$GITHUB_STEP_SUMMARY"
          echo "- **Number:** ${{ steps.fetch.outputs.number }}" >> "$GITHUB_STEP_SUMMARY"
          echo "- **Title:** ${{ steps.fetch.outputs.title }}" >> "$GITHUB_STEP_SUMMARY"
          echo "- **Difficulty:** ${{ steps.fetch.outputs.difficulty }}" >> "$GITHUB_STEP_SUMMARY"
          echo "- **Topics:** ${{ steps.fetch.outputs.topics }}" >> "$GITHUB_STEP_SUMMARY"
          echo "- **Link:** [${{ steps.fetch.outputs.slug }}](https://leetcode.com/problems/${{ steps.fetch.outputs.slug }}/)" >> "$GITHUB_STEP_SUMMARY"
          echo "- **File:** \`problems/${{ steps.fetch.outputs.dir_name }}/problem.md\`" >> "$GITHUB_STEP_SUMMARY"
          echo "" >> "$GITHUB_STEP_SUMMARY"
          if [ "${{ steps.fetch.outputs.already_exists }}" == "true" ]; then
            echo "ℹ️ **Note:** Problem.md already existed. No API call was made." >> "$GITHUB_STEP_SUMMARY"
          fi
```

**Key Features:**

- Simple manual trigger for ad-hoc problem fetching
- Validates premium problems
- Commits only if new data was fetched
- Shows summary in workflow UI
- No solution generation (just metadata storage)

---

## 4. Implementation Plan

### 4.1 Phase 1: Create Supporting Scripts

#### Task 1.1: Create convert-problem-json-to-md.py

**File:** `scripts/convert-problem-json-to-md.py`

**Implementation:**

````python
#!/usr/bin/env python3
"""
Convert problem.json (from fetch-problem.sh) to problem.md format.

Usage:
    python3 scripts/convert-problem-json-to-md.py problem.json problems/0001-two-sum/problem.md
"""
import json
import sys
import os
from datetime import datetime, timezone
from pathlib import Path

try:
    import html2text
    HTML2TEXT_AVAILABLE = True
except ImportError:
    HTML2TEXT_AVAILABLE = False
    print("Warning: html2text not installed. Using basic HTML conversion.", file=sys.stderr)


def basic_html_to_markdown(html):
    """Fallback HTML to markdown conversion using regex."""
    import re

    # Remove HTML comments
    text = re.sub(r'<!--.*?-->', '', html, flags=re.DOTALL)

    # Convert common tags
    text = re.sub(r'<strong>(.*?)</strong>', r'**\1**', text)
    text = re.sub(r'<b>(.*?)</b>', r'**\1**', text)
    text = re.sub(r'<em>(.*?)</em>', r'*\1*', text)
    text = re.sub(r'<i>(.*?)</i>', r'*\1*', text)
    text = re.sub(r'<code>(.*?)</code>', r'`\1`', text)

    # Convert pre blocks to code blocks
    text = re.sub(r'<pre>(.*?)</pre>', r'```\n\1\n```', text, flags=re.DOTALL)

    # Convert paragraphs
    text = re.sub(r'<p>(.*?)</p>', r'\1\n\n', text, flags=re.DOTALL)

    # Convert lists
    text = re.sub(r'<ul>(.*?)</ul>', lambda m: '\n'.join(f'- {li}' for li in re.findall(r'<li>(.*?)</li>', m.group(1))), text, flags=re.DOTALL)
    text = re.sub(r'<ol>(.*?)</ol>', lambda m: '\n'.join(f'{i+1}. {li}' for i, li in enumerate(re.findall(r'<li>(.*?)</li>', m.group(1)))), text, flags=re.DOTALL)

    # Remove remaining HTML tags
    text = re.sub(r'<[^>]+>', '', text)

    # Decode HTML entities
    text = text.replace('&nbsp;', ' ')
    text = text.replace('&lt;', '<')
    text = text.replace('&gt;', '>')
    text = text.replace('&amp;', '&')
    text = text.replace('&quot;', '"')

    # Clean up whitespace
    text = re.sub(r'\n{3,}', '\n\n', text)
    text = text.strip()

    return text


def html_to_markdown(html):
    """Convert HTML to clean markdown."""
    if HTML2TEXT_AVAILABLE:
        h = html2text.HTML2Text()
        h.body_width = 0  # Don't wrap lines
        h.ignore_links = False
        h.ignore_images = False
        h.ignore_emphasis = False
        return h.handle(html).strip()
    else:
        return basic_html_to_markdown(html)


def convert_json_to_md(json_path, md_path):
    """Convert problem.json to problem.md."""
    # Read JSON
    with open(json_path, 'r', encoding='utf-8') as f:
        data = json.load(f)

    # Extract fields
    frontend_id = str(data['questionFrontendId'])
    number = frontend_id.zfill(4)
    title = data['title']
    slug = data['titleSlug']
    difficulty = data['difficulty']
    topics = data.get('topicTags', [])
    ac_rate = round(data['acRate'] * 100, 1)  # Convert 0.492 to 49.2
    is_premium = data['isPaidOnly']
    date = data.get('date', datetime.now(timezone.utc).strftime('%Y-%m-%d'))
    link = f"https://leetcode.com/problems/{slug}/"
    content_html = data['content']

    # Convert HTML to markdown
    content_md = html_to_markdown(content_html)

    # Determine timestamps
    current_time = datetime.now(timezone.utc).isoformat()

    # Check if file exists to preserve created_at
    created_at = current_time
    if os.path.exists(md_path):
        try:
            with open(md_path, 'r', encoding='utf-8') as f:
                existing_content = f.read()
                # Extract existing created_at from YAML frontmatter
                import re
                match = re.search(r'^created_at:\s*"([^"]+)"', existing_content, re.MULTILINE)
                if match:
                    created_at = match.group(1)
        except Exception:
            pass  # Use current time if can't read existing file

    fetched_at = current_time

    # Build YAML frontmatter
    frontmatter = f"""---
number: "{number}"
frontend_id: "{frontend_id}"
title: "{title}"
slug: "{slug}"
difficulty: "{difficulty}"
topics:
{chr(10).join(f'  - "{topic}"' for topic in topics)}
acceptance_rate: {ac_rate}
is_premium: {str(is_premium).lower()}
created_at: "{created_at}"
fetched_at: "{fetched_at}"
link: "{link}"
date: "{date}"
---
"""

    # Build markdown body
    body = f"""# {number}. {title}

{content_md}
"""

    # Write to file
    Path(md_path).parent.mkdir(parents=True, exist_ok=True)
    with open(md_path, 'w', encoding='utf-8') as f:
        f.write(frontmatter + '\n' + body)

    print(f"Created {md_path}", file=sys.stderr)


def main():
    if len(sys.argv) < 3:
        print("Usage: convert-problem-json-to-md.py <input.json> <output.md>", file=sys.stderr)
        sys.exit(1)

    json_path = sys.argv[1]
    md_path = sys.argv[2]

    if not os.path.exists(json_path):
        print(f"Error: Input file not found: {json_path}", file=sys.stderr)
        sys.exit(1)

    convert_json_to_md(json_path, md_path)


if __name__ == '__main__':
    main()
````

**Testing:**

```bash
# Test with existing problem.json
bash scripts/fetch-problem.sh "https://leetcode.com/problems/two-sum/" > problem.json
python3 scripts/convert-problem-json-to-md.py problem.json test-problem.md
cat test-problem.md
```

#### Task 1.2: Create parse-problem-md.py

**File:** `scripts/parse-problem-md.py`

**Implementation:**

```python
#!/usr/bin/env python3
"""
Parse problem.md and output JSON matching fetch-problem.sh format.

Usage:
    python3 scripts/parse-problem-md.py problems/0001-two-sum/problem.md > problem.json
"""
import json
import sys
import os
import re

try:
    import yaml
    YAML_AVAILABLE = True
except ImportError:
    YAML_AVAILABLE = False
    print("Warning: pyyaml not installed. Using regex fallback.", file=sys.stderr)


def parse_yaml_frontmatter_regex(content):
    """Fallback YAML parser using regex."""
    match = re.match(r'^---\n(.*?)\n---', content, re.DOTALL)
    if not match:
        return None, content

    frontmatter_str = match.group(1)
    body = content[match.end():].strip()

    # Parse YAML manually (basic parsing)
    metadata = {}
    current_key = None
    current_list = []

    for line in frontmatter_str.split('\n'):
        # List item
        if line.startswith('  - '):
            item = line[4:].strip().strip('"').strip("'")
            current_list.append(item)
        # Key-value
        elif ': ' in line:
            if current_key and current_list:
                metadata[current_key] = current_list
                current_list = []

            key, value = line.split(': ', 1)
            key = key.strip()
            value = value.strip().strip('"').strip("'")

            # Parse boolean
            if value.lower() == 'true':
                value = True
            elif value.lower() == 'false':
                value = False
            # Parse number
            elif value.replace('.', '', 1).isdigit():
                value = float(value) if '.' in value else int(value)

            current_key = key
            metadata[key] = value

    # Add last list if any
    if current_key and current_list:
        metadata[current_key] = current_list

    return metadata, body


def parse_problem_md(md_path):
    """Parse problem.md and return JSON object."""
    with open(md_path, 'r', encoding='utf-8') as f:
        content = f.read()

    # Parse YAML frontmatter
    if YAML_AVAILABLE:
        match = re.match(r'^---\n(.*?)\n---', content, re.DOTALL)
        if not match:
            print("Error: No YAML frontmatter found", file=sys.stderr)
            sys.exit(1)

        frontmatter_str = match.group(1)
        body = content[match.end():].strip()
        metadata = yaml.safe_load(frontmatter_str)
    else:
        metadata, body = parse_yaml_frontmatter_regex(content)
        if metadata is None:
            print("Error: No YAML frontmatter found", file=sys.stderr)
            sys.exit(1)

    # Convert to fetch-problem.sh format
    output = {
        "date": metadata.get('date', ''),
        "link": f"/problems/{metadata['slug']}/",
        "questionId": metadata.get('number', metadata.get('frontend_id', '')).lstrip('0'),
        "questionFrontendId": metadata.get('frontend_id', ''),
        "title": metadata['title'],
        "titleSlug": metadata['slug'],
        "difficulty": metadata['difficulty'],
        "isPaidOnly": metadata.get('is_premium', False),
        "acRate": metadata.get('acceptance_rate', 0) / 100,  # Convert 49.2 to 0.492
        "topicTags": metadata.get('topics', []),
        "content": body  # Use markdown body as content (build-prompt.py can handle it)
    }

    return output


def main():
    if len(sys.argv) < 2:
        print("Usage: parse-problem-md.py <problem.md>", file=sys.stderr)
        sys.exit(1)

    md_path = sys.argv[1]

    if not os.path.exists(md_path):
        print(f"Error: File not found: {md_path}", file=sys.stderr)
        sys.exit(1)

    result = parse_problem_md(md_path)
    print(json.dumps(result, indent=2))


if __name__ == '__main__':
    main()
```

**Testing:**

```bash
# Test round-trip conversion
python3 scripts/parse-problem-md.py test-problem.md > recovered.json
diff -u problem.json recovered.json
```

---

### 4.2 Phase 2: Create Composite Action

#### Task 2.1: Create action.yml

**File:** `.github/actions/fetch-leetcode-problem/action.yml`

**Implementation:**

```yaml
name: "Fetch LeetCode Problem"
description: "Fetches a LeetCode problem and stores it as problem.md with metadata"
author: "khoahuynhdev"

inputs:
  problem_url:
    description: "LeetCode problem URL (leave empty for daily challenge)"
    required: false
    default: ""
  force_refetch:
    description: "Force re-fetch even if problem.md exists (updates fetched_at)"
    required: false
    default: "false"

outputs:
  dir_name:
    description: "Problem directory name (e.g., 0001-two-sum)"
    value: ${{ steps.metadata.outputs.dir_name }}
  is_premium:
    description: "Whether the problem is premium-only (true/false)"
    value: ${{ steps.metadata.outputs.is_premium }}
  problem_md_path:
    description: "Path to problem.md file"
    value: ${{ steps.metadata.outputs.problem_md_path }}
  number:
    description: "Zero-padded problem number (4 digits)"
    value: ${{ steps.metadata.outputs.number }}
  frontend_id:
    description: "LeetCode frontend ID (not zero-padded)"
    value: ${{ steps.metadata.outputs.frontend_id }}
  title:
    description: "Problem title"
    value: ${{ steps.metadata.outputs.title }}
  slug:
    description: "Problem slug"
    value: ${{ steps.metadata.outputs.slug }}
  difficulty:
    description: "Problem difficulty (Easy/Medium/Hard)"
    value: ${{ steps.metadata.outputs.difficulty }}
  topics:
    description: "Comma-separated list of topics"
    value: ${{ steps.metadata.outputs.topics }}
  ac_rate:
    description: "Acceptance rate percentage"
    value: ${{ steps.metadata.outputs.ac_rate }}
  date:
    description: "Problem date (YYYY-MM-DD)"
    value: ${{ steps.metadata.outputs.date }}
  date_compact:
    description: "Problem date (YYYYMMDD)"
    value: ${{ steps.metadata.outputs.date_compact }}
  already_exists:
    description: "Whether problem.md already existed (true/false)"
    value: ${{ steps.check.outputs.exists }}
  was_fetched:
    description: "Whether API was called (true/false)"
    value: ${{ steps.fetch.outputs.fetched }}

runs:
  using: "composite"
  steps:
    - name: Install Python dependencies
      shell: bash
      run: |
        pip3 install html2text pyyaml 2>/dev/null || echo "Warning: Some Python packages not installed"

    - name: Determine directory name (preliminary)
      id: prelim
      shell: bash
      run: |
        problem_url="${{ inputs.problem_url }}"

        if [ -n "$problem_url" ]; then
          # Extract slug from URL
          slug=$(echo "$problem_url" | sed -E 's|.*/problems/([^/]+).*|\1|')
          echo "slug=$slug" >> "$GITHUB_OUTPUT"
          echo "is_daily=false" >> "$GITHUB_OUTPUT"
        else
          echo "slug=" >> "$GITHUB_OUTPUT"
          echo "is_daily=true" >> "$GITHUB_OUTPUT"
        fi

    - name: Check if problem.md already exists
      id: check
      shell: bash
      run: |
        # For daily challenges, we don't know the slug yet, so check will happen after fetch
        # For specific URLs, we can check now

        if [ -n "${{ steps.prelim.outputs.slug }}" ]; then
          # Find problem directory by slug pattern
          dir_pattern="problems/*-${{ steps.prelim.outputs.slug }}"
          matching_dirs=($(compgen -G "$dir_pattern" || true))

          if [ ${#matching_dirs[@]} -gt 0 ]; then
            dir_name="${matching_dirs[0]#problems/}"
            problem_md="problems/$dir_name/problem.md"

            if [ -f "$problem_md" ]; then
              echo "exists=true" >> "$GITHUB_OUTPUT"
              echo "dir_name=$dir_name" >> "$GITHUB_OUTPUT"
              echo "problem_md_path=$problem_md" >> "$GITHUB_OUTPUT"
              echo "::notice::Problem directory exists: $dir_name"
            else
              echo "exists=false" >> "$GITHUB_OUTPUT"
              echo "dir_name=$dir_name" >> "$GITHUB_OUTPUT"
            fi
          else
            echo "exists=false" >> "$GITHUB_OUTPUT"
            echo "dir_name=" >> "$GITHUB_OUTPUT"
          fi
        else
          # Daily challenge - defer check to after fetch
          echo "exists=false" >> "$GITHUB_OUTPUT"
          echo "dir_name=" >> "$GITHUB_OUTPUT"
        fi

    - name: Fetch problem from LeetCode API
      id: fetch
      if: steps.check.outputs.exists != 'true' || inputs.force_refetch == 'true'
      shell: bash
      run: |
        set +e
        bash scripts/fetch-problem.sh "${{ inputs.problem_url }}" > problem.json 2>fetch.log
        exit_code=$?
        set -e

        cat fetch.log >&2

        if [ $exit_code -eq 1 ]; then
          echo "::error::Failed to fetch problem from LeetCode API after retries"
          exit 1
        fi

        # Check for premium (exit code 2)
        if [ $exit_code -eq 2 ]; then
          echo "::warning::Premium problem detected"
        fi

        echo "fetched=true" >> "$GITHUB_OUTPUT"
        echo "exit_code=$exit_code" >> "$GITHUB_OUTPUT"

    - name: Check if problem already exists (post-fetch for daily)
      id: check_post
      if: steps.fetch.outputs.fetched == 'true' && steps.check.outputs.exists != 'true'
      shell: bash
      run: |
        # Now we have problem.json, determine directory name
        frontend_id=$(jq -r '.questionFrontendId' problem.json)
        padded_id=$(printf "%04d" "$frontend_id")
        slug=$(jq -r '.titleSlug' problem.json)
        dir_name="${padded_id}-${slug}"
        problem_md="problems/$dir_name/problem.md"

        echo "dir_name=$dir_name" >> "$GITHUB_OUTPUT"
        echo "problem_md_path=$problem_md" >> "$GITHUB_OUTPUT"

        if [ -f "$problem_md" ] && [ "${{ inputs.force_refetch }}" != "true" ]; then
          echo "exists=true" >> "$GITHUB_OUTPUT"
          echo "::notice::Problem.md already exists at $problem_md"
        else
          echo "exists=false" >> "$GITHUB_OUTPUT"
        fi

    - name: Convert JSON to Markdown
      id: convert
      if: steps.fetch.outputs.fetched == 'true' && (steps.check_post.outputs.exists != 'true' || inputs.force_refetch == 'true')
      shell: bash
      run: |
        dir_name="${{ steps.check_post.outputs.dir_name }}"
        problem_md="problems/$dir_name/problem.md"

        python3 scripts/convert-problem-json-to-md.py problem.json "$problem_md"

        echo "::notice::Created $problem_md"

    - name: Extract and output metadata
      id: metadata
      shell: bash
      run: |
        # Determine which path to use
        if [ -n "${{ steps.check_post.outputs.dir_name }}" ]; then
          dir_name="${{ steps.check_post.outputs.dir_name }}"
          problem_md="${{ steps.check_post.outputs.problem_md_path }}"
        elif [ -n "${{ steps.check.outputs.dir_name }}" ]; then
          dir_name="${{ steps.check.outputs.dir_name }}"
          problem_md="${{ steps.check.outputs.problem_md_path }}"
        else
          echo "::error::Could not determine problem directory"
          exit 1
        fi

        # Read metadata from problem.md or problem.json
        if [ -f "$problem_md" ]; then
          # Parse from problem.md
          python3 scripts/parse-problem-md.py "$problem_md" > problem.json
        fi

        # Extract all output fields
        frontend_id=$(jq -r '.questionFrontendId' problem.json)
        padded_id=$(printf "%04d" "$frontend_id")
        slug=$(jq -r '.titleSlug' problem.json)

        echo "dir_name=$dir_name" >> "$GITHUB_OUTPUT"
        echo "problem_md_path=$problem_md" >> "$GITHUB_OUTPUT"
        echo "is_premium=$(jq -r '.isPaidOnly' problem.json)" >> "$GITHUB_OUTPUT"
        echo "number=$padded_id" >> "$GITHUB_OUTPUT"
        echo "frontend_id=$frontend_id" >> "$GITHUB_OUTPUT"
        echo "title=$(jq -r '.title' problem.json)" >> "$GITHUB_OUTPUT"
        echo "slug=$slug" >> "$GITHUB_OUTPUT"
        echo "difficulty=$(jq -r '.difficulty' problem.json)" >> "$GITHUB_OUTPUT"
        echo "topics=$(jq -r '[.topicTags[]] | join(", ")' problem.json)" >> "$GITHUB_OUTPUT"
        echo "ac_rate=$(jq -r '.acRate | . * 100 | round / 100' problem.json)" >> "$GITHUB_OUTPUT"
        echo "date=$(jq -r '.date' problem.json)" >> "$GITHUB_OUTPUT"
        echo "date_compact=$(jq -r '.date | gsub("-"; "")' problem.json)" >> "$GITHUB_OUTPUT"
```

**Testing:**

```bash
# Test composite action locally (requires act or manual workflow trigger)
# See Section 5.2 for testing strategy
```

---

### 4.3 Phase 3: Modify Daily Workflow

#### Task 3.1: Update daily-leetcode.yml

**File:** `.github/workflows/daily-leetcode.yml`

**Changes:**

Replace lines 28-63 (Fetch LeetCode problem step) with:

```yaml
- name: Setup Python
  uses: actions/setup-python@v5
  with:
    python-version: "3.11"

- name: Fetch LeetCode problem metadata
  id: fetch
  uses: ./.github/actions/fetch-leetcode-problem
  with:
    problem_url: ${{ inputs.problem_url }}
    force_refetch: false
```

Replace lines 97-102 (Build Claude prompt step) with:

```yaml
- name: Build Claude prompt
  if: steps.fetch.outputs.is_premium != 'true'
  run: |
    # Generate problem.json from problem.md for build-prompt.py
    python3 scripts/parse-problem-md.py "problems/${{ steps.fetch.outputs.dir_name }}/problem.md" > problem.json

    IS_DUPLICATE="${{ steps.dup.outputs.is_duplicate }}" \
    DATE_COMPACT="${{ steps.fetch.outputs.date_compact }}" \
    python3 scripts/build-prompt.py > prompt.txt
```

Update commit step (lines 181-194) to include problem.md:

```yaml
- name: Commit and push solution
  if: steps.fetch.outputs.is_premium != 'true' && steps.existing_pr.outputs.exists != 'true'
  run: |
    dir="problems/${{ steps.fetch.outputs.dir_name }}"
    if [ ! -d "$dir" ] || [ -z "$(ls -A "$dir" 2>/dev/null)" ]; then
      echo "::error::No solution files to commit"
      exit 1
    fi
    git config user.name "github-actions[bot]"
    git config user.email "github-actions[bot]@users.noreply.github.com"
    git remote set-url origin "https://x-access-token:${{ secrets.GITHUB_TOKEN }}@github.com/${{ github.repository }}.git"
    git add "$dir"

    # Check if problem.md was newly created (not just solution files)
    if [ -f "$dir/problem.md" ] && ! git diff --cached --quiet "$dir/problem.md" 2>/dev/null; then
      commit_msg="feat: add solution for ${{ steps.fetch.outputs.number }}. ${{ steps.fetch.outputs.title }}"
    else
      commit_msg="feat: add solution for ${{ steps.fetch.outputs.number }}. ${{ steps.fetch.outputs.title }}"
    fi

    git commit -m "$commit_msg"
    git push -u origin "$BRANCH_NAME"
```

**Full diff summary:**

- Add Python setup step before fetch
- Replace fetch step with composite action call
- Add parse-problem-md.py call before build-prompt.py
- Include problem.md in git add (if created)

---

### 4.4 Phase 4: Create Manual Workflow

#### Task 4.1: Create fetch-problem.yml

**File:** `.github/workflows/fetch-problem.yml`

(Already detailed in Section 3.4)

---

### 4.5 Phase 5: Update build-prompt.py (Optional Enhancement)

#### Task 5.1: Support Markdown Content

**File:** `scripts/build-prompt.py`

**Optional Enhancement:** Modify to accept markdown content directly from `problem.md` instead of HTML.

**Current behavior:** Reads `content` field from JSON (HTML)

**New behavior:** Check if content is HTML or markdown, use as-is if markdown

**Implementation:**

```python
# Add after line 31
def is_html(text):
    """Detect if text is HTML."""
    return bool(re.search(r'<[^>]+>', text))

# Modify line 31
content = data["content"]
if not is_html(content):
    # Content is already markdown, use as-is
    pass
```

**Priority:** Low (can be done in a follow-up PR)

---

## 5. Testing Strategy

### 5.1 Unit Tests

#### Test 5.1.1: convert-problem-json-to-md.py

```bash
# Test with sample problem
bash scripts/fetch-problem.sh "https://leetcode.com/problems/two-sum/" > test-data/problem.json
python3 scripts/convert-problem-json-to-md.py test-data/problem.json test-data/problem.md

# Verify output
cat test-data/problem.md

# Check frontmatter format
head -n 15 test-data/problem.md | grep -E "^(number|title|difficulty):"

# Check timestamps exist
grep -E "^(created_at|fetched_at):" test-data/problem.md
```

**Expected:** Valid markdown file with YAML frontmatter and clean content

#### Test 5.1.2: parse-problem-md.py

```bash
# Round-trip test
python3 scripts/parse-problem-md.py test-data/problem.md > test-data/recovered.json

# Compare key fields
jq '{title, slug, difficulty, isPaidOnly}' test-data/problem.json > original-fields.json
jq '{title, titleSlug, difficulty, isPaidOnly}' test-data/recovered.json > recovered-fields.json

# Verify slugs match
diff -u original-fields.json recovered-fields.json
```

**Expected:** No differences in key fields

#### Test 5.1.3: HTML to Markdown Conversion

````bash
# Test with a complex problem (has code blocks, lists, formatting)
bash scripts/fetch-problem.sh "https://leetcode.com/problems/median-of-two-sorted-arrays/" > test-data/complex-problem.json
python3 scripts/convert-problem-json-to-md.py test-data/complex-problem.json test-data/complex-problem.md

# Verify no HTML tags in body
tail -n +15 test-data/complex-problem.md | grep -E "<[^>]+>" || echo "✓ No HTML tags found"

# Verify code blocks exist
grep -E "^```" test-data/complex-problem.md || echo "✗ Code blocks missing"
````

**Expected:** Clean markdown with no HTML tags

### 5.2 Integration Tests

#### Test 5.2.1: Composite Action (New Problem)

```yaml
# Manual test workflow: .github/workflows/test-composite-action.yml
name: Test Composite Action
on: workflow_dispatch

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-python@v5
        with:
          python-version: "3.11"
      - run: pip3 install html2text pyyaml

      - name: Test fetch new problem
        id: fetch
        uses: ./.github/actions/fetch-leetcode-problem
        with:
          problem_url: "https://leetcode.com/problems/valid-parentheses/"

      - name: Verify outputs
        run: |
          echo "Directory: ${{ steps.fetch.outputs.dir_name }}"
          echo "Is Premium: ${{ steps.fetch.outputs.is_premium }}"
          echo "Already Exists: ${{ steps.fetch.outputs.already_exists }}"
          echo "Was Fetched: ${{ steps.fetch.outputs.was_fetched }}"

          # Verify file exists
          test -f "problems/${{ steps.fetch.outputs.dir_name }}/problem.md"
          echo "✓ problem.md created"
```

**Expected:** Action succeeds, problem.md created, outputs correct

#### Test 5.2.2: Composite Action (Existing Problem)

```yaml
- name: Test fetch existing problem
  id: fetch2
  uses: ./.github/actions/fetch-leetcode-problem
  with:
    problem_url: "https://leetcode.com/problems/valid-parentheses/"

- name: Verify idempotency
  run: |
    test "${{ steps.fetch2.outputs.already_exists }}" == "true"
    test "${{ steps.fetch2.outputs.was_fetched }}" == "false"
    echo "✓ Idempotent behavior confirmed"
```

**Expected:** Action skips fetch, uses existing file

#### Test 5.2.3: Composite Action (Force Refetch)

```yaml
- name: Test force refetch
  id: fetch3
  uses: ./.github/actions/fetch-leetcode-problem
  with:
    problem_url: "https://leetcode.com/problems/valid-parentheses/"
    force_refetch: true

- name: Verify refetch
  run: |
    test "${{ steps.fetch3.outputs.was_fetched }}" == "true"
    echo "✓ Force refetch works"
```

**Expected:** Action refetches, updates fetched_at timestamp

#### Test 5.2.4: Daily Workflow End-to-End

```bash
# Trigger manual run with a test problem
gh workflow run daily-leetcode.yml -f problem_url="https://leetcode.com/problems/climbing-stairs/"

# Wait for completion
gh run watch

# Verify PR created
gh pr list --head "feat/daily-*-climbing-stairs"

# Verify problem.md exists in PR
gh pr view --json files | jq '.files[].path' | grep "problem.md"
```

**Expected:** Full workflow succeeds, PR contains problem.md + solution files

### 5.3 Edge Case Tests

#### Test 5.3.1: Premium Problem

```bash
# Fetch a known premium problem
bash scripts/fetch-problem.sh "https://leetcode.com/problems/design-in-memory-file-system/" > premium.json

# Verify exit code 2
echo $?  # Should be 2

# Verify isPaidOnly is true
jq '.isPaidOnly' premium.json  # Should be true
```

**Expected:** Script exits with code 2, JSON contains isPaidOnly: true

#### Test 5.3.2: Duplicate Problem (Daily Challenge)

```bash
# Manually create a problem directory
mkdir -p problems/0070-climbing-stairs
echo "existing" > problems/0070-climbing-stairs/solution.go

# Run workflow with same problem
gh workflow run daily-leetcode.yml -f problem_url="https://leetcode.com/problems/climbing-stairs/"

# Verify _daily_ suffix used
ls problems/0070-climbing-stairs/solution_daily_*.go
```

**Expected:** Duplicate detected, solution files use \_daily_YYYYMMDD suffix

#### Test 5.3.3: API Failure with Retry

```bash
# Temporarily block LeetCode (requires network manipulation)
# Or test retry logic manually by modifying fetch-problem.sh

# Verify retries happen
bash scripts/fetch-problem.sh 2>&1 | grep "Attempt.*failed"
```

**Expected:** Script retries with exponential backoff

### 5.4 Validation Checklist

Before merging:

- [ ] All unit tests pass
- [ ] Composite action works for new problems
- [ ] Composite action is idempotent (skip existing)
- [ ] Force refetch updates timestamps correctly
- [ ] Daily workflow creates problem.md
- [ ] Daily workflow uses existing problem.md
- [ ] Manual fetch workflow works
- [ ] Premium problems handled correctly
- [ ] Duplicate detection works with new flow
- [ ] No regression in existing functionality
- [ ] Documentation updated (README if needed)

---

## 6. Migration Path

### 6.1 Rollout Strategy

**Phase A: Implementation (1 PR)**

1. Create all new files:
   - `scripts/convert-problem-json-to-md.py`
   - `scripts/parse-problem-md.py`
   - `.github/actions/fetch-leetcode-problem/action.yml`
   - `.github/workflows/fetch-problem.yml`

2. Modify existing files:
   - `.github/workflows/daily-leetcode.yml`

3. Add tests:
   - Create `test-data/` directory with sample inputs
   - Add test workflow (optional)

**Phase B: Testing**

1. Test manually with `fetch-problem.yml`:
   - Fetch a new problem
   - Verify problem.md format
   - Verify round-trip (md → json → prompt)

2. Test daily workflow:
   - Trigger with a test problem URL
   - Verify full flow works
   - Check PR contains problem.md

**Phase C: Production Rollout**

1. Let scheduled daily runs use new flow
2. Monitor for issues
3. Verify problem.md files accumulate in repository

**Phase D: Backfill (Optional - Future Work)**

Create script to backfill problem.md for existing 107 problems:

```bash
#!/bin/bash
# scripts/backfill-problem-md.sh
for dir in problems/*/; do
  if [ ! -f "$dir/problem.md" ]; then
    slug=$(basename "$dir" | sed 's/^[0-9]*-//')
    echo "Backfilling $slug..."
    bash scripts/fetch-problem.sh "https://leetcode.com/problems/$slug/" > problem.json
    python3 scripts/convert-problem-json-to-md.py problem.json "$dir/problem.md"
    sleep 1  # Rate limiting
  fi
done
```

**Note:** Backfill is not required for MVP. Can be done incrementally.

### 6.2 Backward Compatibility

**Guaranteed:**

- Existing workflows continue to work
- `fetch-problem.sh` unchanged
- `build-prompt.py` unchanged (reads problem.json)
- Output fields from composite action match previous step outputs

**Breaking Changes:**

- None (fully backward compatible)

**Deprecation Plan:**

- None (no features deprecated)

### 6.3 Rollback Plan

If critical issues arise:

1. Revert `.github/workflows/daily-leetcode.yml` to previous version
2. Keep composite action and scripts (no harm if unused)
3. Investigate and fix issues
4. Re-deploy when ready

**Risk:** Low - Changes are additive, not replacing existing logic

---

## 7. Examples

### 7.1 Example: Daily Challenge (New Problem)

**Scenario:** Daily challenge on 2026-02-13 is "Container With Most Water" (#11)

**Workflow Execution:**

```
1. Scheduled trigger at 00:30 UTC
2. Composite action called (no problem_url)
3. fetch-problem.sh fetches daily challenge
4. Returns: questionFrontendId=11, slug=container-with-most-water
5. Check: problems/0011-container-with-most-water/problem.md doesn't exist
6. convert-problem-json-to-md.py creates problem.md
7. Outputs: dir_name=0011-container-with-most-water, is_premium=false
8. build-prompt.py generates prompt (using problem.json from step 3)
9. Claude generates analysis.md, solution.go, solution_test.go
10. Git commits all files including problem.md
11. PR created: "daily-2026-02-13-0011-container-with-most-water"
```

**Result:**

```
problems/0011-container-with-most-water/
├── problem.md          # NEW - persisted metadata
├── analysis.md
├── solution.go
└── solution_test.go
```

**PR Description:**

```markdown
## Summary

- **Problem:** 0011. Container With Most Water
- **Difficulty:** Medium
- **Topics:** Array, Two Pointers, Greedy
- **Link:** https://leetcode.com/problems/container-with-most-water/

## Analysis Summary

See `analysis.md` in the PR files for progressive hints and detailed approach.

## Test Results

✅ All tests passed

---

🤖 Generated with Claude (Anthropic API) using dsa-coach analysis
```

### 7.2 Example: Daily Challenge (Duplicate Problem)

**Scenario:** Daily challenge on 2026-03-01 is "Two Sum" (#1), which already exists

**Workflow Execution:**

```
1. Scheduled trigger at 00:30 UTC
2. Composite action called
3. fetch-problem.sh fetches daily challenge
4. Returns: slug=two-sum
5. Check: problems/0001-two-sum/problem.md EXISTS
6. Action skips API call, reads from existing problem.md
7. Outputs: dir_name=0001-two-sum, already_exists=true
8. Workflow detects duplicate (dir exists)
9. Sets IS_DUPLICATE=true, DATE_COMPACT=20260301
10. build-prompt.py generates prompt with _daily_20260301 suffix
11. Claude generates files with suffix
12. Git commits ONLY solution files (problem.md unchanged)
13. PR created
```

**Result:**

```
problems/0001-two-sum/
├── problem.md              # Unchanged
├── analysis.md             # Original
├── solution.go             # Original
├── solution_test.go        # Original
├── solution_daily_20260301.go        # NEW
└── solution_daily_20260301_test.go   # NEW
```

**No redundant API call made!**

### 7.3 Example: Manual Problem Fetch

**Scenario:** User wants to manually fetch "Merge Two Sorted Lists" (#21)

**Steps:**

```bash
# 1. Go to GitHub Actions UI
# 2. Select "Fetch LeetCode Problem" workflow
# 3. Click "Run workflow"
# 4. Enter: https://leetcode.com/problems/merge-two-sorted-lists/
# 5. force_refetch: false
# 6. Run
```

**Workflow Execution:**

```
1. Manual trigger with problem_url
2. Composite action called with URL
3. fetch-problem.sh fetches problem
4. convert-problem-json-to-md.py creates problem.md
5. Git commits problem.md
6. Workflow completes
```

**Result:**

```
problems/0021-merge-two-sorted-lists/
└── problem.md          # Committed to main branch
```

**Output Summary:**

```
## Problem Fetched Successfully

- **Number:** 0021
- **Title:** Merge Two Sorted Lists
- **Difficulty:** Easy
- **Topics:** Linked List, Recursion
- **Link:** [merge-two-sorted-lists](https://leetcode.com/problems/merge-two-sorted-lists/)
- **File:** `problems/0021-merge-two-sorted-lists/problem.md`
```

### 7.4 Example: Premium Problem Handling

**Scenario:** Daily challenge is a premium problem (unlikely but possible)

**Workflow Execution:**

```
1. Composite action fetches problem
2. fetch-problem.sh detects isPaidOnly=true, exits with code 2
3. Action outputs: is_premium=true
4. Daily workflow detects premium
5. Creates GitHub issue
6. Stops (no solution generation)
7. problem.md still created (with is_premium: true in frontmatter)
```

**Result:**

- GitHub issue created: "Daily Challenge Skipped: Premium Problem - {Title}"
- No PR created
- problem.md exists in temp files but not committed (workflow stops)

### 7.5 Example: Force Refetch

**Scenario:** Problem #100 was fetched 6 months ago, user wants to update it

**Steps:**

```bash
# Trigger fetch-problem.yml with force_refetch=true
gh workflow run fetch-problem.yml \
  -f problem_url="https://leetcode.com/problems/same-tree/" \
  -f force_refetch=true
```

**Workflow Execution:**

```
1. Composite action called with force_refetch=true
2. Check: problems/0100-same-tree/problem.md EXISTS
3. Action ignores existing file, fetches from API
4. convert-problem-json-to-md.py updates problem.md
5. Preserves created_at (2025-08-15T...)
6. Updates fetched_at (2026-02-12T...)
7. Git commits updated problem.md
```

**Result:**

```yaml
# problems/0100-same-tree/problem.md
---
number: "0100"
title: "Same Tree"
created_at: "2025-08-15T10:23:45Z" # Original timestamp preserved
fetched_at: "2026-02-12T14:30:00Z" # Updated timestamp
acceptance_rate: 58.7 # Potentially updated (if changed)
# ... rest of frontmatter
---
```

**Use case:** Refresh acceptance rates, detect problem description changes

---

## 8. Acceptance Criteria

### 8.1 Functional Requirements

- [x] Composite action fetches problem data from LeetCode API
- [x] Composite action converts HTML to clean markdown
- [x] Composite action stores data in problem.md with YAML frontmatter
- [x] Composite action is idempotent (skips if file exists)
- [x] Composite action supports force refetch
- [x] Composite action outputs all required metadata
- [x] Manual workflow allows ad-hoc problem fetching
- [x] Daily workflow uses composite action
- [x] Daily workflow reads from existing problem.md when available
- [x] Premium problems are detected and handled
- [x] Duplicate problems reuse existing problem.md
- [x] Timestamps (created_at, fetched_at) are tracked

### 8.2 Quality Requirements

- [ ] All unit tests pass (scripts work correctly)
- [ ] Integration tests pass (workflows work end-to-end)
- [ ] No regression in existing functionality
- [ ] Code follows repository conventions (bash, python, yaml)
- [ ] Error handling is robust (API failures, missing files, etc.)
- [ ] Logging is clear and actionable
- [ ] Documentation is complete and accurate

### 8.3 Performance Requirements

- [ ] API calls reduced for duplicate daily challenges
- [ ] Workflow execution time < 5 minutes (unchanged from current)
- [ ] HTML to markdown conversion completes in < 5 seconds

### 8.4 Security Requirements

- [ ] No secrets exposed in logs
- [ ] No hardcoded credentials
- [ ] File permissions set correctly (no executable markdown files)
- [ ] Input validation for URLs (prevent injection attacks)

---

## 9. Open Questions & Future Work

### 9.1 Open Questions

None - all decisions made.

### 9.2 Future Enhancements

1. **Backfill Script:** Create problem.md for all 107 existing problems
2. **Static Site Generator:** Use problem.md files to generate a browsable website
3. **Problem Diff Tracking:** Detect when LeetCode changes problem statements
4. **Hints Caching:** Store LeetCode's official hints in problem.md
5. **Company Tags:** Add company frequency data to frontmatter
6. **Similar Problems:** Link related problems in frontmatter
7. **Problem Search:** CLI tool to search problems by topic/difficulty
8. **Problem Stats:** Track solve times, submission counts in metadata

### 9.3 Known Limitations

1. **HTML Conversion Quality:** html2text may not perfectly convert all HTML (manual review recommended)
2. **Rate Limiting:** No built-in rate limiting for API calls (LeetCode may throttle)
3. **API Changes:** LeetCode GraphQL API is unofficial and may change without notice
4. **Premium Detection:** Relies on isPaidOnly field; LeetCode may change access model
5. **Timestamp Preservation:** Relies on regex parsing; may break if frontmatter format changes

---

## 10. Appendix

### 10.1 File Size Estimates

| File                            | Lines | Size | Complexity |
| ------------------------------- | ----- | ---- | ---------- |
| `convert-problem-json-to-md.py` | ~150  | 5 KB | Medium     |
| `parse-problem-md.py`           | ~120  | 4 KB | Medium     |
| `action.yml`                    | ~200  | 7 KB | High       |
| `fetch-problem.yml`             | ~60   | 2 KB | Low        |
| `problem.md` (avg)              | ~80   | 3 KB | N/A        |

**Total new code:** ~530 lines, ~18 KB

### 10.2 Dependencies

**Python Packages:**

- `html2text` (HTML to markdown conversion)
- `pyyaml` (YAML parsing, optional)
- `json`, `sys`, `os`, `re`, `datetime` (standard library)

**Bash Tools:**

- `jq` (JSON parsing, already used)
- `curl` (HTTP requests, already used)
- `sed`, `grep`, `awk` (text processing, already available)

**GitHub Actions:**

- `actions/checkout@v4` (already used)
- `actions/setup-python@v5` (new dependency)

### 10.3 Risk Assessment

| Risk                           | Likelihood | Impact | Mitigation                                 |
| ------------------------------ | ---------- | ------ | ------------------------------------------ |
| API changes break fetch        | Medium     | High   | Test regularly, monitor LeetCode changelog |
| HTML conversion errors         | Medium     | Medium | Manual review problem.md, add test cases   |
| Composite action bugs          | Low        | High   | Thorough testing, gradual rollout          |
| File conflicts in git          | Low        | Low    | Atomic commits, proper file locking        |
| Performance degradation        | Very Low   | Low    | Profile workflows, optimize if needed      |
| Premium problem false positive | Very Low   | Medium | Verify isPaidOnly field accuracy           |

**Overall Risk Level:** Low-Medium

### 10.4 References

- [LeetCode GraphQL API (unofficial)](https://github.com/akarsh1995/leetcode-graphql-queries)
- [html2text documentation](https://github.com/Alir3z4/html2text)
- [GitHub Actions: Composite Actions](https://docs.github.com/en/actions/creating-actions/creating-a-composite-action)
- [YAML Frontmatter spec](https://jekyllrb.com/docs/front-matter/)
- [Markdown spec](https://commonmark.org/)

---

## Document History

| Version | Date       | Author      | Changes               |
| ------- | ---------- | ----------- | --------------------- |
| 1.0     | 2026-02-12 | Claude Code | Initial specification |

---

**End of Specification**
