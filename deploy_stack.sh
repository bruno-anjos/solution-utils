#!/bin/bash

bash $(dirname $0)/remove_stack.sh

set -e

SERVICE_NAME=""
OPTIONS=""
PORT=""

function build() {
	cd $SERVICE_NAME || exit
	bash build.sh
	cd ..
}

function run() {
	docker run -d --network=scheduler-network --name=$SERVICE_NAME -p $PORT:$PORT \
	$OPTIONS brunoanjos/$SERVICE_NAME:latest
}

function deploy() {
	build
	run
}

docker system prune -f
docker network create scheduler-network

SERVICE_NAME="archimedes"
PORT="50000"
deploy &

SERVICE_NAME="scheduler"
PORT="50001"
OPTIONS="-v /var/run/docker.sock:/var/run/docker.sock"
deploy &

SERVICE_NAME="deployer"
PORT="50002"
ALTERNATIVES_FILE="$(pwd)/deployer/alternatives.txt"
OPTIONS="--mount type=bind,source=$ALTERNATIVES_FILE,target=alternatives.txt"
deploy &

wait