# Get latest posts

Returns a list of the latest posts from all the feeds you are subscribed to, ordered by publication date in descending order.

The number of posts returned by the endpoint is determined by the `API_RETURN_POSTS` setting in the `.env` file.

**Endpoint:** `GET` `/v1/posts`

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
    "feed": "Feed 1",
    "title": "Tea-licious Herbs And Where To Find Them",
    "description": "An in-depth analysis of all the newest herbs being used in teas around the world.",
    "url": "https://www.feed1.com/posts/tea-licious-herbs-and-where-to-find-them-192394/",
    "published_at": "2025-04-08T00:00:00Z"
  },
  {
    "feed_id": "93f955c4-99f8-4846-8380-1d8c44132991",
    "feed": "Feed 2",
    "title": "Why you are not sleeping well",
    "description": "Wondering why you're not sleeping as well as you did before? All that caffeine is to blame.",
    "url": "https://www.feed2.com/articles/health/why-you-are-not-sleeping-well/",
    "published_at": "2025-03-09T00:00:00Z"
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
