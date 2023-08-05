all: build start


build: build-database build-containers migrate-database

build-containers:
	docker build -t paperless-accounting .
	docker run paperless-accounting

build-database:
	cd backend/database && ~/go/bin/sqlc generate

migrate-database:
	cd backend/database/migrations && ~/go/bin/goose sqlite3 ../../src/:memory up

start: start-redis start-containers

start-containers:
	docker-compose up -d
