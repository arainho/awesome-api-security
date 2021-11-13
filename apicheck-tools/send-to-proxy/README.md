# APICheck Send to proxy

This tool sends the APICheck Request to a remote proxy.

## Quick start

Install the APICheck tools:

- send-to-proxy
- apicheck-curl

```bash
acp install send-to-proxy
acp install apicheck-curl
```

Finally, run the tool:

```console
$ acurl http://my-company.com/api/entry-point | send-to-proxy http://localproxy:9000
[*] Request sent: 'http://my-company.com/api/entry-point'
```


## Using the quiet mode

`Send-to-proxy` accepts a parameter containing the URL of the proxy to send requests to, and the `-q` (`--quiet`) option to suppress output data.

```bash
acurl http://my-company.com/api/entry-point | send-to-proxy -q http://localproxy:9000
```
