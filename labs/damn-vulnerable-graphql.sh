#!/usr/bin/env bash

git clone --depth=1 git@github.com:dolevf/Damn-Vulnerable-GraphQL-Application.git && cd Damn-Vulnerable-GraphQL-Application
docker build -t dvga .
docker run -t -p 5000:5000 -e WEB_HOST=0.0.0.0 dvga

curl http://localhost:5000
