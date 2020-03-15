#!/usr/bin/env bash

URL=https://localhost:8003/v1/orders
MSG_TYPE=Order
PROTO_FILE=model/pb/order.proto
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1ODQyOTEzNTEsImlzcyI6IjhkNzAwODUyLTliMGEtNDA5Zi1hZjFkLWE4NDFkNWNhN2Y0MCIsImp0aSI6ImUxMzczNzJhLTM5ZjItNDM5OC05NjRlLWM3MjJjNWI0ODM1MiJ9.NHFLbWC1QaZtqRuLapjy6VLutjgIgNSbS0W-zd2ro3o

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
        -H "Content-Type: application/protobuf" \
        -H "Authorization: Bearer $TOKEN" \
        --data-binary @- \
        $URL \
    | protoc --decode=$MSG_TYPE $PROTO_FILE