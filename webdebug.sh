#!/bin/bash

docker-compose -f docker-compose.dev.yml -f docker-compose.debug.yml build go
docker-compose -f docker-compose.dev.yml -f docker-compose.debug.yml up -d go
