#!/bin/bash

find . -name "interfaces.go" | while read -r i
do
  # shellcheck disable=SC2206
  SPLIT=(${i//// })
  mockgen --source="$i" --destination=./mocks/"${SPLIT[1]}".go --package=mocks
done
find . -name "*.pb.go" | while read -r i
do
  # shellcheck disable=SC2206
  SPLIT=(${i//// })
  # shellcheck disable=SC2206
  FILE=(${SPLIT[2]//./ })
  mockgen --source="$i" --destination=./mocks/"proto_${FILE[0]}".go --package=mocks
done
