

cd `dirname $0`

docker run --rm -d \
  --name fluentd \
  -p 9880:9880 \
  --network=test_net \
  -v $(pwd)/etc:/fluentd/etc \
  -v $(pwd)/etc:/var/lib/docker/containers \
  quipper/td-agent:latest -c /fluentd/etc/fluentd.conf