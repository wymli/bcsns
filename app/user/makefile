help:
	@echo "run 'make genapi, make genrpc, make genall, make rmall' to generate code from *.api and *.proto"
genall: 
	goctl api go -api ./api/idl/*.api -dir ./api  -style=go_zero
	cd ./rpc/idl; goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=go_zero
genapi:
	goctl api go -api ./api/idl/*.api -dir ./api  -style=go_zero
genrpc:
	cd ./rpc/idl; goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=go_zero
rmall:
	# rm  `ls ./api/idl | grep -v .api`
	# rm  `ls ./rpc/idl | grep -v .proto`
	mv ./api/idl /tmp/api-idl
	mv ./rpc/idl /tmp/rpc-idl
	rm -rf ./api/*
	rm -rf ./rpc/*
	mv /tmp/api-idl ./api/idl
	mv /tmp/rpc-idl ./rpc/idl
pwd:
	@pwd
