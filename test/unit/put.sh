#!/usr/bin/env bash

order_id=9053fe45-1e64-4cf5-a35c-d173327a3b5a

URL=https://localhost:8000/v1/orders/$order_id
MSG_TYPE=Order
PROTO_FILE=model/pb/order.proto
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1ODQ0NTkwNDIsImlzcyI6IjhkNzAwODUyLTliMGEtNDA5Zi1hZjFkLWE4NDFkNWNhN2Y0MCIsImp0aSI6IjM5MTUyYTM0LWRiM2ItNGIxOC05ZmY3LTFiMmYwZjZlY2QyOCJ9.bPkmXjCdWdXpTQITqT9XOh5lzvTfV6cEQ9Nm7kbvXyM


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