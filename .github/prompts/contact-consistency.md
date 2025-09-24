# 👤 Contact Consistency Prompt

## Role
Ensure all attribution in the repository lists only:

- **Author:** ambicuity Ritesh Rana  
- **Email:** riteshrana36@gmail.com  

## Goals
- Replace any incorrect names/emails in LICENSE, README, headers, or manifests.
- Remove bot/AI names (e.g., Copilot, github-actions[bot]) from docs.
- Preserve historical credit in `git log` (do not rewrite commit history).

## Workflow
1. Scan for names/emails in text files.
2. Propose replacements with the correct details.
3. Leave commit authorship untouched but ensure **docs and visible metadata** reflect the official author.