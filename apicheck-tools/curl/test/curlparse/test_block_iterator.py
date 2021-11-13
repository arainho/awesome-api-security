import gurl.curlparse as cp
import uuid


def test_none():
    inp = None
    res = cp.curl_trace_block_iterator(inp)

    assert len(list(res)) == 0


def test_empty():
    inp = ""
    res = cp.curl_trace_block_iterator(inp)

    assert len(list(res)) == 0


def test_one_of_kind():
    imp = f"""
== Meta info block 1
== Meta info block 2
=> Send SSL data, 5 bytes (0x5)
0000: 16 03 01 02 00                                  .....
<= Recv SSL data, 5 bytes (0x5)
0000: 16 03 03 00 7a                                  ....z
=> Send header, 76 bytes (0x4c)
0000: 47 45 54 20 2f 20 48 54 54 50 2f 32 0d 0a 48 6f GET / HTTP/2..Ho
0010: 73 74 3a 20 77 77 77 2e 67 6f 6f 67 6c 65 2e 63 st: www.google.c
0020: 6f 6d 0d 0a 75 73 65 72 2d 61 67 65 6e 74 3a 20 om..user-agent: 
0030: 63 75 72 6c 2f 37 2e 36 39 2e 31 0d 0a 61 63 63 curl/7.69.1..acc
0040: 65 70 74 3a 20 2a 2f 2a 0d 0a 0d 0a             ept: */*....
=> Send data, 5 bytes (0x5)
0000: 17 03 03 02 21                                  ....!
<= Recv header, 13 bytes (0xd)
0000: 48 54 54 50 2f 32 20 32 30 30 20 0d 0a          HTTP/2 200 ..
<= Recv data, 5 bytes (0x5)
0000: 17 03 03 00 1a                                  .....
    """

    expected = [
        cp.Block(
            "Send",
            "SSL data",
            5,
            bytes.fromhex(b"16 03 01 02 00".decode("utf-8")),
            ["Meta info block 1","Meta info block 2"]
        ),
        cp.Block(
            "Recv",
            "SSL data",
            5,
            bytes.fromhex(b"16 03 03 00 7a".decode("utf-8"))
        ),
        cp.Block(
            "Send",
            "header",
            76,
            bytes.fromhex(b"47 45 54 20 2f 20 48 54 54 50 2f 32 0d 0a 48 6f 73 74 3a 20 77 77 77 2e 67 6f 6f 67 6c 65 2e 63 6f 6d 0d 0a 75 73 65 72 2d 61 67 65 6e 74 3a 20 63 75 72 6c 2f 37 2e 36 39 2e 31 0d 0a 61 63 63 65 70 74 3a 20 2a 2f 2a 0d 0a 0d 0a            ".decode("utf-8"))
        ),
        cp.Block(
            "Send",
            "data",
            5,
            bytes.fromhex(b"17 03 03 02 21".decode("utf-8"))
        ),
        cp.Block(
            "Recv",
            "header",
            13,
            bytes.fromhex(b"48 54 54 50 2f 32 20 32 30 30 20 0d 0a".decode("utf-8"))
        ),
        cp.Block(
            "Recv",
            "data",
            5,
            bytes.fromhex(b"17 03 03 00 1a".decode("utf-8"))
        )
    ]

    res = list(cp.curl_trace_block_iterator(bytes(imp, encoding="utf-8")))

    assert res[0] == expected[0], "0 Block error"
    assert res[1] == expected[1], "1 Block error"
    assert res[2] == expected[2], "2 Block error"
    assert res[3] == expected[3], "3 Block error"
    assert res[4] == expected[4], "4 Block error"
    assert res[5] == expected[5], "5 Block error"


