# Sunny Gaming API Service

# Public API

## Game

### GET /games

#### Request

```http
GET https://<service_domain>/v1/games HTTP/2.0
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
    "access_token": "<Access_Token>",
    "token_type": "Bearer",
    "service_id": "<Service_ID>",
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

## Authentication

```protobuf
message User{
	string userId     = 1;
	string userName   = 2;
	int64 balance     = 3;
}
```

| Parameter | Type   | Optional | Description              |
| --------- | ------ | -------- | ------------------------ |
| userId    | string | false    | The user's identifier    |
| username  | string | false    | The user's name          |
| balance   | int64  | false    | The user current balance |

### GET /auth

#### Request

```http
GET https://<service_domain>/v1/auth HTTP/2.0
authorization: Bearer <Access-Token>
```

#### Respoonse

A successful request returns the HTTP `200 OK` status code.
Return `User`

## Orders

```protobuf
message Order{
    string  orderId         =   1;
    string  state           =   2;
    int64   bet             =   3;
    string  gameId          =   4;
    string  userId          =   5;
    string  createdAt       =   6;
    string  updatedAt       =   7;
    string  completedAt     =   8;
}
```

| Parameter   | Type    | Optional | Description                    |
| ----------- | ------- | -------- | ------------------------------ |
| orderId     | string  | true     | The order's identifier         |
| state       | string  | true     | Current order state            |
| bet         | float64 | false    | The bet of this order          |
| gameId      | string  | false    | The game's identifier          |
| userId      | string  | false    | The user's identifier          |
| createdAt   | string  | true     | Time when this order created   |
| updatedAt   | string  | true     | Time when this order updated   |
| completedAt | string  | true     | Time when this order completed |

### POST /orders

#### Request

```http
POST https://<service_domain>/v1/orders HTTP/2.0
content-type: application/x-google-protobuf
authorization: Bearer <Access-Token>
```

| Parameter | Type   | Description           |
| --------- | ------ | --------------------- |
| userId    | string | The user's identifier |
| gameId    | string | The game's identifier |
| bet       | int64  | The game bet          |

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

| Parameter    | Type   | Description                    |
| ------------ | ------ | ------------------------------ |
| completed_at | string | Time when this order completed |

#### Respoonse

A successful request returns the HTTP `202 Accepted` status code.
Return `Order`

### POST /orders/:order_id/sub_orders/

#### Request

#### Respoonse
