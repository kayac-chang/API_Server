#!/bin/bash

data=$(<$1)

echo $data

curl \
	-X POST \
	-H "Content-Type: application/json" \
	-d $data \
	http://localhost:8080/user
