FROM golang:1.21-alpine3.18 as builder
WORKDIR /app
COPY .. /app
ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN apk add -U --no-cache ca-certificates
RUN mkdir -p release/linux/amd64 && go build -o release/linux/amd64/plugin .

FROM alpine:3.18
ENV GODEBUG netdns=go

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder release/linux/amd64 /bin

ENTRYPOINT ["/bin/plugin"]