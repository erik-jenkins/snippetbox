#!/bin/bash

# set working directory to project root
cd "$(dirname "$0")"/..

# determine how google chrome should be opened (OS-specific)
launch_chrome_command=""
if [[ "$OSTYPE" == "darwin"* ]]; then
  launch_chrome_command='/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --incognito "https://localhost:4000" "http://localhost:9001?server=mysql&username=snippetlord&db=snippetbox"'
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
  launch_chrome_command='google-chrome --incognito "https://localhost:4000" "http://localhost:9001?server=mysql&username=snippetlord&db=snippetbox"'
fi

# prepare some cleanup stuff (stop docker)
trap cleanup INT
trap cleanup TERM

function cleanup() {
  docker-compose -f docker-compose.dev.yml down -v
}

eval $launch_chrome_command
docker-compose -f docker-compose.dev.yml up -V
