#!/bin/bash

./scripts/build.sh
docker build -f build/api/Dockerfile -t 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-api .
docker build -f build/contract-manager/Dockerfile -t 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-manager .
docker build -f build/signer/Dockerfile -t 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/signer .
docker build -f build/transaction-manager/Dockerfile -t 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/transaction-manager .
docker build -f build/web/Dockerfile -t 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-web .

docker push 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-api
docker push 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/contract-manager
docker push 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/signer
docker push 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/transaction-manager
docker push 227429870588.dkr.ecr.us-west-2.amazonaws.com/platform/web
