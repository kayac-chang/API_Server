# API_Server

# Games

## `GET /games`

    Find All games

- **Success**

  Code:

  200 OK

  Content:

  ```json
  [
    {
      "id": "<game_id>",
      "name": "<game_name>",
      "href": "<game_href>"
    }

    // ...
  ]
  ```

- **Error**

  Code:

  Content:

- **Test**

  ```bash
  curl -i \
      http://localhost:8080/games
  ```

## `GET /games/:id`

    Find game by ID

- **Success**

  Code:

  200 OK

  Content:

  ```json
  {
    "id": "<game_id>",
    "name": "<game_name>",
    "href": "<game_href>"
  }
  ```

- **Error**

  Code:

  Content:

- **Test**

  ```bash
  curl -i \
      http://localhost:8080/games/20
  ```

## `POST /games`

    Create new games

- **Data**

  ```json
  {
    "name": "<game_name>",
    "href": "<game_href>"
  }
  ```

- **Success**

  Code:

  201 Created

  Content:

  ```json
  {
    "id": "<game_id>",
    "name": "<game_name>",
    "href": "<game_href>"
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
      -d '{ "name": "test", "href": "http://test.com" }' \
      http://localhost:8080/games
  ```

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
      -d '{ "email":"test@gmail.com", "password":"123" }' \
      http://localhost:8080/users
  ```
