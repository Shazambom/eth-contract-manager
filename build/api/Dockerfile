FROM alpine:latest
RUN apk update && apk add --no-cache git && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY ./cmd/listener/listener /listener
CMD ["/listener"]