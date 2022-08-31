postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root flight_booking

dropdb: 
	docker exec -it postgres14 dropdb flight_booking

migrateup:
	migrate -path db/migration -database "postgressql://root:root@localhost:5432/flight_booking?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgressql://root:root@localhost:5432/flight_booking?sslmode=disable" -verbose d

.PHONY: postgres createdb dropdb migrateup migratedown