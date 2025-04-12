# API documentation

## Authentication
Some endpoints are authenticated. To use them, you will need to create a user using the appropriate endpoint *(see the list of endpoints below)*.

The user creation endpoint will generate and return an API secret / key.

Once you have it, you will need to provide it to authenticated endpoints by adding an `Authorization` header to your requests in the form of `Bearer {{SECRET}}`.

## Available endpoints

> [!NOTE]
> Endpoints that require authentication are marked with `[A]`

### User
- [**Get current user**](./user/get.md): `GET` `/v1/user` `[A]`
- [**Create a new user**](./user/create.md): `POST` `/v1/user`

### Feeds
- [**Get available feeds**](./feeds/get.md): `GET` `/v1/feeds` `[A]`
- [**Create a new feed**](./feeds/create.md): `POST` `/v1/feeds` `[A]`

### Subscriptions
- [**Get active subscriptions**](./subscriptions/get.md): `GET` `/v1/subscriptions` `[A]`
- [**Subscribe to a feed**](./subscriptions/subscribe.md): `POST` `/v1/subscriptions/:feed_id` `[A]`
- [**Unsubscribe from a feed**](./subscriptions/unsubscribe.md): `DELETE` `/v1/subscriptions/:feed_id` `[A]`

### Posts
- [**Get latest posts**](./posts/get.md): `GET` `/v1/posts` `[A]`

### Misc
- [**Status / health check**](./misc/healthz.md): `GET` `/v1/healthz`
