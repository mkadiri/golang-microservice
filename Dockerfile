FROM alpine:3.7

RUN \
    apk --update add \
	    mysql-client \
        ca-certificates \
        curl && \
	update-ca-certificates && \
	curl -L https://github.com/golang-migrate/migrate/releases/download/v3.4.0/migrate.linux-amd64.tar.gz | tar xvz && \
	mv migrate.linux-amd64 migrate

COPY /bin/app /app
RUN chmod +x /app

COPY database/migrations /migrations
RUN chmod -R +x /migrations

COPY database/dev-fixtures /dev-fixtures
RUN chmod -R +x /dev-fixtures

EXPOSE 8000

COPY docker/startup.sh /startup.sh
RUN chmod +x /startup.sh

CMD ["/startup.sh"]