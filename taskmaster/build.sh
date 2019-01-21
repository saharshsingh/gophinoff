#!/bin/bash

cleanup() {

    # Remove working container from buildah
    if [ "$tm_container" != "" ]; then buildah rm $tm_container; fi

    # Go back to user directory
    popd
}

# Define script constants
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Process CLI args
USE_BUILDAH=0
if [ "$1" == "--using-buildah" ]; then
    USE_BUILDAH=1
fi

# Change directory to code location
set -e
pushd $SCRIPT_DIR
trap cleanup EXIT INT TERM

# Test code
go test ./...

# Build Go binary
time GOOS=linux CGO_ENABLED=0 go build

# Build container
if [ $USE_BUILDAH -eq 1 ]; then
    tm_container=$(buildah from scratch)
    buildah copy $tm_container taskmaster /
    buildah config --port 8080 --entrypoint '["/taskmaster"]' --label "maintainer=Saharsh Singh" $tm_container
    buildah commit $tm_container saharshsingh/taskmaster
else
    docker build -t saharshsingh/taskmaster $SCRIPT_DIR
fi
