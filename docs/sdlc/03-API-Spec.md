# API Specification

## Tada Learn English

| Field | Value |
|---|---|
| **Base URL** | `https://api.tada-english.dangddt.io.vn` |
| **Version** | 1.0.0 |

## 1. Common Patterns

### Authentication
All endpoints (except register/login) require: `Authorization: Bearer <JWT>`

### Response Envelope
```json
{
  "success": true|false,
  "data": { ... },
  "error": { "code": "STRING", "message": "..." },
  "meta": { "page": 1, "per_page": 20, "total": 150 }
}
```

### HTTP Status Codes
| Code | Meaning |
|---|---|
| 200 | Success |
| 201 | Created |
| 204 | No Content (delete) |
| 400 | Validation error |
| 401 | Unauthorized |
| 403 | Forbidden (not owner) |
| 404 | Not Found |
| 409 | Conflict (duplicate) |
| 500 | Internal Error |

## 2. Authentication Endpoints

### POST /api/v1/auth/register
Register new user.
```json
{ "email": "user@example.com", "password": "Secure123!", "name": "Tam Dang" }
```
Response (201): `{ "success": true, "data": { "id": "uuid", "email": "...", "name": "..." } }`

### POST /api/v1/auth/login
Login and receive JWT.
```json
{ "email": "user@example.com", "password": "Secure123!" }
```
Response (200): `{ "success": true, "data": { "access_token": "eyJ...", "refresh_token": "eyJ...", "expires_in": 3600 } }`

### POST /api/v1/auth/refresh
Refresh access token.
```json
{ "refresh_token": "eyJ..." }
```

## 3. Words Endpoints

### GET /api/v1/words
List/search vocabulary.

Query: `?page=1&per_page=20&q=example&cefr_level=B2&srs_band=learning&sort_by=created_at&sort_dir=desc`

Response (200):
```json
{
  "success": true,
  "data": [{
    "id": "uuid",
    "word": "ephemeral",
    "pronunciation": "/ɪˈfem.ər.əl/",
    "ipa": "ɪˈfemərəl",
    "meaning": "tạm thời, phù du",
    "part_of_speech": "adjective",
    "example_sentences": ["Fame is ephemeral."],
    "cefr_level": "C1",
    "tags": ["ielts"],
    "srs_band": "learning",
    "srs_next_review": "2026-07-10T08:00:00Z"
  }],
  "meta": { "page": 1, "per_page": 20, "total": 150 }
}
```

### POST /api/v1/words
Create word.
```json
{
  "word": "ephemeral",
  "ipa": "ɪˈfemərəl",
  "meaning": "tạm thời, phù du",
  "part_of_speech": "adjective",
  "example_sentences": ["Fame is ephemeral."],
  "cefr_level": "C1",
  "tags": ["ielts"]
}
```
Validation: word required (1-255 chars, unique per user), meaning required (1-1000 chars).

### PUT /api/v1/words/{id}
Update word (only sent fields change).

### DELETE /api/v1/words/{id}
Soft-delete word. Returns 204.

### POST /api/v1/words/import
Bulk import from CSV (multipart/form-data).

CSV format: `word,meaning,part_of_speech,cefr_level,example`

## 4. SRS Endpoints

### GET /api/v1/srs/queue
Get today's review queue. Query: `?limit=20`

```json
{
  "success": true,
  "data": {
    "due_count": 45,
    "queue": [{
      "id": "uuid",
      "word": "ephemeral",
      "meaning": "tạm thời",
      "srs_band": "learning",
      "times_reviewed": 3
    }]
  }
}
```

### POST /api/v1/srs/review
Submit review result.
```json
{ "word_id": "uuid", "rating": "easy" }
```
Rating: easy (promote, ×2.5), medium (stay, ×1.0), hard (demote, ×0.5).

### GET /api/v1/srs/stats
Get SRS statistics.
```json
{
  "success": true,
  "data": {
    "bands": { "new": 120, "learning": 85, "reviewing": 200, "mature": 500, "mastered": 150 },
    "total_words": 1055,
    "reviewed_today": 25,
    "streak_days": 7,
    "accuracy_rate": 0.85
  }
}
```

## 5. Stats & TTS Endpoints

### GET /api/v1/stats/dashboard
Dashboard summary with CEFR distribution, daily activity, SRS bands.

### GET /api/v1/stats/export?format=csv
Export vocabulary as CSV or JSON.

### GET /api/v1/tts/{word}?accent=us
Get pronunciation audio. Returns audio/mpeg.