#!/bin/bash

sudo ./scripts/build_tag_docker.sh

OLDCONTEXT=$(kubectl config current-context)
minikube delete && minikube start
eval $(minikube docker-env)
kubectl config use-context minikube
echo -n "Using the kubernetes context:"
kubectl config current-context
echo -n "....... "
sleep 5
echo "starting."

minikube image load signer --daemon --overwrite
minikube image load contract-manager --daemon --overwrite
minikube image load transaction-manager --daemon --overwrite
minikube image load contract-api --daemon --overwrite
minikube image load contract-web --daemon --overwrite

helm repo update

kubectl create namespace pong-dev


helm install contour bitnami/contour -n pong-dev --set installCRDs=true
helm install dynamodb -f ./helm/dynamodb/values.yaml ./helm/dynamodb -n pong-dev
sleep 3
helm install signer -f ./helm/signer/values-dev.yaml ./helm/signer -n pong-dev
helm install contract-manager -f ./helm/contract-manager/values-dev.yaml ./helm/contract-manager -n pong-dev
helm install transaction-manager -f ./helm/transaction-manager/values-dev.yaml ./helm/transaction-manager -n pong-dev
helm install contract-api -f ./helm/api/values-dev.yaml ./helm/api -n pong-dev
helm install contract-web -f ./helm/web/values-dev.yaml ./helm/web -n pong-dev

echo "Opening ports..."
sleep 10

kubectl port-forward service/signer 8081:8081 -n pong-dev &
kubectl port-forward service/contract-manager 8082:8082 -n pong-dev &
kubectl port-forward service/transaction-manager 8083:8083 -n pong-dev &
kubectl port-forward service/contract-web 8084:8084 -n pong-dev &
kubectl port-forward service/contract-api 8085:8085 -n pong-dev &

kubectl config use-context "$OLDCONTEXT"