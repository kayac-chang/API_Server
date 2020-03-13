#!/usr/bin/env bash

URL=localhost:8000/auth
MSG_TYPE=User
PROTO_FILE=model/pb/user.proto
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1ODQxMTcwOTEsImlzcyI6InNlcnZpY2UiLCJqdGkiOiI3ZTE0NzdiMi1iMWRjLTRmMDgtODgwNy00ZmIxYjIxYTViZmYifQ.Tqx2BDfkhfUuokH5hl5K5ms5OyL_IsnpeqRJ81BRZY8

curl \
    --verbose \
    -H "Content-Type: application/protobuf" \
    -H "Authorization: Bearer $TOKEN" \
    $URL \
    | protoc --decode=$MSG_TYPE $PROTO_FILE