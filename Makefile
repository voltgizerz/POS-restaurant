run:
	@gofmt -w ./.. && go run ./cmd/main.go

run-air:
	@air

format:
	@gofmt -w ./..

build:
	@go build -o main ./cmd

test:
	@go test ./... -coverprofile coverage.out -covermode count && go tool cover -func coverage.out

encrypt-gpg:
	@gpg -c .env

decrypt-gpg:
	@gpg -d .env.gpg

mock-gen:
	@mockgen -source=./internal/app/ports/auth_ports.go -destination=./internal/mocks/mocks_auth.go -package=mocks
	@mockgen -source=./internal/app/ports/user_ports.go -destination=./internal/mocks/mocks_user.go -package=mocks

changelog-gen:
	@auto-changelog

# Target to apply migrations
up:
	@goose -dir=./internal/database/migrations mysql "root@tcp(localhost:3306)/db_pos?parseTime=true" up

# Target to reset migrations (if needed)
down:
	@goose -dir=./internal/database/migrations mysql "root@tcp(localhost:3306)/db_pos?parseTime=true" reset

status:
	@goose -dir=./internal/database/migrations mysql "root@tcp(localhost:3306)/db_pos?parseTime=true" status
