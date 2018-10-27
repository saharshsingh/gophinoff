#!/bin/bash

# Define script constants
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
COVERAGE_OUT="$SCRIPT_DIR/coverage.out"

# Process CLI args
LAUNCH_HTML_REPORT=0
if [ "$1" == "--show-html" ]; then
    LAUNCH_HTML_REPORT=1
fi

# Change to script directory and turn off exit on error
pushd $SCRIPT_DIR
set +e

# Run tests and capture pass/fail
go test -coverprofile "$COVERAGE_OUT" ./...
TESTS_PASSED=$?

# If configured, launch coverage report HTML
if [[ $TESTS_PASSED -eq 0 && $LAUNCH_HTML_REPORT -eq 1 ]]; then
    go tool cover -html "$COVERAGE_OUT"
fi

# Turn exit on error back on and restore original directory
set -e
popd
