build-test:
	@echo "build golang test image"
	cd docker/tester/ && \
	docker build -t mkadiri/golang-tester .

run-test:
	docker rm -f mkadiri-golang-tester mkadiri-mysql
#	docker-compose -f docker-compose-test.yml up -d
#	docker logs -f mkadiri-golang-tester
#	docker-compose down