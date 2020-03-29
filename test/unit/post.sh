#!/usr/bin/env bash

source $(dirname $0)/var.sh

URL=$HOST/v1/orders
MSG_TYPE=Order
PROTO_FILE=model/pb/order.proto

user_id=db780439d285e8aba7bf64daba277ec8
game_id=3e8e582e6f059a5a02c1fb4f05b1a11f
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