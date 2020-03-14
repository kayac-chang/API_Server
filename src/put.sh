#!/usr/bin/env bash

order_id=35b6244e-47c8-42aa-8f76-9f0dbd8ecc40

URL=localhost:8001/orders/$order_id
MSG_TYPE=Order
PROTO_FILE=model/pb/order.proto
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1ODQxNTY5NTksImlzcyI6InNlcnZpY2UiLCJqdGkiOiI4YjRmZjMzMC03OTFkLTQwYmItYTJiZC0zNTAyYjU5OThkN2YifQ.HegLldSNYMJCngANxjCM9PCjemqY5g03EfYGBLeVFt8


req=$(
cat << EOF
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