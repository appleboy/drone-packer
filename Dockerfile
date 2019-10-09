FROM plugins/base:amd64

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Drone Packer" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache ca-certificates \
  wget && \
  rm -rf /var/cache/apk/*

ENV PACKER_VERSION 1.4.4
ENV PACKER_ARCH amd64
RUN wget -q https://releases.hashicorp.com/packer/${PACKER_VERSION}/packer_${PACKER_VERSION}_linux_${PACKER_ARCH}.zip -O packer.zip && \
  unzip packer.zip -d /bin && \
  rm -f packer.zip

ADD release/linux/amd64/drone-packer /bin/
ENTRYPOINT ["/bin/drone-packer"]
