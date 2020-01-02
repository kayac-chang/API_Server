#!/bin/bash

[ -z "$1" ] && echo "username must be supplied" && exit 1

name=${PWD##*/}

username=$1

project=$username/$name

docker login

docker tag $name $project:latest

docker push $project

echo "Publish Successed..."
