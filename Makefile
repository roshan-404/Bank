postgres: 
	docker run --name postgres14 -p 5432:5432  -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb: 
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb: 
	docker exec -it postgres14 dropdb simple_bank

migrateup: 
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
  
migrateup1:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go simpleBank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test mock 