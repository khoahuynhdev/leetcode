#!/usr/bin/env bash
set -euo pipefail

# submit.sh
# Submits a Go solution to LeetCode's submission API.
# Usage:
#   ./scripts/submit.sh <problem> [--file <filename>]
#
#   <problem> can be:
#     - A number (e.g. 1382) — searches problems/ for matching directory
#     - A directory name (e.g. 1382-balance-a-binary-search-tree)
#     - A path (e.g. problems/1382-balance-a-binary-search-tree/)
#
# Exit codes:
#   0 - Accepted
#   1 - Non-Accepted verdict (Wrong Answer, TLE, Runtime Error, etc.)
#   2 - Usage/argument error
#   3 - Authentication error
#   4 - API/network error

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
PROBLEMS_DIR="$REPO_ROOT/problems"

# ---------------------------------------------------------------------------
# Argument parsing
# ---------------------------------------------------------------------------

usage() {
  echo "Usage: $0 <problem> [--file <filename>]" >&2
  echo "  <problem>  Problem number (e.g. 1382) or directory name" >&2
  echo "  --file     Solution file to submit (default: solution.go)" >&2
  exit 2
}

PROBLEM=""
SOLUTION_FILE="solution.go"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --file)
      [[ $# -lt 2 ]] && { echo "Error: --file requires a filename argument." >&2; exit 2; }
      SOLUTION_FILE="$2"
      shift 2
      ;;
    -*)
      echo "Error: Unknown option '$1'" >&2
      usage
      ;;
    *)
      [[ -n "$PROBLEM" ]] && { echo "Error: Unexpected argument '$1'" >&2; usage; }
      PROBLEM="$1"
      shift
      ;;
  esac
done

[[ -z "$PROBLEM" ]] && { echo "Error: Missing <problem> argument." >&2; usage; }

# ---------------------------------------------------------------------------
# Resolve problem directory
# ---------------------------------------------------------------------------

