FROM golang:1.20

WORKDIR /usr/src

COPY ../backend .

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# migrate database
RUN cd database/migrations && goose sqlite3 ../../src/:memory up

RUN go build -v -o /usr/local/bin/paccounting ./src/...


EXPOSE 8080

CMD ["paccounting"]