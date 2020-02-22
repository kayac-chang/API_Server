#!/usr/bin/env bash

# Parse .env
export $(grep -v '^#' .env | xargs -0)

# Check mkcert
brew list mkcert || brew install mkcert

# Make Cert file
dir='./src/.private'

[[ -d $dir ]] || mkdir $dir
(
    cd $dir

    if [[ ! -e $keyfile ]] || [[ ! -e $certfile ]]; then

        mkcert -install

        mkcert \
            -key-file $keyfile \
            -cert-file $certfile \
            example.com localhost 127.0.0.1 ::1
    fi
)

# Build
docker-compose up --build
