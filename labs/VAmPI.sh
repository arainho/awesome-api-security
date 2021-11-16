#!/usr/bin/env bash

git clone --depth=1 https://github.com/erev0s/VAmPI
cd VAmPI
docker build -t vampi_docker:latest .
docker run -d -p 5000:5000 vampi_docker:latest

curl http://localhost:5000
