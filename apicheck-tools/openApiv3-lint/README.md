# OpenAPI v3 Linter

This tool lints the endpoint provided using OpenAPI v3 specification.

**Tool type**: action

## Tool description

Almost all APIs are published using the OpenAPI specification. In order to
provide a good API definition some best practices should be followed.

A linter is a tool that analizes files written in a specific language looking
for incorrect constructions and suspicious or incorrect code.

This tool accepts an endpoint definition written using OpenAPI v3 as input
(provided as an argument or read from standard input) and will output the same
definition, if no errors are encounterd, or the list of errors detected.

In case of error, a code of 1 is returned, 0 otherwise.

## Quick start

Install the tool:

```console
$ acp install openapiv3-lint
[*] Fetching Docker image for tool 'openapiv3-lint'

    Using default tag: latest
    latest: Pulling from bbvalabs/openapiv3-lint
    aad63a933944: Already exists
    dc24e89b59ec: Already exists
    810779e0b9c3: Already exists
    ...
    Status: Downloaded newer image for bbvalabs/openapiv3-lint:latest
    docker.io/bbvalabs/openapiv3-lint:latest

[*] filling environment alias file
```

Finally, activate the default environment and run the tool:

```console
$ curl http://my-company.com/api/entry-point-v3.yml | openapiv3-lint

openapi: "3.0.0"
info:
  title: Endpoints Example
  version: 2.0.0
paths:
  /:
    get:
      operationId: listVersionsv2
      summary: List API versions
      responses:
...
```

In this case, since the API definition satisfies the OpenAPI v3 specification,
the tool returns a code of 0 and writes again the definition to its standard
output. In case of detecting any error, the code will be 1 and will write to
its standard error something like:

```console
$ curl http://my-company.com/api/entry-point-v3.yml | openapiv3-lint
Specification contains lint errors: 5

#/info  R: info-contact  D: info object should contain contact object
expected Object { title: 'Endpoints Example', version: '2.0.0' } to have property contact

More information: https://speccy.io/rules/1-rulesets#info-contact
...
```
