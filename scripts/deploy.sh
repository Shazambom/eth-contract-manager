#!/bin/bash

function cleanup {
  docker-compose -f ./deployments/docker-compose.yaml down
  docker system prune -f
}
trap cleanup EXIT


./scripts/build.sh

docker-compose -f ./deployments/docker-compose.yaml up --force-recreate --build

