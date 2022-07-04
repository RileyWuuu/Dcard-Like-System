package config

import (
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once
var Conf *config

type config struct {
	MySql *mysql `yaml:"MySql"`
	Mongo *mongo `yaml:"Mongo"`
	Redis *redis `yaml:"Redis"`
}

type mysql struct {
	Addr       string `yaml:"Addr"`
	DriverName string `yaml:"DriverName"`
	UserName   string `yaml:"UserName"`
	Password   string `yaml:"Password"`
	DbName     string `yaml:"DbName"`
}

type mongo struct {
	Addr string `yaml:"Addr"`
}

type redis struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
}

func NewFromViper() (*config, error) {
	var c config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func SetConfig(c *config) {
	once.Do(func() {
		Conf = c
	})
}
