#!/bin/bash
#
# Copyright arxanfintech.com. 2017 All Rights Reserved.
#

#set -eux

PROTO_ROOT_DIR="$GOPATH/src/github.com/arxanchain/sdk-go-common/protos"
PROTO_DIRS=`find $GOPATH/src/github.com/arxanchain/sdk-go-common/protos -mindepth 1 -maxdepth 1 -type d`

for dir in $PROTO_DIRS; do
  echo Working on dir $dir
  protoc --proto_path="$PROTO_ROOT_DIR" --go_out=plugins=grpc:$GOPATH/src "$dir"/*.proto
done

# inject_tags
# export PATH=$PATH:$GOPATH/bin
# for i in "wallet/tx.pb.go"; do
#  protoc-go-inject-tag -input="$GOPATH/src/github.com/arxanchain/sdk-go-common/protos/$i"
# done
