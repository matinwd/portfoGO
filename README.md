# Amirhossein Portfolio (Go + Postgres)

This project turns the original static portfolio into a dynamic Go web app backed by PostgreSQL, while preserving all existing styles and layout.

## Features
- Go server renders HTML templates with the same structure and CSS classes
- PostgreSQL stores all content (about, showcase, blog, research, resume)
- Dockerized setup for app + database

## Run with Docker (recommended)

```bash
docker compose up --build
```

Then open:
- http://localhost:8080

## Run locally (without Docker)

1) Start PostgreSQL (or use your existing instance).

2) Set the database URL and start the server:

```bash
export DATABASE_URL='postgres://postgres:postgres@localhost:5432/portfolio?sslmode=disable'
export PORT=8080

go run ./cmd/server
```

## Database
- Schema: `db/schema.sql`
- Seed data: `db/seed.sql`

On startup the server:
1) Creates tables if they do not exist
2) Seeds the database if `site_settings` is empty

## Project layout
- `cmd/server`: Go entrypoint
- `internal/store`: DB access layer
- `templates`: HTML templates
- `assets`, `styles.css`: unchanged static assets
- `db`: schema + seed

## Notes
- Routes keep the original `.html` URLs (`/blog.html`, `/research-comparing-concurrency.html`, etc.)
- Static assets are served from `/assets` and `/styles.css`
- Requires Go 1.24+ for local builds (matches the Docker image)
