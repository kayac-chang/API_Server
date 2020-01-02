#!/usr/bin/env bash

source "${0%/*}/.env"

docker build -t $PROJECT_NAME . && \

docker rmi $(docker images -f "dangling=true" -q) && \

docker images && \

echo "Build Successed..."
