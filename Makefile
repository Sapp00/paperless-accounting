all: build start


build: build-database build-containers migrate-database

build-frontend:
	cd frontend && pnpm build

build-containers:
	docker build -t paperless-accounting .

build-database:
	cd backend/database && ~/go/bin/sqlc generate

migrate-database:
	cd backend/database/migrations && ~/go/bin/goose sqlite3 ../../src/:memory up

start: start-containers

start-frontend:
	cd frontend && pnpm dev

start-containers:
	docker-compose up -d
