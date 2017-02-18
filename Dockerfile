FROM alpine:3.3

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY paydaybeer /usr/src/app

ENTRYPOINT ["/usr/src/app/paydaybeer"]
