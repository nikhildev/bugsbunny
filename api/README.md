# BugsBunny

Bug tracking and management application on AI steroids

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

### Database and vector store (optional)

Start PostgreSQL and Weaviate with Docker from the `api/` directory:

```bash
cd api && docker compose up -d
```

Use `api/.env` or `DB_*` / `DATABASE_URL` environment variables for database config.

## Architecture

```
                            ┌─────────────────────┐
                            │      main.go         │
                            │   cli.Execute()      │
                            └─────────┬───────────┘
                                      │
                            ┌─────────▼───────────┐
                            │    Cobra CLI         │
                            │  server │ migrate    │
                            └─────────┬───────────┘
                                      │ "server" cmd
          ┌───────────────────────────▼───────────────────────────┐
          │                    STARTUP (server.go)                │
          │  1. Load config (.env / env vars)                     │
          │  2. Connect PostgreSQL (GORM)                         │
          │  3. Init Embedding client (EmbeddingGemma)            │
          │  4. Init Weaviate (vector DB)                         │
          │  5. Setup routes + middleware                         │
          │  6. http.ListenAndServe()                             │
          └───────────────────────────┬───────────────────────────┘
                                      │
                      ┌───────────────▼───────────────┐
                      │   Middleware Chain             │
                      │  Recovery → Logging → CORS    │
                      └───────────────┬───────────────┘
                                      │
                ┌─────────────────────▼─────────────────────┐
                │              http.ServeMux (router.go)     │
                ├───────────┬──────────┬──────────┬─────────┤
                │           │          │          │         │
           /health    /projects   /issues   /search   /simulate
                │           │          │       /knowledge    /rag
                │           │          │          │         │
                │      ┌────▼────┐ ┌──▼───┐  ┌───▼───┐ ┌──▼──┐
                │      │ CRUD    │ │ CRUD │  │Vector │ │ RAG │
                │      │ + Sync  │ │      │  │Search │ │Test │
                │      │ to      │ │      │  │       │ │     │
                │      │Weaviate │ │      │  │       │ │     │
                │      └────┬────┘ └──┬───┘  └───┬───┘ └──┬──┘
                │           │         │          │        │
                │      ┌────▼─────────▼──┐  ┌────▼────────▼───┐
                │      │   PostgreSQL    │  │   Weaviate      │
                │      │   (GORM ORM)   │  │   (Vectors)     │
                │      │                │  │                  │
                │      │ Users          │  │ KnowledgeEntry   │
                │      │ Projects       │  │  - projectId     │
                │      │ Issues         │  │  - content       │
                │      │ Comments       │  │  - vector[768]   │
                │      │ Changes        │  │                  │
                │      └────────────────┘  └────────┬─────────┘
                │                                   │
                │                          ┌────────▼─────────┐
                │                          │  EmbeddingGemma  │
                │                          │  (self-hosted)   │
                │                          │  OpenAI-compat   │
                │                          └──────────────────┘
```

## RAG architecture

BugsBunny uses a self-hosted RAG (Retrieval-Augmented Generation) pipeline for bot knowledge search.

### Components

| Component | How to run | Role |
|---|---|---|
| Weaviate | `docker compose up -d` (see `docker-compose.yaml`) | Vector database — stores objects and pre-computed vectors |
| EmbeddingGemma | `docker compose up -d` (pulls `ai/embeddinggemma` on first start) | Embedding model — OpenAI-compatible API at port `12434` |
| PostgreSQL | `docker compose up -d` | Primary data store |

### How it works

Vectors are computed by the Go application before being written to Weaviate. Weaviate is configured with `vectorizer: none` and stores only the vectors it is given.

1. **Ingestion** — when a project is created or updated with `bot_knowledge` entries, the app calls the EmbeddingGemma API (`POST /embeddings`) for each entry to obtain a 768-dimensional vector, then batch-inserts the objects into Weaviate with those vectors attached.

2. **Retrieval** — `GET /search/knowledge?q=<query>` embeds the query text via the same API and issues a `nearVector` search in Weaviate, returning the closest knowledge entries ranked by cosine distance.

3. **Simulation** — `POST /simulate/rag` calls the embedding API directly and returns the raw vector and its dimension count without writing anything to Weaviate.

### Prerequisites

All three services (PostgreSQL, Weaviate, EmbeddingGemma) run via Docker Compose. On first start, the `embeddinggemma` container pulls the `ai/embeddinggemma` model from Docker Hub automatically and caches it in a named volume.

```bash
cd api && docker compose up -d
```

The embedding API is available at `http://localhost:12434/v1`.

### Configuration

| Environment variable | Default | Description |
|---|---|---|
| `EMBEDDING_BASE_URL` | — | Base URL of the embedding API (e.g. `http://localhost:12434/engines/llama.cpp/v1`) |
| `EMBEDDING_MODEL` | — | Model name passed in API requests (e.g. `ai/embeddinggemma`) |
| `WEAVIATE_HOST` | — | Host and port of Weaviate (e.g. `localhost:8000`) |
| `WEAVIATE_SCHEME` | `http` | Scheme used to connect to Weaviate |

### Endpoints

```
GET  /search/knowledge?q=<query>[&project_id=<id>][&top_k=<n>]
POST /simulate/rag   { "text": "..." }
```
