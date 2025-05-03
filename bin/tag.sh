#!/bin/bash

# Load environment variables
source $(dirname "$0")/environment.sh

# Validate required environment variables
: "${VERSION:?Usage: VERSION=1.0.0 GIT_REPO_URL=your-repo-url bash bin/tag.sh}" "${GIT_REPO_URL:?Usage: VERSION=1.0.0 GIT_REPO_URL=your-repo-url bash bin/tag.sh}"

# Create Git tag
echo "#:[.'.]:> Creating Git tag: $VERSION"
git tag -a "$VERSION" -m "[TAG:$VERSION] Version $VERSION released"
git push origin "$VERSION"
echo "#:[.'.]:> Git tag $VERSION created and pushed successfully! 🚀"