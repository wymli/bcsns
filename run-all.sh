
# todo: use docker-compose

docker network create -d bridge test_net

# run redis
cd `dirname $0`
sh deploy/redis/run.sh


# run es/kibana
cd `dirname $0`
sh deploy/es/run.sh


# run fluentd
cd `dirname $0`
sh deploy/fluentd/run.sh


# run dgraph
cd `dirname $0`
sh deploy/dgraph/run.sh

sleep 5
sh deploy/dgraph/create.sh

docker ps


# docker run --rm --name user-center-rpc --network=host -v $(pwd)/app/user_center/etc:/app/etc   user-center-rpc
# docker run --rm --name auth-rpc --network=host -v $(pwd)/app/auth_rpc/etc:/app/etc   auth-rpc
# docker run --rm --name grpc-gateway --network=host -v $(pwd)/app/grpc_gateway/etc:/app/etc   grpc-gateway

# curl -XPOST -d '{"phone":8989898989,"nickname":"hello_world"}' localhost:8888/api/v1/user