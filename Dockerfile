FROM ekaraplatform/ansible-docker-alpine:latest

RUN mkdir -p /opt/ekara/bin
COPY ./go/installer /opt/ekara/bin/installer

RUN mkdir -p /opt/ekara/ansible
WORKDIR /opt/ekara/ansible

ENTRYPOINT ["/opt/ekara/bin/installer"]



