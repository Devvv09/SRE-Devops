SHELL := /bin/zsh
include ./config/app.env

run: 
	go run *.go

build:
	go build -o bin/main main.go

build_and_run: build
			   ./bin/main

clean:
	rm bin/main

create_migration:
	migrate create -ext sql -dir ./migration/ -seq init_mg

migration_up:
	migrate -path ./migration/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up
 
migration_down:
	migrate -path ./migration/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down
	