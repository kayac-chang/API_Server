#!/bin/bash

target=$1

data=$(<$2)

curl \
	-X POST \
	-H "Content-Type: application/json" \
	-d "$data" \
	$target
