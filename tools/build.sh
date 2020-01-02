#!/bin/bash

name=${PWD##*/}

docker build -t $name . && \

docker rmi $(docker images -f "dangling=true" -q) && \

docker images && \

echo "Build Successed..."
