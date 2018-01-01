#!/usr/bin/env bash

#
#
# This script pulls down dep which is a Go package manager and installs
# the necessary dependencies in the project
#
#

set -o nounset
set -o errexit -o errtrace

readonly base_path="https://github.com/golang/dep/releases/download"
readonly utility_version="v0.3.2"
readonly utility_binary="dep-linux-amd64"
readonly dep="${PWD}/.travis.d/dep"

# Retrieve the latest binary
wget $base_path/$utility_version/$utility_binary \
      -O $dep

# Update access permissions of the utility
chmod +x $dep

sudo $dep ensure
