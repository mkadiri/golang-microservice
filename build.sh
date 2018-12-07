#!/bin/sh

BASE_DIRECTORY=$PWD
GOLANG_BUILDER=mkadiri/golang-builder
GOLANG_TESTER=mkadiri/golang-tester
GOLANG_MICROSERVICE_IMAGE=mkadiri/golang-microservice
PROJECT_URL=github.com/mkadiri/golang-microservice

########################################################################################################################

if [[ "$(docker images -q $GOLANG_TESTER 2> /dev/null)" == "" ]]; then
    echo "************************************************************"
    echo "$GOLANG_TESTER image doesn't exist, build it"
    echo "************************************************************"
    cd $BASE_DIRECTORY/docker/tester
    docker build -t $GOLANG_TESTER .
    cd $BASE_DIRECTORY
fi

echo "************************************************************"
echo "Run tests on $GOLANG_TESTER container"
echo "************************************************************"

docker-compose -f docker-compose-test.yml up -d
docker logs -f mkadiri-golang-tester

########################################################################################################################

if [[ "$(docker images -q $GOLANG_BUILDER 2> /dev/null)" == "" ]]; then
    echo "************************************************************"
    echo "$GOLANG_BUILDER image doesn't exist, build it"
    echo "************************************************************"
    cd $BASE_DIRECTORY/docker/builder
    docker build -t $GOLANG_BUILDER .
fi

########################################################################################################################

echo "************************************************************"
echo "Package binary using $GOLANG_BUILDER container"
echo "************************************************************"

cd $BASE_DIRECTORY

# Mount source folder exactly, e.g. go/:go/ so that we cache imports
# CD to your project directory so that we build from there
# mount a bin directory in the project so that we can copy the file to generate a docker file
docker run -ti --rm --name go-builder \
    -v "$HOME"/go:/go \
    -v "$BASE_DIRECTORY"/bin:/go-bin \
    -e PROJECT_URL=$PROJECT_URL \
    $GOLANG_BUILDER sh


echo "Build docker image"
docker build -t $GOLANG_MICROSERVICE_IMAGE .