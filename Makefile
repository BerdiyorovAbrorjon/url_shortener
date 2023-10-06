.PHONY: postgres createdb opendb dropdb migrate-up migrate-down

postgres:
	docker run --name urlshortener-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it urlshortener-postgres createdb --username=root --owner=root url_shortener

opendb:
	docker exec -it urlshortener-postgres psql -U root -d url_shortener

dropdb:
	docker exec -it postgres12 dropdb  url_shortener

migrate-up:
	migrate -path repository/pgstore/migration -database "postgresql://root:secret@localhost:5432/url_shortener?sslmode=disable" -verbose up

migrate-down:
	migrate -path repository/pgstore/migration -database "postgresql://root:secret@localhost:5432/url_shortener?sslmode=disable" -verbose down

server:
	go run cmd/main.go