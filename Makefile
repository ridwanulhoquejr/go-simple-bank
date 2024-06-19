
# run a postgres container
postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine

# create a database
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

# drop a database
dropdb:
	docker exec -it postgres12 dropdb simple_bank

# run a psql shell
psqlshell:
	docker exec -it postgres12 psql -U root simple_bank


# up migrations
migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

# down migrations
migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

# generate sqlc
sqlc:
	sqlc generate


.PHONY: postgres createdb dropdb psqlshell migrateup sqlc