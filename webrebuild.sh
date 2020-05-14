#!/bin/bash

docker-compose -f docker-compose.dev.yml build go
docker-compose -f docker-compose.dev.yml up -d go
