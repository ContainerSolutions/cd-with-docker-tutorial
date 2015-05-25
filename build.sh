#!/bin/sh
docker build -t hello-world-build -f Dockerfile.build .
docker run hello-world-build > build.tar.gz
docker build -t hello-world -f Dockerfile.dist .