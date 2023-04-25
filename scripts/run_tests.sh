#!/bin/bash

function cleanup {
  docker-compose -f ./deployments/docker-compose.yaml down
  docker system prune -f
}
trap cleanup EXIT
set -eu


./scripts/build_tag_docker.sh

docker-compose -f ./deployments/docker-compose.yaml up -d --force-recreate && go test ./... -coverprofile=./deployments/coverage.out
go tool cover -html=./deployments/coverage.out -o=./deployments/coverage.html