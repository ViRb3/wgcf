FROM golang:1.27.0-alpine AS builder

WORKDIR /src
COPY . .

RUN apk add --no-cache git && \
    go mod download && \
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o "wgcf"

FROM scratch

WORKDIR /

COPY --from=builder "/etc/ssl/cert.pem" "/etc/ssl/cert.pem"
COPY --from=builder "/src/wgcf" "/"

ENTRYPOINT ["/wgcf"]
