#!/bin/sh

cd $PROJECT_URL

echo "Get project imports"
go get


echo "Build go binary"
go build -v -o /go-bin/app


echo "Build go test binary"
#go test -c -o /go-bin/app-test
go test -c ./... -o /go-bin/app-test

echo "Finished"