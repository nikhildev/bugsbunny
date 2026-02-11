# BugsBunny

Bug tracking and management application.

## Project structure

- **`api/`** – Main application (CLI and HTTP API)
  - **`commands/`** – Cobra CLI: `run server`, `migrate`
  - **`models/`** – Domain models (issue, component, user, comment)
  - **`routes/`** – HTTP handlers (issue, component, health)
  - **`clients/`** – External clients (database, OpenAI)
  - **`main.go`** – Entry point

Run all commands from the **repository root**.

## How to run

### Run without building

```bash
go run ./api run server
```

```bash
go run ./api migrate
```

To run migrations and seed sample users, components, and issues:

```bash
go run ./api migrate --autopopulate
```

### Build, then run

```bash
go build -o bugsbunny ./api
./bugsbunny run server
./bugsbunny migrate
./bugsbunny migrate --autopopulate   # with sample data
```

### Database (optional)

Start PostgreSQL with Docker from the `api/` directory:

```bash
cd api && docker compose up -d
```

Use `api/.env` or `DB_*` / `DATABASE_URL` environment variables for database config.
