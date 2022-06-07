#!/bin/bash

function cleanup {
  docker-compose -f ./deployments/docker-compose.yaml down
  docker system prune -f
}
trap cleanup EXIT


./scripts/build.sh

docker-compose -f ./deployments/docker-compose.yaml up -d --force-recreate --build

sleep 10

go test ./... -coverprofile=./deployments/coverage.out
go tool cover -html=./deployments/coverage.out -o=./deployments/coverage.html