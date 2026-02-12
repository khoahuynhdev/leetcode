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
    number = str(metadata.get('number', metadata.get('frontend_id', '')))
    frontend_id = str(metadata.get('frontend_id', ''))

    output = {
        "date": metadata.get('date', ''),
        "link": f"/problems/{metadata['slug']}/",
        "questionId": number.lstrip('0'),
        "questionFrontendId": frontend_id,
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
