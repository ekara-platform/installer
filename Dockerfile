FROM ekaraplatform/ansible-docker-alpine:1.0.0-beta1

RUN mkdir -p /opt/ekara/bin
COPY ./go/installer /opt/ekara/bin/installer

RUN mkdir -p /opt/ekara/ansible
WORKDIR /opt/ekara/ansible

ENTRYPOINT ["/opt/ekara/bin/installer"]



