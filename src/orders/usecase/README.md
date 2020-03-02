# Usecase

### Create User

- The user id generate by `md5( username )`

```go
	user.ID = utils.MD5(user.Username)

	// len(user.ID) => 32
```

- The user password must `hash( password )`.

```go
	user.Password = utils.Hash(user.Password)

	// len(user.Password) => 64
```

### Auth

1. Check platform if the player exists or not

```go
	if !exist(ctx, user) {
		return nil, system.ErrAuthFailure
	}
```

2. If not found user in service database, create one

```go
	user.ID = utils.MD5(user.Username)

	user.Password = utils.Hash(user.Password)

	err := it.repo.FindByID(ctx, user)

	if err != nil {

		err = it.repo.Insert(ctx, user)

		if err != nil {

			return nil, err
		}
	}
```

3. Generate new JWT Token, and return

```go
	return genToken(user)
```

#### The JWT token within 4 public claims:

1. `iss`: stands for issuer.

2. `iat`: time for this token issue.

3. `exp`: time for this token expire.

4. `user`: service user id.

And will Hash by HS256.

```go
	token := jwt.NewWithClaims(

		jwt.SigningMethodHS256,

		jwt.MapClaims{
			"iss": serviceName,
			"iat": time.Now().Unix(),
			"exp": exp,
			"user": user.ID,
		})
```

4. If Success, status 200 will return

The response body contains:

1. `access_token (required)`:  
   The access token string as issued by the authorization server.

2. `token_type (required)`:  
   The type of token this is, typically just the string “bearer”.

3. `expires_in (recommended)`:  
   If the access token expires, the server should reply with the duration of time the access token is granted for.

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sun, 01 Mar 2020 15:47:55 GMT
Content-Length: 218
Connection: close

{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODMwODEyNzUsImlhdCI6MTU4MzA3NzY3NSwiaXNzIjoic3VubnkuY29tIn0.RAJ7I9Zx0ThkbVj6FkSU7S6GUR7cxMnQrtIKxwCBvDg",
  "token_type": "Bearer",
  "expires_in": 1583081275
}
```
