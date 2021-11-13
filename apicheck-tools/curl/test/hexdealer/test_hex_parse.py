import gurl.hexdealer as hd


def test_hex_line():
    try:
        inp = b'''
0000: 48 54 54 50 2f 31 2e 31 20 32 30 30 20 4f 4b 0d HTTP/1.1 200 OK.
0010: 0a                                              .
        '''
        lines = inp.strip().split(b'\n')

        l0 = hd.extract_hex_from_curl_line(lines[0])
        assert l0 == b"48 54 54 50 2f 31 2e 31 20 32 30 30 20 4f 4b 0d"
        bytes.fromhex(l0.decode("utf-8"))

        l1 = hd.extract_hex_from_curl_line(lines[1])
        assert l1 == b"0a"
        bytes.fromhex(l1.decode("utf-8"))
    except Exception as e:
        assert False, f"{e}"
    

def test_hex_block():
    try:
        inp = b'''
0000: 48 54 54 50 2f 31 2e 31 20 32 30 30 20 4f 4b 0d HTTP/1.1 200 OK.
0010: 0a                                              .
        '''

        expected = b"48 54 54 50 2f 31 2e 31 20 32 30 30 20 4f 4b 0d 0a"
        res = hd.extract_hex_from_curl(inp)

        assert res == expected
        bytes.fromhex(res.decode("utf-8"))
    except Exception as e:
        assert False, f"{e}"


def test_block_restrictec_content():
    try:
        inp = b'0010: 48 6f 73 74 3a 20 77 77 77 2e 67 6f 6f 67 6c 65 Host: www.google'

        expected = b"48 6f 73 74 3a 20 77 77 77 2e 67 6f 6f 67 6c 65"
        res = hd.extract_hex_from_curl(inp)

        assert res == expected
        bytes.fromhex(res.decode("utf-8"))
    except Exception as e:
        assert False, f"{e}"
