FROM golang:1.18-alpine as builer

LABEL maintainer="Sabir_Ali"

WORKDIR /app

COPY . .

# RUN go mod init example.com/go-productmgmt-grpc/productmgmt

CMD l

RUN go run productmgmt_server/productmgmt_server.go

