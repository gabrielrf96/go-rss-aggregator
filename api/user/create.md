# Create a new user

**Endpoint:** `POST` `/v1/user`

**URL params:** None

**Authenticated:** No

## Request

```json
{
    "name": "Jane Doe"
}
```

Constraints:
- `name=[string]`, `required`

## Response

### Success:

`[201 Created]`
```json
{
  "name": "Jane Doe",
  "secret": "secret_key",
  "created_at": "2025-04-12T14:27:32.611602Z"
}
```

**Note:** save the generated secret, as this is the only time you will see it.

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
    "Parameter 'name' is required"
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
