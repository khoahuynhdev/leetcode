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
