package main

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	mrand "math/rand"
	"reflect"
	"strconv"
	"sync"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	ecies "github.com/ecies/go"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gocql/gocql"
	"github.com/wsddn/go-ecdh"
	seq "github.com/wymli/bcsns/app/seq_service/pb"
	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type User struct {
	Phone     int    `json:"phone"`
	Nickname  string `json:"nickname"`
	Sex       int    `json:"sex"`
	Age       int    `json:"age"`
	Avatar    string `json:"avatar"`
	Address   string `json:"address"`
	PublicKey string `json:"public_key"`
	Password  string `json:"password"`
}

var (
	testPhone = 11122133
	testPwd   = "123456789"
)

func main() {
	// uid := condMutate()
	// query(uid)
	// checkpwd(uid)
	// testEncrypt()
	// testEtherEncrypt()
	// testCassandra()
	// testECIES()
	testCompress()
}

func testCompress() {
	buf := bytes.NewBuffer(nil)
	nums := []uint64{}
	for i := 0; i < 500; i++ {
		nums = append(nums, mrand.Uint64()%10000000)
	}

	for _, num := range nums {
		err := binary.Write(buf, binary.BigEndian, num)
		if err != nil {
			panic(err)
		}
	}

	outBuf := bytes.NewBuffer(nil)
	w, err := flate.NewWriter(outBuf, flate.BestCompression)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	n, err := w.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
	w.Flush()
	fmt.Println("n:", n)

	fmt.Println("after compress: len:", len(outBuf.Bytes()))

	r := flate.NewReader(outBuf)
	defer r.Close()

	for {
		var num uint64
		err := binary.Read(r, binary.BigEndian, &num)
		if errors.Is(err, io.ErrUnexpectedEOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		// fmt.Println(num)
	}

	fmt.Println("using putVarint:")
	bbb := make([]byte, 10000)
	x := bbb[:]
	cnt := 0
	for _, num := range nums {
		n := binary.PutUvarint(x, num)
		x = x[n:]
		cnt += n
	}
	fmt.Println("len:", cnt)
	rrr := bytes.NewReader(bbb[:cnt])

	for {
		_, err := binary.ReadUvarint(rrr)
		if errors.Is(err, io.EOF) {
			break
		}
		// fmt.Println(num)
	}

	outBufGzip := bytes.NewBuffer(nil)
	ww, err := gzip.NewWriterLevel(outBufGzip, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	ww.Write(buf.Bytes())
	ww.Flush()
	defer ww.Close()

	fmt.Println("using gzip, after compress: len:", len(outBufGzip.Bytes()))

	// rr, err := gzip.NewReader(outBufGzip)
	// if err != nil {
	// 	panic(err)
	// }

}

func testECIES() {
	k, err := ecies.GenerateKey()
	if err != nil {
		panic(err)
	}
	log.Println("key pair has been generated")

	ciphertext, err := ecies.Encrypt(k.PublicKey, []byte("THIS IS THE TEST"))
	if err != nil {
		panic(err)
	}
	log.Printf("plaintext encrypted: %v\n", ciphertext)

	plaintext, err := ecies.Decrypt(k, ciphertext)
	if err != nil {
		panic(err)
	}
	log.Printf("ciphertext decrypted: %s\n", string(plaintext))
}

func testCassandra() {
	clusterConf := gocql.NewCluster("localhost")
	clusterConf.Keyspace = "bcsns"
	clusterConf.Consistency = gocql.LocalOne

	sess, err := clusterConf.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	// insert is an insert or update, mean upsert
	if err := sess.Query(`INSERT INTO messages (uid, send_uid, room_id, server_msg_id, send_msg_id, content_type, data) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		1, 2, 999, 2022, 2021, "text", "i love u").Exec(); err != nil {
		log.Fatal(err)
	}

	type MSG struct {
		uid           uint64
		send_uid      uint64
		room_id       uint64
		server_msg_id int64
		send_msg_id   int64
		content_type  string
		data          []byte
	}
	msg := MSG{}

	iter := sess.Query("SELECT uid, send_uid, room_id, server_msg_id, send_msg_id, content_type, data FROM messages WHERE server_msg_id > ? allow filtering", 2020).Iter()
	ok := iter.Scan(&msg.uid, &msg.send_uid, &msg.room_id, &msg.server_msg_id, &msg.send_msg_id, &msg.content_type, &msg.data)
	fmt.Printf("msg:%#v data:%v ok:%v\n", msg, string(msg.data), ok)

	msg = MSG{}
	iter = sess.Query("SELECT uid, send_uid, room_id, server_msg_id, send_msg_id, content_type, data FROM messages WHERE uid = ? ORDER BY server_msg_id DESC", 1).Iter()
	ok = iter.Scan(&msg.uid, &msg.send_uid, &msg.room_id, &msg.server_msg_id, &msg.send_msg_id, &msg.content_type, &msg.data)
	fmt.Printf("msg:%#v data:%v ok:%v\n", msg, string(msg.data), ok)

	msg = MSG{}
	iter = sess.Query("SELECT uid, send_uid, room_id, server_msg_id, send_msg_id, content_type, data FROM messages WHERE uid = ? AND server_msg_id > ?", 1, 1).Iter()
	ok = iter.Scan(&msg.uid, &msg.send_uid, &msg.room_id, &msg.server_msg_id, &msg.send_msg_id, &msg.content_type, &msg.data)
	fmt.Printf("msg:%#v data:%v ok:%v\n", msg, string(msg.data), ok)
}

func testEncrypt() {
	e := ecdh.NewEllipticECDH(crypto.S256())
	priK1, pubK1, err := e.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	priK2, pubK2, err := e.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	k1, err := e.GenerateSharedSecret(priK1, pubK2)
	if err != nil {
		log.Fatal(err)
	}

	k2, err := e.GenerateSharedSecret(priK2, pubK1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(k1, k2, reflect.DeepEqual(k1, k2))
}

func testEtherEncrypt() {
	curve := crypto.S256()
	priK1, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	pubK1 := priK1.PublicKey

	priK2, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	pubK2 := priK2.PublicKey

	fmt.Println(priK1, pubK1)
	fmt.Println(priK2, pubK2)

	x1, _ := curve.ScalarMult(pubK1.X, pubK1.Y, priK2.D.Bytes())
	k1 := x1.Bytes()

	x2, _ := curve.ScalarMult(pubK2.X, pubK2.Y, priK1.D.Bytes())
	k2 := x2.Bytes()
	fmt.Println(k1, k2, reflect.DeepEqual(k1, k2))
}

func testSeq() {
	c := seq.NewSeqClient(zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "localhost:9009",
	}))
	done := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		done.Add(1)
		ii := i
		go func() {
			defer done.Done()
			l := []uint64{}
			for j := 0; j < 1; j++ {
				rsp, err := c.GetSeqId(context.TODO(), &seq.GetSeqIdReq{
					UserId: uint64(ii),
				})
				if err != nil {
					panic(err)
				}
				l = append(l, rsp.SeqId)
				fmt.Println(ii, rsp, err)
			}
			for x := range l[:len(l)-1] {
				if l[x] > l[x+1] {
					panic(fmt.Sprintf("bad:%v", l))
				}
			}
		}()
	}
	done.Wait()
}

func checkpwd(uid string) {
	qLogin := `
query check($phone: int, $password: string){
  check(func: eq(phone, $phone)) {
		uid
    pass: checkpwd(password, $password)
  }
}
`
	c := newDgraphClient("localhost:9080")

	fmt.Printf("check uid %v\n", uid)

	txn := c.NewTxn()
	res, err := txn.QueryWithVars(context.Background(), qLogin,
		map[string]string{"$phone": strconv.Itoa(testPhone), "$password": testPwd})
	if err != nil {
		fmt.Println(err)
	}
	var decode struct {
		Check []struct {
			Uid  string
			Pass bool
		}
	}
	fmt.Printf("%s\n", string(res.Json))

	err = json.Unmarshal(res.Json, &decode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", decode)
}

func query(uid string) {
	q := `query all($a: string) {
    all(func: uid($a)) {
			uid
      phone
			nickname
			sex
			age
			avatar
			address
			public_key
			password
    }
  }`

	c := newDgraphClient("localhost:9080")

	txn := c.NewTxn()
	fmt.Println("querying uid ", uid)
	res, err := txn.QueryWithVars(context.Background(), q, map[string]string{"$a": uid})
	if err != nil {
		fmt.Println(err)
	}
	var decode struct {
		All []User
	}
	fmt.Printf("%s\n", string(res.Json))

	err = json.Unmarshal(res.Json, &decode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", decode)
	// {"all":[{"uid":"0x1a","phone":1,"nickname":"apple","sex":1,"age":12,"avatar":"s","address":"asdfsadf","public_key":"asdfasdfs"}]}
	// struct { All []main.User }{All:[]main.User{main.User{Phone:1, Nickname:"apple", Sex:1, Age:12, Avatar:"s", Address:"asdfsadf", PublicKey:"asdfasdfs", Password:""}}}
}

func condMutate() string {
	user := User{
		Phone:     testPhone,
		Nickname:  "apx",
		Sex:       1,
		Age:       12,
		Avatar:    "s",
		Address:   "asdfsadf",
		PublicKey: "asdfasdfs",
		Password:  testPwd,
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("failed to marshall json, err:%v", err)
		return ""
	}

	qFormat := `
  query {
      user as var(func: eq(phone, "%d"))
  }`
	query := fmt.Sprintf(qFormat, user.Phone)
	mu := &api.Mutation{
		Cond: `@if(eq(len(user), 0))`, // Only mutate if "wrong_email@dgraph.io" belongs to single user.
		// SetNquads: []byte(`uid(user) <email> "correct_email@dgraph.io" .`),
		SetJson: userJson,
	}
	apiReq := &api.Request{
		Query:     query,
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}

	c := newDgraphClient("localhost:9080")

	// Update email only if exactly one matching uid is found.
	res, err := c.NewTxn().Do(context.Background(), apiReq)
	if err != nil {
		// failed to do txn, err:rpc error: code = Unknown desc = Password too short, i.e. should have at least 6 chars<nil>
		fmt.Printf("failed to do txn, err:%v", err)
	}
	fmt.Println(res)
	// get uid
	// 如果插入成功.会返回uid,否则uid是空
	var uid string
	for k, v := range res.Uids {
		fmt.Println("uid_k:", k, "  uid_v:", v)
		uid = v
	}
	// json:"{}" txn:<start_ts:866 commit_ts:867 preds:"1-0-address" preds:"1-0-age" preds:"1-0-avatar" preds:"1-0-nickname" preds:"1-0-password" preds:"1-0-phone" preds:"1-0-public_key" preds:"1-0-sex" > latency:<parsing_ns:1313800 processing_ns:78688900 encoding_ns:100600 assign_timestamp_ns:834600 total_ns:81114700 > metrics:<num_uids:<key:"" value:1 > num_uids:<key:"_total" value:11 > num_uids:<key:"mutation_cost" value:9 > num_uids:<key:"phone" value:0 > num_uids:<key:"user" value:1 > > uids:<key:"dg.3488257281.20" value:"0x1a" > hdrs:<key:"content-type" value:<value:"application/grpc" > > hdrs:<key:"dgraph-toucheduids" value:<value:"11" > >
	return uid
}

func newDgraphClient(endpoint string) *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial(endpoint, grpc.WithInsecure())
	logx.FatalIfErrf(err, "failed to dial dgraph endpoint:%s", endpoint)

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}
