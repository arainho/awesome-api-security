import os
import re
import sys
import json
import base64
import binascii

from typing import List, Tuple, Any

import yaml
import requests
import argparse

from python_pipes import read_stdin_lines
from sanic import Sanic, response, request

HERE = os.path.dirname(__file__)


class FormatErrorException(Exception):
    pass


def flatten_dict(target: Any,
                 path=None) -> Tuple[Tuple[str], str, str or int or bool]:
    """
    This function flatten a dictionary.

    Return format: [(PATH, DICT KEY, DICT VALUE)]

    Usage example:

    >> inp = {
        "name": "Jonh",
        "surname": "Doe",
        "age": 33,
        "meta": {
            "since": "2001-01-01"
        }
    }
    >> expected = [
        (None,"name","Jonh"),
        (None,"surname","Doe"),
        (None,"age",33),
        (("meta",), "since", "2001-01-01")
    ]
    >> _recurse(inp)
    >> res = list(_recurse(inp))
    >> assert res == expected
    """

    def _path_add(item):
        if path:
            return *path, item
        return item,

    if not target:
        yield None, None, None

    elif target.__class__.__name__ == "dict":
        for k, v in target.items():
            for res in flatten_dict(v, _path_add(k)):
                yield res
    elif target.__class__.__name__ == "list":
        for i, v in enumerate(target):
            for res in flatten_dict(v, _path_add(i)):
                yield res
    elif len(path) == 0:
        # error, this can't happen
        print("nah")
    elif len(path) == 1:
        yield None, path[0], target
    else:
        yield path[:-1], path[-1], target


def _load_rules(args: argparse.Namespace) -> List[dict]:
    """Load rules files from local o remote"""
    default_rules_path = os.path.join(HERE, "rules.yaml")

    with open(default_rules_path, "r") as f:
        rules = yaml.safe_load(f.read())

    rules_files = []

    if env_rules := os.environ.get("SENSITIVE_RULES", None):
        rules.append(env_rules)

    if args.rules_file:
        rules_files = rules_files + args.rules_file

    for rule_file in rules_files:
        if rule_file.startswith("http"):
            # Load from remote URL
            rules.extend(
                yaml.safe_load(requests.get(rule_file).content)
            )
        else:
            # Load from local file
            real_file_path = os.path.join(os.getcwd(), rule_file)
            with open(real_file_path, "r") as f:
                rules.extend(yaml.safe_load(f.read()))

    return rules


def _load_ignore_ids(args: argparse.Namespace) -> List[str]:
    """Load ignores files from local o remote"""

    ignores = []

    if args.ignore_rule:
        for x in args.ignore_rule:
            ignores.extend(x.split(","))

    ignore_file = args.ignore_file or []
    if env_ignore := os.environ.get("SENSITIVE_IGNORES", None):
        ignore_file.append(env_ignore)

    if ignore_file:
        for rule_file in ignore_file:
            if rule_file.startswith("http"):
                # Load from remote URL
                ignores.extend(requests.get(rule_file).content.splitlines())
            else:
                # Load from local file
                real_file_path = os.path.join(os.getcwd(), rule_file)
                with open(real_file_path, "r") as f:
                    ignores.extend(
                        [x.replace("\n", "") for x in f.readlines()])

    return ignores


def _check_input_data(data: dict) -> bool or FormatErrorException:
    if "request" not in data:
        raise FormatErrorException("Missing key 'request'")
    if "response" not in data:
        raise FormatErrorException("Missing key 'request'")

    return True


def decode_body(data: str) -> Tuple[str, dict]:
    """
    This function try to decode body and return the type of body

    :return: tuple("text|dict", content)
    """
    try:
        _data = data.encode("UTF-8")
    except AttributeError:
        _data = data

    # Try to decode base64 body
    try:
        _body = base64.decodebytes(_data).decode("UTF-8")
    except binascii.Error:
        _body = _data

    # Try con convert to json
    try:
        return "dict", json.loads(_body)
    except json.decoder.JSONDecodeError:
        return "text", _body


def search_in_dict(body: dict, rule, where, url) -> list:
    issues = []

    if not body:
        return issues

    for (path, key, value) in flatten_dict(body):

        for v in (key, value):
            if regex := re.search(rule["regex"], v):

                issues.append({
                    "rule": rule["id"],
                    "where": where,
                    "url": url,
                    "description": rule["description"],
                    "sensitiveData": regex.group()
                })

    return issues


