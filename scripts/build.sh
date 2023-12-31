#!/bin/sh

go mod tidy

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/api/api ./cmd/api/main.go ./cmd/api/config.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/web/web ./cmd/web/main.go ./cmd/web/config.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/contract-manager/contract-manager ./cmd/contract-manager/main.go  ./cmd/contract-manager/config.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/signer/signer ./cmd/signer/main.go  ./cmd/signer/config.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/transaction-manager/transaction-manager ./cmd/transaction-manager/main.go  ./cmd/transaction-manager/config.go


