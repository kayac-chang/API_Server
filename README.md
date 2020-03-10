# Sunny Gaming API Service

### OK

1. POST /token

# Public API

## Game

### GET /games

#### Request

```http
GET https://<service_domain>/v1/games HTTP/2.0
```

#### Respoonse

A successful request returns the HTTP `201 Created` status code.

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

## User

### GET /users

#### Request

#### Respoonse

## Token

### POST /token

#### Request

```http
POST https://<service_domain>/v1/token HTTP/2.0
content-type: application/json

{
  "game": "catpunch",
  "username": "kayac"
}
```

| Parameter | Type   | Description                             |
| --------- | ------ | --------------------------------------- |
| game      | string | The game which user want to access      |
| username  | string | The username or email for user identity |

#### Respoonse

A successful request returns the HTTP `200 OK` status code.

```json
{
  "data": {
    "access_token": "<access_token>",
    "token_type": "Bearer",
    "service_id": "<service_ID>",
    "issued_at": "<issued_at>",
    "links": [
      {
        "rel": "access",
        "method": "GET",
        "href": "https://<game_domain>"
      },
      {
        "rel": "reauthorize",
        "method": "POST",
        "href": "https://<service_domain>/v1/token"
      }
    ]
  }
}
```

| Parameter    | Type   | Description                          |
| ------------ | ------ | ------------------------------------ |
| access_token | string | The jwt token for authentication     |
| token_type   | string | The token type                       |
| service_id   | string | The service if who issued this token |

# Private API

Private API for game service internal network, use `protobuf`.

## User and Authentication

### GET /auth

#### Request

```http
GET https://<service_domain>/v1/auth HTTP/2.0
authorization: Bearer <Access-Token>
```

#### Respoonse

A successful request returns the HTTP `200 OK` status code. Return `User`

| Parameter | Type   | Description              |
| --------- | ------ | ------------------------ |
| user_id   | string | The user's identifier    |
| username  | string | The user's name          |
| balance   | uint64 | The user current balance |

## Orders

### POST /orders

#### Request

```http
POST https://<service_domain>/v1/orders HTTP/2.0
content-type: application/x-google-protobuf
authorization: Bearer <Access-Token>
```

| Parameter | Type   | Description           |
| --------- | ------ | --------------------- |
| user_id   | string | The user's identifier |
| game_id   | string | The game's identifier |
| bet       | uint64 | The game bet          |

#### Respoonse

A successful request returns the HTTP `201 Created` status code.
Return `Order`

### PUT /orders/:order_id

#### Request

```http
PUT https://<service_domain>/v1/orders/:order_id HTTP/2.0
content-type: application/x-google-protobuf
authorization: Bearer <Access-Token>
```

| Parameter    | Type      | Description                    |
| ------------ | --------- | ------------------------------ |
| completed_at | Timestamp | Time when this order completed |

#### Respoonse

A successful request returns the HTTP `202 Accepted` status code.
Return `Order`

### POST /orders/:order_id/sub_orders/

#### Request

#### Respoonse

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
    }
    State state = 2;

    uint64 bet = 3;
    string game_id = 4;
    string user_id = 5;
    google.protobuf.Timestamp created_at = 6;
    google.protobuf.Timestamp updated_at = 7;
    google.protobuf.Timestamp completed_at = 8;
}
```

| Parameter    | Type      | Optional | Description                    |
| ------------ | --------- | -------- | ------------------------------ |
| orderId      | string    | false    | The order's identifier         |
| state        | State     | false    | Current order state            |
| bet          | uint64    | false    | The bet of this order          |
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
