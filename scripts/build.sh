#!/bin/bash

if [ -d /build ]; then
  cd /build

  echo "Building..."
  CGO_ENABLED=0 /gopath/bin/gb build

  if [ -d /export ]; then
    echo "Preparing for export..."
    cp bin/* /export
  fi
else
  echo "Build directory not found..."
  exit 1
fi
