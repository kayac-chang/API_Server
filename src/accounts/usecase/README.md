# Usecase

### Create account

The account id generate by `hash( email + username )`

```go
	account.ID = utils.MD5(
		account.Email + account.Username,
	)
```

The account password must be hash.

```go
    account.Password = utils.MD5(account.Password)
```
