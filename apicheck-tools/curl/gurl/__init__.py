from functools import reduce
import base64
import io
import operator

import httptools

from gurl.RequestResponseCallbacks import RequestResponseCallbacks
import gurl.curlparse as cp
import gurl.hexdealer as hd


def _parse_raw_http(parser_builder, raw_http, parser_extract=None):
    callbacks = RequestResponseCallbacks()
    parser = parser_builder(callbacks)
    parser.feed_data(raw_http)
    callbacks.data["version"] = parser.get_http_version()
    if "body" in callbacks.data:
        # body must be base64
        callbacks.data["body"] = base64.encodebytes(callbacks.data["body"]).decode("utf-8")
    if parser_extract:
        parser_data, parser_meta = parser_extract(parser)
        data = reduce(_dict_reducer, [callbacks.data, parser_data], {})
        meta = reduce(_dict_reducer, [callbacks.meta, parser_meta], {})
        return data, meta
    return callbacks.data, callbacks.meta


def _request(req):
    def _build_parser(callbacks):
        return httptools.HttpRequestParser(callbacks)
    def _extract_from_parser(parser):
        return {"method": parser.get_method().decode("utf-8")}, {}
    request_data, request_meta = _parse_raw_http(
            _build_parser, req, _extract_from_parser)

    return request_data, request_meta


def _response(res):
    def _build_parser(callbacks):
        return httptools.HttpResponseParser(callbacks)
    def _extract_from_parser(parser):
        return {"status": parser.get_status_code()}, {}
    request_data, request_meta = _parse_raw_http(
        _build_parser, res, _extract_from_parser
    )

    return request_data, request_meta


def _dict_reducer(x, acc):
    acc.update(x)
    return acc


def parse_binary(raw_request, raw_response):
    if not raw_request or not raw_response:
        return None

    request_data, request_meta = _request(raw_request)
    response_data, response_meta = _response(raw_response)
    meta = reduce(_dict_reducer, [request_meta, response_meta], {})

    # rebuild original url to be compilance with reqres format
    request_data["url"] = request_data["headers"]["Host"] + request_data["url"]

    return {
        "_meta": meta,
        "request": request_data,
        "response": response_data
    }


def _bytes_reduce(a:bytearray, b:bytes):
    a.extend(b)
    return a


_just_data = operator.attrgetter('data')
_just_req = operator.attrgetter('req')
_just_res = operator.attrgetter('res')


def _extract_bin_block_from_multipart(from_attr, what_attr):
    def _ext(target):
        raw = map(what_attr, from_attr(target))
        return reduce(_bytes_reduce, raw, bytearray())
    return _ext


_extract_req = _extract_bin_block_from_multipart(_just_req, _just_data)
_extract_res = _extract_bin_block_from_multipart(_just_res, _just_data)


def parse_curl_trace(curl_trace_content):
    if not curl_trace_content:
        return None
    
    log = []
    req = bytearray()
    res = bytearray()
    is_https = False

    blocks = cp.curl_trace_block_iterator(curl_trace_content)
    
    for meta_req_res in cp.curl_trace_reqres_iterator(blocks):
        req_bytes = _extract_req(meta_req_res)
        res_bytes = _extract_res(meta_req_res)

        reqres = parse_binary(req_bytes, res_bytes)

        if not "_meta" in reqres:
            reqres["_meta"] = {}
        reqres["_meta"]["curl_log"] = meta_req_res.meta
        if is_https:
            reqres["request"]["url"] = f"https://{reqres['request']['url']}"
        else:
            reqres["request"]["url"] = f"http://{reqres['request']['url']}"
        
        yield reqres
