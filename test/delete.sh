#!/bin/bash

target=$1

curl \
	-X DELETE \
	-H "Content-Type: application/json" \
	$target
