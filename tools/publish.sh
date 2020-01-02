#!/usr/bin/env bash

[ -z "$1" ] && echo "username must be supplied" && exit 1

source "${0%/*}/.env"

USER_NAME=$1

PROJECT=$USER_NAME/$PROJECT_NAME

docker login && \

docker tag $PROJECT_NAME $PROJECT:latest && \

docker push $PROJECT && \

echo "Publish Successed..."
