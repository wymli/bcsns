cd `dirname $0`

# create dockerfile
# goctl docker --go *.go

# build
# cd ../..
# docker build -t user-center-rpc -f app/user_center/Dockerfile .

# run
docker run --rm --name user-center-rpc -p 9991:9991 --network=test_net user-center-rpc

# docker run --rm --name user-center-rpc -p 9991:9991 --network=test_net -v $(pwd)/app/user_center/etc:/app/etc  user-center-rpc