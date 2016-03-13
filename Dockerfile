FROM gliderlabs/alpine:edge
MAINTAINER Greg Poirier <greg@opsee.co>

RUN apk add --update bash curl && \
    rm -rf /var/cache/apt/* && \
    mkdir -p /opt/bin

COPY target/linux/amd64/* /opt/bin/

CMD ["/bin/bash"]
