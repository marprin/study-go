FROM golang:latest
MAINTAINER Marprin <marprin93@gmail.com>

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/gin-gonic/gin