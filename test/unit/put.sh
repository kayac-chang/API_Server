#!/usr/bin/env bash

source $(dirname $0)/var.sh

order_id=39fe7749-7df1-412e-a954-13867f20b436

URL=$HOST/v1/orders/$order_id
MSG_TYPE=Order
PROTO_FILE=model/pb/order.proto

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