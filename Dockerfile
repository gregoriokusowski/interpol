FROM golang:1.9-alpine
MAINTAINER Gregorio Kusowski

RUN mkdir -p /go/src/github.com/gregoriokusowski/interpol
ENV GOPATH /go

COPY . /go/src/github.com/gregoriokusowski/interpol/
WORKDIR /go/src/github.com/gregoriokusowski/interpol/cmd/interpol/

RUN go build -i .

RUN mkdir /interpol
WORKDIR /interpol

CMD ["/go/src/github.com/gregoriokusowski/interpol/cmd/interpol/interpol"]
