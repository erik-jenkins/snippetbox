#!/bin/bash

# set working directory to project root
cd "$(dirname "$0")"/..

# create tls directory if it doesn't exist
if [[ ! -d "tls" ]]; then
  mkdir tls
fi

# change to tls directory
cd tls

# generate self signed cert
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost