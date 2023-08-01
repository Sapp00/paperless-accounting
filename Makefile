all: build start


build: build-database build-containers

build-containers:
	docker build -t paperless-accounting .
	docker run paperless-accounting

build-database:
	cd backend/database && ~/go/bin/sqlc generate

start: start-redis start-containers

start-containers:
	

start-redis:
	docker run --rm -it --name pa-reddis -p 6379:6379 -d redis /bin/bash