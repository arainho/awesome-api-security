# jwt-checker

Checks that JWT tokens (only JWT, does not support JWS or JWE yet) have been
issued following the best practices.

**Tool type**: action

## Tool description

This tool accepts a JWT token, either provided as an argument or from a
Request/Response object read from standard input, and runs some checks in order
to verify if the best practices have been followed for its issue. The checks are
run over the header and standard claims fields in addition to signature
verification and can be selected by using of the jwt-checker's options. The
following options are supported:

- -h                 show help message and exit
- -V                 show version info and exit
- -unsig             don't raise an error if token is unsigned (algorithm "none")
- -allowAlg value    raise an error if token has different signing algorithm. Provide several options if you want to allow more than one algorithm
- -audience string   raise an error if token has different audience
- -expiresAt string  raise an error if the expiration date is after this. Format: YYYY-MM-DDThh:mm:ss
- -issuer string     raise an error if token has different issuer
- -notBefore string  raise an error if the not before date is priot to this. Format: YYYY-MM-DDThh:mm:ss
- -secret string     use this secret to validate the sign. Base64 encoded (mandatory when validating signed tokens)
- -subject string    raise an error if token has different subject

In case of error, a code of 1 is returned, 0 otherwise.

## Quick start

You can use `jwt-checker` from the `APICheck Package Manager`:

```bash
acp install jwtchk
```

## Usage example

```console
$ acurl http://my-company.com/api/entry-point | jwtchk -allowAlg HS256 -allowAlg HS384 -issuer bbva-iam -subject subject-id -secret bXlTZWNyZXRQYXNzd29yZG15U2VjcmV0UGFzc3dvcmQK

Issuer claim doesn't match. Expected: bbva-iam, got: other-iam
Subject claim doesn't match. Expected: subject-id, got: other-id
```
