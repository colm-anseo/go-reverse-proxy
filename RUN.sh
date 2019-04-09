#!/bin/bash

cd `dirname "$0"`

# common ENV VAR list fed into both local and docker runs
ENV_FILE="./env.list"

LOCAL_DIR=`pwd`

run_docker() {
	echo "
	docker run \
		-d \
		--env-file "./env.list" \
		-v "${LOCAL_DIR}/tls:/tls" \
		-p 9999:9999 \
		--restart unless-stopped \
		go-proxy:v1
"
	docker run \
		-d \
		--env-file "./env.list" \
		-v "${LOCAL_DIR}/tls:/tls" \
		-p 9999:9999 \
		--restart unless-stopped \
		go-proxy:v1

	# -d : detach
	# -it : interactive with TTY
}

run_local() {
	source "$ENV_FILE"
	export $(cut -d= -f1 ${ENV_FILE})
	echo -e "\n./go-reverse-proxy\n"
	      ./go-reverse-proxy
}

MODE="local"

if [ "$1" == "docker" ]; then
	MODE="docker"
fi

echo -e "\nrun mode: $MODE"

if [ "$MODE" == "local" ]; then
	echo "
Note: to perform a docker run: $0 docker
"
	run_local
elif [ "$MODE" == "docker" ]; then
	run_docker
else
	echo "unrecognized run mode: '$MODE' - expecting 'local' or 'docker'"
fi
