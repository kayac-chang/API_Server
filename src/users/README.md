# Token

### POST /token

Generate game service token.

```bash
curl \
    # --http2 \ => optional
    --request POST \
    --url https://{{service_domain}}/token \
    --header 'content-type: application/json' \
    --data '{"username": "kayac","password": "123456"}'
```

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sun, 01 Mar 2020 16:06:22 GMT
Content-Length: 218
Connection: close

{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODMwODIzODIsImlhdCI6MTU4MzA3ODc4MiwiaXNzIjoic3VubnkuY29tIn0.NVmRKuoBd2ZK3o6hKKfWr0OIkJ42yJQjny2-_hDYb8k",
  "token_type": "Bearer",
  "expires_in": 1583082382
}
```
