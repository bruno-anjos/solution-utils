#!/bin/bash


docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker network rm scheduler-network

set -e

docker system prune -f
docker network create scheduler-network

cd archimedes || exit
bash ../archimedes/build.sh
cd ..

cd scheduler || exit
bash build.sh
cd ..

cd deployer || exit
bash build.sh
cd ..

docker run -d --network=scheduler-network --name=archimedes -p 50000:50000 brunoanjos/archimedes:latest
docker run -d --network=scheduler-network --name=scheduler -p 50001:50001 -v /var/run/docker.sock:/var/run/docker.sock brunoanjos/scheduler:latest
docker run -d --network=scheduler-network --name=deployer -p 50002:50002 brunoanjos/deployer:latest
