test:
	go clean -testcache && go test  ./...
testv:
	go clean -testcache && go test -v ./... 
postgres:
	docker compose up -d
run:
	go run cmd/main.go
prun:
	GIN_MODE=release go run cmd/main.go
sast:
	@echo "Running Snyk security analysis..."
	@SNYK_TOKEN=${SNYK_TOKEN} snyk code test