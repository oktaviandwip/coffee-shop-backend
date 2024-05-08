APP=goback
BUILD="./build/$(APP)"
DB_DRIVER=postgres
DB_SOURCE="postgresql://oktav:postgres@localhost/coffee?sslmode=disable&search_path=public"
MIGRATIONS_DIR=./migrations
# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


install:
	go get -u ./... && go mod tidy

build:
	set CGO_ENABLED=0 
	set GOOS=linux 
	go build -o ${BUILD} ./cmd/main.go

test:
	go test -cover -v ./...

migrate-init:
	migrate create -dir ${MIGRATIONS_DIR} -ext sql $(name)

migrate-up:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose up

migrate-down:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose down

migrate-fix:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} force 0