def search_issues(content_json: dict, rules: list, ignores: set) -> List[dict]:
    issues = []

    # Matching
    for rule in rules:

        if rule["id"] in ignores:
            continue

        #
        # Search in Request / Response
        #
        content_to_search = {}
        include_keys = rule.get("includeKeys", False)

        where_to_find_location = {"Headers", "Request", "Response"}
        where_to_find = set()

        #
        # Check and set places where search
        #
        if rule["searchIn"] == "All":
            where_to_find.update(where_to_find_location)
        elif rule["searchIn"] in where_to_find_location:
            where_to_find.add(rule["searchIn"])
        else:
            raise FormatErrorException(
                f"Invalid 'searchIn' value. Allowed values are: "
                f"'{','.join(where_to_find_location)}'"
            )

        #
        # Collecting data
        #
        if "Request" in where_to_find:
            if body := content_json["request"].get("body", None):
                content_to_search["request"] = decode_body(body)

        if "Response" in where_to_find:
            if body := content_json["response"].get("body", None):
                content_to_search["response"] = decode_body(body)

        if "Headers" in where_to_find:
            _headers_request = content_json["request"].get("headers", {})
            _headers_response = content_json["response"].get("headers", {})

            content_to_search["requestHeaders"] = "dict", _headers_request
            content_to_search["responseHeaders"] = "dict", _headers_response

        url = content_json["request"]['url']

        for where, (body_type, body_content) in content_to_search.items():

            if body_type == "dict":
                issues.extend(
                    search_in_dict(body_content, rule, where, url)
                )
            else:

                if regex := re.search(rule["regex"], body_content):

                    issues.append({
                        "rule": rule["id"],
                        "where": where,
                        "url": content_json[where],
                        "description": rule["description"],
                        "sensitiveData": regex.group()
                    })

    return issues


def cli_analyze(args: argparse.Namespace):
    quiet = args.quiet

    # -------------------------------------------------------------------------
    # Read info by stdin or parameter
    # -------------------------------------------------------------------------
    for has_stdin_pipe, has_stdout_pipe, json_line in read_stdin_lines():
        if not has_stdin_pipe:
            raise FileNotFoundError(
                "Input data must be entered as a UNIX pipeline. For example: "
                "'cat info.json | tool-name'")

        rules = _load_rules(args)
        ignores = set(_load_ignore_ids(args))

        # this var contains JSON data in APICheck format
        content_json: dict = json.loads(json_line)

        found_issues = search_issues(content_json, rules, ignores)

        # You're being piped or redirected
        if has_stdout_pipe:

            #
            # Dump content as APICheck format
            #
            if not hasattr(content_json, "_meta"):
                content_json["_meta"] = {}

            if type(content_json["_meta"]) is not dict:
                content_json["_meta"] = {}

            content_json["_meta"]["sensitive-json"] = found_issues

            output_apicheck_data = json.dumps(content_json)

            # Info for next pip command
            sys.stdout.write(f"{output_apicheck_data}\n")
            sys.stdout.flush()

        # If not quiet also display in console. If has output pipe -> write
        # console into stderr, otherwise write in stdout
        if has_stdout_pipe:
            console_print = sys.stderr.write
            console_flush = sys.stderr.flush
        else:
            console_print = sys.stdout.write
            console_flush = sys.stdout.flush

        if not quiet:
            console_print(f"\n")

            for issue in found_issues:
                url = content_json['request']['url']
                console_print(f"{url}\n")
                console_print(f"{'-' * len(url)}\n\n")

                for x, y in issue.items():
                    console_print(f" > {x.ljust(15)}-> {y}\n")

                console_print(f"\n")
                console_flush()


def server(args: argparse.Namespace):
    app = Sanic("apicheck-sensiteve-data")

    @app.route("/apicheck/sensitive-data", methods=["POST"])
    def home_analyze(_request: request.Request):
        _config = _request.app.config
        _input_data = _request.json

        try:
            _check_input_data(_input_data)
        except FormatErrorException as e:
            return response.json({"message": str(e)}, status=400)

        issues = search_issues(_input_data,
                               app.config.RULES,
                               app.config.IGNORES)

        if _config.CONSOLE_MODE:
            print(issues, flush=True)

        if _config.DONT_CHECK:
            return response.json({"message": "Ok"})
        else:
            return response.json(issues)

    app.config.RULES = _load_rules(args)
    app.config.IGNORES = set(_load_ignore_ids(args))
    app.config.CONSOLE_MODE = args.show_in_console
    app.config.DONT_CHECK = args.dont_check

    listen_addr, listen_port = args.server.split(":")

    app.run(host=listen_addr, port=int(listen_port), debug=False,
            access_log=False)


def main():
    parser = argparse.ArgumentParser(
        description='Analyze a HTTP Request / Response searching for '
                    'sensitive data'
    )
    parser.add_argument('-q', '--quiet',
                        default=False,
                        action="store_true",
                        help="quiet mode")
    parser.add_argument('-F', '--ignore-file',
                        action="append",
                        help="file with ignores rules")
    parser.add_argument('-i', '--ignore-rule',
                        action="append",
                        help="rule to ignore")
    parser.add_argument('-r', '--rules-file',
                        action="append",
                        help="rules file. One rule ID per line")
    parser.add_argument('--server',
                        default=None,
                        help="launch in server mode listening at "
                             "localhost:8000")

    group = parser.add_argument_group('Server mode options')
    group.add_argument("-C", "--show-in-console",
                       default=False,
                       action="store_true",
                       help="show results in console")
    group.add_argument("-D", "--dont-check",
                       default=False,
                       action="store_true",
                       help="always returns OK although a rule matches")

    parsed_cli = parser.parse_args()

    if not parsed_cli.server:
        try:
            cli_analyze(parsed_cli)
        except Exception as e:
            print("\n", f"[!!] {e}", "\n", file=sys.stderr)
            exit(1)
    else:
        server(parsed_cli)


if __name__ == '__main__':
    main()
