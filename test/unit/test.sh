#!/usr/bin/env bash

URL=https://localhost:8002/v1/auth
MSG_TYPE=User
PROTO_FILE=model/pb/user.proto
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1ODQyODU4MTEsImlzcyI6IjhkNzAwODUyLTliMGEtNDA5Zi1hZjFkLWE4NDFkNWNhN2Y0MCIsImp0aSI6IjYyODEyNmU2LWM4OWQtNDZhOS1hZDczLTBiOGM2NmIxMGNlNCJ9.CYRzZinL-LSnmhoFm2mSOpMwAg9MaClMm4W2QxsRJN8

curl \
    -H "Content-Type: application/protobuf" \
    -H "Authorization: Bearer $TOKEN" \
    $URL \
    | protoc --decode=$MSG_TYPE $PROTO_FILE