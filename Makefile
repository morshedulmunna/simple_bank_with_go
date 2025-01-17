postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank_db

dropdb:
	docker exec -it postgres16 dropdb simple_bank_db

create_migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank_db?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank_db?sslmode=disable" -verbose down
install_sqlc:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

sqlc:
	sqlc generate

install_dependency: 
	go mod tidy

.PHONY: postgres16 createdb dropdb migrateup migratedown sqlc