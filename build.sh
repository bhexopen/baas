#!/bin/bash -eu

root=`pwd`/`dirname $0`
rm -rf $root/build
mkdir $root/build
docker build -t apidoc $root/docker
docker run -v $root/source:/slate/source -v $root/build:/slate/build apidoc \
  sh -c 'cd /slate && cp -nr source_orig/* source && exec bundle exec middleman build'