resolve_problem_dir() {
  local input="$1"

  # If it looks like a bare number, glob for it
  if [[ "$input" =~ ^[0-9]+$ ]]; then
    local padded
    padded=$(printf "%04d" "$input")
    local matches=("$PROBLEMS_DIR"/${padded}-*/)

    if [[ ! -d "${matches[0]}" ]]; then
      echo "Error: No problem directory found matching '${padded}-*' in problems/." >&2
      exit 2
    fi
    if [[ ${#matches[@]} -gt 1 ]]; then
      echo "Error: Multiple directories match '${padded}-*':" >&2
      printf "  %s\n" "${matches[@]}" >&2
      exit 2
    fi
    echo "${matches[0]}"
    return
  fi

  # Strip trailing slash and "problems/" prefix if present
  input="${input%/}"
  input="${input#problems/}"

  local dir="$PROBLEMS_DIR/$input"
  if [[ ! -d "$dir" ]]; then
    echo "Error: Directory not found: $dir" >&2
    exit 2
  fi
  echo "$dir"
}

PROBLEM_DIR="$(resolve_problem_dir "$PROBLEM")"
SOLUTION_PATH="$PROBLEM_DIR/$SOLUTION_FILE"

if [[ ! -f "$SOLUTION_PATH" ]]; then
  echo "Error: Solution file not found: $SOLUTION_PATH" >&2
  exit 2
fi

# ---------------------------------------------------------------------------
# Authentication
# ---------------------------------------------------------------------------

if [[ -z "${LEETCODE_SESSION:-}" ]]; then
  echo "Error: LEETCODE_SESSION environment variable is not set." >&2
  echo "Export it from your browser's cookies (Application > Cookies > leetcode.com)." >&2
  exit 3
fi

if [[ -z "${LEETCODE_CSRF:-}" ]]; then
  echo "Error: LEETCODE_CSRF environment variable is not set." >&2
  echo "Export it from your browser's cookies (Application > Cookies > leetcode.com)." >&2
  exit 3
fi

# ---------------------------------------------------------------------------
# Derive slug
# ---------------------------------------------------------------------------

DIR_NAME="$(basename "$PROBLEM_DIR")"
SLUG="${DIR_NAME#[0-9][0-9][0-9][0-9]-}"

echo "Problem:   $DIR_NAME" >&2
echo "File:      $SOLUTION_FILE" >&2
echo "Slug:      $SLUG" >&2
echo "" >&2

# ---------------------------------------------------------------------------
# Minify solution for LeetCode submission
# ---------------------------------------------------------------------------

# LeetCode provides these types in its Go environment. Submitting them
# causes "type redeclared" compile errors.
LEETCODE_TYPES="TreeNode|ListNode|Node"

minify_solution() {
  local code="$1"

  # Strip package declaration
  code=$(echo "$code" | sed '/^package /d')

  # Strip LeetCode-provided struct definitions and preceding doc comments.
  # Handles both single-line (// ...) and block (/* ... */) comments that
  # sit directly above "type <Name> struct {".
  code=$(echo "$code" | awk -v types="$LEETCODE_TYPES" '
    BEGIN { skip = 0; buf_count = 0 }

    # Accumulate comment lines and blank lines that might precede a type definition
    (/^[[:space:]]*(\/\/|\/\*|\*|\*\/)/ || /^[[:space:]]*$/) && skip == 0 {
      buf[buf_count++] = $0
      next
    }

    # If we hit a LeetCode type definition, discard buffered comments and skip the struct body
    /^type[[:space:]]+('"$LEETCODE_TYPES"')[[:space:]]+struct[[:space:]]*\{/ {
      buf_count = 0
      skip = 1
      next
    }

    # While skipping a struct body, wait for the closing brace
    skip == 1 {
      if ($0 ~ /^\}/) skip = 0
      next
    }

    # Flush any buffered lines (they were not before a LeetCode type)
    {
      for (i = 0; i < buf_count; i++) print buf[i]
      buf_count = 0
      print
    }

    END {
      for (i = 0; i < buf_count; i++) print buf[i]
    }
  ')

  # Collapse leading blank lines
  code=$(echo "$code" | sed '/./,$!d')

  echo "$code"
}

# ---------------------------------------------------------------------------
# Helpers
# ---------------------------------------------------------------------------

LEETCODE_GRAPHQL="https://leetcode.com/graphql"
CACHE_DIR="$REPO_ROOT/.leetcode-cache"
CACHE_FILE="$CACHE_DIR/problems.json"
MAX_RETRIES=3
INITIAL_DELAY=2

graphql_request() {
  local query="$1"
  local delay=$INITIAL_DELAY

  for attempt in $(seq 1 $MAX_RETRIES); do
    local response
    local http_code

    response=$(curl -s -w "\n%{http_code}" -X POST "$LEETCODE_GRAPHQL" \
      -H "Content-Type: application/json" \
      -H "Accept: application/json" \
      -H "Referer: https://leetcode.com" \
      -H "Origin: https://leetcode.com" \
      -d "{\"query\": $(echo "$query" | jq -Rs .)}" \
      2>/dev/null) || true

    http_code=$(echo "$response" | tail -1)
    local body
    body=$(echo "$response" | sed '$d')

    if [[ "$http_code" == "200" ]] && echo "$body" | jq -e '.data' >/dev/null 2>&1; then
      echo "$body"
      return 0
    fi

    if [[ $attempt -lt $MAX_RETRIES ]]; then
      echo "Attempt $attempt failed (HTTP $http_code). Retrying in ${delay}s..." >&2
      sleep "$delay"
      delay=$((delay * 2))
    fi
  done

  echo "All $MAX_RETRIES attempts failed." >&2
  return 1
}

# ---------------------------------------------------------------------------
# Fetch question ID (with cache)
# ---------------------------------------------------------------------------

mkdir -p "$CACHE_DIR"
if [[ ! -f "$CACHE_FILE" ]] || ! jq -e '.' "$CACHE_FILE" >/dev/null 2>&1; then
  echo '{}' > "$CACHE_FILE"
fi

QUESTION_ID=$(jq -r --arg slug "$SLUG" '.[$slug] // empty' "$CACHE_FILE")

if [[ -z "$QUESTION_ID" ]]; then
  echo "Fetching question ID for '$SLUG'..." >&2

  query=$(cat <<EOF
    query questionData {
      question(titleSlug: "${SLUG}") {
        questionId
        questionFrontendId
      }
    }
EOF
  )

  response=$(graphql_request "$query") || { echo "Error: Failed to fetch question ID. Check that the slug '$SLUG' is correct." >&2; exit 4; }

  QUESTION_ID=$(echo "$response" | jq -r '.data.question.questionId')
  if [[ -z "$QUESTION_ID" || "$QUESTION_ID" == "null" ]]; then
    echo "Error: No question found for slug '$SLUG'. Check the directory name." >&2
    exit 4
  fi

  # Atomic cache update
  tmp=$(mktemp)
  jq --arg slug "$SLUG" --arg id "$QUESTION_ID" '. + {($slug): $id}' "$CACHE_FILE" > "$tmp" && mv "$tmp" "$CACHE_FILE"
  echo "Cached question ID: $QUESTION_ID" >&2
else
  echo "Question ID (cached): $QUESTION_ID" >&2
fi

# ---------------------------------------------------------------------------
# Read solution and submit
# ---------------------------------------------------------------------------

RAW_CODE=$(<"$SOLUTION_PATH")
TYPED_CODE=$(minify_solution "$RAW_CODE")

echo "Submitting to LeetCode..." >&2

submit_response=$(curl -s -w "\n%{http_code}" -X POST "https://leetcode.com/problems/${SLUG}/submit/" \
  -H "Content-Type: application/json" \
  -H "X-Requested-With: XMLHttpRequest" \
  -H "X-CSRFToken: $LEETCODE_CSRF" \
  -H "Referer: https://leetcode.com/problems/${SLUG}/" \
  -H "Origin: https://leetcode.com" \
  -H "Cookie: LEETCODE_SESSION=$LEETCODE_SESSION; csrftoken=$LEETCODE_CSRF" \
  -d "$(jq -n --arg lang "golang" --arg qid "$QUESTION_ID" --arg code "$TYPED_CODE" \
    '{lang: $lang, question_id: $qid, typed_code: $code}')" \
  2>/dev/null) || true

submit_http=$(echo "$submit_response" | tail -1)
submit_body=$(echo "$submit_response" | sed '$d')

if [[ "$submit_http" != "200" ]]; then
  echo "Error: Submit request failed (HTTP $submit_http)." >&2
  echo "$submit_body" >&2
  exit 4
fi

SUBMISSION_ID=$(echo "$submit_body" | jq -r '.submission_id')
if [[ -z "$SUBMISSION_ID" || "$SUBMISSION_ID" == "null" ]]; then
  echo "Error: No submission_id in response." >&2
  echo "$submit_body" >&2
  exit 4
fi

echo "Submission ID: $SUBMISSION_ID" >&2

# ---------------------------------------------------------------------------
# Poll for results
# ---------------------------------------------------------------------------

MAX_POLLS=15
POLL_INTERVAL=2

for i in $(seq 1 $MAX_POLLS); do
  sleep $POLL_INTERVAL

  check_response=$(curl -s "https://leetcode.com/submissions/detail/${SUBMISSION_ID}/check/" \
    -H "X-Requested-With: XMLHttpRequest" \
    -H "X-CSRFToken: $LEETCODE_CSRF" \
    -H "Referer: https://leetcode.com/problems/${SLUG}/" \
    -H "Origin: https://leetcode.com" \
    -H "Cookie: LEETCODE_SESSION=$LEETCODE_SESSION; csrftoken=$LEETCODE_CSRF" \
    2>/dev/null) || true

  state=$(echo "$check_response" | jq -r '.state')

  if [[ "$state" == "SUCCESS" ]]; then
    break
  fi

  if [[ "$state" != "PENDING" && "$state" != "STARTED" ]]; then
    echo "Error: Unexpected state '$state'." >&2
    echo "$check_response" >&2
    exit 4
  fi

  echo "Polling... ($i/$MAX_POLLS)" >&2
done

if [[ "$state" != "SUCCESS" ]]; then
  echo "Error: Timed out waiting for submission result (${MAX_POLLS}x${POLL_INTERVAL}s)." >&2
  exit 4
fi

# ---------------------------------------------------------------------------
# Report results
# ---------------------------------------------------------------------------

status_msg=$(echo "$check_response" | jq -r '.status_msg')
total_correct=$(echo "$check_response" | jq -r '.total_correct // empty')
total_testcases=$(echo "$check_response" | jq -r '.total_testcases // empty')
runtime=$(echo "$check_response" | jq -r '.status_runtime // empty')
memory=$(echo "$check_response" | jq -r '.status_memory // empty')
runtime_pct=$(echo "$check_response" | jq -r '.runtime_percentile // empty')
memory_pct=$(echo "$check_response" | jq -r '.memory_percentile // empty')

echo ""
echo "Problem:  ${DIR_NAME}"
echo "Verdict:  ${status_msg}"

if [[ -n "$total_correct" && -n "$total_testcases" ]]; then
  echo "Tests:    ${total_correct}/${total_testcases} passed"
fi

if [[ "$status_msg" == "Accepted" ]]; then
  if [[ -n "$runtime_pct" ]]; then
    printf "Runtime:  faster than %.2f%% (%s)\n" "$runtime_pct" "$runtime"
  fi
  if [[ -n "$memory_pct" ]]; then
    printf "Memory:   less than %.2f%% (%s)\n" "$memory_pct" "$memory"
  fi
  exit 0
else
  last_testcase=$(echo "$check_response" | jq -r '.last_testcase // empty')
  expected=$(echo "$check_response" | jq -r '.expected_output // empty')
  code_output=$(echo "$check_response" | jq -r '.code_output // empty')

  if [[ -n "$last_testcase" ]]; then
    echo ""
    echo "Failing input:"
    echo "$last_testcase"
  fi
  if [[ -n "$expected" ]]; then
    echo "Expected: $expected"
  fi
  if [[ -n "$code_output" ]]; then
    echo "Got:      $code_output"
  fi
  exit 1
fi
