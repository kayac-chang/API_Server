#!/bin/bash

curl \
	-i \
	-X POST \
	-H "Content-Type: application/json" \
	-d "@$2" \
	$1
