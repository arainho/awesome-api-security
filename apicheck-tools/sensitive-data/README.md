# APICheck Sensitive Data

This tool analyzes a Request / Response content and headers and tries to find sensitive data
in both the request and the response.

## Motivation

Sometimes, APIs can return sensitive data for some entry-points. Sensitive data
could be user information, some data from the business logic, internal IPs, and
such.

Detecting this data is a hard task as it depends on the application's business
logic. With `APICheck Sensitive Data`, you can configure a set of rules for
analyzing the Request / Response object of an application.

Rules are provided in a simple `YAML` file that could be hosted in a remote
place or in the local filesystem.

## Quick start

Install APICheck tools:

- sensitive-data
- acurl

```bash
acp install sensitive-data
acp install acurl
```

Finally run tools:

```console
$ acurl http://my-company.com/api/entry-point | sensitive-data

http://my-company.com
---------------------

 > rule           -> core-001
 > where          -> request
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password

```

You can check how to install `APICheck Package Manager <https://bbva.github.io/apicheck/docs/quick-start>`

# Rules format

Rules are provided in a YAML file with the following format:

```yaml
- id: core-001
  description: Find 'password' keyword in flow data
  regex: '([pP][aA][sS][sS][wW][oO][rR][dD])'
  severity: Medium  # Allowed values: Low, Medium, High
  searchIn: All  # Allowed values: Response, Request, Headers, All
- id: core-002
  description: HTTP Headers contains information about backedn
  regex: '([Xx]-[Pp][Oo][Ww][Ee][Rr][Ee][Dd]-[Bb][Yy])'
  severity: Low  # Allowed values: Low, Medium, High
  searchIn: Headers  # Allowed values: Response, Request, Headers, All
```

The above example is from the core rules file: `core.yaml`:

- Severity values allowed are: High, Medium, Low
- searchIn: Allows to search in the HTTP Request / Response and HTTP Headers in Request / Response  

# Running as a Service

`APIcheck Sensitive Data` can be run as a service, but only when running as a
standalone docker container

```console
$ docker run --rm -p 9000:9000 bbvalabs/sensitive-data --server 0.0.0.0:9000
[2020-05-08 10:18:01 +0000] [1] [INFO] Goin' Fast @ http://0.0.0.0:9000
[2020-05-08 10:18:01 +0000] [1] [INFO] Starting worker [1]
```

It provides a single entry-point (`/apicheck/sensitive-data`) that can be
accessed by using the HTTP POST method.

POST data must be a valid APICheck data object.

Example:

```bash
acurl http://myservice.com > query.json
curl -X POST -H "Content-Type: application/json" -X POST --d @query.json http://localhost:9000/apicheck/sensitive-data
```

# Examples

## Common examples

You can run these examples by using both the `APICheck Package Manager` or
directly with Docker.

### Core rules

By default, the tool has a set of embedded rules: **core rules**. Unless you
provide a rules file by leveraging the `-r` parameter, these core rules will be
used:

```console
$ acurl http://my-company.com/api/entry-point | sensitive-data

http://my-company.com
---------------------

 > rule           -> core-001
 > where          -> request
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password

```

### With a remote rules file

You can also use rules stored in remote files. `APIcheck Sensitive Data` will
download the rules file prior to execution.

```console
$ acurl http://my-company.com/api/entry-point | sensitive-data -r http://127.0.0.1:9999/rules/rules.yaml

http://my-company.com
---------------------

 > rule           -> core-001
 > where          -> request
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password

```

### With many rules files

```console
$ acurl http://my-company.com/api/entry-point | sensitive-data -r http://127.0.0.1:9999/rules/java.yaml -r http://127.0.0.1:9999/rules/credentials.yaml

http://my-company.com
---------------------

 > rule           -> java-001
 > where          -> request
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password

```

### Filtering false positives

Sometimes, rules generate false positives. In these cases, we can remove them by
using the `-i` parameter and a comma-separated list of `rule ID`, or by
providing an ignore file, containing one `rule ID` per line, with the `-F`:

For instance:

```console
$ acurl http://my-company.com/api/entry-point | sensitive-data -r http://127.0.0.1:9999/rules/rules2.yaml

http://my-company.com
---------------------

 > rule           -> core-001
 > where          -> request
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password

http://my-company.com
---------------------

 > rule           -> myrules2-001
 > where          -> response
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password

http://my-company.com
---------------------

 > rule           -> myrules2-002
 > where          -> responseHeaders
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> Another password

```

If you want to remove the errors generated by **myrules2-002** and **core-001**
rules, you can do so with the `-i` parameter:

```console
$ acurl http://my-company.com/api/entry-point | sensitive-data -r http://127.0.0.1:9999/rules/rules2.yaml -it core-001,myrules2-002

http://my-company.com
---------------------

 > rule           -> core-001
 > where          -> request
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password

```

Or using an ignore file:

```console
$ cat ignore-file
core-001
myrules2-002
$ acurl http://my-company.com/api/entry-point | sensitive-data -r http://127.0.0.1:9999/rules/rules2.yaml -F ignore-file

http://my-company.com
---------------------

 > rule           -> core-001
 > where          -> request
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password

```

## Running with Docker

You can also use `APIcheck Sensitive Data` by running by hand with Docker:

### Running without parameters

```console

$ acurl http://my-company.com/api/entry-point | docker run --rm -it bbvalabs/sensitive-data

http://my-company.com
---------------------

 > rule           -> core-001
 > where          -> request
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password

```

### Running with parameters

```console
$ acurl http://my-company.com/api/entry-point | docker run --rm -it bbvalabs/sensitive-data -r http://127.0.0.1:9999/rules/credentials.yaml

http://my-company.com
---------------------

 > rule           -> core-001
 > where          -> request
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password

```

### Defining rules and ignore files via env vars

You can set these environment vars to set rule file or ignore file:

- SENSITIVE_RULES
- SENSITIVE_IGNORES

```console
$ acurl http://my-company.com/api/entry-point | docker run --rm -it -e SENSITIVE_RULES=/home/john/rules.yaml -e SENSITIVE_IGNORES=http://myserver.com/ignores bbvalabs/sensitive-data

http://my-company.com
---------------------

 > rule           -> core-001
 > where          -> request
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password

```

## Mixing with other APICheck tools

When `APIcheck Sensitive Data` detects an output pipe, it writes an APICheck-compatible output data to allow the connection with other APICheck tools.

```console
$ acurl http://my-company.com/api/entry-point | sensitive-data | send-to-proxy http://my-proxy-addr:9000

http://my-company.com
---------------------

 > rule           -> core-001
 > where          -> request
 > url            -> http://my-company.com/api/entry-point
 > description    -> Find 'password' keyword in flow data
 > sensitiveData  -> password
```

In this case, the `--quiet` (`-q`) flag proves useful:

```bash
acurl http://my-company.com/api/entry-point | sensitive-data -q | send-to-proxy http://my-proxy-addr:9000
```
