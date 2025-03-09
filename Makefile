test:
	bash bin/test.sh
testv:
	bash bin/testv.sh
postgres:
	bash bin/postgres.sh
run:
	bash bin/run.sh
prun:
	bash bin/prun.sh
security:
	@echo "Running Go security analysis..."
	@echo "Running Snyk security analysis..."
	@SNYK_TOKEN=${SNYK_TOKEN} snyk code test
	@echo "Running Gitleaks security analysis..."
	@gitleaks detect .