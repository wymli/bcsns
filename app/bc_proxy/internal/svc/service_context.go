package svc

import (
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/segmentio/kafka-go"
	"github.com/wymli/bcsns/app/bc_proxy/internal/config"
	bcsns "github.com/wymli/bcsns/smart_contract"

	"github.com/wymli/bcsns/common/logx"
)

type ServiceContext struct {
	Config         config.Config
	EtherClient    *ethclient.Client
	BcsnsClient    *bcsns.Bcsns
	KafkaProducer  *kafka.Writer
	BCAuth         *bind.TransactOpts
	SysAccountAddr common.Address // 系统区块链账户地址
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Init(c.Log)

	// 以太网节点客户端
	ethClient, err := ethclient.Dial(c.BlockChain.NodeEndpoint)
	if err != nil {
		panic(err)
	}

	// 合约客户端
	bcsnsClient, err := bcsns.NewBcsns(common.HexToAddress(c.BlockChain.ContractAddr), ethClient)
	if err != nil {
		panic(err)
	}

	godPrivKeyStr, addr, err := KeystoreToPrivateKey(c.BlockChain.KeyStoreFile, "")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(godPrivKeyStr)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:      c,
		EtherClient: ethClient,
		BcsnsClient: bcsnsClient,
		KafkaProducer: &kafka.Writer{
			Addr:     kafka.TCP(c.Kafka.Broker.Endpoints...),
			Balancer: &kafka.LeastBytes{},
		},
		BCAuth:         bind.NewKeyedTransactor(privateKey),
		SysAccountAddr: common.HexToAddress(addr),
	}
}

func KeystoreToPrivateKey(privateKeyFile, password string) (privKey string, address string, err error) {
	keyjson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		return "", "", err
	}
	unlockedKey, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		return "", "", err
	}
	privKey = common.Bytes2Hex(unlockedKey.PrivateKey.D.Bytes())
	addr := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)

	return privKey, addr.String(), nil
}
