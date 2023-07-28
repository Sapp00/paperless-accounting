all:
	docker build -t paperless-accounting .
	docker run paperless-accounting

tailwind:
	cd frontend && pnpm dlx tailwind init -p