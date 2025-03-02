test:
	go test ./...
testv:
	go test -v ./...
postgres:
	docker compose up -d
run:
	go run cmd/main.go
prun:
	GIN_MODE=release go run cmd/main.go	