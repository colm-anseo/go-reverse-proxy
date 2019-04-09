#!/bin/bash

cd `dirname "$0"`

build_docker() {
	echo -e "$ docker build -t go-proxy:v1 .\n"
    	           docker build -t go-proxy:v1 .
}

build_local() {
	echo -e "go build:\n"
	go build
	ls -al go-reverse-proxy
	du -kh go-reverse-proxy
}

MODE="local"

if [ "$1" == "docker" ]; then
	MODE="docker"
fi

echo "
build mode: $MODE
"

if [ "$MODE" == "local" ]; then
	echo -e "Note: for a docker build: $0 docker\n"
	build_local
elif [ "$MODE" == "docker" ]; then
	build_docker
else
	echo "unrecognized build mode: '$MODE' - expecting 'local' or 'docker'"
fi
