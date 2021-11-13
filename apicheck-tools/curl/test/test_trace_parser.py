import os
import json

import gurl


HERE = os.path.dirname(__file__)


def test_empty():
    res = gurl.parse_curl_trace(None)

    assert not list(res)


def test_google():
    with open(os.path.join(HERE, "tracefiles", "google"), "rb") as f:
        res = next(gurl.parse_curl_trace(f.read()), None)
    assert res is not None
    assert "_meta" in res
    assert "request" in res
    assert "response" in res

    assert res["_meta"]["status_text"] == "OK"

    assert res["request"]["method"] == "GET"
    assert res["request"]["url"] == "http://www.google.com/"
    assert res["request"]["version"] == "1.1"
    assert "headers" in res["request"]
    assert res["request"]["headers"]["Host"] == "www.google.com"
    assert res["request"]["headers"]["User-Agent"] == "curl/7.69.1"
    assert res["request"]["headers"]["Accept"] == "*/*"

    assert res["response"]["status"] == 200
    assert "headers" in res["response"]
    assert res["response"]["headers"]["Date"] == "Tue, 28 Apr 2020 21:07:56 GMT"
    assert res["response"]["headers"]["Expires"] == "-1"
    assert res["response"]["headers"]["Cache-Control"] == "private, max-age=0"
    assert res["response"]["headers"]["Content-Type"] == "text/html; charset=ISO-8859-1"
    assert res["response"]["headers"]["P3P"] == 'CP="This is not a P3P policy! See g.co/p3phelp for more info."'
    assert res["response"]["headers"]["Server"] == "gws"
    assert res["response"]["headers"]["X-XSS-Protection"] == "0"
    assert res["response"]["headers"]["X-Frame-Options"] == "SAMEORIGIN"
    assert res["response"]["headers"]["Set-Cookie"][0] == "1P_JAR=2020-04-28-21; expires=Thu, 28-May-2020 21:07:56 GMT; path=/; domain=.google.com; Secure"
    assert res["response"]["headers"]["Set-Cookie"][1] == "NID=203=qNRKJGTSC5khJwy8FPbtcFuAEE4J30KuYJr6cnusf-p6Sy9Px7b0Nx6DYPoJrwsf2KIxClAIcprR4oLksTvyBt3DLuTwIyBsW94XBghjaORq2GBSOTiTyT4yMFkLXhDCaOxn2cjj4YjR6RkyGsdAeWPSOID9vhEVIWUSfjcCCfU; expires=Wed, 28-Oct-2020 21:07:56 GMT; path=/; domain=.google.com; HttpOnly"
    assert res["response"]["headers"]["Accept-Ranges"] == "none"
    assert res["response"]["headers"]["Vary"] == "Accept-Encoding"
    assert res["response"]["headers"]["Transfer-Encoding"] == "chunked"
    try:
        json.dumps(res)
    except Exception as e:
        assert False, e


def test_https_google():
    try:
        with open(os.path.join(HERE, "tracefiles", "httpsgoogle"), "rb") as f:
            res = gurl.parse_curl_trace(f.read())
    except Exception as e:
        pass # TODO: add http2 support


def test_yahoo():
    with open(os.path.join(HERE, "tracefiles", "yahoo"), "rb") as f:
        res = list(gurl.parse_curl_trace(f.read()))

    assert len(res) == 3
    try:
        json.dumps(res)
    except Exception as e:
        assert False, e
