#!/bin/bash
CONTAINER_NAME="nlquery-db"

# Check if the container is running
if [ "$(docker ps -q -f name=${CONTAINER_NAME})" ]; then
    echo "Container '${CONTAINER_NAME}' is running. Stopping it now..."
    
    # goose down
    echo "Running goose down migrations..."
    goose postgres "postgres://nlquery:postgres@localhost:5432/nlquery" -dir "server/sql/schema" down

    # Stop the container
    docker stop "${CONTAINER_NAME}"
    
    # Optionally, remove the container
    # Uncomment the following line if you want to remove the container after stopping
    docker rm "${CONTAINER_NAME}"

    echo "Container '${CONTAINER_NAME}' has been stopped and removed."
else
    echo "Container '${CONTAINER_NAME}' is not running."
fi

# Check if the container is running
CONTAINER_NAME = "nlquery-backend"

if [ "$(docker ps -q -f name=${CONTAINER_NAME})" ]; then
    echo "Container '${CONTAINER_NAME}' is running. Stopping it now..."
    
    # Stop the container
    docker container stop "${CONTAINER_NAME}"
    
    # Optionally, remove the container
    docker rm "${CONTAINER_NAME}"

    echo "Container '${CONTAINER_NAME}' has been stopped and removed."
else
    echo "Container '${CONTAINER_NAME}' is not running."
fi

# docker container stop nlquery-backend nlquery-db

# docker rm nlquery-backend nlquery-db
docker rmi nlquery-app
echo "Deleted docker images and containers"
