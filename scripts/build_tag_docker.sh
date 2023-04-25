#!/bin/bash

./scripts/build.sh

docker build -f ./build/signer/Dockerfile -t "signer" --no-cache .
docker build -f ./build/contract-manager/Dockerfile -t "contract-manager" --no-cache .
docker build -f ./build/transaction-manager/Dockerfile -t "transaction-manager" --no-cache .
docker build -f ./build/api/Dockerfile -t "contract-api" --no-cache .
docker build -f ./build/web/Dockerfile -t "contract-web" --no-cache .
