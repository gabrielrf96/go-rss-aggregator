# Get current user

**Endpoint:** `GET` `/v1/user`

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
  "name": "Jane Doe",
  "created_at": "2025-04-12T14:27:32.611602Z",
  "updated_at": "2025-04-12T14:27:32.611602Z"
}
```

### Error:
`[500 Internal Server Error]`
```json
{
  "error": "An unexpected error has occurred, please try again later"
}
```
