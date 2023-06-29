FROM alpine:3.18.0
RUN apk update && apk add --upgrade libcrypto3 libssl3

COPY node-collector /usr/local/bin/node-collector

ENTRYPOINT ["node-collector"]