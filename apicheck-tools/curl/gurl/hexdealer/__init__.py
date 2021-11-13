

def extract_hex_from_curl_line(curlline):
    if not curlline:
        return None
    dots = curlline.index(b":")
    hex_part = curlline[dots+1:53]
    return hex_part.strip()


def extract_hex_from_curl(curlbody):
    if not curlbody:
        return None
    res = []
    for line in curlbody.strip().split(b'\n'):
        hex_line = extract_hex_from_curl_line(line)
        res.append(hex_line)
    return b" ".join(res)

