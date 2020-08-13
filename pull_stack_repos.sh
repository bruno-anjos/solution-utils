#!/bin/bash

set -e

cd archimedes || exit
git pull
cd ..

cd archimedesHTTPClient || exit
git pull
cd ..

cd deployer || exit
git pull
cd ..

cd deployer-cli-client || exit
git pull
cd ..

cd scheduler || exit
git pull
cd ..

cd solution-utils || exit
git pull
cd ..