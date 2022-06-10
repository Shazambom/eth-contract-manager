#!/bin/bash
set -x

echo "Constructing Contract Table"
aws $AWS_ENDPOINT \
    dynamodb create-table \
      --region us-east-1 \
      --table-name Contracts \
      --attribute-definitions \
        AttributeName=Address,AttributeType=S \
        AttributeName=Owner,AttributeType=S \
      --key-schema AttributeName=Address,KeyType=HASH \
      --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
      --global-secondary-indexes \
            "[
              {
                \"IndexName\": \"Owner\",
                \"KeySchema\": [{\"AttributeName\":\"Owner\",\"KeyType\":\"HASH\"}],
                \"Projection\":{
                    \"ProjectionType\":\"ALL\"
                },
                 \"ProvisionedThroughput\": {
                    \"ReadCapacityUnits\": 5,
                    \"WriteCapacityUnits\": 5
                }
              }
            ]"

aws $AWS_ENDPOINT \
    dynamodb describe-table --region us-east-1  --table-name Contracts --output table

echo "Constructing Private Key Table"
aws $AWS_ENDPOINT \
    dynamodb create-table \
      --region us-east-1 \
      --table-name ContractPrivateKeyRepository \
      --attribute-definitions \
        AttributeName=ContractAddress,AttributeType=S \
      --key-schema AttributeName=ContractAddress,KeyType=HASH \
      --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

aws $AWS_ENDPOINT \
    dynamodb describe-table --table-name ContractPrivateKeyRepository --output table --region us-east-1

set +x