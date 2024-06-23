.PHONY: createdb dropdb migrateup migratedown dockerup dockerdown run generatebinary clean runbinary all

DB_NAME = simple-bank
ifeq ($(ENV), test)
    DB_NAME = simple-bank-test
endif

createdb:
	docker exec -it postgres-container createdb --username=root --owner=root $(DB_NAME)

dropdb:
	docker exec -it postgres-container dropdb -U root --if-exists $(DB_NAME) --force

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down

dockerup:
	@echo "\n>>>>>>>>>> Initiating docker build"
	docker compose up -d

dockerdown:
	docker compose down

test:
	# go test -v -cover ./...
	go test -v -cover ./cmd/internals/services/

run:
	go run $(PWD)/cmd/api

genmock:
	@echo "\n>>>>>>>>>> Generating mocks"
	mockgen -source=./cmd/internals/services/account.service.go -destination=./mock/account.service.go IAccountService

generatebinary:
	@echo "\n>>>>>>>>>> Generating binary"
	go build -o bankapi $(PWD)/cmd/api

generatebinaryfordebug:
	@echo "\n>>>>>>>>>> Generating binary"
	go build -gcflags="all=-N -l" -o bankapi $(PWD)/cmd/api

dlv:
	echo "\n>>>>>>>>>> Starting a debugger session"
	dlv exec ./bankapi --headless --listen=:2345 --api-version=2 --accept-multiclient > /tmp/dlv_logs.txt 2>&1 &

dlvconnect:
	echo "\n>>>>>>>>>> Connecting to dlv"
	dlv connect 127.0.0.1:2345

dlvexit:
	echo "\n>>>>>>>>>> Exiting the session"
	echo "exit" > /tmp/dlv_commands.txt && (echo "Y" | dlv connect 127.0.0.1:2345 --init=/tmp/dlv_commands.txt)

clean:
	@echo "\n>>>>>>>>>> Deleting existing binary"
	rm -f $(PWD)/bankapi

runbinary:
	@echo "\n>>>>>>>>>> Starting the app from binary"
	$(PWD)/bankapi

# Define a new target to run multiple scripts sequentially
build: clean generatebinary

startdebuggersession: clean generatebinaryfordebug dlv


start: clean generatebinary runbinary

cleandb: dropdb createdb migrateup

