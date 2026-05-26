# Project Agent Instructions

## Project

Telegram AI assistant for YouTube creators.

The product is not a generic AI chat. It is a Telegram-based YouTube producer assistant with methodology, prompts, knowledge base, onboarding, and structured workflows.

## Tech stack

- Language: Go
- Bot: Telegram Bot API
- Database: PostgreSQL
- Migrations: goose or golang-migrate
- Config: environment variables and `.env` locally
- Logging: Go `slog`
- Knowledge base: markdown first, pgvector/RAG later
- Local dev: Docker Compose for PostgreSQL

## Architecture

Keep the project layered:

- `/cmd/bot` — entrypoint
- `/internal/bot` — Telegram handlers, messages, keyboards
- `/internal/app` — use cases and business logic
- `/internal/ai` — AI gateway and model providers
- `/internal/youtube` — YouTube Data API client
- `/internal/knowledge` — knowledge base and later RAG
- `/internal/storage` — PostgreSQL repositories
- `/internal/billing` — subscriptions, credits, limits
- `/internal/config` — config loading
- `/migrations` — database migrations
- `/prompts` — prompt templates
- `/knowledge` — methodology docs, checklists, examples
- `/docs` — project documentation

Do not put business logic inside Telegram handlers.

## Development rules

Work in small steps.

Do not implement the whole bot at once.

For every feature:
1. Make the smallest working version.
2. Add or update tests where useful.
3. Keep errors explicit.
4. Avoid hidden global state.
5. Avoid hardcoded API keys, model names, prices, tokens, or secrets.

## AI rules

All AI calls must go through `/internal/ai`.

Do not call OpenAI, Anthropic, Gemini, or other providers directly from handlers or use cases.

Model names, prices, and provider choice must be configurable.

Log AI usage at least conceptually:
- user id
- feature
- provider
- model
- input tokens if available
- output tokens if available
- estimated cost if available
- timestamp

## Prompt rules

Prompts live in `/prompts`.

Do not hardcode long prompts inside Go files.

Prompt output should use structured sections where possible.

Important prompt categories:
- channel analysis
- competitor analysis
- idea analyzer
- script writer
- metadata writer
- thumbnail brief
- free YouTube question

## Product rules

MVP priority:
1. `/start`
2. onboarding
3. save user profile
4. YouTube channel fetch
5. AI channel analysis
6. idea analyzer
7. script writer
8. title/description/tags
9. credits and limits

Do not add payment, image generation, complex admin UI, or full RAG before the core flow works.

## Testing rules

Add tests especially for:
- YouTube URL parsing
- state transitions
- credit calculation
- prompt formatting
- storage repositories
- AI gateway with mock provider

Do not require real external APIs in unit tests.

## Safety and secrets

Never commit:
- `.env`
- API keys
- Telegram bot tokens
- database passwords
- production credentials

Use `.env.example` for required variables.

## Codex behavior

Before making large changes, inspect the existing structure.

Prefer minimal patches.

Do not rename folders or change architecture without a clear reason.

After changes, say what was changed and what command should be run to verify it.
