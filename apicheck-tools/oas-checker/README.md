# OpenAPI v3 Checker

This tool checks the endpoint provided using OpenAPI v3 specification
using a set of rules capable of security and semantic checks.

This tool is based on stoplight/spectral.

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
$ acp install oas-checker
[*] Fetching Docker image for tool 'oas-checker'

    Using default tag: latest
    latest: Pulling from bbvalabs/oas-checker
    aad63a933944: Already exists
    dc24e89b59ec: Already exists
    810779e0b9c3: Already exists
    ...
    Status: Downloaded newer image for bbvalabs/oas-checker:latest
    docker.io/bbvalabs/oas-checker:latest

[*] filling environment alias file
```

Finally, activate the default environment and run the tool:

```console
$ curl http://my-company.com/api/entry-point-v3.yml | oas-checker

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

In this case, since al the rules are satisfied,
the tool returns a code of 0 and writes again the definition to its standard
output. In case of detecting any error, the code will be 1 and will write to
its standard error something like:

```console
$ curl http://my-company.com/api/entry-point-v3.yml | oas-checker
Specification contains lint errors: 5

OpenAPI 3.x detected

/tmp/content.oas3
 1:1    error  has-contact         API MUST reference a contact, either url or email: {{path}}
 1:1    error  has-termsOfService  API MUST reference the URL of the Terms of Service {{path}}
 1:1    error  has-x-api-id        API must have an unique identifier in x-api-id {{path}}
 1:1    error  has-x-summary       API must have an one-liner summary field in x-summary {{path}}
 1:1  warning  info-contact        Info object should contain `contact` object.
 1:1  warning  info-description    OpenAPI object info `description` must be present and non-empty string.
 1:1  warning  oas3-api-servers    OpenAPI `servers` must be present and non-empty array.
 1:1    error  oas3-schema         Object should have required property `info`.
 1:1  warning  openapi-tags        OpenAPI object should have non-empty `tags` array.

...
```
