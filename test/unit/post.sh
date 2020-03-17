#!/usr/bin/env bash

URL=https://localhost:8000/v1/orders
MSG_TYPE=Order
PROTO_FILE=model/pb/order.proto
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1ODQ0NTkwNDIsImlzcyI6IjhkNzAwODUyLTliMGEtNDA5Zi1hZjFkLWE4NDFkNWNhN2Y0MCIsImp0aSI6IjM5MTUyYTM0LWRiM2ItNGIxOC05ZmY3LTFiMmYwZjZlY2QyOCJ9.bPkmXjCdWdXpTQITqT9XOh5lzvTfV6cEQ9Nm7kbvXyM

user_id=db780439d285e8aba7bf64daba277ec8
game_id=b5ac49be5d3f76cb878671003dbb62ed
bet=10

req=$(
cat << EOF
    user_id: "$user_id"
    game_id: "$game_id"
    bet: $bet
EOF
)

echo $req \
    | protoc --encode=$MSG_TYPE $PROTO_FILE \
    | curl \
        -v \
        -H "Content-Type: application/protobuf" \
        -H "Authorization: Bearer $TOKEN" \
        --data-binary @- \
        $URL \
    | protoc --decode=$MSG_TYPE $PROTO_FILE