cd `dirname $0`
goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=go_zero