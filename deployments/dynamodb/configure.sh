#!/bin/bash
echo "########### Creating profile ###########"
aws configure set aws_access_key_id xxx
aws configure set aws_secret_access_key yyy
aws configure set region us-east-1

echo "########### Listing profile ###########"
aws configure list