postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretpass -d postgres:12-alpine
restartdb:
	sudo docker start postgres12
createdb:
	sudo docker exec -it postgres12 createdb --username=root --owner=root simple_bank
dropdb:
	sudo docker exec -it postgres12 dropdb simple_bank
migrateup:
	sudo migrate -path db/migrations -database "postgresql://root:secretpass@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	sudo migrate -path db/migrations -database "postgresql://root:secretpass@localhost:5432/simple_bank?sslmode=disable" -verbose down
migraterestart:
	sudo migrate -path db/migrations -database "postgresql://root:secretpass@localhost:5432/simple_bank?sslmode=disable" -verbose force 1
sqlc:
	sudo systemctl start apparmor
	snap run sqlc generate
test:
	go test -v -cover ./...
.PHONY: createdb	