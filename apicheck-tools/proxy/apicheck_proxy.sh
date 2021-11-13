#!/bin/sh

APICHECK_PROXY_LISTEN_ADDR=${APICHECK_PROXY_LISTEN_ADDR:=0.0.0.0}
APICHECK_PROXY_LISTEN_PORT=${APICHECK_PROXY_LISTEN_PORT:=8080}
APICHECK_PROXY_ALLOWED_HOST=${APICHECK_PROXY_ALLOWED_HOST:=".*"}

mitmdump -s /addons/apicheck_addon.py -q --flow-detail 0 --cert *=/data/certificates/apicheck.pem --listen-host ${APICHECK_PROXY_LISTEN_ADDR} --listen-port ${APICHECK_PROXY_LISTEN_PORT} --allow-hosts ${APICHECK_PROXY_ALLOWED_HOST}

