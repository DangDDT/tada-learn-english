# Tada Learn English — Agent Instructions

## Context
This is a Go backend + Next.js frontend English vocabulary learning app.
- Backend: Go 1.22, chi router, sqlc, pgx, PostgreSQL 16
- Frontend: Next.js 14, TypeScript, Tailwind, shadcn/ui
- Deploy: Docker Compose

## SDLC Rules
1. ALWAYS create a GitHub issue BEFORE writing code
2. Branch off main: `feat/issue-N-description` or `fix/issue-N-description`
3. SRS/API-Spec/SDD in docs/sdlc/ are the Source of Truth
4. Commit format: `type: concise subject` (feat/fix/refactor/docs/chore)
5. No direct commits to main — PR required
6. Run tests before creating PR

## File Structure
- backend/ — Go API server
- frontend/ — Next.js app (to be scaffolded)
- docs/sdlc/ — SDLC documentation (SRS, API-Spec, SDD, etc.)
- .github/workflows/ — CI/CD pipelines
- docker-compose.yml — Container orchestration

## Database
- PostgreSQL 16 with pgvector, pg_trgm extensions
- Schema migrations in backend/migrations/
- Use sqlc for type-safe SQL queries
