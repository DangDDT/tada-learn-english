# Product Requirements Document (PRD)

## Tada Learn English

| Field | Value |
|-------|-------|
| **Project** | Tada Learn English |
| **Version** | 1.0.0 |
| **Status** | Draft |
| **Author** | DangDDT |

## 1. Executive Summary

Tada Learn English is a self-hosted web application for English vocabulary learning and management, inspired by [LearnMyWords.com](https://learnmywords.com). It provides intelligent vocabulary storage, Spaced Repetition System (SRS) review, multiple learning modes, interactive games, and progress analytics — all self-hosted on the user's own infrastructure.

The product targets Vietnamese learners who want full control over their vocabulary data and study process, without relying on third-party services that may limit features, charge subscriptions, or expose private data.

## 2. Problem Statement

| Problem | Impact |
|---------|--------|
| Existing vocabulary apps (Duolingo, Memrise) lock vocabulary behind subscription paywalls | Users pay recurring fees for basic features |
| Third-party services own user vocabulary data | Privacy concerns, no data portability |
| Generic SRS algorithms don't adapt to individual learning patterns | Wasted review time on known words |
| No Vietnamese-first vocabulary tools with CEFR alignment | Learners can't track progress against international standards |
| Most tools don't support bulk import from real-world reading | Manual entry of every word is tedious |

## 3. Target Audience

### Primary: Vietnamese English Learners

| Profile | Description |
|---------|-------------|
| **Students** | High school / university students preparing for IELTS/TOEIC |
| **Professionals** | Working adults improving English for career |
| **Self-learners** | Independent learners who read English content daily |

### Secondary: Self-Hosted Enthusiasts

| Profile | Description |
|---------|-------------|
| **Tech-savvy users** | Users who run their own servers (homelab, VPS) |
| **Privacy-conscious** | Users who don't want their data on third-party services |

## 4. User Goals

| Goal | Frequency | Importance |
|------|-----------|------------|
| Add new words encountered while reading | Daily | Critical |
| Review vocabulary with SRS scheduling | Daily | Critical |
| Practice with flashcards and quizzes | Weekly | High |
| Test pronunciation with audio | Weekly | Medium |
| Track CEFR progress over time | Monthly | Medium |
| Export vocabulary for backup/sharing | Monthly | Low |
| Play word games for fun | Weekly | Low |

## 5. Success Metrics

### North Star Metric
**Words retained at Mature/Mastered band after 30 days**

### Key Performance Indicators

| Metric | Target (MVP) | Stretch |
|--------|-------------|---------|
| **Daily active users (DAU)** | 1 (single user) | 10+ (multi-user) |
| **Words added per user** | 500 / month | 2000 / month |
| **Review completion rate** | > 80% due reviews done daily | > 95% |
| **SRS accuracy rate** | > 70% Easy/Medium ratings | > 85% |
| **Flashcard flip rate** | > 50% words recalled from memory | > 75% |
| **P95 API response time** | < 200ms | < 100ms |
| **Uptime** | 99% (self-hosted) | 99.9% |

### Quality Gates

| Gate | Criteria |
|------|----------|
| **Build** | `go build` passes with zero errors |
| **Lint** | `golangci-lint` passes with zero issues |
| **Tests** | All unit + integration tests pass |
| **API Tests** | Bruno collection passes end-to-end |
| **Security** | No secrets in repo, HTTPS enforced, bcrypt passwords |
| **Docs** | SRS, API Spec, SDD updated with every feature |

## 6. Competitive Landscape

| Feature | Tada | LearnMyWords | Anki | Memrise | Duolingo |
|---------|------|-------------|------|---------|----------|
| Self-hosted | ✅ | ❌ | ❌ | ❌ | ❌ |
| Free & open source | ✅ | ❌ | ✅ | Freemium | Freemium |
| SRS | ✅ SM-2 custom | ✅ | ✅ SM-2 | ✅ | ✅ |
| CEFR tracking | ✅ | ❌ | ❌ | ❌ | ❌ |
| Vietnamese UI | ✅ | ❌ | ❌ | ❌ | ❌ |
| Bulk import CSV | ✅ | ❌ | ✅ | ❌ | ❌ |
| TTS audio | ✅ | ❌ | ✅ (plugin) | ✅ | ✅ |
| Word games | ✅ | ❌ | ❌ | ❌ | ✅ |
| Mobile app | Future | ✅ | ✅ | ✅ | ✅ |

## 7. Release Criteria

### MVP (Sprint 1-2)
- [x] User can register and login
- [x] User can add, edit, search, delete words
- [x] SRS automatically schedules reviews
- [x] Flashcard review with rating works
- [x] Vocabulary quiz works
- [x] Backend deployed on VPS

### Beta (Sprint 3-4)
- [ ] Dictation and translation practice work
- [ ] Text analyzer extracts words with CEFR
- [ ] Word games functional
- [ ] Frontend fully styled and responsive

### GA (Sprint 5)
- [ ] Dashboard with charts
- [ ] Data export
- [ ] Multi-user isolation
- [ ] Performance meets NFR targets

## 8. Risks and Mitigation

| Risk | Likelihood | Impact | Mitigation |
|------|-----------|--------|------------|
| Single developer bottleneck | High | High | Focus on vertical slices, AI-assisted development |
| No frontend expertise | High | High | Use shadcn/ui + Tailwind for rapid prototyping |
| SRS algorithm too complex | Medium | Medium | Start with simple SM-2, iterate based on feedback |
| Scope creep | High | Medium | Strict sprint boundaries, no gold-plating |
| VPS resource exhaustion | Low | Medium | Docker resource limits, monitoring alerts |
| User adoption (multi-user) | Low | Low | Validate single-user first, add multi-user later |
