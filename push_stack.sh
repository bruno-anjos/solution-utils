#!/bin/bash

set -e

cd archimedes || exit
bash build.sh
docker push brunoanjos/archimedes:latest
cd ..

cd scheduler || exit
bash build.sh
docker push brunoanjos/scheduler:latest
cd ..

cd deployer || exit
bash build.sh
docker push brunoanjos/deployer:latest
cd ..