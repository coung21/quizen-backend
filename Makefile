include .env

server:
	go run main.go

test:
	go test -v -cover -short ./...

newmigrate:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DB)" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DB)" -verbose down

migrateforce:
	migrate -path db/migration -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DB)" -verbose force $(version)

migrateup1:
	migrate -path db/migration -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DB)" -verbose up 1

migratedown1:
	migrate -path db/migration -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DB)" -verbose down 1

.PHONY: server test migrateup migratedown migrateforce migrateup1 migratedown1 newmigrate