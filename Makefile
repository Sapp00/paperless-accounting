all: build start


build: build-database build-containers migrate-database

build-frontend:
	cd frontend && pnpm build
start-frontend:
	cd frontend && pnpm dev

build-containers:
	docker build -t paperless-accounting .
start-containers:
	docker-compose up -d

build-database:
	cd backend/database && ~/go/bin/sqlc generate
migrate-database:
	cd backend/database/migrations && ~/go/bin/goose sqlite3 ../../src/:memory up

build-graphql:
	cd backend/src/ && go run github.com/99designs/gqlgen generate