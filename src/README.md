# Sunny Gaming API Service

# Public API

## Game

### Get Games Links

#### Request

```http
GET http://{{host}}/v1/games HTTP/1.1
```

#### Respoonse

A successful request returns the HTTP `200 OK` status code.

```json
{
  "data": {
    "links": [
      {
        "rel": "self",
        "method": "GET",
        "href": "https://<service_domain>/v1/games"
      },
      {
        "rel": "<game_name>",
        "method": "GET",
        "href": "https://<game_domain>/"
      }
      // ...
    ]
  }
}
```

### Get Game By Name

#### Request

```http
GET http://{{host}}/v1/games/{{game}} HTTP/1.1
```

#### Respoonse

A successful request returns the HTTP `200 OK` status code.

```json
{
  "data": {
    "game_id": "b5ac49be5d3f76cb878671003dbb62ed",
    "name": "catpunch",
    "href": "https://catpunch",
    "category": "slot"
  }
}
```

### Create Game

#### Request

```http
POST http://{{host}}/v1/games HTTP/1.1
Content-Type: application/json
Authorization: Bearer {{token}}

{
    "name": {{game}},
    "href": {{href}},
    "category": {{category}}
}
```

#### Respoonse

A successful request returns the HTTP `201 Created` status code.

```json
{
  "data": {
    "game_id": "b5ac49be5d3f76cb878671003dbb62ed",
    "name": "catpunch",
    "href": "https://catpunch",
    "category": "slot"
  }
}
```

### Modify Game By Name

#### Request

```http
PUT http://{{host}}/v1/games/{{game}} HTTP/1.1
Content-Type: application/json
Authorization: Bearer {{token}}

{
    "name": {{game}},
    "href": {{href}},
    "category": {{category}}
}
```

#### Respoonse

A successful request returns the HTTP `202 Accepted` status code.

```json
{
  "data": {
    "game_id": "b5ac49be5d3f76cb878671003dbb62ed",
    "name": "catpunch",
    "href": "https://catpunch.io",
    "category": "slot"
  }
}
```

## Token

### Create new Token

#### Request

```http
POST http://{{host}}/v1/tokens HTTP/1.1
Content-Type: application/json
Session: {{session}}

{
    "game": {{game}},
    "username": {{username}}
}
```

- Header

| Parameter | Type   | Description          |
| --------- | ------ | -------------------- |
| session   | string | The session identity |

- Body

| Parameter | Type   | Description                             |
| --------- | ------ | --------------------------------------- |
| game      | string | The game which user want to access      |
| username  | string | The username or email for user identity |

#### Respoonse

A successful request returns the HTTP `201 Created` status code.

```json
{
  "data": {
    "links": [
      {
        "rel": "access",
        "method": "GET",
        "href": "https://<game_domain>?access_token=<token>"
      },
      {
        "rel": "reauthorize",
        "method": "POST",
        "href": "https://<service_domain>/v1/token"
      }
    ],
    "token": {
      "access_token": "<access_token>",
      "token_type": "Bearer",
      "service_id": "<service_ID>",
      "issued_at": "<issued_at>"
    }
  }
}
```

| Parameter    | Type   | Description                          |
| ------------ | ------ | ------------------------------------ |
| access_token | string | The jwt token for authentication     |
| token_type   | string | The token type                       |
| service_id   | string | The service if who issued this token |
| issued_at    | string | Time when this token created         |

## Admin

### Create an Admin account

#### Request

```http
POST http://{{host}}/v1/admins HTTP/1.1
Content-Type: application/json

{
    "secret": {{secret}},
    "email": {{email}},
    "username": {{username}},
    "password": {{password}},
    "organization": {{organization}}
}
```

- Header

| Parameter    | Type   | Description                             |
| ------------ | ------ | --------------------------------------- |
| Content-Type | string | indicate the media type of the resource |

- Body

| Parameter    | Type   | Description                                    |
| ------------ | ------ | ---------------------------------------------- |
| secret       | string | The secret key dispatch by game provider       |
| email        | string | The account's email                            |
| username     | string | The account's username                         |
| password     | string | The account's password                         |
| organization | string | The organization which the account attached to |

#### Respoonse

A successful request returns the HTTP `201 Created` status code.

```json
{
  "data": {
    "admin_id": "04773c00cef59479d30bf9f026c8fed9",
    "email": "egg734631@gmail.com",
    "username": "kayac",
    "organization": "sunny"
  }
}
```

- data

