#!/usr/bin/bash

git clone --depth=1 https://github.com/snoopysecurity/dvws-node.git
cd dvws-node
docker-compose up -d

curl http://localhost:80
curl http://localhost;9090

