
# docker network create -d bridge test_net

docker run -d --rm --name elasticsearch \
  -p 9200:9200 \
  -p 9300:9300 \
  -e "discovery.type=single-node" \
  --network=test_net \
  docker.elastic.co/elasticsearch/elasticsearch:7.13.4


docker run -d --rm --name kibana \
      -e elasticsearch.hosts=http://elasticsearch:9200 \
      -e TZ=Asia/Shanghai  \
      --network=test_net  \
      -p 5601:5601 \
      docker.elastic.co/kibana/kibana:7.13.4