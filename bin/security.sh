#!/bin/bash

# Load environment variables
source $(dirname "$0")/environment.sh

# Validate required environment variables
: "${SNYK_TOKEN:?Usage: SNYK_TOKEN=your-snyk-token bash bin/security.sh}"

# Run security analysis
echo "#:[.'.]:> Running security analysis..."
echo "#:[.'.]:> Running Snyk analysis..."
SNYK_TOKEN="$SNYK_TOKEN" snyk code test
SNYK_TOKEN="$SNYK_TOKEN" snyk test --all-projects --detection-depth=10
echo "#:[.'.]:> Running Gitleaks analysis..."
gitleaks detect --source . --verbose
echo "#:[.'.]:> Security analysis completed successfully! 🚀"