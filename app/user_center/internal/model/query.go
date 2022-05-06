package model

import (
	"context"
	"encoding/json"

	"github.com/dgraph-io/dgo/v210"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/utils"
)

const (
	QFriends = `query all($uid: string) {
		G as var(func: uid($uid)) {
			F as follows
		}
	
		all(func: uid(G)) @normalize {
			friends : ~follows @filter(uid(F)) {
				uid: uid
				nickname: nickname
				sex: sex
				age: age
				avatar: avatar
				address: address
				public_key: public_key
			}
		}
	}`

	QUserInfo = `query all($uid: string) {
		all(func: uid($uid)) {
			uid
			nickname
			sex
			age
			phone
			avatar
			address
			public_key
		}
	}`

	QGroupInfo = `query all($uid: string) {
		all(func: uid($uid)) {
			uid
			group_name
			group_avatar
			members {
				uid
				nickname
				sex
				age
				avatar
				address
				public_key
			}
		}
	}`

	QGroupMemberUid = `query all($uid: string) {
		all(func: uid($uid)) {
			members {
				uid
			}
		}
	}`

	QFollows = `query all($uid: string) {
		all(func: uid($uid)) @normalize {
			follows{
				uid: uid
				nickname: nickname
				sex: sex
				age: age
				avatar: avatar
				address: address
				public_key: public_key
			}
		}
	}`

	QFans = `query all($uid: string) {
		all(func: uid($uid)) @normalize {
			~follows{
				uid: uid
				nickname: nickname
				sex: sex
				age: age
				avatar: avatar
				address: address
				public_key: public_key
			}
		}
	}`

	QFansUid = `query all($uid: string) {
		all(func: uid($uid)) @normalize {
			~follows{
				uid: uid
			}
		}
	}`

	QMyGroups = `query all($uid: string) {
		all(func: uid($uid)) @normalize {
			~members {
				uid: uid
				group_name: group_name
				group_avatar: group_avatar
			}
		}
	}`

	QMembers = `query all($uid: string) {
		all(func: uid($uid)) @normalize  {
			members {
				uid: uid
				nickname: nickname
				sex: sex
				age: age
				avatar: avatar
				address: address
				public_key: public_key
			}
		}
	}`

	QLogin = `
query check($phone: int, $password: string){
  check(func: eq(phone, $phone)) {
    pass: checkpwd(password, $password)
  }
}
`
)

type (
	FriendsDecode struct {
		All []User
	}
	UserInfoDecode struct {
		All []User
	}
	GroupInfoDecode struct {
		All []Group
	}
	GroupMemberUidDecode struct {
		All []Group
	}
	FollowsDecode struct {
		All []User
	}
	FansDecode struct {
		All []User
	}
	FansUidDecode struct {
		All []User
	}
	MyGroupsDecode struct {
		All []Group
	}
	MembersDecode struct {
		All []User
	}
	LoginDecode struct {
		Check []struct {
			Uid  string
			Pass bool
		}
	}
)

func Query(client *dgo.Dgraph, qFormat string, decode interface{}, uid uint64) error {
	res, err := client.NewTxn().QueryWithVars(context.Background(), qFormat, map[string]string{"$uid": utils.StringUid(uid)})
	if err != nil {
		return errx.Wrapf(errx.ERROR_DB, "failed to query uid:%v from dgraph", uid)
	}

	err = json.Unmarshal(res.Json, decode)
	if err != nil {
		return errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to unmarshall json: %v, err:%v", string(res.Json), err)
	}

	return nil
}
