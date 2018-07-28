# drone-packer

Drone plugin for build Automated machine images with [Packer](https://www.packer.io/). For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-packer/).

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-packer
docker build --rm -t appleboy/drone-packer .
```
