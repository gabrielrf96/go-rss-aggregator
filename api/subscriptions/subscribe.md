# Subscribe to a feed

**Endpoint:** `POST` `v1/subscriptions/:feed_id`

**URL params:**
- `:feed_id=[string]`, `required`, `valid UUID`

**Authenticated:** Yes

## Request

```json
{}
```

## Response

### Success:

`[201 Created]`
```json
{
  "feed_id": "bce17462-f9cb-46d6-a627-353a6b84e322",
  "name": "Feed 1",
  "url": "https://www.feed1.com/index.xml",
  "created_at": "2025-04-13T14:27:32.611602Z"
}
```

### Error:
`[400 Bad Request]`
```json
{
  "error": "Failed parsing feed ID, make sure you provided a valid ID"
}
```

---

`[404 Not Found]`
```json
{
  "error": "The provided feed does not exist"
}
```

---

`[409 Conflict]`
```json
{
  "error": "You are already subscribed to that feed"
}
```

---

`[500 Internal Server Error]`
```json
{
  "error": "An unexpected error has occurred, please try again later"
}
```
