#!/bin/bash

# Define script constants
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Test code
pushd $SCRIPT_DIR
set +e
go test ./...

# Exit if tests failed
TESTS_PASSED=$?
popd
if [ $TESTS_PASSED -ne 0 ]; then exit $TESTS_PASSED; fi

# Build
set -e
time CGO_ENABLED=0 go build -a $SCRIPT_DIR
docker build -t saharsh/taskmaster $SCRIPT_DIR
