#!/usr/bin/env bash

URL=localhost:8001/orders
MSG_TYPE=Order
PROTO_FILE=model/pb/order.proto
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1ODQxMTk5ODYsImlzcyI6InNlcnZpY2UiLCJqdGkiOiI0YzFlZTk1YS05MjkyLTQxMmQtYjBkYS1lMjA3MWM5NjIwMzIifQ.HZeKLzHQZ4578cyDO84Suc4RucZWx9hopKLUe-tyKNY

user_id=db780439d285e8aba7bf64daba277ec8
game_id=b5ac49be5d3f76cb878671003dbb62ed
bet=600270

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
        -X POST \
        -H "Content-Type: application/protobuf" \
        -H "Authorization: Bearer $TOKEN" \
        --data-binary @- \
        $URL \
    | protoc --decode=$MSG_TYPE $PROTO_FILE