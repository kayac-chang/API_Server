#!/usr/bin/env bash
set -e

# Parse .env
export $(grep -v '^#' .env | xargs -0)

# Gen Cert
./tools/cert-gen.sh \
    && echo 'Generate certification success...'

# Gen nginx.conf
(cd tools; envsubst < template.conf > nginx.conf) \
    && echo 'Generate nginx.conf success...'

docker-compose up --build \
    && rm ./tools/nginx.conf