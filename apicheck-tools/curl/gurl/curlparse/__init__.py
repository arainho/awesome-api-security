from dataclasses import dataclass
import re

import gurl.hexdealer as he

"""
A block is a send or receive item
A block represent bytes, so bytearray must be intended
A block can has meta data associated
"""
@dataclass
class Block:
    kind: str
    item: str
    bytes_size: int
    data: bytes
    meta:list=None


def _extract_item(line):
    regex = r"(=>|<=)\s+(Send|Recv)\s+([\w\s]+),\s+(\d+)"
    match = re.search(regex, line)
    kind = match.groups()[1]
    item = match.groups()[2]
    size = int(match.groups()[3])
    return kind, item, size


def curl_trace_block_iterator(curl_content:bytes):
    if not curl_content or curl_content == '':
        yield from ()
        return
    kind = None
    item = None
    size = 0
    content = bytearray()
    meta = []
    last_info = False
    for line in curl_content.strip().split(b'\n'):
        if line.startswith(b'=='):
            last_info = True
            clean = line.replace(b"== ", b"")
            meta.append(clean.decode("utf-8"))
        elif line.startswith(b'=>') or line.startswith(b'<='):
            last_info = False
            if content:
                if meta:
                    yield Block(kind, item, size, bytes(content), meta)
                else:
                    yield Block(kind, item, size, bytes(content))
                content = bytearray()
                meta = []
            kind, item, size = _extract_item(line.decode("utf-8"))
        elif last_info: # append to info
            meta[-1] = meta[-1] + "\n" + line.decode("utf-8")
        else: # append to block
            hex_part = he.extract_hex_from_curl(line)
            b = bytes.fromhex(hex_part.decode("utf-8"))
            content.extend(b)
    if content and meta:
        yield Block(kind, item, size, bytes(content), meta)
    elif content:
        yield Block(kind, item, size, bytes(content))


@dataclass
class MetaReqRes:
    meta: list
    req: list
    res: list
    is_ssl: bool = False


def curl_trace_reqres_iterator(blocks_iterator):
    """
    SSL negotiation can occur any time this method ignores it. But keep
        info about it.
    """
    meta, req, res = ([],[],[])
    receiving = False
    is_ssl = False
    
    for block in blocks_iterator: 
        if "ssl" in block.item.lower():
            is_ssl = True
            if block.meta:
                meta.extend(block.meta)
        elif block.kind.lower() == "recv":
            receiving = True
            res.append(block)
        elif block.kind.lower() == "send":
            if receiving:
                yield MetaReqRes(meta, req, res, is_ssl)
                meta, req, res = ([],[],[])
                receiving = False
            req.append(block)
    
    if meta or req or res:
        yield MetaReqRes(meta, req, res, is_ssl)
