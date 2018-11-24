FROM golang:latest

ENV \
    GOPATH="/go" \
    CGO_ENABLED=0

RUN mkdir -p /go-bin

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

WORKDIR /go/src/

ENTRYPOINT ["/entrypoint.sh"]