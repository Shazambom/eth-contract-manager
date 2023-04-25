#!/bin/bash

./scripts/build.sh
docker build -f build/api/Dockerfile -t 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-service/api:local .
docker build -f build/contract-manager/Dockerfile -t 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-service/contract-manager:local .
docker build -f build/signer/Dockerfile -t 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-service/signer:local .
docker build -f build/transaction-manager/Dockerfile -t 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-service/transaction-manager:local .
docker build -f build/web/Dockerfile -t 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-service/web:local .

docker push 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-service/api:local
docker push 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-service/contract-manager:local
docker push 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-service/signer:local
docker push 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-service/transaction-manager:local
docker push 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-service/web:local
