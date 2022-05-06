#!/bin/env bash

cd `dirname $0`

dir=.
out_dir=.
protos=./*.proto

protoc -I $dir \
  -I .. \
  --go_out $out_dir --go_opt paths=source_relative \
  $protos