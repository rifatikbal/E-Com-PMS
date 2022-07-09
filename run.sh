#! /bin/bash

#docker-compose up -d

set -e
set -o pipefail

until $(curl --output /dev/null --silent --fail http://0.0.0.0:8500/v1/kv); do
    echo 'waiting for consul'
    sleep 6
done

curl --request PUT --data-binary @config.example.yml http://localhost:8500/v1/kv/pms

GO111MODULE=on CGO_ENABLED=0 go build -v .

export PMS_CONSUL_URL="127.0.0.1:8500"
export PMS_CONSUL_PATH="pms"

# start application
./E-Com-PMS serve
