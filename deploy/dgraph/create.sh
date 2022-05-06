# https://play.dgraph.io/?latest#
# note: field name 必须全局唯一

curl -X POST localhost:8080/alter -d \
'
type User {
    nickname
    avatar
    sex
    age
    phone
    password
    public_key
    address
    follows
}
nickname: string .
avatar: string .
sex: int .
age: int .
phone: int @index(int) .
password: password .
public_key: string .
address: string .
follows: [uid] @count @reverse .
'

curl -X POST localhost:8080/alter -d \
'
type Group {
    group_name
    group_avatar
    members
}
group_name: string .
group_avatar: string .
members: [uid] @count @reverse .
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



# 查询双向关系
# var qfriend = `query all($uid: string) {
# 	G as var(func: uid($uid)) {
# 		F as follows
# 		F2 as ~follows 
# 	}

# 	all(func: uid(G)) {
# 		uid
# 		requestsLeft : follows @filter(NOT uid(F2)) {
# 			uid
# 		}
# 		accepted : ~follows @filter(uid(F)) {
# 			uid
# 		}
# 		allFriends : follows {
# 			uid
# 		}
# 	}
# }`