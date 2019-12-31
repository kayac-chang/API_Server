#!/bin/bash

target=$1

curl \
	-X GET \
	-H "Content-Type: application/json" \
	$target
