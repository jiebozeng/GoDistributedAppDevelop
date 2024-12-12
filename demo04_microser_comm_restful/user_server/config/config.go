package configs

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

type ElasticSearch struct {
	Host string `mapstructure:"host" json:"host"`
}

// the cache info
type Cache struct {
	//UserBlackListKey  string `mapstructure:"user_black_list_key" json:"user_black_list_key"`
	//AdminBlackListKey string `mapstructure:"user_black_list_key" json:"user_black_list_key"`
}

// the user's token info
type UserToken struct {
	SginKey      string `mapstructure:"sign_key" json:"sign_key"`
	STTL         int    `mapstructure:"sttl" json:"sttl"`
	BlackListKey string `mapstructure:"black_list_key" json:"black_list_key"`
	TTL          int    `mapstructure:"ttl" json:"ttl"`
}

// the admin's token info
type AdminToken struct {
	SginKey      string `mapstructure:"sign_key" json:"sign_key"`
	STTL         int    `mapstructure:"sttl" json:"sttl"`
	BlackListKey string `mapstructure:"black_list_key" json:"black_list_key"`
	TTL          int    `mapstructure:"ttl" json:"ttl"`
	Issuer       string `mapstructure:"issuser" json:"issuser"`
}

// the system's info
type System struct {
	Port          string `mapstructure:"port" json:"port"`
	Mode          string `mapstructure:"mode" json:"mode"`
	LogPath       string `mapstructure:"logPath" json:"logPath"`
	LogMaxSize    string `mapstructure:"logMaxSize" json:"logMaxSize"`
	LogMaxBackups string `mapstructure:"logMaxBackups" json:"logMaxBackups"`
	LogMaxAge     string `mapstructure:"logMaxAge" json:"logMaxAge"`
	LogCompress   string `mapstructure:"logCompress" json:"logCompress"`
}

// tls
type Tls struct {
	Enable bool   `mapstructure:"enable" json:"enable"`
	Cert   string `mapstructure:"cert" json:"cert"`
	Key    string `mapstructure:"key" json:"key"`
}

// logger
type Logger struct {
	Stdout    bool     `mapstructure:"stdout" json:"stdout"`
	Level     string   `mapstructure:"level" json:"level"`
	Dir       string   `mapstructure:"dir" json:"dir"`
	Rotation  bool     `mapstructure:"rotation" json:"rotation"`
	LogMaxAge int      `mapstructure:"logMaxAge" json:"logMaxAge"`
	LogTypes  []string `mapstructure:"logTypes" json:"logTypes"`
}

// daylog
type Daylog struct {
	Filepath string `mapstructure:"filepath" json:"filepath"`
	Rotation bool   `mapstructure:"rotation" json:"rotation"`
	Daily    bool   `mapstructure:"daily" json:"daily"`
	Hourly   bool   `mapstructure:"hourly" json:"hourly"`
}

type HuaWeiOBS struct {
	SsgfimagesAccessUrl   string `mapstructure:"SsgfimagesAccessUrl" json:"ssgfimagesaccessurl"`
	OBSImageBasePath      string `mapstructure:"OBSImageBasePath" json:"obsimagebasepath"`
	OBSImageBasePathDebug string `mapstructure:"OBSImageBasePathDebug" json:"obsimagebasepathdebug"`
}

// all configs's info
type Config struct {
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql"`
	Redis      Redis      `mapstructure:"redis" json:"redis"`
	AdminToken AdminToken `mapstructure:"admin_token" json:"admin_token"`
	//UserToken  UserToken  `mapstructure:"user_token" json:"user_token"`
	System System `mapstructure:"system" json:"system"`
	//Cache  Cache  `mapstructure:"cache" json:"cache"`
	Tls           Tls           `mapstructure:"tls" json:"tls"`
	Logger        Logger        `mapstructure:"logger" json:"logger"`
	Daylog        Daylog        `mapstructure:"daylog" json:"daylog"`
	Kafka         Kafka         `mapstructure:"kafka" json:"kafka"`
	ElasticSearch ElasticSearch `mapstructure:"elasticsearch" json:"elasticsearch"`
	Env           Env           `mapstructure:"env" json:"env"`
	HuaWeiOBS     HuaWeiOBS     `mapstructure:"HuaWeiOBS" json:"huaweiobs"`
}

var (
	CONFIG Config
	VP     *viper.Viper
)

// the init
func Init() {
	// todo read the configs.json file
	v := viper.New()
	v.SetConfigFile("configs/config.yaml")
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
