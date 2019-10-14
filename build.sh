#!/usr/bin/env bash
set -e

docker run --rm -v "$(dirname "$PWD")":/go/src -w /go/src/installer golang:1.13-alpine go build -o installer

if [[ "$1" = "" ]]
then
    docker build -t ekaraplatform/installer .     
else
	echo http_proxy \$1:  $1
	echo https_proxy \$2:  $2

	if [[ "$2" = "" ]]
	then
		echo "   using only http_proxy..."
		docker build --build-arg http_proxy="$1" --build-arg https_proxy="$1" -t ekaraplatform/installer .     
	else
		echo "   using http_proxy and https_proxy..."
		docker build --build-arg http_proxy="$1" --build-arg https_proxy="$2" -t ekaraplatform/installer .     
	fi		
fi
