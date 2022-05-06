#!/bin/env bash
cd `dirname $0`
# goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=go_zero


# only gen pb
mkdir /tmp/ttt
goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=/tmp/ttt --style=go_zero -I ~/bcsns/dependency/pb -I .
rm -rf /tmp/ttt