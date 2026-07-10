# API Specification

## Tada Learn English

| Field | Value |
|-------|-------|
| **Base URL** | `https://api.tada-english.dangddt.io.vn` |
| **Version** | 1.0.1 |
| **Protocol** | REST over HTTPS |
| **Content Type** | `application/json` (except TTS: `audio/mpeg`) |

## 1. Common Patterns

### Authentication

All endpoints except register/login require:
```
Authorization: Bearer <JWT>
```

### Response Envelope

Success:
```json
{
  "success": true,
  "data": { ... },
  "meta": { "page": 1, "per_page": 20, "total": 150 }
}
```

Error:
```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable description"
  },
  "data": null
}
```

### Error Codes

| Code | HTTP Status | Meaning |
|------|-------------|---------|
| `VALIDATION_ERROR` | 400 | Invalid input |
| `UNAUTHORIZED` | 401 | Missing or expired token |
| `FORBIDDEN` | 403 | Not resource owner |
| `NOT_FOUND` | 404 | Resource doesn't exist |
| `DUPLICATE_ENTRY` | 409 | Resource already exists |
| `RATE_LIMITED` | 429 | Too many requests |
| `INTERNAL_ERROR` | 500 | Unexpected server error |

### Error Response Examples

```json
// 400 Validation Error
{ "success": false, "error": { "code": "VALIDATION_ERROR", "message": "word is required (1-255 chars)" }, "data": null }

// 401 Unauthorized
{ "success": false, "error": { "code": "UNAUTHORIZED", "message": "Missing or invalid authorization header" }, "data": null }

// 404 Not Found
{ "success": false, "error": { "code": "NOT_FOUND", "message": "Word not found" }, "data": null }

// 409 Conflict
{ "success": false, "error": { "code": "DUPLICATE_ENTRY", "message": "Word 'ephemeral' already exists" }, "data": null }

// 500 Internal Error
{ "success": false, "error": { "code": "INTERNAL_ERROR", "message": "An unexpected error occurred" }, "data": null }
```

### HTTP Status Codes Summary

| Code | Usage |
|------|-------|
| 200 | Success (GET, PUT, POST review) |
| 201 | Created (POST new resource) |
| 204 | No content (DELETE) |
| 400 | Validation error |
| 401 | Missing or invalid auth |
| 403 | Not resource owner |
| 404 | Resource not found |
| 409 | Duplicate entry |
| 429 | Rate limit exceeded |
| 500 | Internal server error |

### Pagination

All list endpoints support:

| Param | Default | Max | Description |
|-------|---------|-----|-------------|
| `page` | 1 | 1000 | Page number (1-indexed) |
| `per_page` | 20 | 100 | Items per page |

Response includes `meta` object: `{ page, per_page, total }`.

---

## 2. Authentication Endpoints

### POST /api/v1/auth/register

Register a new user account.

**Request:**
```json
{
  "email": "user@example.com",
  "password": "Secure123!",
  "name": "Tam Dang"
}
```

**Validation:**
- `email`: valid email format, max 255 chars
- `password`: min 8 chars, at least 1 uppercase, 1 lowercase, 1 digit
- `name`: 1-100 chars

**Response (201):**
```json
{
  "success": true,
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "name": "Tam Dang",
    "created_at": "2026-07-10T08:00:00Z"
  }
}
```

**Errors:**
- 400: Validation error (weak password, invalid email)
- 409: Email already registered

---

### POST /api/v1/auth/login

Authenticate and receive JWT tokens.

**Request:**
```json
{
  "email": "user@example.com",
  "password": "Secure123!"
}
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
    "expires_in": 3600
  }
}
```

**Notes:**
- `access_token`: valid for 3600s (1 hour)
- `refresh_token`: valid for 604800s (7 days)
- `expires_in`: TTL of access_token in seconds

**Errors:**
- 401: Invalid email or password

---

### POST /api/v1/auth/refresh

Exchange a refresh token for a new access token.

