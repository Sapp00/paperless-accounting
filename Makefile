all: build start


build: build-database build-containers

build-containers:
	docker build -t paperless-accounting .
	docker run paperless-accounting

build-database:
	cd backend/database && ~/go/bin/sqlc generate

start: start-redis start-containers

start-containers:
	docker-compose up -d
	