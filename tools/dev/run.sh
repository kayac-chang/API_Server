#!/usr/bin/env sh

docker-compose -f docker-compose.yml -f production.yml up --build -d
