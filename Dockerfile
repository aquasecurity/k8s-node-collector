FROM alpine:3.16.2

COPY node-collector /usr/local/bin/node-collector

ENTRYPOINT ["node-collector"]