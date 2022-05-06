
docker pull cassandra:3.11
docker run --name cassandra --network host -d cassandra:3.11


# cqlsh
# docker exec -it cassandra cqlsh