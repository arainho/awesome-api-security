APICheck from cURL
==================

This tool provide a binary `curl` command (`acurl`) that will be translated into
a valid `reqres` object.

*Tool type:* generator

## Tool description

You can use it this way:

The binary `curl` command accepts the same parameters than the `curl` command. For
further documentation you can check `curl` man:

```bash
man curl
```

This tool adds a `curl_log` field to the `_meta` data. You can find all the
`curl` internal log info in this field.

## Quick start

As simple as using `curl`:

```bash
apc install acurl
acurl www.google.com
```

You can use it to retrieve also ssl protected URLs:

```bash
acurl https://www.google.com
```

