FROM golang:1.18-alpine as builer

LABEL maintainer="Sabir_Ali"

WORKDIR /app

CMD apt-get install bash

COPY . .

# RUN go mod init example.com/go-productmgmt-grpc/productmgmt

#CMD l

WORKDIR /app/productmgmt_server

RUN DATABASE_URL=postgresql://postgresql:mypass@172.17.0.2:5432/products?sslmode=disable go run .

#RUN go run productmgmt_server/productmgmt_server.go

