#!/usr/bin/env bash

URL=localhost:8000/auth
MSG_TYPE=User
PROTO_FILE=model/pb/user.proto
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1ODM4NTc4ODIsImlzcyI6InNlcnZpY2UiLCJqdGkiOiJiYzAxZDU0Ny01MTQxLTQyYzctOTc0ZC1lYmVlYThjOGVlYWEifQ.OjdDe-rrEe_uIzBJ9UFNqvMlKbVnyoct0EldOEBlc-4

curl \
    --verbose \
    -X GET \
    -H "Content-Type: application/protobuf" \
    -H "Authorization: Bearer $TOKEN" \
    $URL \
    | protoc --decode=$MSG_TYPE $PROTO_FILE