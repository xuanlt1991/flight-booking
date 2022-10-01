postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root flight_booking

dropdb: 
	docker exec -it postgres14 dropdb flight_booking

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/flight_booking?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/flight_booking?sslmode=disable" -verbose d

customerproto:
	protoc -Iproto --go_out=paths=source_relative:./pb --go-grpc_out=paths=source_relative:./pb proto/customer.proto

flightproto:
	protoc -Iproto --go_out=paths=source_relative:./pb --go-grpc_out=paths=source_relative:./pb proto/flight.proto

bookingproto:
	protoc -Iproto --go_out=paths=source_relative:./pb --go-grpc_out=paths=source_relative:./pb proto/booking.proto

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown customerproto flightproto bookingproto proto server