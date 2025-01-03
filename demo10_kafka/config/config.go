package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Env string

const (
	ConfigEnvDebug   Env = "debug"
	ConfigEnvRelease Env = "release"
)

// mysql connection's info
type Mysql struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         string `mapstructure:"port" json:"port"`
	Database     string `mapstructure:"database" json:"database"`
	User         string `mapstructure:"user" json:"user"`
	Password     string `mapstructure:"password" json:"password"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns"`
}

// redisdb connection's info
type Redis struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Database string `mapstructure:"database" json:"database"`
	Password string `mapstructure:"password" json:"password"`
}

// kafka connection's info
type Kafka struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
}

//type ElasticSearch struct {
//	Host string `mapstructure:"host" json:"host"`
//}

// the system's info
type System struct {
	Port string `mapstructure:"port" json:"port"`
	Mode string `mapstructure:"mode" json:"mode"`
}

// tls
//type Tls struct {
//	Enable bool   `mapstructure:"enable" json:"enable"`
//	Cert   string `mapstructure:"cert" json:"cert"`
//	Key    string `mapstructure:"key" json:"key"`
//}

// logger
type Logger struct {
	Stdout    bool     `mapstructure:"stdout" json:"stdout"`
	Level     string   `mapstructure:"level" json:"level"`
	Dir       string   `mapstructure:"dir" json:"dir"`
	Rotation  bool     `mapstructure:"rotation" json:"rotation"`
	LogMaxAge int      `mapstructure:"logMaxAge" json:"logMaxAge"`
	LogTypes  []string `mapstructure:"logTypes" json:"logTypes"`
}

// all configs's info
type Config struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis"`
	System System `mapstructure:"system" json:"system"`
	Logger Logger `mapstructure:"logger" json:"logger"`
	Kafka  Kafka  `mapstructure:"kafka" json:"kafka"`
}

var (
	CONFIG Config
	VP     *viper.Viper
)

// the init
func Init() {
	// todo read the configs.json file
	v := viper.New()
	v.SetConfigFile("../config/config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		// todo log error to be handle with unity
		fmt.Println(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("configs file changed:", e.Name)
		if err = v.Unmarshal(&CONFIG); err != nil {
			// todo log error to be handle with unity
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&CONFIG); err != nil {
		// todo log error to be handle with unity
		fmt.Println(err)
	}
	VP = v
}
