default: run

run:
	go run cmd/main.go

test:
	go test -coverprofile=c.out ./...
	go tool cover -func=c.out
