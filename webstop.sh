#!/bin/bash

docker-compose -f docker-compose.dev.yml down
docker-compose -f docker-compose.dev.yml -f docker-compose.debug.yml down
