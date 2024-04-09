postgres:
	docker run --name xyfin -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine


createdb:
	docker exec -it xyfin createdb --username=root --owner=root xyfin

dropdb:
	docker exec -it xyfin dropdb xyfin

migrateup:
	migrate -path database/migrations -database "postgresql://root:secret@localhost:5432/xyfin?sslmode=disable" -verbose up

migrateup1:
	migrate -path database/migrations -database "postgresql://root:secret@localhost:5432/xyfin?sslmode=disable" -verbose up 1

migratedown:
	migrate -path database/migrations -database "postgresql://root:secret@localhost:5432/xyfin?sslmode=disable" -verbose down

migratedown1:
	migrate -path database/migrations -database "postgresql://root:secret@localhost:5432/xyfin?sslmode=disable" -verbose down 1


sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mock_database -destination database/mock/store.go  github.com/arya2004/Xyfin/database/sqlc Store

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

evans:
	evans -r repl --host localhost --port 9090 --package pb --service Xyfin

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock proto