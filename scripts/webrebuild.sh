#!/bin/bash

# set working directory to project root
cd "$(dirname "$0")"/..

docker-compose -f docker-compose.dev.yml build go
docker-compose -f docker-compose.dev.yml up -d go
