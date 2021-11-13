

class RequestResponseCallbacks:
    def __init__(self):
        self.data = {}
        self.meta = {}

    def on_message_begin(self):
        pass

    def on_url(self, url: bytes):
        self.data["url"] = url.decode("utf-8")

    def on_header(self, name: bytes, value: bytes):
        key = name.decode("utf-8")
        val = value.decode("utf-8")
        if not "headers" in self.data:
            self.data["headers"] = {}
        if key in self.data["headers"] and self.data["headers"][key].__class__.__name__ == "list":
            l = self.data["headers"][key]
            l.append(val)
            self.data["headers"][key] = l
        elif key in self.data["headers"]:
            l = [self.data["headers"][key], val]
            self.data["headers"][key] = l
        else:
            self.data["headers"][key] = val

    def on_header_field(self):
        pass

    def on_headers_complete(self):
        pass

    def on_body(self, body: bytes):
        if "body" not in self.data:
            self.data["body"] = bytearray()
        self.data["body"].extend(body)

    def on_message_complete(self):
        pass

    def on_chunk_header(self):
        pass

    def on_chunk_complete(self):
        pass

    def on_status(self, status: bytes):
        s_status = status.decode("utf-8")
        self.meta["status_text"] = s_status