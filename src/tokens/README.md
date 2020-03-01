# /accounts

### POST /accounts

Create an account for client.

> - username must be unique.
> - one email can registe multiple username.

```http
POST https://localhost:8081/accounts HTTP/2.0
content-type: application/json

{
    "username": "test",
    "email": "test@gmail.com",
    "password": "123456"
}
```
