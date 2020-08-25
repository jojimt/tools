FROM alpine:3.10.2
RUN apk upgrade --no-cache && \
  apk --no-cache add tcpdump
COPY ./reqtrace /usr/local/bin/

ENTRYPOINT exec /usr/local/bin/reqtrace
