# Sprint Plan — Tada Learn English

> SDLC Document: Master Sprint Plan
> Version: 1.0.1 — Last Updated: 2026-07-10
> [PRD](00-PRD.md) | [SRS](01-SRS.md) | [Progress](../progress.md) | [Plans](../plans/INDEX.md)

## Sprint Summary

| Sprint | Title | Focus | Issues | FR Refs | Target Date | Status |
|--------|-------|-------|--------|---------|-------------|--------|
| Sprint 1 | Backend MVP | Auth, CRUD Words, Docker, CI | #1–#8 | FR-1.1→1.4, FR-6.1→6.3 | Jul 10 ✅ | Done |
| Sprint 2 | SRS + Frontend | SRS Engine, Frontend scaffold, Learning UI | #9–#18 | FR-2.x, FR-3.1→3.2, FR-7.1, FR-1.5 | Aug 11 | 🔴 Active |
| Sprint 3 | Practice Modes | Dictation, Translation, Text Analyzer | #19–#21 | FR-3.3→3.6, FR-7.2, FR-1.6 | Aug 25 | Pending |
| Sprint 4 | Games | Word Chain, Word Builder, Unscramble | #22 | FR-4.1→4.3 | Sep 8 | Pending |
| Sprint 5 | Analytics & Multi-User | Dashboard Charts, Export, User Isolation | #23–#25 | FR-5.x, FR-6.4 | Sep 22 | Pending |
| OPS | Deploy | Production deployment VPS | #26 | — | TBD | Pending |

---

## Sprint 1 — Backend MVP ✅ DONE

**Goal:** Go REST API with auth + word CRUD + Docker + CI

**Actual completion:** Jul 10, 2026 (completed early, backend only)

| Issue | Title | Status | Delivered |
|-------|-------|--------|-----------|
| #1 | [BUG] Update() only patches word & meaning | ✅ Closed | Dynamic field update |
| #2 | [BUG] List() returns all words without pagination | ✅ Closed | Pagination + filter + sort |
| #3 | [BUG] Race condition Create() | ✅ Closed | ON CONFLICT + unique constraint |
| #4 | [BUG] Import() not transactional | ✅ Closed | Transactional CSV import |
| #5 | [BUG] Password reset stubs | ✅ Closed | Working forgot/reset flow |
| #6 | [BUG] Update() allows duplicate words | ✅ Closed | Uniqueness check on rename |
| #7 | [OPS] Dockerfile version mismatch | ✅ Closed | Go 1.25 Dockerfile, .dockerignore |
| #8 | [CHORE] Remove DEBUG_MARKER | ✅ Closed | Code cleanup |

**What was built:**
- Go REST API with chi router
- JWT auth (register, login, refresh, forgot/reset password)
- Word CRUD (create, list with pagination/filter/sort, get, update with partial fields, soft-delete, import CSV)
- Auth middleware (Bearer JWT validation)
- Health check endpoint
- Swagger docs (swaggo)
- PostgreSQL schema + migrations (password_reset_tokens)
- Dockerfile + docker-compose.yml
- Bruno API tests (11 tests) + GitHub Actions CI

**What was NOT built (carried over to Sprint 2):**
- ❌ Next.js frontend scaffold
- ❌ Frontend auth pages
- ❌ Frontend word management pages

---

## Sprint 2 — SRS & Frontend 🔴 Active

**Goal:** SRS engine + frontend scaffold + learning UI

**Timeline:** Jul 10 – Aug 11, 2026

**Carried over from Sprint 1:**
- Frontend scaffold (Next.js + Tailwind + shadcn/ui)
- Frontend auth pages (login, register, forgot/reset password)
- Frontend word management pages

**Backend (new):**
| Issue | Title | FR Ref | Priority |
|-------|-------|--------|----------|
| #9 | SRS Engine — 5-band model, auto-schedule, rating | FR-2.1, 2.2, 2.3 | Must |
| #10 | Daily review queue API | FR-2.4 | Must |
| #11 | SRS statistics API | FR-2.5 | Should |
| #12 | Flashcard review API | FR-3.1 | Must |
| #13 | Vocabulary Quiz API | FR-3.2 | Must |
| #14 | Word pronunciation TTS | FR-7.1 | Should |
| #15 | Bulk import CSV/JSON | FR-1.5 | Should |

**Frontend:**
| Issue | Title | Priority |
|-------|-------|----------|
| #13 (was #16) | Frontend scaffold + Auth pages | Must |
| #14 (was #17) | Dashboard & Word management pages | Must |
| #15 (was #18) | Flashcard & Quiz pages | Must |

**Recommended order:**
1. Frontend scaffold (Next.js + Tailwind + shadcn/ui + API client)
2. SRS engine backend (#9, #10)
3. Frontend auth pages
4. Flashcard API + frontend (#12, #15)
5. Quiz API + frontend (#13, #15)
6. TTS + CSV import + SRS stats (#14, #15, #11)
7. Frontend word management pages (#14)

---

## Sprint 3 — Practice Modes

**Goal:** Advanced learning modes

| Issue | Title | FR Ref | Priority |
|-------|-------|--------|----------|
| #19 | Spelling & Sentence Dictation API | FR-3.3, 3.4 | Should |
| #20 | Translation VI→EN & Text Analyzer API | FR-3.5, 3.6, 7.2, 1.6 | Should/Could |
| #21 | Frontend — Dictation, Translation & Analyzer pages | — | Should |

---

## Sprint 4 — Games

**Goal:** Interactive vocabulary games

| Issue | Title | FR Ref | Priority |
|-------|-------|--------|----------|
| #22 | Word Chain, Word Builder, Unscramble | FR-4.1, 4.2, 4.3 | Could |

---

## Sprint 5 — Analytics & Multi-User

**Goal:** Progress dashboard, data export, user isolation

| Issue | Title | FR Ref | Priority |
|-------|-------|--------|----------|
| #23 | Dashboard Charts & Data Export API | FR-5.1→5.4 | Must/Should |
| #24 | Multi-User Isolation & Profile | FR-6.4 | Should |
| #25 | Frontend — Dashboard Charts & Export | — | Should |

---

## OPS — Deployment

| Issue | Title | Priority |
|-------|-------|----------|
| #26 | Deploy Docker stack to VPS production | Must |

---

## Links

| Document | Purpose |
|----------|---------|
| [PRD](00-PRD.md) | Product requirements & success metrics |
| [SRS](01-SRS.md) | Functional & non-functional requirements |
| [Use Cases](02-Use-Cases.md) | Use cases & user stories |
| [API Spec](03-API-Spec.md) | API endpoint specifications |
| [SDD](04-SDD.md) | Software design & architecture |
| [DB Schema](05-Database-Schema.md) | Database schema & migrations |
| [Progress](../progress.md) | Live progress dashboard |
| [GitHub Issues](https://github.com/DangDDT/tada-learn-english/issues) | Issue tracker |
