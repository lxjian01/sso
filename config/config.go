package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	Version    string        `yaml:"version"`
	Env        string        `yaml:"env"`
	Server     ServerConfig  `yaml:"server"`
	LinuxUser  string        `yaml:"linuxUser"`
	PoolNum    int           `yaml:"poolNum"`
	Log        LogConfig     `yaml:"log"`
	Mysql      MysqlConfig   `yaml:"mysql"`
	Redis      RedisConfig   `yaml:"redis"`
}

type ServerConfig struct {
	Host string
	Port int
}

type LogConfig struct {
	Dir       string
	Name      string
	Format    string
	RetainDay int8
	Level     string
}

type MysqlConfig struct {
	Host        string
	Port        int
	Db          string
	User        string
	Password    string
	Charset     string
}

type RedisConfig struct {
	Host        string
	Port        int
	PoolSize    int
	MaxRetries  int
}

func InitConfig() *AppConfig {
	var appConf *AppConfig
	if err :=viper.Unmarshal(&appConf); err !=nil{
		fmt.Errorf("Unmarshal config error by %v \n",err)
		os.Exit(1)
	}
	return appConf
}