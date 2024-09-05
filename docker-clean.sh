#!/bin/bash

docker container stop nlquery-backend nlquery-db

docker rm nlquery-backend nlquery-db
docker rmi nlquery-app
echo "Deleted docker images and containers"
