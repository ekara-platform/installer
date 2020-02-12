FROM ekaraplatform/base:latest

RUN mkdir -p /opt/ekara/bin
COPY ./installer /opt/ekara/bin/installer

ENTRYPOINT ["/opt/ekara/bin/installer"]
