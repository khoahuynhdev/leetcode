The GH_TOKEN env var is used by the gh CLI (for PR creation, issue creation,
etc.), but git push doesn't read it. Git needs credentials configured
through either a credential helper or a token embedded in the remote URL.
The error message "Invalid username or token. Password authentication is not
supported" confirms that git was falling back to plain password auth, which
GitHub no longer supports.

The fix replaces the remote URL with one that includes the GITHUB_TOKEN as
an x-access-token — this is the standard pattern for authenticating git
operations in GitHub Actions. I also removed the unnecessary GH_TOKEN env
var from this step since it doesn't use gh.
