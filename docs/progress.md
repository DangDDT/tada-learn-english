# Project Progress

> Live progress dashboard for tada-learn-english.
> Last updated: 2026-07-10 18:55 ICT

## Active Sprint: Sprint 2 — SRS & Frontend

> Current Task: SDLC docs revision (PRD, SRS acceptance criteria, Sprint Plan/Backlog, SDD, DB Schema, API Spec) — all 7 documents updated and aligned with reality.

## Overall Status

```
Sprint 1 [████████████████████] 100% (8/8 closed, backend only)
Sprint 2 [····················]   0% (0/11)
Sprint 3 [····················]   0% (0/8)
Sprint 4 [····················]   0% (0/6)
Sprint 5 [····················]   0% (0/10)
OPS      [····················]   0% (0/1)
─────────────────────────────────────────
Total    [██··················]   8% (8/44)
```

## Sprint 1 — MVP Core ✅ (Jul 14-28)

**Status: Completed**

| Issue | Title | Status |
|-------|-------|--------|
| #1 | Update() only patches word & meaning — other fields silently ignored | ✅ Fixed |
| #2 | List() returns all words without pagination or filtering | ✅ Fixed |
| #3 | Race condition in Create() — EXISTS check + INSERT not atomic | ✅ Fixed |
| #4 | Import() not transactional — partial failures create inconsistent state | ✅ Fixed |
| #5 | Password reset endpoints are stubs — return success but do nothing | ✅ Fixed |
| #6 | Update() allows duplicate word entries — no uniqueness check on rename | ✅ Fixed |
| #7 | Dockerfile Go version mismatch + missing .dockerignore | ✅ Fixed |
| #8 | Remove leftover DEBUG_MARKER comment in words.go | ✅ Fixed |

## Sprint 2 — SRS & Learning ❌ (Jul 28-Aug 11)

**10 open issues — pending implementation**

| Issue | Title | Priority |
|-------|-------|----------|
| #9 | SRS Engine API (SM-2 algorithm) | High |
| #10 | SRS Review API (scheduling, next review) | High |
| #11 | Flashcard + Quiz API | Medium |
| #12 | TTS API (Text-to-Speech) | Medium |
| #13 | Frontend scaffold (Next.js, Tailwind, shadcn/ui) | High |
| #14 | Frontend auth pages (login, register, forgot/reset password) | High |
| #15 | Frontend word management pages | High |
| #16 | Frontend flashcard review page | High |
| #17 | Frontend quiz page | Medium |
| #18 | Frontend flashcard & quiz pages | Medium |

## Sprint 3 — Practice Modes ❌ (Aug 11-25)

| Issue | Title | Priority |
|-------|-------|----------|
| #19 | Spelling & Sentence Dictation API | Medium |
| #20 | Translation VI→EN & Text Analyzer API | Medium |
| #21 | Frontend — Dictation, Translation & Text Analyzer pages | Medium |

## Sprint 4 — Games ❌ (Aug 25-Sep 8)

| Issue | Title | Priority |
|-------|-------|----------|
| #22 | Word Chain, Word Builder, Unscramble (backend + frontend) | Low |

## Sprint 5 — Analytics & Multi-User ❌ (Sep 8-22)

| Issue | Title | Priority |
|-------|-------|----------|
| #23 | Dashboard Charts & Data Export API | Medium |
| #24 | Multi-User Isolation & Profile Management | Medium |
| #25 | Frontend — Dashboard Charts & Data Export pages | Medium |

## OPS ❌

| Issue | Title | Priority |
|-------|-------|----------|
| #26 | Deploy Docker stack to VPS production | Medium |

## Project Health

| Metric | Status |
|--------|--------|
| Backend Go build | ✅ Pass |
| Bruno API tests | ✅ 11 tests pass |
| GitHub Actions CI | ✅ Configured |
| Go unit tests | ❌ None |
| Go linter (golangci-lint) | ❌ Not configured |
| Frontend scaffold | ❌ Not started |
| .env.example | ❌ Missing |
| Pre-commit hooks | ❌ Not configured |
| Production deploy | ❌ Not deployed |
| SDLC Docs (PRD/SRS/SDD/API/DB) | ✅ Updated 2026-07-10 |
| GitHub Project Board | ✅ Active, 28 items tracked |
| Progress Sync Flow | ✅ Defined in AGENTS.md |
