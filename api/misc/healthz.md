# Status / health check

This endpoint can be used to quickly check if the server is running.

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
{
  "status": "ok"
}
```

### Error:

`[Connection refused by server]`

If the server isn't running.
