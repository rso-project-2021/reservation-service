postgres:
	docker run --name reservation_service_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it reservation_service_postgres createdb --username=root --owner=root reservation_service

dropdb:
	docker exec -it reservation_service_postgres dropdb reservation_service

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/reservation_service?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/reservation_service?sslmode=disable" -verbose down

test: 
	go test -cover ./...

.PHONY: postgres, createdb, dropdb, migrateup, migratedown, test