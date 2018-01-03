#!/bin/bash

set -e

export GO15VENDOREXPERIMENT=1
echo -n "Obtaining list of tests to run.."
PKGS=`go list github.com/arxanchain/sdk-go-common/... | grep -v /vendor/ | grep -v /examples/ `
echo "DONE!"

echo "Running tests..."
go test -cover -p 1 -timeout=20m $PKGS
