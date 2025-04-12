# Get available feeds

**Endpoint:** `GET` `/v1/feeds`

**URL params:** None

**Authenticated:** Yes

## Request

```json
{}
```

## Response

### Success:

`[200 OK]`
```json
[
  {
    "id": "bce17462-f9cb-46d6-a627-353a6b84e322",
    "name": "Feed 1",
    "url": "https://www.feed1.com/index.xml",
    "created_at": "2025-04-10T00:00:00.294858Z",
    "fetched_at": "2025-04-12T16:54:11.135276Z"
  },
  {
    "id": "93f955c4-99f8-4846-8380-1d8c44132991",
    "name": "Feed 2",
    "url": "https://www.feed2.com/rss.xml",
    "created_at": "2025-04-10T15:55:52.613684Z",
    "fetched_at": "2025-04-12T16:54:11.13534Z"
  },
  [...]
]
```

### Error:
`[500 Internal Server Error]`
```json
{
  "error": "An unexpected error has occurred, please try again later"
}
```
