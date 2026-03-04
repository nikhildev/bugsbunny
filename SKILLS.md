# Skills

Custom Claude Code skills and commands available in this project.

## Commands

### /create-github-pr

Automatically create a pull request based on the current branch and commits. Analyzes git history to generate a meaningful title and description with a test plan template.

Options: `--draft`, `--base <branch>`, `--title <title>`, `--body <body>`

### /commit

Smart commit message generation based on staged changes.

### /review

Code review helper for the current branch or PR.

### /code-review

Comprehensive PR code review. Audits CLAUDE.md compliance, detects bugs via shallow code analysis, reviews git history for context, and posts formatted reviews to GitHub. Uses confidence scoring to filter false positives (only reports issues with confidence >= 80).

### /feature-dev

Guided feature development with a structured 7-phase process:

1. **Discovery** - understand requirements
2. **Codebase exploration** - launch explorer agents to map architecture
3. **Clarifying questions** - gather specifics from the user
4. **Architecture design** - generate implementation approaches
5. **Implementation** - build the feature
6. **Quality review** - automated code review
7. **Summary** - document what was done

Spawns specialized sub-agents (code-explorer, code-architect, code-reviewer) automatically during each phase.

### /cost-history

Fetch and summarize Claude API usage costs for the current period.