**Request:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Response (200):** Same as login response (new access_token + new refresh_token)

**Errors:**
- 401: Expired or revoked refresh token

---

### POST /api/v1/auth/forgot-password

Request a password reset token. Always returns 200 (even if email doesn't exist — security best practice).

**Request:**
```json
{
  "email": "user@example.com"
}
```

**Response (200):**
```json
{
  "success": true,
  "data": { "message": "If the email exists, a reset link has been sent" }
}
```

---

### POST /api/v1/auth/reset-password

Reset password using a reset token.

**Request:**
```json
{
  "token": "reset-token-uuid",
  "password": "NewSecure456!"
}
```

**Response (200):**
```json
{
  "success": true,
  "data": { "message": "Password has been reset successfully" }
}
```

**Errors:**
- 400: Invalid or expired token, weak new password
- 400: Token already used

---

## 3. Words Endpoints

### GET /api/v1/words

List/search vocabulary with pagination and filters.

**Query Parameters:**

| Param | Type | Default | Description |
|-------|------|---------|-------------|
| `page` | int | 1 | Page number |
| `per_page` | int | 20 | Items per page (max 100) |
| `q` | string | — | Search query (fuzzy on word + meaning) |
| `cefr_level` | string | — | Filter: A1, A2, B1, B2, C1, C2 |
| `srs_band` | string | — | Filter: new, learning, reviewing, mature, mastered |
| `tag` | string | — | Filter by tag |
| `sort_by` | string | created_at | Sort: word, created_at, updated_at |
| `sort_dir` | string | desc | Direction: asc, desc |

**Response (200):**
```json
{
  "success": true,
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "word": "ephemeral",
      "pronunciation": "/ɪˈfem.ər.əl/",
      "ipa": "ɪˈfemərəl",
      "meaning": "tạm thời, phù du",
      "part_of_speech": "adjective",
      "example_sentences": ["The fame is ephemeral."],
      "cefr_level": "C1",
      "tags": ["ielts", "academic"],
      "srs_band": "learning",
      "times_reviewed": 3,
      "srs_last_review": "2026-07-09T08:00:00Z",
      "srs_next_review": "2026-07-12T08:00:00Z",
      "created_at": "2026-07-09T10:00:00Z",
      "updated_at": "2026-07-09T10:00:00Z"
    }
  ],
  "meta": { "page": 1, "per_page": 20, "total": 150 }
}
```

**Errors:**
- 400: Invalid page/per_page values (negative, zero, or exceeding max)

---

### GET /api/v1/words/{id}

Get a single word by ID.

**Response (200):** Single word object (same shape as list item)

**Errors:**
- 404: Word not found or belongs to another user

---

### POST /api/v1/words

Create a new word.

**Request:**
```json
{
  "word": "ephemeral",
  "ipa": "ɪˈfemərəl",
  "meaning": "tạm thời, phù du",
  "part_of_speech": "adjective",
  "example_sentences": ["The fame is ephemeral."],
  "cefr_level": "C1",
  "tags": ["ielts"]
}
```

**Validation:**
- `word`: required, 1-255 chars, unique per user
- `meaning`: required, 1-1000 chars
- `cefr_level`: optional, must be A1/A2/B1/B2/C1/C2
- `tags`: optional, array of strings, max 10 tags
- `example_sentences`: optional, array of strings, max 10 sentences

**Response (201):** Created word object

**Errors:**
- 400: Missing required fields, invalid CEFR level
- 409: Word already exists for this user

---

### PUT /api/v1/words/{id}

Update a word (partial update — only sent fields change).

**Request:** Same shape as create, all fields optional
```json
{
  "meaning": "updated meaning",
  "tags": ["ielts", "updated-tag"]
}
```

**Response (200):** Updated word object

**Errors:**
- 400: Invalid field values
- 404: Word not found
- 409: Rename to an existing word

---

### DELETE /api/v1/words/{id}

Soft-delete a word.

**Response (204):** No content

**Errors:**
- 404: Word not found

---

### POST /api/v1/words/import

Bulk import words from CSV file.

**Request:** `multipart/form-data`
- Field: `file` — CSV file with header: `word,meaning,part_of_speech,cefr_level,example`

**Response (200):**
```json
{
  "success": true,
  "data": {
    "created": 45,
    "skipped": 3,
    "errors": ["Row 12: missing meaning"]
  }
}
```

**Notes:**
- Duplicate words (same user + word) are silently skipped
- Validation errors logged per-row, don't fail entire import
- Max file size: 5MB

---

## 4. SRS Endpoints (Planned Sprint 2)

### GET /api/v1/srs/queue

Get today's review queue.

**Query:** `?limit=20`

**Response (200):**
```json
{
  "success": true,
  "data": {
    "due_count": 45,
    "queue": [
      {
        "id": "uuid",
        "word": "ephemeral",
        "meaning": "tạm thời",
        "ipa": "ɪˈfemərəl",
        "srs_band": "learning",
        "times_reviewed": 3,
        "last_reviewed_at": "2026-07-09T08:00:00Z"
      }
    ]
  }
}
```

**Errors:**
- 400: Invalid limit (negative or > 100)

---

### POST /api/v1/srs/review

Submit a review result for a word.

**Request:**
```json
{
  "word_id": "uuid",
  "rating": "easy"
}
```

**Rating Values:**

| Rating | Effect | Interval Multiplier | Band Change |
|--------|--------|--------------------|-------------|
| `easy` | Promote | ×2.5 | Up one band |
| `medium` | Stay | ×1.0 | Same band |
| `hard` | Demote | ×0.5 | Down one band |

**Response (200):**
```json
{
  "success": true,
  "data": {
    "word_id": "uuid",
    "new_band": "reviewing",
    "interval_days": 7,
    "next_review_at": "2026-07-17T08:00:00Z",
    "times_reviewed": 4
  }
}
```

**Errors:**
- 400: Invalid rating value
- 404: Word not found

---

### GET /api/v1/srs/stats

Get SRS statistics for the authenticated user.

**Response (200):**
```json
{
  "success": true,
  "data": {
    "bands": {
      "new": 120,
      "learning": 85,
      "reviewing": 200,
      "mature": 500,
      "mastered": 150
    },
    "total_words": 1055,
    "reviewed_today": 25,
    "streak_days": 7,
    "accuracy_rate": 0.85
  }
}
```

---

## 5. Study Endpoints (Planned Sprint 2-3)

### GET /api/v1/study/flashcard/queue

Get flashcard review queue.

**Response (200):**
```json
{
  "success": true,
  "data": [
    {
      "id": "uuid",
      "word": "ephemeral",
      "pronunciation": "/ɪˈfem.ər.əl/",
      "ipa": "ɪˈfemərəl",
      "meaning": "tạm thời, phù du",
      "part_of_speech": "adjective",
      "example_sentences": ["The fame is ephemeral."],
      "srs_band": "learning"
    }
  ]
}
```

---

### POST /api/v1/study/quiz

Start a quiz session or get next question.

**Response (200):**
```json
{
  "success": true,
  "data": {
    "session_id": "uuid",
    "question": {
      "type": "meaning-to-word",
      "prompt": "tạm thời, phù du",
      "hint": "e________"
    }
  }
}
```

---

### POST /api/v1/study/quiz/check

Check a quiz answer.

**Request:**
```json
{
  "session_id": "uuid",
  "answer": "ephemeral"
}
```

**Response (200):**
```json
{
  "success": true,
  "data": {
    "correct": true,
    "correct_answer": "ephemeral",
    "next_question": { ... }
  }
}
```

---

## 6. TTS Endpoint (Planned Sprint 2)

### GET /api/v1/tts/{word}

Get pronunciation audio.

**Query:** `?accent=us` (default: us, options: us, uk)

**Response (200):** `audio/mpeg` binary stream

**Errors:**
- 400: Word contains non-alphabetic characters
- 404: TTS unavailable for this word

---

## 7. Stats Endpoints (Planned Sprint 5)

### GET /api/v1/stats/dashboard

Dashboard summary.

**Response (200):**
```json
{
  "success": true,
  "data": {
    "total_words": 1055,
    "reviewed_today": 25,
    "streak_days": 7,
    "words_added_today": 3,
    "accuracy_today": 0.88,
    "cefr_distribution": {
      "A1": 50, "A2": 120, "B1": 300,
      "B2": 350, "C1": 180, "C2": 55
    }
  }
}
```

---

### GET /api/v1/stats/daily-activity

Daily activity for charts.

**Query:** `?days=30`

**Response (200):**
```json
{
  "success": true,
  "data": [
    { "date": "2026-06-10", "reviewed": 20, "added": 5 },
    { "date": "2026-06-11", "reviewed": 15, "added": 3 }
  ]
}
```

---

### GET /api/v1/export

Export vocabulary data.

**Query:** `?format=csv` (default: json, options: csv, json)

**Response (200):**
- `format=csv`: `text/csv` with headers
- `format=json`: `application/json` array of word objects

---

## 8. Health Endpoint

### GET /api/v1/health

Health check (no auth required).

**Response (200):**
```json
{
  "status": "ok",
  "database": "connected",
  "uptime_seconds": 3600
}
```

**Note:** Returns 503 if database is unreachable.

---

## 9. Appendix: Planned Endpoints Summary

| Endpoint | Method | Sprint | Status |
|----------|--------|--------|--------|
| `/api/v1/health` | GET | 1 | ✅ Implemented |
| `/api/v1/auth/register` | POST | 1 | ✅ Implemented |
| `/api/v1/auth/login` | POST | 1 | ✅ Implemented |
| `/api/v1/auth/refresh` | POST | 1 | ✅ Implemented |
| `/api/v1/auth/forgot-password` | POST | 1 | ✅ Implemented |
| `/api/v1/auth/reset-password` | POST | 1 | ✅ Implemented |
| `/api/v1/words` | GET | 1 | ✅ Implemented |
| `/api/v1/words/{id}` | GET | 1 | ✅ Implemented |
| `/api/v1/words` | POST | 1 | ✅ Implemented |
| `/api/v1/words/{id}` | PUT | 1 | ✅ Implemented |
| `/api/v1/words/{id}` | DELETE | 1 | ✅ Implemented |
| `/api/v1/words/import` | POST | 1 | ✅ Implemented |
| `/api/v1/srs/queue` | GET | 2 | 🔜 Planned |
| `/api/v1/srs/review` | POST | 2 | 🔜 Planned |
| `/api/v1/srs/stats` | GET | 2 | 🔜 Planned |
| `/api/v1/study/flashcard/queue` | GET | 2 | 🔜 Planned |
| `/api/v1/study/quiz` | POST | 2 | 🔜 Planned |
| `/api/v1/study/quiz/check` | POST | 2 | 🔜 Planned |
| `/api/v1/tts/{word}` | GET | 2 | 🔜 Planned |
| `/api/v1/study/dictation/spelling` | POST | 3 | 🔜 Planned |
| `/api/v1/study/dictation/sentence` | POST | 3 | 🔜 Planned |
| `/api/v1/study/translate` | POST | 3 | 🔜 Planned |
| `/api/v1/study/analyze` | POST | 3 | 🔜 Planned |
| `/api/v1/games/word-chain/*` | POST | 4 | 🔜 Planned |
| `/api/v1/games/word-builder/*` | POST | 4 | 🔜 Planned |
| `/api/v1/games/unscramble/*` | POST | 4 | 🔜 Planned |
| `/api/v1/stats/dashboard` | GET | 5 | 🔜 Planned |
| `/api/v1/stats/daily-activity` | GET | 5 | 🔜 Planned |
| `/api/v1/export` | GET | 5 | 🔜 Planned |
