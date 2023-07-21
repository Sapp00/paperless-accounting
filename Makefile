all:
	docker build -t paperless-accounting .
	docker run paperless-accounting