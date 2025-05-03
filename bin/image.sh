#!/bin/bash

# Load environment variables
source $(dirname "$0")/environment.sh

# Validate required environment variables
: "${VERSION:?Usage: VERSION=1.0.0 BOILERPLATES_IMAGE_NAME=your-image-name bash bin/image.sh}" "${BOILERPLATES_IMAGE_NAME:?Usage: VERSION=1.0.0 BOILERPLATES_IMAGE_NAME=your-image-name bash bin/image.sh}"

# Build Docker image
echo "#:[.'.]:> Building Docker image version: $VERSION"
docker build -t "$BOILERPLATES_IMAGE_NAME:$VERSION" -t "$BOILERPLATES_IMAGE_NAME:latest" .
echo "#:[.'.]:> Docker image $BOILERPLATES_IMAGE_NAME:$VERSION built successfully! 🚀"