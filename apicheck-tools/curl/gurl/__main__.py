import argparse
import json
import sys

import gurl


def main():
    par = argparse.ArgumentParser(
        description="Convert curl traces to reqres object, provide by -f or stdin"
    )
    par.add_argument(
        "-f", "--file", help="Trace file to convert"
    )

    cli = par.parse_args()

    if cli.file:
        with open(cli.file, "rb") as f:
            resps = gurl.parse_curl_trace(f.read())
            for res in resps:
                print(json.dumps(res))
    else:
        std_input = sys.stdin.buffer.read()
        if std_input:
            resps = gurl.parse_curl_trace(std_input)
            for res in resps:
                print(json.dumps(res))
        else:
            sys.stderr.write("No input, please use stdin of -f to provide")


if __name__ == '__main__':
    main()
