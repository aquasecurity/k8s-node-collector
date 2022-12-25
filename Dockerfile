FROM alpine:3.16.2

COPY node-info /usr/local/bin/node-info

ENTRYPOINT ["node-info"]