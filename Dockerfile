FROM alpine:3.7

RUN apk add --update --no-cache \
    ca-certificates \
    && rm -rf /var/cache/apk/*

COPY toddler /usr/local/bin/toddler
COPY docker-entrypoint /usr/local/bin/

WORKDIR /usr/local/var/toddler

EXPOSE 8080 8081
VOLUME ["/opt", "/usr/local/etc/toddler"]

ENTRYPOINT ["/usr/local/bin/docker-entrypoint"]
