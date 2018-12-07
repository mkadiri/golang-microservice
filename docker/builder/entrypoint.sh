#!/bin/sh

cd $PROJECT_URL

echo "Get project imports"
go get

echo "Build go binary"
go build -v -o /go-bin/app

echo "Finished"