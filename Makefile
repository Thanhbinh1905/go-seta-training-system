include .env
export

initdb:
	docker run --name seta-training -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -p 5432:5432 -d postgres:13.21-alpine3.21

createdb:
	docker exec -it seta-training createdb --username=root --owner=root training-system

dropdb:
	docker exec -it seta-training dropdb --username=root --if-exists training-system

migrateup:
	migrate -path internal/db/migrations -database ${DATABASE_URL} -verbose up

migratedown:
	migrate -path internal/db/migrations -database ${DATABASE_URL} -verbose down

sqlc:
	sqlc generate

gqlgen:
	gqlgen generate

.PHONY: initdb createdb dropdb migrateup migratedown sqlc gqlgen
