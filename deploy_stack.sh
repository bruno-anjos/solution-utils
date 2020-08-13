#!/bin/bash

bash $(dirname $0)/remove_stack.sh

set -e

docker system prune -f
docker network create scheduler-network

cd archimedes || exit
bash build.sh
cd ..

cd scheduler || exit
bash build.sh
cd ..

cd deployer || exit
bash build.sh
cd ..

docker run -d --network=scheduler-network --name=archimedes -p 50000:50000 brunoanjos/archimedes:latest
sleep 2
docker run -d --network=scheduler-network --name=scheduler -p 50001:50001 -v /var/run/docker.sock:/var/run/docker.sock brunoanjos/scheduler:latest
sleep 2
docker run -d --network=scheduler-network --name=deployer -p 50002:50002 brunoanjos/deployer:latest
