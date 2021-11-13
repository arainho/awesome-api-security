from datetime import timedelta
import functools as f
import itertools as it
import json
import os

import subprocess
from hypothesis import settings
from hypothesis import given
from hypothesis import strategies as st

BASE = os.getcwd()

def gurl(url, headers):
    header_params = f.reduce(it.chain, it.product(["-H"], [f"{k}: {v}" for k, v in headers.items()]))
    return subprocess.check_output([os.path.join(BASE, "gcurl"), *header_params, url])

@settings(deadline=timedelta(seconds=3))
@given(headers=st.dictionaries(
    keys=st.from_regex(
        r"[a-zA-Z0-9_\-]+",
        fullmatch=True
        ).filter(lambda x: len(x.strip()) > 0), 
    values=st.from_regex(
        r"[a-zA-Z0-9_\-]+",
        fullmatch=True
        ).filter(lambda x: len(x.strip()) > 0))
    )
def test_decode_random_headers(headers):
    headers.update({'Accept': '*/*', 'Host': 'httpbin.org', 'User-Agent': 'curl/7.70.0'})
    reqres_text = gurl("http://httpbin.org/headers", headers=headers)
    reqres = json.loads(reqres_text)
    assert reqres['request']['headers'] == headers
