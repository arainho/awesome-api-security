import json
import base64

from mitmproxy import http


def to_apicheck_format(flow: http.HTTPFlow, is_error: bool = False) -> str:

    _request = flow.request
    _response = flow.response

    data = {}

    scheme = f"{_request.scheme}:" if _request.scheme else ""
    url = f"{scheme}//{_request.host}:{_request.port}{_request.path or ''}"

    if getattr(_request, "content", None):
        request_content: bytes = _request.content
    else:
        request_content: bytes = _request.text.encode("UTF-8")

    data["request"] = {
        "url": url,
        "method": _request.method,
        "version": _request.http_version,
        "headers": dict(_request.headers),
        "body": base64.b64encode(request_content).decode("UTF-8"),
    }

    if _response:

        if getattr(_response, "content", None):
            response_content: bytes = _response.content
        else:
            response_content: bytes = _response.text.encode("UTF-8")

        data["response"] = {
            "status": _response.status_code,
            "reason": _response.reason,
            "headers": dict(_response.headers),
            "body": base64.b64encode(response_content).decode("UTF-8"),
        }

    if is_error:
        try:
            meta: dict = data["_meta"]
        except KeyError:
            meta = {}
            data["_meta"] = meta

        meta["apicheck_proxy"] = {"error": "can't connect to server"}

    return json.dumps(data)


class APICheck:

    def response(self, flow: http.HTTPFlow):
        print(to_apicheck_format(flow, is_error=False), flush=True)

    def error(self, flow: http.HTTPFlow):
        print(to_apicheck_format(flow, is_error=True), flush=True)

addons = [
    APICheck()
]
