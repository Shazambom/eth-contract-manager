#!/bin/bash

go mod tidy

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/captcha/captcha ./cmd/captcha/main.go ./cmd/captcha/signed_request.go ./cmd/captcha/google.go ./cmd/captcha/util.go ./cmd/captcha/response.go ./cmd/captcha/redis.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/listener/listener ./cmd/listener/main.go ./cmd/listener/s3.go ./cmd/listener/redis.go


