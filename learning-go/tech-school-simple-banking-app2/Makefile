createdb:
	docker exec -it postgres-container createdb --username=root --owner=root simple-bank

dropdb:
	docker exec -it postgres-container dropdb -U root --if-exists simple-bank --force

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/simple-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/simple-bank?sslmode=disable" -verbose down

start:
	go run main.go

.PHONY: createdb