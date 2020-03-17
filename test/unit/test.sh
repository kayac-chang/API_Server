#!/usr/bin/env bash

URL=https://localhost:8000/v1/auth
MSG_TYPE=User
PROTO_FILE=model/pb/user.proto
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1ODQ0NTg1MjYsImlzcyI6IjhkNzAwODUyLTliMGEtNDA5Zi1hZjFkLWE4NDFkNWNhN2Y0MCIsImp0aSI6IjdjZjE4M2YzLTE4YTAtNGI2ZC04ZDNhLTExMmJjYWMyMWUxYSJ9.uy9wOzmqCYheLuvbZRzadM7MAnZOFjoX-CE6lCKcFzg

curl \
    -H "Content-Type: application/protobuf" \
    -H "Authorization: Bearer $TOKEN" \
    $URL \
    | protoc --decode=$MSG_TYPE $PROTO_FILE