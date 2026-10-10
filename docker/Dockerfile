ARG BUILDPLATFORM
FROM --platform=$BUILDPLATFORM golang:1.26.9-alpine AS packer-build

ARG TARGETOS
ARG TARGETARCH

# Rebuild the pinned release until upstream ships the Go and x/net security fixes.
WORKDIR /src/packer
RUN apk add --no-cache git && \
  git clone --depth 1 --branch v1.16.1 https://github.com/hashicorp/packer.git . && \
  test "$(git rev-parse HEAD)" = "8d3617fc870d5f08a598806a43d49ed6ee4d08ab" && \
  go get golang.org/x/net@v0.60.0
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -trimpath \
  -ldflags="-s -w -X github.com/hashicorp/packer/version.GitCommit=8d3617fc870d5f08a598806a43d49ed6ee4d08ab" \
  -o /out/packer .

FROM alpine:3.23

ARG TARGETOS
ARG TARGETARCH

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Drone Packer" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

ENV PACKER_VERSION=1.16.1
ENV PACKER_ARCH=${TARGETARCH}
COPY --from=packer-build /out/packer /bin/packer
COPY --from=packer-build /src/packer/LICENSE /usr/share/licenses/packer/LICENSE
RUN apk add --no-cache ca-certificates cdrkit && \
  rm -rf /var/cache/apk/*

COPY release/${TARGETOS}/${TARGETARCH}/drone-packer /bin/

ENTRYPOINT ["/bin/drone-packer"]
