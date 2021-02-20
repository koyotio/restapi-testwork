init:
	make build && make migrate

build:
	docker-compose build app

listen:
	docker-compose up -d

restart:
	docker-compose restart app

test:
	go test -v ./...

migrate:
	docker-compose run --rm app migrate -path ./migrations -database 'postgres://postgres:qwerty@db:5432/postgres?sslmode=disable' up

shutdown:
	docker-compose down