#!/bin/bash

# This script will start mysql and adminer in a container and run the go server
# locally.

# prepare some cleanup stuff (stop docker)
trap cleanup INT
trap cleanup TERM

function cleanup() {
  docker-compose down -v
}

docker-compose -f docker-compose.yml up -V -d
google-chrome --incognito "http://localhost:4000" "http://localhost:9001?username=root&db=snippetbox"
go run ./cmd/web
