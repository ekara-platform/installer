FROM lagoonplatform/ansible-docker-alpine:alpha3

RUN mkdir -p /opt/lagoon/bin
COPY ./go/installer /opt/lagoon/bin/installer

RUN mkdir -p /opt/lagoon/ansible
WORKDIR /opt/lagoon/ansible

ENTRYPOINT ["/opt/lagoon/bin/installer"]



