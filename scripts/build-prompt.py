#!/usr/bin/env python3
"""
Reads problem.json and scripts/prompt-template.md, substitutes placeholders,
and writes the completed prompt to stdout.

Environment variables:
  IS_DUPLICATE  - "true" if problem directory already exists
  DATE_COMPACT  - YYYYMMDD string for duplicate file suffixes
"""
import json
import os
import sys


def main():
    problem_path = sys.argv[1] if len(sys.argv) > 1 else "problem.json"
    template_path = sys.argv[2] if len(sys.argv) > 2 else "scripts/prompt-template.md"

    with open(problem_path) as f:
        data = json.load(f)

    with open(template_path) as f:
        template = f.read()

    number = data["questionFrontendId"].zfill(4)
    slug = data["titleSlug"]
    title = data["title"]
    difficulty = data["difficulty"]
    topics = ", ".join(data["topicTags"])
    ac_rate = f"{data['acRate']:.1f}"
    content = data["content"]
    dir_name = f"{number}-{slug}"

    is_duplicate = os.environ.get("IS_DUPLICATE", "false") == "true"
    date_compact = os.environ.get("DATE_COMPACT", "")

    if is_duplicate:
        duplicate_notice = (
            f"\n**Note:** This problem already exists in the repository at "
            f"`problems/{dir_name}/`. Create files with the `_daily_{date_compact}` "
            f"suffix instead."
        )
        file_note = (
            f"Name the file `solution_daily_{date_compact}.go` and place it in "
            f"`problems/{dir_name}/`"
        )
        test_note = (
            f"Name the file `solution_daily_{date_compact}_test.go` and place it in "
            f"`problems/{dir_name}/`"
        )
    else:
        duplicate_notice = ""
        file_note = f"Name the file `solution.go` and place it in `problems/{dir_name}/`"
        test_note = f"Name the file `solution_test.go` and place it in `problems/{dir_name}/`"

    replacements = {
        "{{PROBLEM_NUMBER}}": number,
        "{{PROBLEM_TITLE}}": title,
        "{{PROBLEM_SLUG}}": slug,
        "{{DIFFICULTY}}": difficulty,
        "{{TOPICS}}": topics,
        "{{AC_RATE}}": ac_rate,
        "{{PROBLEM_CONTENT}}": content,
        "{{DUPLICATE_NOTICE}}": duplicate_notice,
        "{{FILE_NAMING_NOTE}}": file_note,
        "{{TEST_FILE_NAMING_NOTE}}": test_note,
    }

    prompt = template
    for placeholder, value in replacements.items():
        prompt = prompt.replace(placeholder, value)

    print(prompt)


if __name__ == "__main__":
    main()
