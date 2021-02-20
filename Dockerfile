FROM golang:1.15

COPY ./app /var/www/app
COPY ./wait-for-db.sh /usr/local/sbin/wait-for-db.sh

# make wait-for-db.sh executable
RUN chmod +x /usr/local/sbin/wait-for-db.sh

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

WORKDIR /var/www/app

# build go app
RUN go mod download

CMD ["go", "run", "main.go"]