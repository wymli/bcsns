

docker exec -it cassandra cqlsh

create keyspace bcsns with replication = {'class':'SimpleStrategy', 'replication_factor' : 3};
use bcsns;

# uid is recv_id

create table message (
  uid bigint,
  send_uid bigint,
  room_id bigint,
  server_msg_id bigint,
  content_type text,
  data blob,
  primary key(uid,server_msg_id)
) WITH CLUSTERING ORDER BY(server_msg_id DESC);

create table failed_message (
  uid bigint,
  send_uid bigint,
  room_id bigint,
  server_msg_id bigint,
  content_type text,
  data blob,
  primary key(uid,server_msg_id)
) WITH CLUSTERING ORDER BY(server_msg_id DESC);

create table moment (
  uid bigint,
  send_uid bigint,
  server_msg_id bigint,
  content_type text,
  data blob,
  primary key(uid,server_msg_id)
) WITH CLUSTERING ORDER BY(server_msg_id DESC);



# 建索引,不过一般用partition key和cluster key就可以,不用单独建
# create index messages_server_msg_id_index on messages(server_msg_id);

# 查看所有keyspace
# describe keyspaces;

# 查看特定keyspace
# describe keyspace bcsns; 或 describe bcsns;