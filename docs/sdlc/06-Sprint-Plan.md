# Sprint Plan — Tada Learn English

> SDLC Document: Master Sprint Plan
> Version: 1.0.0 — Last Updated: 2026-07-10

## 🏆 Milestone Summary

| Sprint | Title | Focus | Issues | SRS Refs | Target Date |
|--------|-------|-------|--------|----------|-------------|
| Sprint 1 | MVP Core | ✅ DONE — Auth, CRUD Words, Docker, API Docs | #1–#8 | FR-1.1→1.4, FR-6.1→6.2 | Jul 28 ✅ |
| Sprint 2 | SRS & Learning | SRS Engine, Flashcard, Quiz, Frontend scaffold | #9–#18 | FR-2.x, FR-3.1→3.2, FR-7.1, FR-1.5 | Aug 11 |
| Sprint 3 | Practice Modes | Dictation, Translation, Text Analyzer | #19–#21 | FR-3.3→3.6, FR-7.2, FR-1.6 | Aug 25 |
| Sprint 4 | Games | Word Chain, Word Builder, Unscramble | #22 | FR-4.1→4.3 | Sep 8 |
| Sprint 5 | Analytics & Multi-User | Dashboard Charts, Export, User Isolation | #23–#25 | FR-5.x, FR-6.4 | Sep 22 |
| OPS | Deploy | Production deployment VPS | #26 | — | TBD |

---

## 🚀 Sprint 1 — MVP Core ✅ DONE

**Goal:** Auth + Word CRUD + Docker + API Docs

| Issue | Title | Status |
|-------|-------|--------|
| #1 | [BUG] Update() only patches word meaning | ✅ Closed |
| #2 | [BUG] List() negative offset bug | ✅ Closed |
| #3 | [BUG] Race condition Create() | ✅ Closed |
| #4 | [BUG] Import() not transactional | ✅ Closed |
| #5 | [BUG] Password reset stubs | ✅ Closed |
| #6 | [BUG] Update() allows duplicate | ✅ Closed |
| #7 | [OPS] Dockerfile version mismatch | ✅ Closed |
| #8 | [CHORE] Remove DEBUG_MARKER | ✅ Closed |

**Deliverables:**
- Go REST API with chi router (JWT auth, word CRUD, swagger docs)
- PostgreSQL 16 database (users, words, srs_states, srs_reviews, refresh_tokens)
- Docker Compose (backend + pgvector)
- GitHub Actions CI (migration + Bruno API tests)

---

## 📚 Sprint 2 — SRS & Learning

**Goal:** SRS engine + learning modes + frontend scaffold

**Backend (9–15):**
| Issue | Title | FR Ref | Priority |
|-------|-------|--------|----------|
| #9 | SRS Engine — 5-band model, auto-schedule, rating | FR-2.1, 2.2, 2.3 | Must |
| #10 | Daily review queue API | FR-2.4 | Must |
| #11 | SRS statistics API | FR-2.5 | Should |
| #12 | Flashcard review API | FR-3.1 | Must |
| #13 | Vocabulary Quiz API | FR-3.2 | Must |
| #14 | Word pronunciation TTS | FR-7.1 | Should |
| #15 | Bulk import CSV/JSON | FR-1.5 | Should |

**Frontend (16–18):**
| Issue | Title | Priority |
|-------|-------|----------|
| #16 | Frontend scaffold + Auth pages | Must |
| #17 | Dashboard & Word management pages | Must |
| #18 | Flashcard & Quiz pages | Must |

---

## 🎯 Sprint 3 — Practice Modes

**Goal:** Advanced learning modes

| Issue | Title | FR Ref | Priority |
|-------|-------|--------|----------|
| #19 | Spelling & Sentence Dictation API | FR-3.3, 3.4 | Should |
| #20 | Translation VI→EN & Text Analyzer API | FR-3.5, 3.6, 7.2, 1.6 | Should/Could |
| #21 | Frontend — Dictation, Translation & Analyzer pages | — | Should |

---

## 🎮 Sprint 4 — Games

**Goal:** Interactive vocabulary games

| Issue | Title | FR Ref | Priority |
|-------|-------|--------|----------|
| #22 | Word Chain, Word Builder, Unscramble | FR-4.1, 4.2, 4.3 | Could |

---

## 📊 Sprint 5 — Analytics & Multi-User

**Goal:** Progress dashboard, data export, user isolation

| Issue | Title | FR Ref | Priority |
|-------|-------|--------|----------|
| #23 | Dashboard Charts & Data Export API | FR-5.1→5.4 | Must/Should |
| #24 | Multi-User Isolation & Profile | FR-6.4 | Should |
| #25 | Frontend — Dashboard Charts & Export | — | Should |

---

## ⚙️ OPS — Deployment

| Issue | Title | Priority |
|-------|-------|----------|
| #26 | Deploy Docker stack to VPS production | Must |

---

## 📌 Định Hướng Sprint Tiếp Theo (Sprint 2)

Đề xuất workflow cho Sprint 2:
1. **Backend trước**: #9 → #10 → #12 → #13 → #14 → #15 → #11
2. **Frontend song song**: #16 (setup) → #17 (pages) → #18 (learning UI)
3. **Backend + Frontend kết nối**: API integration testing

**Tổng story points (Sprint 2):** ~42 points
**Target:** 2 weeks (Jul 28 → Aug 11)

---

## 🔗 Links

- [SRS](01-SRS.md)
- [API Spec](03-API-Spec.md)
- [SDD](04-SDD.md)
- [Use Cases](02-Use-Cases.md)
- [GitHub Issues](https://github.com/DangDDT/tada-learn-english/issues)
