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

changelog-gen:
	@auto-changelog
