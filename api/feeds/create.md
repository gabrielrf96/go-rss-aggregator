# Create a new feed

**Endpoint:** `POST` `/v1/feeds`

**URL params:** None

**Authenticated:** Yes

## Request

```json
{
    "name": "A feed",
    "url": "https://www.example.com/rss.xml"
}
```

Constraints:
- `name=[string]`, `required`
- `url=[string]`, `required`, `valid URL`

## Response

### Success:

`[201 Created]`
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "A feed",
  "url": "https://www.example.com/rss.xml"
  "created_at": "2025-04-12T14:27:32.611602Z"
}
```

### Error:

`[400 Bad Request]`
```json
{
  "error": "Error parsing JSON"
}
```
OR
```json
{
  "error": "Empty request body, missing parameters"
}
```

---

`[422 Unprocessable Entity]`
```json
{
  "error": "Request parameter validation failed, see suggested corrections",
  "corrections": [
    "Parameter 'name' is required",
    "Parameter 'url' must be a valid URL"
  ]
}
```

---

`[500 Internal Server Error]`
```json
{
  "error": "An unexpected error has occurred, please try again later"
}
```
