help:
	@echo "run 'make genapi, make genrpc, make genall, make rmall' to generate code from *.api and *.proto"
gen: genall
genall: genapi genrpc
genapi:
	goctl api go -api ./api/idl/*.api -dir ./api  -style=go_zero
genrpc:
	cd ./rpc/idl; goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=go_zero
gentcp:
	cd ./tcp/idl;rm -rf ../pb; mkdir -p ../pb;  protoc -I . --go_out ../pb --go_opt paths=source_relative  ./*.proto
rmapi:
	mkdir -p /tmp/api-idl
	mv ./api/idl/*.api /tmp/api-idl/
	rm -rf ./api/*
	mv /tmp/api-idl ./api/idl
rmrpc:
	mkdir -p /tmp/rpc-idl
	mv ./rpc/idl/*.proto /tmp/rpc-idl/
	rm -rf ./rpc/*
	mv /tmp/rpc-idl ./rpc/idl
rm: rmall
rmall: rmapi rmrpc
build:
	cd ./api ; go build -o ../bin/api
	cd ./rpc ; go build -o ../bin/rpc
run:
	# ./bin/rpc
	# ./bin/api
pwd:
	@pwd
