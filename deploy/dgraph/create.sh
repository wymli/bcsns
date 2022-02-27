

curl -X POST localhost:8080/alter -d \
'
type User {
    nickname
    password
    follows
    publickey
    address
    phone
}
nickname: string .
password: password .
follows: [uid] @count @reverse .
publickey: string .
address: string .
phone: int @index(int) .
'

# curl -H "Content-Type: application/json" localhost:8080/mutate -XPOST -d '
# {
#     "set":[
#       {
#         "xid": "1",
#         "nickname": "apple"
#       },
#       {
#         "xid":"2",
#         "nickname": "banana"
#       }
#     ]
# }'

# curl -H "Content-Type: application/rdf" localhost:8080/mutate -XPOST -d '
# {
#     "set":[
#       {
#         "xid": "1",
#         "follows":{
#           "xid":"2"
#         }
#       }
#     ]
# }'

# curl -H "Content-Type: application/rdf" localhost:8080/mutate -XPOST -d '
# {
#     "set":{
#       <>
#     }
# }'

# curl -X POST localhost:8080/alter -d '{"drop_all": true}'

# func uidFromxid(xid string) uint64 {
# 	f := fnv.New64a()
# 	f.Write([]byte(xid))
# 	return f.Sum64()
# }