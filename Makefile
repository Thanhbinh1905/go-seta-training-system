initdb:
	docker run --name seta-training -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -p 5432:5432 -d postgres:13.21-alpine3.21

createdb:
	docker exec -it seta-training createdb --username=root --owner=root training-system

dropdb:
	docker exec -it seta-training dropdb --username=root --if-exists training-system

gensql:
	go run ./cmd/gen-sql

gqlgen:
	gqlgen generate

migratediff:
	atlas migrate diff --env gorm "init_schema"

migrateup:
	atlas migrate apply --env gorm

migratedown:
	atlas migrate apply --env gorm --to "last(-1)"

server:
	go run ./cmd/server/server.go

.PHONY: initdb createdb dropdb migrateup migratedown sqlc gqlgen gensql migratediff server
