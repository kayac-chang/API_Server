# Build

FROM golang:latest as builder

COPY ./src /bin

WORKDIR /bin

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o run

# Main

FROM ubuntu:latest

WORKDIR /root/

COPY --from=builder /bin .

CMD ["./run"]

