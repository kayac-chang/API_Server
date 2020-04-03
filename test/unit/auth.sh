#!/usr/bin/env bash

source $(dirname $0)/var.sh

URL=$HOST/v1/users/$TOKEN
MSG_TYPE=User
PROTO_FILE=model/pb/user.proto

curl \
    -H "Content-Type: application/protobuf" \
    $URL \
    | protoc --decode=$MSG_TYPE $PROTO_FILE