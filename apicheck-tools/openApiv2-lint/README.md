# OpenAPI v2 Linter

This tool lints the endpoint provided using OpenAPI v2 specification.

**Tool type**: action

## Tool description

Almost all APIs are published using the OpenAPI specification. In order to
provide a good API definition some best practices should be followed.

A linter is a tool that analizes files written in a specific language looking
for incorrect constructions and suspicious or incorrect code.

This tool accepts an endpoint definition written using OpenAPI v2 as input
(provided as an argument or read from standard input) and will output the same
definition, if no errors are encounterd, or the list of errors detected.

In case of error, a code of 1 is returned, 0 otherwise.

## Quick start

Install the tool:

```console
$ acp install openapiv2-lint
[*] Fetching Docker image for tool 'openapiv2-lint'

    Using default tag: latest
    latest: Pulling from bbvalabs/openapiv2-lint
    aad63a933944: Already exists
    dc24e89b59ec: Already exists
    810779e0b9c3: Already exists
    ...
    Status: Downloaded newer image for bbvalabs/openapiv2-lint:latest
    docker.io/bbvalabs/openapiv2-lint:latest

[*] filling environment alias file
```

Finally, activate the default environment and run the tool:

```console
$ curl http://my-company.com/api/entry-point.yml | openapiv2-lint

# [START swagger]
swagger: "2.0"
info:
  description: "A simple test endpoints API example."
  title: "Endpoints Example"
  version: "1.0.0"
host: "api.endpoints.my-company.com"
# [END swagger]
consumes:
- "application/json"
produces:
- "application/json"
schemes:
# Uncomment the next line if you configure SSL for this API.
#- "https"
- "http"
paths:
...
```

In this case, since the API definition satisfies the OpenAPI v2 specification,
the tool returns a code of 0 and writes again the definition to its standard
output. In case of detecting any error, the code will be 1 and will write to
its standard error something like:

```console
$ curl http://my-company.com/api/entry-point.yml | openapiv2-lint

OpenAPI Specification document is not valid!

ERRORS
  #/paths/~1echo/post: Additional properties not allowed: securit,produce
...
```
