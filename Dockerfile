# Build
FROM golang:alpine as builder

COPY ./src /bin

WORKDIR /bin

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o run

# Main

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /bin .

CMD ["./run"]