| Parameter    | Type   | Description                                    |
| ------------ | ------ | ---------------------------------------------- |
| admin_id     | string | The account's id                               |
| email        | string | The account's email                            |
| username     | string | The account's username                         |
| organization | string | The organization which the account attached to |

### Create Admin Token

#### Request

```http
POST http://{{host}}/v1/admins/tokens HTTP/1.1
Content-Type: application/json

{
    "secret": {{secret}},
    "email": {{email}},
    "password": {{password}},
}
```

#### Respoonse

A successful request returns the HTTP `201 Created` status code.

```json
{
  "data": {
    "access_token": "<access_token>",
    "token_type": "Bearer",
    "service_id": "<service_ID>",
    "issued_at": "<issued_at>"
  }
}
```

- data

| Parameter    | Type   | Description                          |
| ------------ | ------ | ------------------------------------ |
| access_token | string | The jwt token for authentication     |
| token_type   | string | The token type                       |
| service_id   | string | The service if who issued this token |
| issued_at    | string | Time when this token created         |

# Internal API

Internal API for game service internal network, use `protobuf` for transform.

## User

### Authenticate Token

#### Request

```http
GET http://{{host}}/v1/users/{{token}} HTTP/1.1
Content-Type: application/protobuf
```

#### Respoonse

A successful request returns the HTTP `200 OK` status code. Return `User`

| Parameter | Type   | Description              |
| --------- | ------ | ------------------------ |
| user_id   | string | The user's identifier    |
| username  | string | The user's name          |
| balance   | uint64 | The user current balance |

## Orders

### Create new Order

#### Request

```http
POST http://{{host}}/v1/orders HTTP/1.1
content-type: application/protobuf
authorization: Bearer {{token}}
```

| Parameter | Type   | Description           |
| --------- | ------ | --------------------- |
| user_id   | string | The user's identifier |
| game_id   | string | The game's identifier |
| bet       | uint64 | The game bet          |

#### Respoonse

A successful request returns the HTTP `201 Created` status code.
Return `Order`

### Modify Order by ID

#### Request

```http
PUT http://{{host}}/v1/orders/{{order_id}} HTTP/1.1
content-type: application/protobuf
authorization: Bearer {{token}}
```

| Parameter    | Type      | Description                    |
| ------------ | --------- | ------------------------------ |
| completed_at | Timestamp | Time when this order completed |

#### Respoonse

A successful request returns the HTTP `202 Accepted` status code.
Return `Order`

# Proto3 Schema

## Order

```protobuf
syntax = "proto3";

import "google/protobuf/timestamp.proto";

message Order {
    string order_id = 1;

    enum State {
        Pending = 0;
        Completed = 1;
        Rejected = 2;
        Issue = 3;
    }
    State state = 2;

    uint64 bet = 3;
    uint64 win = 4;

    string game_id = 5;
    string user_id = 6;
    google.protobuf.Timestamp created_at = 7;
    google.protobuf.Timestamp updated_at = 8;
    google.protobuf.Timestamp completed_at = 9;
}
```

| Parameter    | Type      | Optional | Description                    |
| ------------ | --------- | -------- | ------------------------------ |
| orderId      | string    | false    | The order's identifier         |
| state        | State     | false    | Current order state            |
| bet          | uint64    | false    | The bet of this order          |
| win          | uint64    | false    | The win of this order          |
| game_id      | string    | false    | The game's identifier          |
| user_id      | string    | false    | The user's identifier          |
| created_at   | Timestamp | true     | Time when this order created   |
| updated_at   | Timestamp | true     | Time when this order updated   |
| completed_at | Timestamp | true     | Time when this order completed |

## User

```protobuf
syntax = "proto3";

import "google/protobuf/timestamp.proto";

message User {
    string user_id = 1;
    string username = 2;
    uint64 balance = 3;
    google.protobuf.Timestamp created_at = 4;
    google.protobuf.Timestamp updated_at = 5;
}
```

| Parameter  | Type      | Optional | Description                 |
| ---------- | --------- | -------- | --------------------------- |
| user_id    | string    | false    | The user's identifier       |
| username   | string    | false    | The user's name             |
| balance    | uint64    | false    | The user current balance    |
| created_at | Timestamp | true     | Time when this user created |
| updated_at | Timestamp | true     | Time when this user created |

## Error

```protobuf
syntax = "proto3";

message Error {
    uint32 code = 1;
    string name = 2;
    string message = 3;
}
```

| Parameter | Type   | Optional | Description       |
| --------- | ------ | -------- | ----------------- |
| code      | uint32 | false    | The error code    |
| name      | string | false    | The error name    |
| message   | string | false    | The error message |
