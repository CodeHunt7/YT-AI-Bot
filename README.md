# YT-AI-Bot

Telegram AI assistant for YouTube creators.

## Local PostgreSQL

Start PostgreSQL via Docker Compose:

```bash
docker compose up -d postgres
```

## Migrations (goose)

Apply migrations:

```bash
goose -dir migrations postgres "$DATABASE_URL" up
```

Rollback one step:

```bash
goose -dir migrations postgres "$DATABASE_URL" down
```

## Tests

Run tests:

```bash
go test ./...
```
