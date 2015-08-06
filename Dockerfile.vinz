FROM gliderlabs/alpine:3.2
MAINTAINER Greg Poirier <greg@opsee.co>

ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin:/opt/bin

RUN apk update && \
    apk add bash && \
    rm -rf /var/cache/apk/* && \
    mkdir -p /opt/bin

COPY export/* /opt/bin/

CMD ["/bin/bash"]
