# Replay

This tool re-sends the requests read from stdin again capturing the new
responses.

It outputs the same pair but substituting the old responses with the new ones.
And storing the old request on _meta for further processess.

# Quick start

Install replay tool:

```bash
apc install replay
```

You need an valid request-response object as input in json line format. You can
grab it from `curl`:

```bash
acurl www.google.com | replay
```

This will generate an output with two responses, the original response made by
`curl` in `_meta/original` field; and the new response made by replay in
`response` field.
