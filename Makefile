DB_URL = postgresql://admin1:XQy0%2A4%7B%5CM%3DR%27UFAR@34.42.79.184/eduwave?sslmode=disable

postgres:
	docker run -d --name postgres --network eduwave-network -p 5432:5432 -e POSTGRES_USER=pasan -e POSTGRES_PASSWORD=12345 postgres:16-alpine

createdb:
	docker exec -it postgres createdb --username=pasan --owner=pasan postgres

migrate:
	migrate create -ext sql -dir db/migration -seq init_mg

dropdb:
	docker exec -it postgres dropdb --username=pasan postgres

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

dbdocs:
	dbdocs build docs/db.dbml

dbschema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

server:
	go run main.go

.PHONY: postgres createdb migrate dropdb migrateup migratedown sqlc server dbdocs dbschema