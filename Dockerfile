FROM gliderlabs/alpine:3.2
MAINTAINER Greg Poirier <greg@opsee.co>

RUN apk update && \
    apk add bash && \
    rm -rf /var/cache/apk/* && \
    mkdir -p /opt/bin

COPY target/linux/amd64/* /opt/bin/

CMD ["/bin/bash"]
