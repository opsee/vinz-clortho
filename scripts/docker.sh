#!/bin/bash

docker build -t gozer-build -f Dockerfile.build .
docker run --rm --volume `pwd`/export:/export gozer-build
docker build -t quay.io/opsee/vinz -f Dockerfile.vinz .

