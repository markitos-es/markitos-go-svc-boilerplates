#!/bin/bash

# Load environment variables
source $(dirname "$0")/environment.sh

# Validate required environment variables
: "${SERVICE_NAME:?Usage: SERVICE_NAME=your-service-name ENTITY_NAME=your-entity-name TEMPLATE_PATH=your-template-path bash bin/caas.sh}" \
  "${ENTITY_NAME:?Usage: SERVICE_NAME=your-service-name ENTITY_NAME=your-entity-name TEMPLATE_PATH=your-template-path bash bin/caas.sh}" \
  "${TEMPLATE_PATH:?Usage: SERVICE_NAME=your-service-name ENTITY_NAME=your-entity-name TEMPLATE_PATH=your-template-path bash bin/caas.sh}"

# Clone as a service
echo "#:[.'.]:> Cloning service: $SERVICE_NAME with entity: $ENTITY_NAME"
bash $(dirname "$0")/clone-caas.sh "$SERVICE_NAME" "$ENTITY_NAME"
echo "#:[.'.]:> Service $SERVICE_NAME cloned successfully! 🚀"