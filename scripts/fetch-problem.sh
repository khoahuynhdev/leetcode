#!/usr/bin/env bash
set -euo pipefail

# fetch-problem.sh
# Fetches a LeetCode problem via GraphQL API.
# Usage:
#   ./fetch-problem.sh                     # Fetch today's daily challenge
#   ./fetch-problem.sh <problem_url>       # Fetch a specific problem by URL
#
# Output: JSON object written to stdout with problem metadata.
# Exit codes:
#   0 - Success
#   1 - API failure after retries
#   2 - Premium problem detected

LEETCODE_GRAPHQL="https://leetcode.com/graphql"
MAX_RETRIES=3
INITIAL_DELAY=2

# ---------------------------------------------------------------------------
# Helpers
# ---------------------------------------------------------------------------

extract_slug_from_url() {
  local url="$1"
  # Handles URLs like:
  #   https://leetcode.com/problems/two-sum/
  #   https://leetcode.com/problems/two-sum/description/
  #   https://leetcode.com/problems/two-sum
  echo "$url" | sed -E 's|.*/problems/([^/]+).*|\1|'
}

graphql_request() {
  local query="$1"
  local variables="${2:-{}}"
  local delay=$INITIAL_DELAY

  for attempt in $(seq 1 $MAX_RETRIES); do
    local response
    local http_code

    response=$(curl -s -w "\n%{http_code}" -X POST "$LEETCODE_GRAPHQL" \
      -H "Content-Type: application/json" \
      -H "Accept: application/json" \
      -H "Referer: https://leetcode.com" \
      -H "Origin: https://leetcode.com" \
      -d "{\"query\": $(echo "$query" | jq -Rs .), \"variables\": $variables}" \
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
# GraphQL Queries
# ---------------------------------------------------------------------------

fetch_daily_challenge() {
  local query='
    query questionOfToday {
      activeDailyCodingChallengeQuestion {
        date
        link
        question {
          questionId
          questionFrontendId
          title
          titleSlug
          difficulty
          isPaidOnly
          acRate
          topicTags { name slug }
          content
        }
      }
    }
  '

  local response
  response=$(graphql_request "$query") || return 1

  echo "$response" | jq '{
    date: .data.activeDailyCodingChallengeQuestion.date,
    link: .data.activeDailyCodingChallengeQuestion.link,
    questionId: .data.activeDailyCodingChallengeQuestion.question.questionId,
    questionFrontendId: .data.activeDailyCodingChallengeQuestion.question.questionFrontendId,
    title: .data.activeDailyCodingChallengeQuestion.question.title,
    titleSlug: .data.activeDailyCodingChallengeQuestion.question.titleSlug,
    difficulty: .data.activeDailyCodingChallengeQuestion.question.difficulty,
    isPaidOnly: .data.activeDailyCodingChallengeQuestion.question.isPaidOnly,
    acRate: .data.activeDailyCodingChallengeQuestion.question.acRate,
    topicTags: [.data.activeDailyCodingChallengeQuestion.question.topicTags[].name],
    content: .data.activeDailyCodingChallengeQuestion.question.content
  }'
}

fetch_problem_by_slug() {
  local slug="$1"
  # Use inline slug instead of GraphQL variables to avoid LeetCode's 499 responses
  local query
  query=$(cat <<EOF
    query questionData {
      question(titleSlug: "${slug}") {
        questionId
        questionFrontendId
        title
        titleSlug
        difficulty
        isPaidOnly
        acRate
        topicTags { name slug }
        content
      }
    }
EOF
  )

  local response
  response=$(graphql_request "$query") || return 1

  local today
  today=$(date -u +%Y-%m-%d)

  echo "$response" | jq --arg date "$today" '{
    date: $date,
    link: ("/problems/" + .data.question.titleSlug + "/"),
    questionId: .data.question.questionId,
    questionFrontendId: .data.question.questionFrontendId,
    title: .data.question.title,
    titleSlug: .data.question.titleSlug,
    difficulty: .data.question.difficulty,
    isPaidOnly: .data.question.isPaidOnly,
    acRate: .data.question.acRate,
    topicTags: [.data.question.topicTags[].name],
    content: .data.question.content
  }'
}

# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------

main() {
  local problem_url="${1:-}"
  local result

  if [[ -n "$problem_url" ]]; then
    local slug
    slug=$(extract_slug_from_url "$problem_url")
    echo "Fetching problem: $slug" >&2
    result=$(fetch_problem_by_slug "$slug") || exit 1
  else
    echo "Fetching today's daily challenge..." >&2
    result=$(fetch_daily_challenge) || exit 1
  fi

  # Check for premium problem
  local is_paid
  is_paid=$(echo "$result" | jq -r '.isPaidOnly')
  if [[ "$is_paid" == "true" ]]; then
    echo "Premium problem detected." >&2
    echo "$result"
    exit 2
  fi

  echo "$result"
}

main "$@"
