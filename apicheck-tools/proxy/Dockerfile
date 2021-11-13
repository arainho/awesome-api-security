FROM mitmproxy/mitmproxy

COPY ./certificates /data/certificates
COPY ./addons /addons
COPY ./apicheck_proxy.sh /usr/local/bin/apicheck_proxy

RUN chmod +x /usr/local/bin/apicheck_proxy

ENTRYPOINT ["/usr/local/bin/apicheck_proxy"]
