# APICheck Proxy

This tool launches a local proxy.

## Tool description

`APICheck-proxy` tool runs the popular `MITM Proxy` under the hoods with a custom
*addon* to add APICheck functionality.

## Quick start

```bash
$ docker run --rm -p 8080:8080 -it bbvalabs/apicheck-proxy
Loading script /addons/apicheck_addon.py
Proxy server listening at http://*:8080
```

## Config

You can configure `APICheck-proxy` by using environment vars. Available variables are:

- **APICHECK_PROXY_LISTEN_ADDR**: Proxy listen address.
- **APICHECK_PROXY_LISTEN_PORT**: Proxy listen port
- **APICHECK_PROXY_ALLOWED_HOST**: Only intercept traffic from this host.

Example:

```console
$ docker run --rm -p 9001:9001 -e APICHECK_PROXY_ALLOWED_HOST=mysite.com -e APICHECK_PROXY_LISTEN_PORT=9001 -it bbvalabs/apicheck-proxy
```
