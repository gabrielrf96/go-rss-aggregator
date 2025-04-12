# Get active subscriptions

Returns the list of feeds you are currently subscribed to, including their ID, name and URL, and
the date you subscribed to each of them.

**Endpoint:** `GET` `/v1/subscriptions`

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
    "feed_id": "bce17462-f9cb-46d6-a627-353a6b84e322",
    "name": "Feed 1",
    "url": "https://www.feed1.com/index.xml",
    "created_at": "2025-04-13T16:54:11.135276Z"
  },
  {
    "feed_id": "93f955c4-99f8-4846-8380-1d8c44132991",
    "name": "Feed 2",
    "url": "https://www.feed2.com/rss.xml",
    "created_at": "2025-04-13T17:08:41.163172Z"
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
