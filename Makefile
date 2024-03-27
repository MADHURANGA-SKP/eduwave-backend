DB_URL = postgresql://pasan:12345@localhost:5432/eduwave?sslmode=disable

postgres:
	docker run -d --name eduwave -p 5432:5432 -e POSTGRES_USER=pasan -e POSTGRES_PASSWORD=12345 postgres:16-alpine

createdb:
	docker exec -it eduwave createdb --username=pasan --owner=pasan eduwave

migrate:
	migrate create -ext sql -dir db/migration -seq init_mg

dropdb:
	docker exec -it eduwave dropdb --username=pasan eduwave

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up


migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose up" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb migrate dropdb migrateup migratedown sqlc server 