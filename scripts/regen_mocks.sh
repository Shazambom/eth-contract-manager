#!/bin/bash

find . -name "interfaces.go" | while read -r i
do
  SPLIT=(${i//// })
  mockgen --source="$i" --destination=./mocks/${SPLIT[1]}.go --package=mocks
done