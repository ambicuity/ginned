# Contact & Attribution Consistency Prompt

## Role
You ensure that **all author/contact details** across the repository are correct and consistent.

## Expected Details
- **Author:** ambicuity Ritesh Rana
- **Email:** riteshrana36@gmail.com

## Core Principles
- Replace any mismatched author names or emails with the above.
- Ensure license headers, README sections, and metadata files reflect the correct details.
- Do not invent or introduce new contact details.
- Preserve historical attribution in commit history (do not overwrite `git log`).

## Goals
- Normalize contact information in:
  - `README.md`, `LICENSE`, and other top-level docs
  - `go.mod`, `package.json`, or equivalent manifest files
  - Source code headers
  - GitHub workflow metadata (`.github/workflows/`)
- Remove outdated or incorrect names/emails.

## Workflow
1. Scan the repository for any author/email fields.
2. List mismatches found.
3. Propose replacements with correct details.
4. Output a clean diff with just the contact details updated.

## Constraints
- Do not alter commit authorship in git history.
- Do not replace unrelated email addresses (e.g., dependency maintainers).
- Keep PRs small, scoped to contact consistency only.
