#!/usr/bin/env bash

git clone --depth=1 https://github.com/OWASP/crAPI.git
cd crAPI
deploy/docker/build-all.sh
docker-compose -f deploy/docker/docker-compose.yml --compatibility up -d

curl http://localhost:8888
