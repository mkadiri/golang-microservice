BASE_DIRECTORY=$PWD
GOLANG_BUILDER=mkadiri/golang-builder
GOLANG_TESTER=mkadiri/golang-tester
GOLANG_MICROSERVICE_IMAGE=mkadiri/golang-microservice
PROJECT_URL=github.com/mkadiri/golang-microservice

build:
	bash build.sh

run:
	docker-compose down && docker-compose up -d && docker logs -f mkadiri-golang-microservice

build-test:
	cd ./docker/tester && docker build -t mkadiri/golang-tester .

run-tests:
	docker-compose -f docker-compose-test.yml up -d && \
	docker logs -f mkadiri-golang-tester && \
	docker rm -f mkadiri-golang-tester