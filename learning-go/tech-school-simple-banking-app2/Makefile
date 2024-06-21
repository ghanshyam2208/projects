.PHONY: createdb dropdb migrateup migratedown dockerup dockerdown run generatebinary clean runbinary all

createdb:
	docker exec -it postgres-container createdb --username=root --owner=root simple-bank

dropdb:
	docker exec -it postgres-container dropdb -U root --if-exists simple-bank --force

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/simple-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/simple-bank?sslmode=disable" -verbose down

dockerup:
	@echo "\n>>>>>>>>>> Initiating docker build"
	docker compose up -d

dockerdown:
	docker compose down

run:
	go run $(PWD)/cmd/api

generatebinary:
	@echo "\n>>>>>>>>>> Generating binary"
	go build -o bankapi $(PWD)/cmd/api

clean:
	@echo "\n>>>>>>>>>> Deleting existing binary"
	rm -f $(PWD)/bankapi

runbinary:
	@echo "\n>>>>>>>>>> Starting the app from binary"
	$(PWD)/bankapi

# Define a new target to run multiple scripts sequentially
build: clean generatebinary runbinary
