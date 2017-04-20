FROM golang

ADD . /go/src/github.com/dcaro/skuservice

RUN go install github.com/dcaro/skuservice

ENTRYPOINT /go/bin/skuservice
