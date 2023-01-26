#!/bin/bash

OLDCONTEXT=(kubectl config current-context)
kubectl config use-context docker-desktop
minikube start
eval $(minikube docker-env)

./scripts/build.sh

docker build -f ./build/signer/Dockerfile -t "signer" .
docker build -f ./build/contract-manager/Dockerfile -t "contract-manager" .
docker build -f ./build/transaction-manager/Dockerfile -t "transaction-manager" .
docker build -f ./build/api/Dockerfile -t "contract-api" .
docker build -f ./build/web/Dockerfile -t "contract-web" .

minikube image load signer --daemon --overwrite
minikube image load contract-manager --daemon --overwrite
minikube image load transaction-manager --daemon --overwrite
minikube image load contract-api --daemon --overwrite
minikube image load contract-web --daemon --overwrite

helm repo update

kubectl create namespace pong-dev

helm install nats nats/nats -n pong-dev
sleep 3
helm install signer -f ./helm/signer/values-dev.yaml ./helm/signer -n pong-dev
helm install contract-manager -f ./helm/contract-manager/values-dev.yaml ./helm/contract-manager -n pong-dev
helm install transaction-manager -f ./helm/transaction-manager/values-dev.yaml ./helm/transaction-manager -n pong-dev
helm install contract-api -f ./helm/api/values-dev.yaml ./helm/api -n pong-dev
helm install contract-web -f ./helm/api/values-dev.yaml ./helm/contract-web -n pong-dev

kubectl config use-context "$OLDCONTEXT"