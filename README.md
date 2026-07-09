# Tada Learn English

> Self-hosted English vocabulary learning application — Go backend, Next.js frontend, PostgreSQL, Docker.

[![GitHub repo](https://img.shields.io/badge/GitHub-DangDDT%2Ftada--learn--english-blue)](https://github.com/DangDDT/tada-learn-english)

## Quick Start
```bash
git clone https://github.com/DangDDT/tada-learn-english.git
cd tada-learn-english
docker compose up -d
```

## SDLC Documentation
| Document | Description |
|---|---|
| [01-SRS.md](docs/sdlc/01-SRS.md) | Software Requirements Specification |
| [02-Use-Cases.md](docs/sdlc/02-Use-Cases.md) | Use Cases & User Stories |
| [03-API-Spec.md](docs/sdlc/03-API-Spec.md) | API Specification (OpenAPI/Swagger) |
| [04-SDD.md](docs/sdlc/04-SDD.md) | Software Design Document (Architecture) |
| [05-Database-Schema.md](docs/sdlc/05-Database-Schema.md) | Database ERD & Schema |
| [06-Sprint-Backlog.md](docs/sdlc/06-Sprint-Backlog.md) | Sprint & Product Backlog |

## Tech Stack
| Layer | Technology |
|---|---|
| **Frontend** | Next.js 14, TypeScript, Tailwind CSS, shadcn/ui |
| **Backend** | Go 1.22+, chi router, sqlc, pgx |
| **Database** | PostgreSQL 16 + pgvector + pg_trgm |
| **API Docs** | Swagger (swaggo/swag) |
| **Deployment** | Docker Compose, Nginx Proxy Manager, Cloudflare |

## Sprint Roadmap
| Sprint | Dates | Goal |
|---|---|---|
| Sprint 1 | Jul 14–28 | MVP Core — Auth + CRUD Words + Docker |
| Sprint 2 | Jul 28–Aug 11 | SRS + Learning Modes |
| Sprint 3 | Aug 11–25 | Practice Modes |
| Sprint 4 | Aug 25–Sep 8 | Games |
| Sprint 5 | Sep 8–22 | Analytics + Multi-User + Release |

## Features
- 📝 CRUD vocabulary with IPA, CEFR levels, tags
- 🧠 Spaced Repetition System (SM-2, 5-band)
- 📇 Flashcard review with audio
- ✍️ Dictation, Translation, Text Analysis
- 🎮 Word Chain, Builder, Unscramble games
- 📊 Progress dashboard with CEFR charts
- 👥 Multi-user with JWT auth
- 🐳 Docker Compose deployment