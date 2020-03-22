#!/usr/bin/env bash

URL=$HOST/v1/auth
MSG_TYPE=User
PROTO_FILE=model/pb/user.proto

curl \
    -H "Content-Type: application/protobuf" \
    -H "Authorization: Bearer $TOKEN" \
    $URL \
    | protoc --decode=$MSG_TYPE $PROTO_FILE