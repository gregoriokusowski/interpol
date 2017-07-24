FROM golang:1.9-alpine as builder
MAINTAINER Gregorio Kusowski

RUN mkdir -p /go/src/github.com/gregoriokusowski/interpol
ENV GOPATH /go

COPY . /go/src/github.com/gregoriokusowski/interpol/
WORKDIR /go/src/github.com/gregoriokusowski/interpol/cmd/interpol/

RUN go build -i .

FROM alpine:latest

RUN mkdir /interpol
WORKDIR /interpol/
COPY --from=builder /go/src/github.com/gregoriokusowski/interpol/cmd/interpol/interpol /root

CMD ["/root/interpol"]
