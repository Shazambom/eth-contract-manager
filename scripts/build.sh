#!/bin/bash

go mod tidy

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/api/api ./cmd/api/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/contract-manager/contract-manager ./cmd/contract-manager/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/listener/listener ./cmd/listener/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/signer/signer ./cmd/signer/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/transaction-manager/transaction-manager ./cmd/transaction-manager/main.go


