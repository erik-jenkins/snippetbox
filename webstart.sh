#!/bin/bash

# This script will start mysql and adminer in a container and run the go server
# locally.

# determine how google chrome should be opened (OS-specific)
launch_chrome_command=""
if [[ "$OSTYPE" == "darwin"* ]]; then
  launch_chrome_command='/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --incognito "http://localhost:4000" "http://localhost:9001?server=mysql&username=snippetlord&db=snippetbox"'
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
  launch_chrome_command='google-chrome --incognito "http://localhost:4000" "http://localhost:9001?server=mysql&username=snippetlord&db=snippetbox"'
fi

# prepare some cleanup stuff (stop docker)
trap cleanup INT
trap cleanup TERM

function cleanup() {
  docker-compose -f docker-compose.dev.yml down -v
}

eval $launch_chrome_command
docker-compose -f docker-compose.dev.yml up -V
