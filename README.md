# API_Server

# User

## `POST /users`

    Create new user

- **Data**

  ```json
  {
    "email": "<user_email>",
    "password": "<user_password>"
  }
  ```

- **Success**

  Code:

  201 Created

  Content:

  ```json
  {
    "_id": "5e0c7b0ce6e6feac6f21873f",
    "email": "<user_email>",
    "password": "<user_password>"
  }
  ```

- **Error**

  Code:

  Content:

- **Test**

  ```bash
  curl -i \
      -X POST \
      -H 'Content-Type: application/json' \
      -d '{"email":"test@gmail.com", "password":"123"}' \
      http://localhost:8080/users
  ```
