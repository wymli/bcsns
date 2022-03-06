# 注意standalone内含ratel 这个ui的镜像
# 实际上https://play.dgraph.io/?latest# 可以直接访问localhost
docker run --rm -d --name dgraph -p 8000:8000 -p 8080:8080 -p 9080:9080 dgraph/standalone:latest