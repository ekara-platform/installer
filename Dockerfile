FROM ekaraplatform/base:v1.0.0

RUN mkdir -p /opt/ekara/bin
COPY ./installer /opt/ekara/bin/installer

RUN mkdir -p /opt/ekara/ansible
WORKDIR /opt/ekara/ansible

ENTRYPOINT ["/opt/ekara/bin/installer"]
