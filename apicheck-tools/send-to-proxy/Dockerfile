FROM python:3.8-alpine as base

RUN apk update && \
    apk upgrade && \
    apk --no-cache --update add openssl ca-certificates

FROM base as builder

RUN apk add --no-cache build-base && \
    python -m pip install --no-cache-dir -U pip wheel

COPY ./ /app/
RUN python -m pip wheel --no-cache-dir --wheel-dir=/root/wheels -r /app/requirements.txt && \
    python -m pip wheel --no-cache-dir --wheel-dir=/root/wheels /app/

FROM base
COPY --from=builder /root/wheels /root/wheels

RUN python -m pip install --no-cache --no-index /root/wheels/*
RUN rm -rf /root/wheels

ENTRYPOINT ["send-to-proxy"]
