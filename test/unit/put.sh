#!/usr/bin/env bash

order_id=9ee9a7bb-f2ae-4ec0-bf17-c2142faa861c

URL=https://localhost:8003/v1/orders/$order_id
MSG_TYPE=Order
PROTO_FILE=model/pb/order.proto
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1ODQzNzM0ODgsImlzcyI6IjhkNzAwODUyLTliMGEtNDA5Zi1hZjFkLWE4NDFkNWNhN2Y0MCIsImp0aSI6IjU5M2VmYjFkLTIyMDAtNGUxMi05ODVjLWY3NmFjMTgzZGIxZiJ9.na_OexBCcGZamPhfOTZgCTR8wiaJN7c2WB7vMbSCqkM


req=$(
cat << EOF
    state: Completed
    completed_at: {
        seconds: 1584156999
        nanos: 386469000
    } 
EOF
)

echo $req \
    | protoc --encode=$MSG_TYPE $PROTO_FILE \
    | curl \
        -X PUT \
        -H "Content-Type: application/protobuf" \
        -H "Authorization: Bearer $TOKEN" \
        --data-binary @- \
        $URL \
    | protoc --decode=$MSG_TYPE $PROTO_FILE