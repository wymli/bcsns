/*
	第三方组件的配置结构
*/

package config

import (
	"fmt"
	"io/ioutil"

	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/metric"
	"github.com/wymli/bcsns/common/tracer"
	"gopkg.in/yaml.v2"
)

type RpcServerConfig struct {
	NormalServerConfig `yaml:",inline"`
	ListenOn           string `yaml:"listen_on,omitempty"`
}

type NormalServerConfig struct {
	ServiceDiscovery ServiceDiscovery `yaml:"service_discovery,omitempty"`
	Trace            tracer.Config    `yaml:"trace,omitempty"`
	Metric           metric.Config    `yaml:"metric,omitempty"`
	Name             string           `yaml:"name,omitempty"`
	Mode             string           `yaml:"mode,omitempty"`
	Log              logx.Config      `yaml:"log,omitempty"`
	MockAll          bool             `yaml:"mock_all,omitempty"`
}

type ServiceDiscovery struct {
	Discovery string   `yaml:"discovery,omitempty"`
	Endpoints []string `yaml:"endpoints,omitempty"`
}

type RpcClientConfig struct {
	Name     string `yaml:"name,omitempty"`
	Endpoint string `yaml:"endpoint,omitempty"`
}

type KafkaConfig struct {
	Broker struct {
		Endpoints []string `yaml:"endpoints,omitempty"`
	} `yaml:"broker,omitempty"`

	GroupId string `yaml:"group_id,omitempty"`
}

type RedisConfig struct {
	Host     string `yaml:"host,omitempty"`
	Type     string `yaml:"type,omitempty"`
	Password string `yaml:"password,omitempty"`
}

type SnowflakeConfig struct {
	NodeNumber int64 `yaml:"node_number,omitempty"`
}

type BlockChainConfig struct {
	NodeEndpoint string `yaml:"node_endpoint,omitempty"`
	KeyStoreFile string `yaml:"key_store_file,omitempty"`
	ContractAddr string `yaml:"contract_addr,omitempty"`
}

type DeployConfig struct {
	Use string `yaml:"use,options=k8s|docker|docker-compose,omitempty"`
}

type CassandraConfig struct {
	Endpoints   []string `yaml:"endpoints,omitempty"`
	KeySpace    string   `yaml:"key_space,omitempty"`
	Consistency uint16   `yaml:"consistency,omitempty"`
}

type DgraphConfig struct {
	Endpoint string `yaml:"endpoint,omitempty"`
}

type JwtConfig struct {
	AccessSecret string `yaml:"access_secret,omitempty"`
	AccessExpire int64  `yaml:"access_expire,omitempty"`
}

func LoadConfig(fileName string, config interface{}) error {
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(in, config)
	if err != nil {
		return err
	}

	return nil
}

func MustLoadConfig(fileName string, config interface{}) {
	if err := LoadConfig(fileName, config); err != nil {
		panic(err)
	}

	fmt.Printf("config: %#v\n", config)
}
