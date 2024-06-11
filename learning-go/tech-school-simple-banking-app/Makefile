createdb:
	docker exec -it postgres-container createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres-container dropdb -U postgres --if-exists simple_bank --force

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

start:
	go run main.go

.PHONY: createdb