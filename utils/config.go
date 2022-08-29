package utils

import (
	"github.com/jinzhu/configor"
)

var Config = ParserConfig{}

type ParserConfig struct {
	AppName       string
	AppDebug      bool
	Coin          string
	OutputDir     string `default:"./outputs"`
	Nodes         map[string][]Node
	StatsDatabase map[string]MySQLDB
	ParseRedis    Redis
	Log           Log
	Pyroscope     Pyroscope
	Start         uint64
}

type MySQLDSN struct {
	Name    string
	DSN     string
	Type    string
	SSHName string
}

// MySQLDB 代表一个 mysql 连接信息
type MySQLDB struct {
	Read     MySQLDSN
	Write    MySQLDSN
	Timezone string
	Region   string
	CoinType string
}

type Node struct {
	Addr  string
	Chain string
	Type  string
}

type Pyroscope struct {
	Enabled bool
	Address string
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
type Log struct {
	Level    string
	Path     string
	Filename string
}

func InitConfig(filename string) {
	err := configor.Load(&Config, filename)
	if err != nil {
		panic(err)
	}
}

func AddDatabaseConfig(value *MySQLDB, configs map[string]MySQLDSN) {
	if value.Read.DSN != "" && value.Read.Name != "" {
		configs[value.Read.Name] = MySQLDSN{DSN: value.Read.DSN, SSHName: value.Read.SSHName, Type: value.Read.Type}
	}
	if value.Write.DSN != "" && value.Write.Name != "" {
		configs[value.Write.Name] = MySQLDSN{DSN: value.Write.DSN, SSHName: value.Read.SSHName, Type: value.Write.Type}
	}
}
