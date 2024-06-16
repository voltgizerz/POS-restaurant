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
	@mockgen -source=./internal/app/ports/auth_ports.go -destination=./internal/app/mocks/mocks_auth.go -package=mocks
	@mockgen -source=./internal/app/ports/user_ports.go -destination=./internal/app/mocks/mocks_user.go -package=mocks

changelog-gen:
	@auto-changelog
