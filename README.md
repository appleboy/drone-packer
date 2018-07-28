# drone-packer

[![GoDoc](https://godoc.org/github.com/appleboy/drone-packer?status.svg)](https://godoc.org/github.com/appleboy/drone-packer)
[![Build Status](http://drone.wu-boy.com/api/badges/appleboy/drone-packer/status.svg)](http://drone.wu-boy.com/appleboy/drone-packer)
[![codecov](https://codecov.io/gh/appleboy/drone-packer/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-packer)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-packer)](https://goreportcard.com/report/github.com/appleboy/drone-packer)
[![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/drone-packer.svg)](https://hub.docker.com/r/appleboy/drone-packer/)
[![](https://images.microbadger.com/badges/image/appleboy/drone-packer.svg)](https://microbadger.com/images/appleboy/drone-packer "Get your own image badge on microbadger.com")
[![Release](https://github-release-version.herokuapp.com/github/appleboy/drone-packer/release.svg?style=flat)](https://github.com/appleboy/drone-packer/releases/latest)
[![Build status](https://ci.appveyor.com/api/projects/status/pmkfbnwtlf1fm45l/branch/master?svg=true)](https://ci.appveyor.com/project/appleboy/drone-packer/branch/master)

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
