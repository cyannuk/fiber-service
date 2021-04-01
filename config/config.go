package config

import (
	"flag"
	"fmt"
)

type databaseConfig struct {
	host           string
	user           string
	password       string
	dbname         string
	ssl            bool
	connectTimeout uint
}

type serverConfig struct {
	address string
	port    uint
	cert    string
	key     string
}

type DatabaseConfig interface {
	ConnectionString() string
}

type ServerConfig interface {
	BindAddress() string
	Certificate() (string, string)
}

func (config *databaseConfig) ConnectionString() string {
	var sslmode string
	if config.ssl {
		sslmode = "enable"
	} else {
		sslmode = "disable"
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%d",
		config.host, config.user, config.password, config.dbname, sslmode, config.connectTimeout)
}

func (config *serverConfig) BindAddress() string {
	return fmt.Sprintf("%s:%d", config.address, config.port)
}

func (config *serverConfig) Certificate() (string, string) {
	return config.cert, config.key
}

func GetDatabaseConfig() DatabaseConfig {
	return &databaseCfg
}

func GetServerConfig() ServerConfig {
	return &serverCfg
}

var databaseCfg databaseConfig
var serverCfg serverConfig

func init() {
	flag.StringVar(&serverCfg.address, "address", "localhost", "Bind host address")
	flag.UintVar(&serverCfg.port, "port", 8080, "Listen port")
	flag.StringVar(&serverCfg.cert, "cert", "", "SSL certificate file")
	flag.StringVar(&serverCfg.key, "key", "", "Certificate key")

	flag.StringVar(&databaseCfg.host, "db_host", "localhost", "Database host address")
	flag.StringVar(&databaseCfg.user, "db_user", "", "Database user name")
	flag.StringVar(&databaseCfg.password, "db_password", "", "Database user password")
	flag.StringVar(&databaseCfg.dbname, "db_name", "", "Database name")
	flag.BoolVar(&databaseCfg.ssl, "ssl", false, "SSL mode")
	flag.UintVar(&databaseCfg.connectTimeout, "connect_timeout", 2, "Database connect timeout")

	flag.Parse()
}
