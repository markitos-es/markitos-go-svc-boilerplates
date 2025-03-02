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
security:
	@echo "Running Go security analysis..."
	@echo "Running Snyk security analysis..."
	@SNYK_TOKEN=${SNYK_TOKEN} snyk code test
	@echo "Running Gitleaks security analysis..."
	@gitleaks detect .