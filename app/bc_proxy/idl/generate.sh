#!/bin/env bash
cd `dirname $0`

dir=.
out_dir="../pb"
googleapis_dir="../../../dependency/pb/googleapis"


goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=go_zero --proto_path=$googleapis_dir --proto_path=$dir


protoc -I $dir \
  -I $googleapis_dir \
  --go_out $out_dir --go_opt paths=source_relative \
  --go-grpc_out $out_dir --go-grpc_opt paths=source_relative \
  --grpc-gateway_out $out_dir --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt generate_unbound_methods=true \
  --grpc-gateway_opt register_func_suffix=GW \
  --grpc-gateway_opt allow_delete_body=true \
  --openapiv2_out $out_dir --openapiv2_opt logtostderr=true \
  ./*.proto