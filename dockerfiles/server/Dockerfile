# Build
FROM golang:alpine as builder

WORKDIR /bin
COPY ./src .
COPY .env .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o run

# Main

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /bin .

ADD https://raw.githubusercontent.com/eficode/wait-for/master/wait-for .
RUN chmod +x wait-for