FROM alpine:latest
RUN apk update && apk add --no-cache git && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN wget -qO /bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.14/grpc_health_probe-linux-arm64 && \
    chmod +x /bin/grpc_health_probe
COPY ./cmd/signer/signer /signer
CMD ["/signer"]