# Unsubscribe from a feed

**Endpoint:** `DELETE` `v1/subscriptions/:feed_id`

**URL params:**
- `:feed_id=[string]`, `required`, `valid UUID`

**Authenticated:** Yes

## Request

```json
{}
```

## Response

### Success:

`[200 OK]`
```json
{
  "feed_id": "bce17462-f9cb-46d6-a627-353a6b84e322",
  "deleted_at": "2025-04-14T14:27:32.611602Z"
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
  "error": "You are not subscribed to the provided feed"
}
```

---

`[500 Internal Server Error]`
```json
{
  "error": "An unexpected error has occurred, please try again later"
}
```
