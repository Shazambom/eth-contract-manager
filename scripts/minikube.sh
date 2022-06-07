#!/bin/bash

minikube start
eval $(minikube docker-env)

./scripts/build.sh

docker build -f ./build/generator/Dockerfile -t "generator" .

docker build -f ./build/pong/Dockerfile -t "pong" .

docker build -f ./build/artieverseAvatarService/Dockerfile -t "artieverse_avatar_service" .



minikube image load pong --daemon --overwrite
minikube image load generator --daemon --overwrite
minikube image load artieverse_avatar_service --daemon --overwrite

helm repo update

kubectl create namespace pong-dev
kubectl create namespace projectcontour

helm install cert-manager jetstack/cert-manager -n cert-manager --create-namespace --version v1.3.1 --set installCRDs=true
helm install my-release -n agones-system agones/agones --version "1.18" --set "gameservers.namespaces={default,pong-dev}" --create-namespace --set "agones.image.sdk.tag=1.18.0-70d56ad-linux_amd64"
helm install linkerd2 --set-file identityTrustAnchorsPEM=ca.crt --set-file identity.issuer.tls.crtPEM=issuer.crt --set-file identity.issuer.tls.keyPEM=issuer.key --set identity.issuer.crtExpiry=$(date -v+8760H +"%Y-%m-%dT%H:%M:%SZ") linkerd/linkerd2
helm install contour -f ./helm/pong/values.yaml bitnami/contour -n projectcontour --set installCRDs=true
sleep 3
helm install pong -f ./helm/pong/values-dev.yaml ./helm/pong -n pong-dev
helm install artieverse-avatar-service -f ./helm/artieverseAvatarService/values.yaml ./helm/artieverseAvatarService -n artieverse-avatar-service --create-namespace
helm install generator -f ./helm/generator/values.yaml ./helm/generator -n generator --create-namespace