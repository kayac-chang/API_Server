# Usecase

### Create account

The account id generate by `md5( email + username )`

```go
	account.ID = utils.MD5(
		account.Email + account.Username,
	)

	// len(account.ID) => 32
```

The account password must be hash.

```go
	account.Password = utils.Hash(account.Password)

	// len(account.Password) => 64
```
