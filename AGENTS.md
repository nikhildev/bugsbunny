## BugsBunny Agents Guide

This document is for AI/codegen agents (Cursor, Claude, etc.) working in this repository.  
Follow these conventions unless the user explicitly asks for something different.

### High‑level overview

- **Project**: Bug tracking and management application with an HTTP API in Go and a React frontend powered by Bun.
- **Back end**: Go 1.26+ app in `api/` (Cobra CLI + HTTP routes + Postgres).
- **Front end**: React + Tailwind + shadcn UI in `frontend/`, built and served with Bun.

### Project layout (important directories)

- **`api/`**: Main Go module.
  - **`commands/`**: Cobra commands (`run server`, `migrate`, etc.).
  - **`routes/`**: HTTP handlers (`issue`, `component`, `comments`, `health`, etc.).
  - **`models/`**: Domain models (`issue`, `component`, `user`, `comment`, `change`, etc.).
  - **`clients/`**: External clients (`database`, `openai`).
  - **`common/`**: Shared helpers, including `test_utils` for Go tests.
  - **`main.go`**: CLI entry point.
- **`frontend/`**: Bun + React app.
  - Uses `bun` scripts from `frontend/package.json` (`dev`, `start`, `build`).

### Tooling & environment

- **Languages & runtimes**
  - **Go**: Target **Go 1.26+** for `api/`.
  - **Bun**: Use **Bun** (not Node/npm/yarn) for everything in `frontend/`.
- **Databases & services**
  - **Postgres**: Run via Docker Compose from `api/` when tests or the API need a database.
  - **OpenAI**: The Go API depends on `github.com/openai/openai-go/v3`; treat keys as secrets and never hardcode them.

### Setup instructions

- **Backend (Go API)**
  - **Dependencies**: Go modules are tracked in `api/go.mod`; normally no manual step is needed beyond:
    - **Install Go** 1.26+.
    - From the repo root, Go tools will auto‑resolve modules on first build/test.
- **Frontend (Bun + React)**
  - From the repo root:
    - **Install Bun** (v1+).
    - Install dependencies:
      - **`cd frontend && bun install`**

### Running the application

- **Database (Postgres via Docker)**
  - From the repo root:
    - **`cd api && docker compose up -d`**
  - Configuration:
    - Use `api/.env` or `DB_*` / `DATABASE_URL` environment variables.

- **Run the Go API without building (from repo root)**
  - **Start HTTP server**:
    - **`go run ./api run server`**
  - **Run migrations**:
    - **`go run ./api migrate`**
  - **Run migrations with sample data**:
    - **`go run ./api migrate --autopopulate`**

- **Build the Go binary, then run (from repo root)**
  - **Build**:
    - **`go build -o bugsbunny ./api`**
  - **Run server**:
    - **`./bugsbunny run server`**
  - **Run migrations**:
    - **`./bugsbunny migrate`**
    - **`./bugsbunny migrate --autopopulate`**

- **Run the frontend (from repo root)**
  - Development:
    - **`cd frontend && bun dev`**
      - Uses `scripts.dev = "bun --hot src/index.ts"` from `frontend/package.json`.
  - Production:
    - **`cd frontend && bun start`**
      - Uses `scripts.start = "NODE_ENV=production bun src/index.ts"`.
  - Build artifacts (if needed by other tooling):
    - **`cd frontend && bun run build`**

### Code generation & conventions

- **UUIDs**
  - **Whenever generating UUIDs, always use UUID7 format.**
  - In Go, prefer the existing UUID library (`github.com/google/uuid`) and extend it with UUID7 helpers or use a compatible UUID7 helper. Do **not** introduce incompatible formats.

- **Go code organization**
  - **Preserve existing layering**:
    - New CLI behavior → `api/commands/`.
    - New API endpoints → `api/routes/<domain>/`.
    - New domain entities → `api/models/`.
    - New external services (e.g. new API clients) → `api/clients/`.
    - Shared logic that isn’t domain‑specific → `api/common/`.
  - Prefer standard library + existing dependencies; avoid adding new frameworks unless necessary and justified.

- **Go test utilities**
  - **Existing rule (preserve)**:  
    - **Whenever generating Go test files, keep common code modular and in a `test_utils` location so it can be reused.**
  - Practically:
    - Put shared helpers in `api/common/test_utils.go` or a dedicated `test_utils` package.
    - Tests should import from `test_utils` instead of duplicating setup/teardown logic.

- **Frontend (React + Bun)**
  - Always:
    - Use **Bun** commands (`bun install`, `bun dev`, `bun start`, `bun run build`) instead of `npm` / `yarn` / `pnpm`.
    - Use modern React (hooks, function components) and the existing Tailwind + shadcn design system.
  - When adding dependencies:
    - Prefer packages consistent with the existing stack (React 19, TanStack Query, Radix, shadcn).
    - Avoid introducing a separate build tool (e.g. Vite, Webpack); use Bun’s built‑in bundler.

### Testing strategy

- **Go API tests**
  - **Run all tests (from repo root)**:
    - **`cd api && go test ./...`**
  - For tests that touch Postgres or other infrastructure:
    - Use **testcontainers-go** (already in `api/go.mod`) for integration tests.
    - Ensure Docker is running; these tests may be slower and should be clearly named to reflect integration behavior.
  - New tests:
    - Place unit tests alongside their packages (e.g. `routes/issue/create_test.go`).
    - Move repeated setup/fixtures into `api/common/test_utils.go` (or similar) rather than duplicating across files.

- **Frontend tests**
  - Prefer Bun’s native test runner for any JavaScript/TypeScript tests:
    - **`cd frontend && bun test`**
  - When creating new tests:
    - Co‑locate tests with components (e.g. `Component.test.tsx`).
    - Use Bun’s `bun:test` APIs (no Jest/Vitest boilerplate unless the project is explicitly migrated).

### When editing or extending the project

- **Backend changes**
  - After modifying Go code:
    - Run **`cd api && go test ./...`**.
    - Ensure builds succeed (`go build ./api/...` or `go build -o bugsbunny ./api`).
  - For changes to CLI commands or routes, update any relevant docs in `api/README.md`.

- **Frontend changes**
  - After modifying React or Bun code:
    - Run **`cd frontend && bun dev`** to verify the app loads and basic flows work.
    - If tests exist or were added, run **`cd frontend && bun test`**.

- **Cross‑cutting behavior**
  - Keep environment‑specific values in `.env` / env vars, not hardcoded in code or committed files.
  - When adding new modules or directories, follow the naming and layout patterns already present in `api/` and `frontend/`.