def test_reqres_block_iterator():
    imp = bytes(f"""
== Meta info block 1
== Meta info block 2
=> Send SSL data, 5 bytes (0x5)
0000: 16 03 01 02 00                                  .....
<= Recv SSL data, 5 bytes (0x5)
0000: 16 03 03 00 7a                                  ....z
=> Send header, 76 bytes (0x4c)
0000: 47 45 54 20 2f 20 48 54 54 50 2f 32 0d 0a 48 6f GET / HTTP/2..Ho
0010: 73 74 3a 20 77 77 77 2e 67 6f 6f 67 6c 65 2e 63 st: www.google.c
0020: 6f 6d 0d 0a 75 73 65 72 2d 61 67 65 6e 74 3a 20 om..user-agent: 
0030: 63 75 72 6c 2f 37 2e 36 39 2e 31 0d 0a 61 63 63 curl/7.69.1..acc
0040: 65 70 74 3a 20 2a 2f 2a 0d 0a 0d 0a             ept: */*....
=> Send data, 5 bytes (0x5)
0000: 17 03 03 02 21                                  ....!
<= Recv header, 13 bytes (0xd)
0000: 48 54 54 50 2f 32 20 32 30 30 20 0d 0a          HTTP/2 200 ..
<= Recv data, 5 bytes (0x5)
0000: 17 03 03 00 1a                                  .....
    """, encoding="utf-8")

    exp = [
        cp.MetaReqRes(
            ["Meta info block 1", "Meta info block 2"],
            [
                cp.Block(
                    "Send",
                    "header",
                    76,
                    bytes.fromhex(b"47 45 54 20 2f 20 48 54 54 50 2f 32 0d 0a 48 6f 73 74 3a 20 77 77 77 2e 67 6f 6f 67 6c 65 2e 63 6f 6d 0d 0a 75 73 65 72 2d 61 67 65 6e 74 3a 20 63 75 72 6c 2f 37 2e 36 39 2e 31 0d 0a 61 63 63 65 70 74 3a 20 2a 2f 2a 0d 0a 0d 0a            ".decode("utf-8"))
                ),
                cp.Block(
                    "Send",
                    "data",
                    5,
                    bytes.fromhex(b"17 03 03 02 21".decode("utf-8"))
                )
            ],
            [
                cp.Block(
                    "Recv",
                    "header",
                    13,
                    bytes.fromhex(b"48 54 54 50 2f 32 20 32 30 30 20 0d 0a".decode("utf-8"))
                ),
                cp.Block(
                    "Recv",
                    "data",
                    5,
                    bytes.fromhex(b"17 03 03 00 1a".decode("utf-8"))
                )
            ],
            True
        )
    ]

    blocks = cp.curl_trace_block_iterator(imp)
    res = list(cp.curl_trace_reqres_iterator(blocks))

    assert res == exp


def test_multiline_info():
    inp = bytes(f"""== Info:   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: none
=> Send header, 76 bytes (0x4c)
0000: 47 45 54 20 2f 20 48 54 54 50 2f 32 0d 0a 48 6f GET / HTTP/2..Ho
0010: 73 74 3a 20 77 77 77 2e 67 6f 6f 67 6c 65 2e 63 st: www.google.c
0020: 6f 6d 0d 0a 75 73 65 72 2d 61 67 65 6e 74 3a 20 om..user-agent: 
0030: 63 75 72 6c 2f 37 2e 36 39 2e 31 0d 0a 61 63 63 curl/7.69.1..acc
0040: 65 70 74 3a 20 2a 2f 2a 0d 0a 0d 0a             ept: */*....""", encoding="utf-8")

    exp = [
        cp.Block(
            "Send",
            "header",
            76,
            bytes.fromhex(b"47 45 54 20 2f 20 48 54 54 50 2f 32 0d 0a 48 6f 73 74 3a 20 77 77 77 2e 67 6f 6f 67 6c 65 2e 63 6f 6d 0d 0a 75 73 65 72 2d 61 67 65 6e 74 3a 20 63 75 72 6c 2f 37 2e 36 39 2e 31 0d 0a 61 63 63 65 70 74 3a 20 2a 2f 2a 0d 0a 0d 0a            ".decode("utf-8")),
            ["Info:   CAfile: /etc/ssl/certs/ca-certificates.crt\n  CApath: none"]
        )
    ]

    res = list(cp.curl_trace_block_iterator(inp))

    assert res == exp